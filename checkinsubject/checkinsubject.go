package checkinsubject

import (
	"database/sql"
)

type CheckinSubjectStudent struct {
	StudentRef      int
	StudentCode     string
	studentNo       int
	StudentFullName string
	StatusNo        int
	StatusName      string
}

type CheckinSubject struct {
	Created          string
	Period           int
	RoomRef          int
	RoomName         string
	GradeRef         int
	GradeName        string
	TeacherRef       int
	TeacherFullName  string
	SubjectGroupRef  int
	SubjectGroupName string
	SubjectRef       int
	SubjectName      string
	Students         []CheckinSubjectStudent
}

type CheckinStudent struct {
	Ref      int
	No       int
	Code     string
	FullName string
	P1       string
	P2       string
	P3       string
	P4       string
	P5       string
	P6       string
	P7       string
	P8       string
	P9       string
	P10      string
}

type Result struct {
	Success bool
}

func Get(urldb string, created string, room_ref string) []CheckinStudent {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			A.student_ref as ref,
			B.no as no,
			A.student_code as code, 
			concat(B.first_name, ' ', B.last_name) as fullname,
			ifnull((select status from checkinsubject_student where created='` + created + `' and period=1 and student_ref=ref), '') as period1, 
			ifnull((select status from checkinsubject_student where created='` + created + `' and period=2 and student_ref=ref), '') as period2, 
			ifnull((select status from checkinsubject_student where created='` + created + `' and period=3 and student_ref=ref), '') as period3, 
			ifnull((select status from checkinsubject_student where created='` + created + `' and period=4 and student_ref=ref), '') as period4, 
			ifnull((select status from checkinsubject_student where created='` + created + `' and period=5 and student_ref=ref), '') as period5, 
			ifnull((select status from checkinsubject_student where created='` + created + `' and period=6 and student_ref=ref), '') as period6, 
			ifnull((select status from checkinsubject_student where created='` + created + `' and period=7 and student_ref=ref), '') as period7, 
			ifnull((select status from checkinsubject_student where created='` + created + `' and period=8 and student_ref=ref), '') as period8, 
			ifnull((select status from checkinsubject_student where created='` + created + `' and period=9 and student_ref=ref), '') as period9, 
			ifnull((select status from checkinsubject_student where created='` + created + `' and period=10 and student_ref=ref), '') as period10 
		from checkinsubject_student as A 
		left join students as B on A.student_ref=B.ref
		where A.created='` + created + `' and A.room_ref=` + room_ref + ` 
		group by cast(no as int);
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	var checkinstudents []CheckinStudent
	var checkinstudent CheckinStudent

	for rows.Next() {
		err = rows.Scan(
			&checkinstudent.Ref,
			&checkinstudent.No,
			&checkinstudent.Code,
			&checkinstudent.FullName,
			&checkinstudent.P1,
			&checkinstudent.P2,
			&checkinstudent.P3,
			&checkinstudent.P4,
			&checkinstudent.P5,
			&checkinstudent.P6,
			&checkinstudent.P7,
			&checkinstudent.P8,
			&checkinstudent.P9,
			&checkinstudent.P10,
		)

		if err != nil {
			panic(err)
		}

		checkinstudent.P1 = convStatus(checkinstudent.P1)
		checkinstudent.P2 = convStatus(checkinstudent.P2)
		checkinstudent.P3 = convStatus(checkinstudent.P3)
		checkinstudent.P4 = convStatus(checkinstudent.P4)
		checkinstudent.P5 = convStatus(checkinstudent.P5)
		checkinstudent.P6 = convStatus(checkinstudent.P6)
		checkinstudent.P7 = convStatus(checkinstudent.P7)
		checkinstudent.P8 = convStatus(checkinstudent.P8)
		checkinstudent.P9 = convStatus(checkinstudent.P9)
		checkinstudent.P10 = convStatus(checkinstudent.P10)

		checkinstudents = append(checkinstudents, checkinstudent)
	}

	return checkinstudents
}

