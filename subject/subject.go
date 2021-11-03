package subject

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Subject struct {
	Ref       int
	GroupRef  int
	GroupName string
	GradeRef  int
	GradeName string
	Code      string
	Name      string
}

func Get(urldb string) []Subject {
	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			A.ref,
			B.ref,
			B.name,
			ifnull(C.ref, 0),
			ifnull(C.name, ''),
			A.code,
			A.name
		from subject as A
		left join subject_group as B on A.group_ref = B.ref
		left join grade as C on A.grade_ref = C.ref
		order by B.name, A.code;
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var subjects []Subject
	var subject Subject

	for rows.Next() {
		err := rows.Scan(
			&subject.Ref,
			&subject.GroupRef,
			&subject.GroupName,
			&subject.GradeRef,
			&subject.GradeName,
			&subject.Code,
			&subject.Name,
		)

		if err != nil {
			panic(err)
		}

		subjects = append(subjects, subject)
	}

	return subjects
}

func GetByGroup(urldb string, group_ref string) []Subject {
	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			A.ref,
			B.ref,
			B.name,
			A.code,
			A.name
		from subject as A
		left join subject_group as B on A.group_ref = B.ref
		where group_ref=` + group_ref + `
		order by B.name, A.code;
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var subjects []Subject
	var subject Subject

	for rows.Next() {
		err := rows.Scan(
			&subject.Ref,
			&subject.GroupRef,
			&subject.GroupName,
			&subject.Code,
			&subject.Name,
		)

		if err != nil {
			panic(err)
		}

		subjects = append(subjects, subject)
	}

	return subjects
}

func GetByGradeGroup(urldb string, grade_ref string, group_ref string) []Subject {
	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			A.ref,
			B.ref,
			B.name,
			A.code,
			A.name
		from subject as A
		left join subject_group as B on A.group_ref = B.ref
		where A.group_ref=` + group_ref + ` and A.grade_ref=` + grade_ref + `
		order by B.name, A.code;
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var subjects []Subject
	var subject Subject

	for rows.Next() {
		err := rows.Scan(
			&subject.Ref,
			&subject.GroupRef,
			&subject.GroupName,
			&subject.Code,
			&subject.Name,
		)

		if err != nil {
			panic(err)
		}

		subjects = append(subjects, subject)
	}

	return subjects
}

func insert(urldb string, subject Subject) {
	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		insert into subject(
			group_ref,
			grade_ref,
			code,
			name
		) values (?,?,?,?);
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(subject.GroupRef, subject.GradeRef, subject.Code, subject.Name)

	if err != nil {
		panic(err)
	}
}

func update(urldb string, subject Subject) {
	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		update subject set
			group_ref=?,
			grade_ref=?, 
			code=?,
			name=?
		where ref=?;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(subject.GroupRef, subject.GradeRef, subject.Code, subject.Name, subject.Ref)

	if err != nil {
		panic(err)
	}
}

func Save(urldb string, subject Subject) []Subject {

	if subject.Ref == 0 {
		insert(urldb, subject)

	} else {
		update(urldb, subject)
	}

	return Get(urldb)
}

func Del(urldb string, subject Subject) []Subject {
	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		delete from subject 
		where ref=?;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(subject.Ref)

	if err != nil {
		panic(err)
	}

	return Get(urldb)
}
