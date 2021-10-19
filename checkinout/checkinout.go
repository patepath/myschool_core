package checkinout

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Checkinout struct {
	Ref             int
	DateTime        string
	StudentCode     string
	StudentFullName string
	CamNo           string
	Temperature     float32
	CutPoint        int
	Status          int
}

func Get(urldb string) []Checkinout {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query(`
		select  
			checkinout.ref, 
			created as datetime, 
			checkinout.idcard as studentcode, 
			concat(ifnull(students.first_name, ''), ' ', ifnull(students.last_name,'')) as studentfullname, 
			camno, 
			temperature, 
			ifnull(cutpoint, 0),
			0 as status 
		from checkinout 
		left join students on checkinout.idcard = students.code
		order by created desc
		limit 4000;
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var checkinouts []Checkinout
	var checkinout Checkinout

	for rows.Next() {
		err := rows.Scan(
			&checkinout.Ref,
			&checkinout.DateTime,
			&checkinout.StudentCode,
			&checkinout.StudentFullName,
			&checkinout.CamNo,
			&checkinout.Temperature,
			&checkinout.CutPoint,
			&checkinout.Status,
		)

		if err != nil {
			panic(err)
		}

		checkinouts = append(checkinouts, checkinout)
	}

	return checkinouts
}

func GetByRef(urldb string, ref string) Checkinout {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select  
			checkinout.ref, 
			created as datetime, 
			checkinout.idcard as studentcode, 
			concat(ifnull(students.first_name, ''), ' ', ifnull(students.last_name,'')) as studentfullname, 
			camno, 
			temperature, 
			ifnull(cutpoint, 0),
			0 as status 
		from checkinout 
		left join students on checkinout.idcard = students.code
		where checkinout.ref = '` + ref + `'
		limit 1;
	`

	row, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer row.Close()

	var checkinout Checkinout

	for row.Next() {
		err := row.Scan(
			&checkinout.Ref,
			&checkinout.DateTime,
			&checkinout.StudentCode,
			&checkinout.StudentFullName,
			&checkinout.CamNo,
			&checkinout.Temperature,
			&checkinout.CutPoint,
			&checkinout.Status,
		)

		if err != nil {
			panic(err)
		}
	}

	return checkinout
}

func GetByCode(urldb string, code string) []Checkinout {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query(`
		select  
			checkinout.ref, 
			created as datetime, 
			checkinout.idcard as studentcode, 
			concat(ifnull(students.first_name, ''), ' ', ifnull(students.last_name,'')) as studentfullname, 
			camno, 
			temperature, 
			ifnull(cutpoint, 0),
			0 as status 
		from checkinout 
		left join students on checkinout.idcard = students.code
		where checkinout.idcard = '` + code + `'
		order by created desc;
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var checkinouts []Checkinout
	var checkinout Checkinout

	for rows.Next() {
		err := rows.Scan(
			&checkinout.Ref,
			&checkinout.DateTime,
			&checkinout.StudentCode,
			&checkinout.StudentFullName,
			&checkinout.CamNo,
			&checkinout.Temperature,
			&checkinout.CutPoint,
			&checkinout.Status,
		)

		if err != nil {
			panic(err)
		}

		checkinouts = append(checkinouts, checkinout)
	}

	return checkinouts
}

func main() {
	const urldb = "admin:35.103232@tcp(localhost:3306)/school"
	fmt.Println("%s", Get(urldb))
}
