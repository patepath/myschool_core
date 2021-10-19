package employee

import (
	//"encoding/json"
	//"fmt"
	//"net/http"
	//"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Ref	int
	Code string
	TitleRef int
	TitleName string
	FirstName string
	LastName string
	DeptRef int
	DeptName string
	GradeRef int
	GradeName string
	RoomRef int
	RoomName string
	FaceImg string
}

func Get(urldb string) []Employee {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select
			A.ref,
			A.code,
			title_ref,
			B.name as title_name,
			first_name,
			last_name,
			dept_ref,
			C.name as dept_name,
			grade_ref,
			E.name as grade_name,
			room_ref,
			D.name as room_name
		from employee as A
		left join title as B on B.ref = A.title_ref
		left join department as C on C.ref = A.dept_ref
		left join classroom as D on D.ref = A.room_ref
		left join grade as E on E.ref = D.grade_ref;
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var employees []Employee
	var employee Employee

	for rows.Next() {

		err := rows.Scan(
			&employee.Ref,
			&employee.Code,
			&employee.TitleRef,
			&employee.TitleName,
			&employee.FirstName,
			&employee.LastName,
			&employee.DeptRef,
			&employee.DeptName,
			&employee.GradeRef,
			&employee.GradeName,
			&employee.RoomRef,
			&employee.RoomName,
		)

		if err != nil {
			panic(err)
		}

		employees = append(employees, employee)
	}

	return employees
}

func Insert(urldb string, employee Employee) {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		insert into employee(
		code,
		title_ref,
		first_name,
		last_name,
		dept_ref,
		room_ref
		) values(?,?,?,?,?,?)
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		employee.Code,
		employee.TitleRef,
		employee.FirstName,
		employee.LastName,
		employee.DeptRef,
		employee.RoomRef,
	)

	if err != nil {
		panic(err)
	}

}


func Update(urldb string, employee Employee) {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		update employee set
			code=?,
			title_ref=?,
			first_name=?,
			last_name=?,
			dept_ref=?,
			room_ref=?
		where ref=?;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		employee.Code,
		employee.TitleRef,
		employee.FirstName,
		employee.LastName,
		employee.DeptRef,
		employee.RoomRef,
		employee.Ref,
	)

	if err != nil {
		panic(err)
	}
}

func Save(urldb string, employee Employee) []Employee {

	if employee.Ref == 0 {
		Insert(urldb, employee)

	} else {
		Update(urldb, employee)
	}

	return Get(urldb)
}
