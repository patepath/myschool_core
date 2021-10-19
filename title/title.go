package title
//package main

import (
	//"encoding/json"
	//"fmt"
	//"net/http"
	//"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Title struct {
	Ref int
	OrderNo int
	Name string
}

func Get(urldb string) []Title {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `select ref, rank, name from title order by rank`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var titles []Title
	var title Title

	for rows.Next() {
		err := rows.Scan(
			&title.Ref,
			&title.OrderNo,
			&title.Name,
		)

		if err != nil {
			panic(err)
		}

		titles = append(titles, title)
	}

	return titles

}

func Save(urldb string, title Title) []Title {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	if title.Ref == 0 {
		insert(db, title)

	} else {
		update(db, title)
	}

	return Get(urldb)
}

func insert(db *sql.DB, title Title) {

	sql:= `
		insert into title (
			rank,
			name
		) value(?,?) `

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		title.OrderNo,
		title.Name)

	if err != nil {
		panic(err)
	}
}

func update(db *sql.DB, title Title) {

	sql := `
		update title set
			rank=?,
			name=?
		where ref=?`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		title.OrderNo,
		title.Name,
		title.Ref)

	if err != nil {
		panic(err)
	}
}

//func main() {
//	const urldb = "admin:35.103232@tcp(localhost:3306)/school"
//
//	var title = Title {
//		Ref: 1,
//		OrderNo: 1,
//		Name: "นาย",
//	}
//
//	Save(urldb, title)
//
//}
