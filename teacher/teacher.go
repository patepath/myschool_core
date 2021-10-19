package teacher

import (
	//"encoding/json"

	//"net/http"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Teacher struct {
	Ref       int
	DeptRef   int
	DeptName  string
	GradeRef  int
	GradeName string
	RoomRef   int
	RoomName  string
	Code      string
	FirstName string
	LastName  string
	Phone     string
}

func Get(urldb string) []Teacher {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			A.ref, 
			ifnull(D.ref, 0), 
			ifnull(D.name, ''), 
			ifnull(B.ref, 0), 
			ifnull(B.name, ''), 
			ifnull(C.ref, 0), 
			ifnull(C.name, ''), 
			A.code, 
			firstname, 
			lastname, 
			phone
		from teacher as A
		left join classroom as B on A.room_ref = B.ref
		left join grade as C on C.ref = B.grade_ref
		left join department as D on A.dept_ref = D.ref;
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var teachers []Teacher
	var teacher Teacher

	for rows.Next() {

		err := rows.Scan(
			&teacher.Ref,
			&teacher.DeptRef,
			&teacher.DeptName,
			&teacher.RoomRef,
			&teacher.RoomName,
			&teacher.GradeRef,
			&teacher.GradeName,
			&teacher.Code,
			&teacher.FirstName,
			&teacher.LastName,
			&teacher.Phone,
		)

		if err != nil {
			panic(err)
		}

		teachers = append(teachers, teacher)
	}

	return teachers
}

func GetByRef(urldb string, ref string) Teacher {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			A.ref, 
			ifnull(D.ref, 0), 
			ifnull(D.name, ''), 
			ifnull(B.ref, 0), 
			ifnull(B.name, ''), 
			ifnull(C.ref, 0), 
			ifnull(C.name, ''), 
			A.code, 
			firstname, 
			lastname, 
			phone
		from teacher as A
		left join classroom as B on A.room_ref = B.ref
		left join grade as C on C.ref = B.grade_ref
		left join department as D on A.dept_ref = D.ref
		where A.ref=` + ref

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var teacher Teacher

	if rows.Next() {

		err := rows.Scan(
			&teacher.Ref,
			&teacher.DeptRef,
			&teacher.DeptName,
			&teacher.RoomRef,
			&teacher.RoomName,
			&teacher.GradeRef,
			&teacher.GradeName,
			&teacher.Code,
			&teacher.FirstName,
			&teacher.LastName,
			&teacher.Phone,
		)

		if err != nil {
			panic(err)
		}
	}

	return teacher
}

func GetByGrade(urldb string, grade string) []Teacher {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			A.ref, 
			ifnull(D.ref, 0), 
			ifnull(D.name, ''), 
			ifnull(B.ref, 0), 
			ifnull(B.name, ''), 
			ifnull(C.ref, 0), 
			ifnull(C.name, ''), 
			A.code, 
			firstname, 
			lastname, 
			phone
		from teacher as A
		left join classroom as B on A.room_ref = B.ref
		left join grade as C on C.ref = B.grade_ref
		left join department as D on A.dept_ref = D.ref
		where C.ref = ` + grade + `
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var teachers []Teacher
	var teacher Teacher

	for rows.Next() {

		err := rows.Scan(
			&teacher.Ref,
			&teacher.DeptRef,
			&teacher.DeptName,
			&teacher.RoomRef,
			&teacher.RoomName,
			&teacher.GradeRef,
			&teacher.GradeName,
			&teacher.Code,
			&teacher.FirstName,
			&teacher.LastName,
			&teacher.Phone,
		)

		if err != nil {
			panic(err)
		}

		teachers = append(teachers, teacher)
	}

	return teachers
}

func insert(urldb string, teacher Teacher) {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		insert into teacher(
			dept_ref,
			room_ref,
			code,
			firstname,
			lastname,
			phone
		) values(?,?,?,?,?,?);
		
	`
	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		teacher.DeptRef,
		teacher.RoomRef,
		teacher.Code,
		teacher.FirstName,
		teacher.LastName,
		teacher.Phone,
	)

	if err != nil {
		panic(err)
	}
}

func update(urldb string, teacher Teacher) {

	fmt.Println(teacher)

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		update teacher set
			dept_ref=?,
			room_ref=?,
			code=?,
			firstname=?,
			lastname=?,
			phone=?
		where ref=?;
	`
	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		teacher.DeptRef,
		teacher.RoomRef,
		teacher.Code,
		teacher.FirstName,
		teacher.LastName,
		teacher.Phone,
		teacher.Ref,
	)

	if err != nil {
		panic(err)
	}
}

func Save(urldb string, teacher Teacher) []Teacher {

	if teacher.Ref == 0 {
		insert(urldb, teacher)
	} else {
		update(urldb, teacher)
	}

	return Get(urldb)
}

func Remove(urldb string, ref string, grade_ref string) []Teacher {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		delete from teacher where ref=?;
	`
	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(ref)

	if err != nil {
		panic(err)
	}

	return GetByGrade(urldb, grade_ref)
}
