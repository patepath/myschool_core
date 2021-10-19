package register

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type RegView struct {
	Ref       int
	IdCard    string
	Grade     string
	Room      string
	Code      string
	FirstName string
	LastName  string
}

func Get(urlregdb string) []RegView {

	db, err := sql.Open("mysql", urlregdb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			ref, 
			grade, 
			room, 
			ifnull(code, ''), 
			student_firstname, 
			student_lastname, 
			student_idcard 
		from students 
		group by student_idcard 
		order by grade, room, code;
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var regViews []RegView
	var regView RegView

	for rows.Next() {
		err := rows.Scan(
			&regView.Ref,
			&regView.Grade,
			&regView.Room,
			&regView.Code,
			&regView.FirstName,
			&regView.LastName,
			&regView.IdCard,
		)

		if err != nil {
			panic(err)
		}

		regViews = append(regViews, regView)
	}

	return regViews
}
