package grade

import (
	//"encoding/json"
	//"fmt"
	//"net/http"
	//"strconv"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Grade struct {
	Ref     int
	OrderNo string
	Name    string
}

func Get(urldb string) []Grade {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `select ref, rank, name from grade order by rank`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var grades []Grade
	var grade Grade

	for rows.Next() {
		err := rows.Scan(
			&grade.Ref,
			&grade.OrderNo,
			&grade.Name,
		)

		if err != nil {
			panic(err)
		}

		grades = append(grades, grade)
	}

	return grades
}

func GetByRef(urldb string, ref string) Grade {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `select ref, rank, name from grade where ref=` + ref

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var grade Grade

	for rows.Next() {
		err := rows.Scan(
			&grade.Ref,
			&grade.OrderNo,
			&grade.Name,
		)

		if err != nil {
			panic(err)
		}

	}

	return grade
}

func Save(urldb string, grade Grade) []Grade {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	if grade.Ref == 0 {
		insert(db, grade)
	} else {
		update(db, grade)
	}

	return Get(urldb)
}

func insert(db *sql.DB, grade Grade) {

	sql := `
		insert into grade (
			rank,
			name
		) values(?,?)`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		grade.OrderNo,
		grade.Name)

	if err != nil {
		panic(err)
	}
}

func update(db *sql.DB, grade Grade) {
	sql := `
		update grade set
			rank=?,
			name=?
		where ref=?`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		grade.OrderNo,
		grade.Name,
		grade.Ref)

	if err != nil {
		panic(err)
	}

}

func Remove(urldb string, grade Grade) []Grade {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		delete from grade 
		where ref=?`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(grade.Ref)

	if err != nil {
		panic(err)
	}

	return Get(urldb)
}
