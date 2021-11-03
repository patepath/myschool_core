package subjectgroup

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type SubjectGroup struct {
	Ref  int
	Name string
}

func Get(urldb string) []SubjectGroup {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := ` select ref, name from subject_group order by name; `

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var subjectgrps []SubjectGroup
	var subjectgrp SubjectGroup

	for rows.Next() {
		err := rows.Scan(&subjectgrp.Ref, &subjectgrp.Name)

		if err != nil {
			panic(err)
		}

		subjectgrps = append(subjectgrps, subjectgrp)
	}

	return subjectgrps
}

func insert(urldb string, subjectgroup SubjectGroup) {
	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `insert into subject_group(name) values('` + subjectgroup.Name + `');`

	_, err = db.Exec(sql)

	if err != nil {
		panic(err)
	}
}

func update(urldb string, subjectgroup SubjectGroup) {
	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `update subject_group set name=? where ref=?;`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(subjectgroup.Name, subjectgroup.Ref)

	if err != nil {
		panic(err)
	}
}

func Save(urldb string, subjectgroup SubjectGroup) []SubjectGroup {

	if subjectgroup.Ref == 0 {
		insert(urldb, subjectgroup)

	} else {
		update(urldb, subjectgroup)
	}

	return Get(urldb)
}

func Del(urldb string, subjectgroup SubjectGroup) []SubjectGroup {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `delete from subject_group where ref=` + strconv.Itoa(subjectgroup.Ref) + `;`

	_, err = db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return Get(urldb)
}
