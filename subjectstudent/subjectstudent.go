package subjectstudent

import (
	"database/sql"
)

type SubjectStudent struct {
	YearEdu          int
	GradeRef         int
	GradeName        string
	RoomRef          int
	RoomName         string
	SubjectGroupRef  int
	SubjectGroupName string
	SubjectRef       int
	SubjectName      string
	StudentRef       int
	StudentCode      string
	StudentName      string
	CanExam          bool
	Point            int
	Gpa              int
}

func Get(
	urldb string,
	yearedu string,
	grade_ref string,
	subject_ref string) []SubjectStudent {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select
			yearedu,
			C.ref as grade_ref,
			C.name as grade_name,
			room_ref,
			B.name as room_name,
			F.ref as subjectgroup_ref,
			F.name as subjectgroup_name,
			subject_ref,
			D.name as subject_name,
			student_ref,
			E.code,
			concat(E.first_name, ' ', E.last_name)  as student_name
		from subjectresult_student as A
		left join classroom as B on A.room_ref=B.ref
		left join grade as C on B.grade_ref=C.ref
		left join subject as D on A.subject_ref=D.ref
		left join students as E on A.student_ref=E.ref
		left join subject_group as F on D.group_ref=F.ref
		where yearedu=` + yearedu + ` and semester=1 and C.ref=` + grade_ref + ` and A.subject_ref=` + subject_ref + ` 
		order by C.name;
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var results []SubjectStudent
	var result SubjectStudent

	for rows.Next() {
		err = rows.Scan(
			&result.YearEdu,
			&result.GradeRef,
			&result.GradeName,
			&result.RoomRef,
			&result.RoomName,
			&result.SubjectGroupRef,
			&result.SubjectGroupName,
			&result.SubjectRef,
			&result.SubjectName,
			&result.StudentRef,
			&result.StudentCode,
			&result.StudentName,
		)

		if err != nil {
			panic(err)
		}

		results = append(results, result)
	}

	return results
}

func GetByRoom(
	urldb string,
	yearedu string,
	room_ref string,
	subject_ref string) []SubjectStudent {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select
			yearedu,
			C.ref as grade_ref,
			C.name as grade_name,
			room_ref,
			B.name as room_name,
			F.ref as subjectgroup_ref,
			F.name as subjectgroup_name,
			subject_ref,
			D.name as subject_name,
			student_ref,
			E.code,
			concat(E.first_name, ' ', E.last_name)  as student_name
		from subjectresult_student as A
		left join classroom as B on A.room_ref=B.ref
		left join grade as C on B.grade_ref=C.ref
		left join subject as D on A.subject_ref=D.ref
		left join students as E on A.student_ref=E.ref
		left join subject_group as F on D.group_ref=F.ref
		where yearedu=` + yearedu + ` and semester=1 and A.room_ref=` + room_ref + ` and A.subject_ref=` + subject_ref + ` 
		order by C.name;
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var results []SubjectStudent
	var result SubjectStudent

	for rows.Next() {
		err = rows.Scan(
			&result.YearEdu,
			&result.GradeRef,
			&result.GradeName,
			&result.RoomRef,
			&result.RoomName,
			&result.SubjectGroupRef,
			&result.SubjectGroupName,
			&result.SubjectRef,
			&result.SubjectName,
			&result.StudentRef,
			&result.StudentCode,
			&result.StudentName,
		)

		if err != nil {
			panic(err)
		}

		results = append(results, result)
	}

	return results
}

func InsertStudent(
	urldb string,
	yearedu string,
	semester int,
	grade_ref string,
	room_ref string,
	subject_ref string,
	student_ref string) {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		insert into subjectresult_student(
			yearedu,
			semester,
			grade_ref,
			room_ref,
			student_ref,
			subject_ref
		) values(?,?,?,?,?,?);
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(yearedu, semester, grade_ref, room_ref, student_ref, subject_ref)

	if err != nil {
		panic(err)
	}
}