func GetByKey(urldb string, created string, period string, room_ref string) CheckinSubject {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select
			A.created,
			A.period,
			C.ref as grade_ref,
			A.room_ref,
			A.teacher_ref as teacher_ref,
			concat(B.firstname, ' ', B.lastname) as teacher_fullname,
			D.group_ref as subjectgroup_ref,
			A.subject_ref as subject_ref
		from checkinsubject as A
		left join teacher as B on A.teacher_ref = B.ref
		left join classroom as C on A.room_ref = C.ref
		left join subject as D on A.subject_ref = D.ref
		where created='` + created + `' and period=` + period + ` and A.room_ref=` + room_ref + `
		limit 1;
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	var checkinsubject CheckinSubject

	if rows.Next() {
		err = rows.Scan(
			&checkinsubject.Created,
			&checkinsubject.Period,
			&checkinsubject.GradeRef,
			&checkinsubject.RoomRef,
			&checkinsubject.TeacherRef,
			&checkinsubject.TeacherFullName,
			&checkinsubject.SubjectGroupRef,
			&checkinsubject.SubjectRef,
		)

		if err != nil {
			panic(err)
		}

		sql = `
			select
				student_ref,
				B.No as student_no,
				student_code,
				concat(B.first_name, ' ', B.last_name) as student_fullname,
				status
			from checkinsubject_student as A
			left join students as B on A.student_ref = B.ref
			where created='` + created + `' and period=` + period + ` and room_ref=` + room_ref + `;
		`
		rows, err = db.Query(sql)

		if err != nil {
			panic(err)
		}

		var checkinstudent CheckinSubjectStudent

		for rows.Next() {
			err = rows.Scan(
				&checkinstudent.StudentRef,
				&checkinstudent.StudentCode,
				&checkinstudent.studentNo,
				&checkinstudent.StudentFullName,
				&checkinstudent.StatusNo,
			)

			if err != nil {
				panic(err)
			}

			switch checkinstudent.StatusNo {
			case 1:
				checkinstudent.StatusName = "มา"
			case 2:
				checkinstudent.StatusName = "ขาด"
			case 3:
				checkinstudent.StatusName = "ลากิจ"
			case 4:
				checkinstudent.StatusName = "ลาป่วย"
			}

			checkinsubject.Students = append(checkinsubject.Students, checkinstudent)
		}
	}

	return checkinsubject
}

func convStatus(status string) string {
	var conv string

	switch status {

	case "1":
		conv = "มา"
	case "2":
		conv = "ขาด"
	case "3":
		conv = "ลากิจ"	
	case "4":
		conv = "ป่วย"	
	}

	return conv
}

func CheckDuplicate(urldb string, created string, period string, room_ref string) Result {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select created 
		from checkinsubject 
		where created=? and period=? and room_ref=?;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	row, err := sttn.Query(created, period, room_ref)

	if err != nil {
		panic(err)
	}

	defer row.Close()

	var result Result

	if row.Next() {
		result.Success = false
	} else {
		result.Success = true
	}

	return result
}

func Save(urldb string, checkinsubject CheckinSubject) Result {
	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `select ref from user where name='` + checkinsubject.TeacherFullName + `';`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var result Result

	if rows.Next() {

		err = rows.Scan(&checkinsubject.TeacherRef)

		if err != nil {
			panic(err)
		}

		sql := `
			insert into checkinsubject(created, period, room_ref, teacher_ref, subject_ref) 
			values(?,?,?,?,?);
		`

		sttn, err := db.Prepare(sql)

		if err != nil {
			panic(err)
		}

		defer sttn.Close()

		_, err = sttn.Exec(
			checkinsubject.Created,
			checkinsubject.Period,
			checkinsubject.RoomRef,
			checkinsubject.TeacherRef,
			checkinsubject.SubjectRef,
		)

		if err != nil {
			panic(err)
		}

		for _, s := range checkinsubject.Students {
			sql = `
			insert into checkinsubject_student(created, period, room_ref, student_ref, student_code, status) 
			values(?,?,?,?,?,?);`

			sttn, err = db.Prepare(sql)

			if err != nil {
				panic(err)
			}

			_, err = sttn.Exec(
				checkinsubject.Created,
				checkinsubject.Period,
				checkinsubject.RoomRef,
				s.StudentRef,
				s.StudentCode,
				s.StatusNo,
			)

			if err != nil {
				panic(err)
			}
		}

		result.Success = true

	} else {
		result.Success = false
	}

	return result
}
