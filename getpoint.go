package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const urldb = "admin:35.103232@tcp(localhost:3306)/school"

func GetPoint(urldb string, datetime string, student_code string) int {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `select checkinprofile from students where code=?`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(student_code)

	if err != nil {
		panic(err)
	}

	var code string
	var point int

	if rows.Next() {
		rows.Scan(&code)
	}

	sql = `
		select time1, time2, time3, point1, point2, point3
		from checkinprofile
		where name=? and date_in=?;
	`

	sttn, err = db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	var time1 string
	var time2 string
	var time3 string

	//var t1 time.Time
	//	var t2 time.Time
	//	var t3 time.Time

	var point1 int
	var point2 int
	var point3 int

	d, err := time.Parse("2006-01-02 15:04:00", datetime)

	if err != nil {
		fmt.Println(err)
	}

	date := d.Format("2006-01-02")

	fmt.Println("date: ", datetime)
	fmt.Println("code: ", code, ", date: ", d)

	rows, err = sttn.Query(code, d.Format("2006-01-02"))

	if err != nil {
		panic(err)
	}

	for rows.Next() {
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
			point = 0
		} else if diff1.Minutes() > 0 && diff2.Minutes() <= 0 {
			point = point1
		} else if diff2.Minutes() > 0 && diff3.Minutes() <= 0 {
			point = point2
		} else {
			point = point3
		}

	}

	return point
}

func main() {
	var point int

	fmt.Println("start")
	point = GetPoint(urldb, "2021-07-05 8:10:00", "36622")

	fmt.Println("Point: ", point)

}