func DeleteStudent(
	urldb string,
	yearedu string,
	subject_ref string,
	student_ref string) {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		delete from subjectresult_student
		where yearedu=` + yearedu + ` and subject_ref=` + subject_ref + ` and student_ref=` + student_ref + `;
	`

	_, err = db.Exec(sql)

	if err != nil {
		panic(err)
	}
}

func InsertAllGrade(
	urldb string,
	yearedu string,
	grade_ref string,
	subject_ref string) []SubjectStudent {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select
			ref,
			code,
			room
		from students
		where grade=` + grade_ref + ` 
		order by grade, room; 
	`
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var student_ref string
	var student_code string
	var room_ref string

	for rows.Next() {
		err = rows.Scan(&student_ref, &student_code, &room_ref)

		if err != nil {
			panic(err)
		}

		InsertStudent(
			urldb,
			yearedu,
			1,
			grade_ref,
			room_ref,
			subject_ref,
			student_ref,
		)

		InsertStudent(
			urldb,
			yearedu,
			2,
			grade_ref,
			room_ref,
			subject_ref,
			student_ref,
		)
	}

	return Get(urldb, yearedu, grade_ref, subject_ref)
}

func DeleteAllGrade(
	urldb string,
	yearedu string,
	grade_ref string,
	subject_ref string) []SubjectStudent {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		delete from subjectresult_student
		where yearedu=` + yearedu + ` and subject_ref=` + subject_ref + ` and grade_ref=` + grade_ref + `; 
	`

	_, err = db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return Get(urldb, yearedu, grade_ref, subject_ref)
}

func InsertByRoom(
	urldb string,
	yearedu string,
	room_ref string,
	subject_ref string) []SubjectStudent {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select
			ref,
			code,
			grade
		from students
		where room=` + room_ref + ` 
		order by grade, room; 
	`
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var student_ref string
	var student_code string
	var grade_ref string

	for rows.Next() {
		err = rows.Scan(&student_ref, &student_code, &grade_ref)

		if err != nil {
			panic(err)
		}

		InsertStudent(
			urldb,
			yearedu,
			1,
			grade_ref,
			room_ref,
			subject_ref,
			student_ref,
		)

		InsertStudent(
			urldb,
			yearedu,
			2,
			grade_ref,
			room_ref,
			subject_ref,
			student_ref,
		)
	}

	return GetByRoom(urldb, yearedu, room_ref, subject_ref)
}

func DeleteByRoom(
	urldb string,
	yearedu string,
	room_ref string,
	subject_ref string) []SubjectStudent {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		delete from subjectresult_student
		where yearedu=` + yearedu + ` and subject_ref=` + subject_ref + ` and room_ref=` + room_ref + `; 
	`

	_, err = db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return GetByRoom(urldb, yearedu, room_ref, subject_ref)
}

func InsertByCode(
	urldb string,
	yearedu string,
	subject_ref string,
	student_code string) []SubjectStudent {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select
			ref,
			grade,
			room
		from students
		where code=` + student_code + `;
	`
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var student_ref string
	var grade_ref string
	var room_ref string

	if rows.Next() {
		err = rows.Scan(&student_ref, &grade_ref, &room_ref)

		if err != nil {
			panic(err)
		}

		InsertStudent(
			urldb,
			yearedu,
			1,
			grade_ref,
			room_ref,
			subject_ref,
			student_ref,
		)

		InsertStudent(
			urldb,
			yearedu,
			2,
			grade_ref,
			room_ref,
			subject_ref,
			student_ref,
		)
	}

	return Get(urldb, yearedu, grade_ref, subject_ref)
}

func DeleteByCode(
	urldb string,
	yearedu string,
	student_code string,
	subject_ref string) []SubjectStudent {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `select ref, grade, room from students where code='` + student_code + `'`

	row, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	var student_ref string
	var grade_ref string
	var room_ref string

	if row.Next() {

		err = row.Scan(&student_ref, &grade_ref, &room_ref)

		if err != nil {
			panic(err)
		}

		sql = `
			delete from subjectresult_student
			where yearedu=` + yearedu + ` and subject_ref=` + subject_ref + ` and room_ref=` + room_ref + `; 
		`

		_, err = db.Exec(sql)

		if err != nil {
			panic(err)
		}
	}

	return Get(urldb, yearedu, grade_ref, subject_ref)
}
