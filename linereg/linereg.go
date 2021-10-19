package linereg

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type StudentReg struct {
	Code   string
	IdCard string
	UserId string
}

type Student struct {
	Ref       int
	Code      string
	FirstName string
	LastName  string
}

type Msg struct {
	Success bool
}

func RegisterStudent(urldb string, payload StudentReg) Msg {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			ref,
			code,
			first_name,
			last_name
		from students 
		where code='` + payload.Code + `' 
		limit 1;
	`
	row, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer row.Close()

	var msg Msg
	var student Student

	if row.Next() {

		err = row.Scan(
			&student.Ref,
			&student.Code,
			&student.FirstName,
			&student.LastName,
		)

		if err != nil {
			panic(err)
		}

		if getUserByCode(urldb, student.Code) == 0 {
			insertUser(urldb, student, payload)
			msg.Success = true

		} else {
			updateUser(urldb, student, payload)
			msg.Success = true
		}

	} else {
		msg.Success = false
	}

	return msg
}

func getUserByCode(urldb string, code string) int {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select ref
		from user
		where name='` + code + `'
		limit 1;
	`
	row, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer row.Close()

	var ref int

	if row.Next() {

		err := row.Scan(&ref)

		if err != nil {
			panic(err)
		}

		return ref
	}

	return 0
}

func insertUser(urldb string, student Student, payload StudentReg) bool {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		insert into user (
			code,
			line_userid,
			idcard,
			name,
			full_name,
			password,
			role
		) values ('',?,?,?,?,md5('123456789'),'student')
	`
	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		payload.UserId,
		payload.IdCard,
		payload.Code,
		student.FirstName+" "+student.LastName,
	)

	if err != nil {
		panic(err)
		return false
	}

	return true
}

func updateUser(urldb string, student Student, payload StudentReg) bool {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		update user set
			line_userid=?,
			idcard=?
		where name=?
	`
	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		payload.UserId,
		payload.IdCard,
		payload.Code,
	)

	if err != nil {
		panic(err)
		return false
	}

	return true
}
