//package main
package room

import (
	//"encoding/json"
	//"fmt"
	//"net/http"
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Room struct {
	Ref        int
	GradeRef   int
	GradeOrder int
	GradeName  string
	Name       string
}

func Get(urldb string) []Room {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			ifnull(classroom.ref, 0) as room_ref, 
			ifnull(grade.ref, 0) as grade_ref, 
			ifnull(grade.rank, 0) as grade_order, 
			ifnull(grade.name, '') as grade_name, 
			ifnull(classroom.name, '') as room_name 
		from classroom 
		left join grade on classroom.grade_ref = grade.ref
		order by grade.rank, classroom.ref`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var rooms []Room
	var room Room

	for rows.Next() {
		err := rows.Scan(
			&room.Ref,
			&room.GradeRef,
			&room.GradeOrder,
			&room.GradeName,
			&room.Name,
		)

		if err != nil {
			panic(err)
		}

		rooms = append(rooms, room)
	}

	return rooms
}

func GetByRef(urldb string, ref string) Room {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			classroom.ref as room_ref, 
			grade.ref as grade_ref, 
			grade.rank as grade_order, 
			grade.name as grade_name, 
			classroom.name as room_name 
		from classroom 
		left join grade on classroom.grade_ref = grade.ref
		where classroom.ref=?
		limit 1;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	param1, _ := strconv.Atoi(ref)

	rows, err := sttn.Query(param1)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var room Room

	if rows.Next() {
		err := rows.Scan(
			&room.Ref,
			&room.GradeRef,
			&room.GradeOrder,
			&room.GradeName,
			&room.Name,
		)

		if err != nil {
			panic(err)
		}

	}

	return room
}

func GetByGrade(urldb string, gradeRef string) []Room {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			classroom.ref as room_ref, 
			grade.ref as grade_ref, 
			grade.rank as grade_order, 
			grade.name as grade_name, 
			classroom.name as room_name 
		from classroom 
		left join grade on classroom.grade_ref = grade.ref
		where grade_ref=?
		order by grade.rank, classroom.ref`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	param1, _ := strconv.Atoi(gradeRef)

	rows, err := sttn.Query(param1)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var rooms []Room
	var room Room

	for rows.Next() {
		err := rows.Scan(
			&room.Ref,
			&room.GradeRef,
			&room.GradeOrder,
			&room.GradeName,
			&room.Name,
		)

		if err != nil {
			panic(err)
		}

		rooms = append(rooms, room)
	}

	return rooms
}

func Save(urldb string, room Room) []Room {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	if room.Ref == 0 {
		insert(db, room)
	} else {
		update(db, room)
	}

	return Get(urldb)
}

func insert(db *sql.DB, room Room) {

	sql := `insert into classroom (
				grade_ref,
				name
			) values(?,?)`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		room.GradeRef,
		room.Name)

	if err != nil {
		panic(err)
	}

}

func update(db *sql.DB, room Room) {

	sql := `
		update classroom
		set grade_ref=?, name=?
		where ref=?;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(room.GradeRef, room.Name, room.Ref)

	if err != nil {
		panic(err)
	}

}

func Remove(urldb string, room Room) []Room {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		delete from classroom
		where ref=?;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(room.Ref)

	if err != nil {
		panic(err)
	}

	return Get(urldb)
}

//func main() {
//	 var urldb = "admin:35.103232@tcp(localhost:3306)/school"
//
//	 GetByGrade(urldb, "1")
// }
