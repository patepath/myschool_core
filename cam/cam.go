package cam

import (
	//"fmt"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Picture struct {
	Camno       int
	Datetime    string
	Code        string
	Temperature float32
	Faceimage   string
}

func GetSec(times string) int64 {

	time := strings.Split(times, ":")

	h, _ := strconv.ParseInt(time[0], 10, 32)
	m, _ := strconv.ParseInt(time[1], 10, 32)
	s, _ := strconv.ParseInt(time[2], 10, 32)

	sec := h*3600 + m*60 + s

	return sec
}

func GetCutPoint(urldb string, datetime string) float32 {

	dt := strings.Fields(datetime)
	//date := strings.Split(dt[0], "-")

	checkin := GetSec(dt[1])
	//fmt.Println("checkin: " , checkin)

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
	select time1, time2, time3, point1, point2, point3
	from checkinprofile
	order by grp, date_in
	limit 1
	`
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var time1 string
	var time2 string
	var time3 string
	var point1 float32
	var point2 float32
	var point3 float32
	var cutpoint float32

	for rows.Next() {
		err := rows.Scan(
			&time1,
			&time2,
			&time3,
			&point1,
			&point2,
			&point3,
		)

		if err != nil {
			panic(err)
		}
	}

	if checkin > GetSec(time1) && checkin <= GetSec(time2) {
		cutpoint = point1

	} else if checkin > GetSec(time2) && checkin <= GetSec(time3) {
		cutpoint = point2

	} else {
		cutpoint = point3
	}

	return cutpoint
}

func GetPoint(urldb string, datetime string, student_code string) (int, int) {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select checkinprofile from students where code=? limit 1;
	`
	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(student_code)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var code string
	var behavior_ref int
	var point int

	if rows.Next() {
		rows.Scan(&code)
	}

	sql = `
		select time1, time2, time3, point1, point2, point3
		from checkinprofile
		where name=? and date_in=?
		limit 1;
	`

	sttn, err = db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	var time1 string
	var time2 string
	var time3 string

	var point1 int
	var point2 int
	var point3 int

	d, err := time.Parse("2006-01-02 15:04:05", datetime)

	if err != nil {
		fmt.Println(err)
	}

	date := d.Format("2006-01-02")

	rows, err = sttn.Query(code, d.Format("2006-01-02"))

	if err != nil {
		panic(err)
	}

	if rows.Next() {

		rows.Scan(&time1, &time2, &time3, &point1, &point2, &point3)

		t1, err := time.Parse("2006-01-02 15:04:00", date+" "+time1)

		if err != nil {
			panic(err)
		}

		t2, err := time.Parse("2006-01-02 15:04:00", date+" "+time2)

		if err != nil {
			panic(err)
		}

		t3, err := time.Parse("2006-01-02 15:04:00", date+" "+time3)

		if err != nil {
			panic(err)
		}

		diff1 := d.Sub(t1)
		diff2 := d.Sub(t2)
		diff3 := d.Sub(t3)

		if diff1.Minutes() <= 0 {
			behavior_ref = 0
			point = 0

		} else if diff1.Minutes() > 0 && diff2.Minutes() <= 0 {
			behavior_ref = 50
			point = point1

		} else if diff2.Minutes() > 0 && diff3.Minutes() <= 0 {
			behavior_ref = 51
			point = point2

		} else {
			behavior_ref = 52
			point = point3
		}

	}

	return point, behavior_ref
}

func Save(urldb string, pic Picture) Picture {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, err := db.Prepare(`
		insert into checkinout (
			camno, 
			created, 
			idcard, 
			cutpoint,
			temperature )
		values (?,?,?,?,?)
	`)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	pic.Datetime = strings.ReplaceAll(strings.ReplaceAll(pic.Datetime, "T", " "), "+07:00", "")

	_, err = stmt.Exec(
		pic.Camno,
		pic.Datetime,
		pic.Code,
		GetCutPoint(urldb, pic.Datetime),
		pic.Temperature,
	)

	if err != nil {
		panic(err)
	}

	dt := strings.Fields(pic.Datetime)

	dates := strings.Split(dt[0], "-")
	year := dates[0]
	month := dates[1]
	day := dates[2]

	path := "imgs/checkinout/" + year + "/" + month + "/" + day + "/"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0777)
	}

	filename := dt[0] + "_" + dt[1] + "_" + pic.Code + ".b64"

	err = ioutil.WriteFile(path+filename, []byte(pic.Faceimage), 0777)

	if err != nil {
		panic(err)
	}

	// student point

	point, behavior_ref := GetPoint(urldb, pic.Datetime, pic.Code)

	if point != 0 {

		sql := `
			select student_code
			from point_students
			where date(created) = date(?) and student_code=?
			limit 1;
		`
		stmt, err = db.Prepare(sql)

		if err != nil {
			panic(err)
		}

		rows, err := stmt.Query(pic.Datetime, pic.Code)

		if err != nil {
			panic(err)
		}

		if !rows.Next() {

			sql = `
				insert into point_students (
					yearedu, 
					created, 
					behavior_ref, 
					student_code, 
					point
				) values(2563, ?,?,?,?);
			`

			stmt, err = db.Prepare(sql)

			if err != nil {
				panic(err)
			}

			_, err = stmt.Exec(pic.Datetime, behavior_ref, pic.Code, point)

			if err != nil {
				panic(err)
			}
		}
	}
	return pic
}
