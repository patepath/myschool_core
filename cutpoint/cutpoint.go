package cutpoint
//package main

import (
	//"fmt"
	"strconv"
	"schoolcore/student"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type StudentPoint struct {
	Ref int
	CreateDate string
	BehaviorRef int
	BehaviorGroupNo int
	BehaviorGroupName string
	BehaviorTopicNo string
	BehaviorTopicName string
	BehaviorPoint int
	StudentRef int
	StudentCode string
	StudentFullName string
	StudentGradeRef int
	StudentGradeName string
	StudentRoomRef int
	StudentRoomName string
}

type Cutpoint struct {
	Ref int
	CreateDate string
	BehaviorGroup int
	BehaviorGroupName string
	BehaviorRef int
	BehaviorTopicName string
	BehaviorPoint int
	Students []student.StudentInfo
}

type Report1 struct {
	IdCard	string
	No		string
	Code	string
	Title		string
	FullName	string
	Grade		string
	Room		string
	CutPoint	int
}

type ReportStudentBehavior struct {
	No	int
	Description	string
	Point	int
}

func Get(urldb string) []StudentPoint {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			B.ref as ref,
			B.createdate,
			B.behavior_ref as behavior_ref,
			C.groupno as behavior_group_no,
			C.groupname as behavior_group_name,
			C.topicno as behavior_topic_no,
			C.topicname as behavior_topic_name,
			C.point as behavior_point,
			A.student_ref as student_ref,
			D.code as student_code,
			concat(D.first_name, ' ', D.last_name) as fullname,
			D.grade as grade_ref,
			E.name as grade_name,
			D.room as room_ref,
			F.name as room_name
		from cutpoint_students as A
		left join cutpoint as B on B.ref = A.cutpoint_ref
		left join behavioral_point as C on C.ref = B.behavior_ref
		left join students as D on D.ref = A.student_ref
		left join grade as E on E.ref = D.grade
		left join classroom as F on F.ref = D.room
		order by B.createdate desc;
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var studentpoints []StudentPoint
	var studentpoint StudentPoint

	for rows.Next() {

		err := rows.Scan(
			&studentpoint.Ref,
			&studentpoint.CreateDate,
			&studentpoint.BehaviorRef,
			&studentpoint.BehaviorGroupNo,
			&studentpoint.BehaviorGroupName,
			&studentpoint.BehaviorTopicNo,
			&studentpoint.BehaviorTopicName,
			&studentpoint.BehaviorPoint,
			&studentpoint.StudentRef,
			&studentpoint.StudentCode,
			&studentpoint.StudentFullName,
			&studentpoint.StudentGradeRef,
			&studentpoint.StudentGradeName,
			&studentpoint.StudentRoomRef,
			&studentpoint.StudentRoomName,
		)

		if err != nil {
			panic(err)
		}

		studentpoints = append(studentpoints,studentpoint)
	}

	return studentpoints
}

func Get_bak(urldb string) []Cutpoint {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			cutpoint.ref as ref, 
			if(isnull(createdate), '', createdate), 
			behavioral_point.groupno as behavior_group, 
			behavioral_point.groupname as behavior_groupname, 
			behavioral_point.Ref as behavior_ref,
			behavioral_point.topicname as behavior_topicname,
			behavioral_point.point as behavior_point
		from cutpoint 
		left join behavioral_point on cutpoint.behavior_ref = behavioral_point.ref
		order by createdate desc 
		limit 100`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var cutpoints []Cutpoint
	var cutpoint Cutpoint

	for rows.Next() {
		err := rows.Scan(
			&cutpoint.Ref,
			&cutpoint.CreateDate,
			&cutpoint.BehaviorGroup,
			&cutpoint.BehaviorGroupName,
			&cutpoint.BehaviorRef,
			&cutpoint.BehaviorTopicName,
			&cutpoint.BehaviorPoint,
		)

		if err != nil {
			panic(err)
		}

		cutpoints = append(cutpoints, cutpoint)
	}

	return cutpoints
}

func GetStudentsByRef(urldb string, strcutpoint_ref string) []student.StudentInfo {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			students.ref as student_ref, 
			students.code, 
			students.first_name, 
			students.last_name, 
			grade.ref,
			grade.name, 
			classroom.ref as classroom_ref, 
			classroom.name 
		from cutpoint_students 
		left join students on cutpoint_students.student_ref = students.ref 
		left join grade on students.grade = grade.ref 
		left join classroom on students.room = classroom.ref 
		where cutpoint_students.cutpoint_ref=?`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	var students []student.StudentInfo
	var student student.StudentInfo
	cutpoint_ref, _ := strconv.Atoi(strcutpoint_ref)

	rows, err := sttn.Query(cutpoint_ref)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(
			&student.Ref,
			&student.Code,
			&student.FirstName,
			&student.LastName,
			&student.GradeNo,
			&student.GradeName,
			&student.RoomNo,
			&student.RoomName,
		)

		students = append(students, student)
	}

	return students
}

func GetByStudent(urldb string, strstudent_ref string) []Cutpoint {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			cutpoint_ref, 
			cutpoint.createdate, 
			behavioral_point.groupno, 
			behavioral_point.groupname, 
			behavioral_point.ref, 
			behavioral_point.topicname, 
			behavioral_point.point 
		from cutpoint_students 
		left join cutpoint on cutpoint_students.cutpoint_ref = cutpoint.ref 
		left join behavioral_point on cutpoint.behavior_ref = behavioral_point.ref
		where cutpoint_students.student_ref=` + strstudent_ref

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var cutpoints []Cutpoint
	var cutpoint Cutpoint

	for rows.Next() {
		err := rows.Scan(
			&cutpoint.Ref,
			&cutpoint.CreateDate,
			&cutpoint.BehaviorGroup,
			&cutpoint.BehaviorGroupName,
			&cutpoint.BehaviorRef,
			&cutpoint.BehaviorTopicName,
			&cutpoint.BehaviorPoint,
		)

		if err != nil {
			panic(err)
		}

		cutpoints = append(cutpoints, cutpoint)
	}

	return cutpoints
}

func GetReport1_bak(urldb string) []Report1 {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
	select 
		A.idcard, 
		ifnull(A.no, 0) as no, 
		A.code, 
		ifnull(E.name, '-') as title, 
		concat(A.first_name, ' ',  A.last_name) as fullname,  
		ifnull(C.name, '-') as grade,
		ifnull(D.name, '-') as room, 
		coalesce(sum(B.cutpoint), 0) as cutpoint  
	from students as A 
	left join (select idcard, cutpoint from checkinout) as B 
		on B.idcard=A.code
	left join grade as C
		on C.ref = A.grade
	left join classroom as D
		on D.ref = A.room
	left join titlestudent as E
		on E.ref = A.title
	group by A.code 
	order by C.rank, D.name, no;
	`
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var reports []Report1
	var report Report1


	for rows.Next() {

		err := rows.Scan(
			&report.IdCard,
			&report.No,
			&report.Code,
			&report.Title,
			&report.FullName,
			&report.Grade,
			&report.Room,
			&report.CutPoint,
		)

		if err != nil {
			panic(err)
		}

		reports = append(reports, report)
	}

	return reports
}

func GetCutPoint(urldb string, code string) int {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select count(*) 
		from (
			select A.date_in
			from checkinprofile as A
			left join (select created, idcard from checkinout where idcard=? group by date(created)) as B on date(B.created)=A.date_in
			where B.idcard is null and yearedu=2564 and grp=0 and 
				date_in between (select date_in from checkinprofile order by date_in limit 1) and curdate()
		) as C;
	`
	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(code)

	if err != nil {
		panic(err)
	}

	var point1 int

	if rows.Next() {
		err := rows.Scan(&point1)

		if err != nil {
			panic(err)
		}

		point1 = point1 * 5
	}

	var point2 int

	sql = `
		select 
			ifnull(sum(case
			when time(B.created) <= A.time1 then 0
			when time(B.created) > A.time1 and time(created) <= A.time2 then A.point1
			when time(B.created) > A.time2 and time(created) <= A.time3 then A.point2
			when time(B.created) > A.time3 then A.point2
			end), 0) as point
		from (
				select * 
				from checkinprofile 
				where yearedu=2564 and grp=0 and date_in between (select date_in from checkinprofile order by date_in limit 1) and curdate()) as A
		left join (select * from checkinout where idcard=?  group by date(created)) as B on date(B.created)=A.date_in
		where B.created is not null;
	`
	sttn, err = db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	rows, err = sttn.Query(code)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		err := rows.Scan(&point2)

		if err != nil {
			panic(err)
		}
	}

	return (point1 + point2)  * -1
}

func GetReport1(urldb string, room_ref string) []Report1 {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			A.idcard, 
			ifnull(A.no, 0) as no, 
			A.code, 
			case
				when A.title = 1 then 'ดช.'
				when A.title = 2 then 'ดญ.'
				when A.title = 3 then 'นาย'
				when A.title = 4 then 'นางสาว'
				else " - "
			end,
			concat(A.first_name, ' ',  A.last_name) as fullname,  
			ifnull(C.name, '-') as grade,
			ifnull(D.name, '-') as room 

		from students as A 

		left join (select idcard, cutpoint from checkinout) as B 
			on B.idcard=A.code

		left join grade as C
			on C.ref = A.grade

		left join classroom as D
			on D.ref = A.room

		where A.room=` + room_ref + `

		group by A.code 

		order by C.rank, D.name, no;
	`
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var reports []Report1
	var report Report1


	for rows.Next() {

		err := rows.Scan(
			&report.IdCard,
			&report.No,
			&report.Code,
			&report.Title,
			&report.FullName,
			&report.Grade,
			&report.Room,
		)

		report.CutPoint = GetCutPoint(urldb, report.Code)

		if err != nil {
			panic(err)
		}

		reports = append(reports, report)
	}

	return reports
}

func GetLatePoint(urldb string, yearedu int, code string) int {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			ifnull(sum(case
			when time(B.created) <= A.time1 then 0
			when time(B.created) > A.time1 and time(created) <= A.time2 then A.point1
			when time(B.created) > A.time2 and time(created) <= A.time3 then A.point2
			when time(B.created) > A.time3 then A.point2
			end), 0) as point
		from (
				select * 
				from checkinprofile 
				where yearedu=? and grp=0 and date_in between (select date_in from checkinprofile order by date_in limit 1) and curdate()) as A
		left join (select * from checkinout where idcard=?  group by date(created)) as B on date(B.created)=A.date_in
		where B.created is not null;
	`

	sttn, err := db.Prepare(sql)

	if err!= nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(yearedu, code)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var point int

	if rows.Next() {
		err := rows.Scan(&point)

		if err != nil {
			panic(err)
		}

	}

	return point * -1
}

func GetMissPoint(urldb string, yearedu int, code string) int {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select count(*) 
		from (
			select A.date_in
			from checkinprofile as A
			left join (select created, idcard from checkinout where idcard=? group by date(created)) as B on date(B.created)=A.date_in
			where B.idcard is null and yearedu=? and grp=0 and 
				date_in between (select date_in from checkinprofile order by date_in limit 1) and curdate()
		) as C;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(code, yearedu)

	if err != nil {
		panic(err)
	}

	defer rows.Close()


	var count int

	if rows.Next() {
		err := rows.Scan(&count)

		if err != nil {
			panic(err)
		}
	}

	return count * -5
}

func GetStudentBehaviorPoint(urldb string, yearedu int, code string) []ReportStudentBehavior {

	var reports []ReportStudentBehavior
	var report ReportStudentBehavior

	report.No = 1
	report.Description = "มาสาย"
	report.Point = GetLatePoint(urldb, yearedu, code)
	reports = append(reports, report)

	report.No = 2
	report.Description = "ขาดเรียน"
	report.Point = GetMissPoint(urldb, yearedu, code)
	reports = append(reports, report)

	return reports
}


func Save(urldb string, cutpoint Cutpoint) []StudentPoint {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	if cutpoint.Ref == 0 {
		insert(db, cutpoint)

	} else {
		update(db, cutpoint)
	}

	return Get(urldb)
}

func insert(db *sql.DB, cutpoint Cutpoint) {

	sql := `
		insert into cutpoint (
			createdate,
			behavior_ref
		) values(curdate(), ?)`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		cutpoint.BehaviorRef)

	if err != nil {
		panic(err)
	}

	//check new ref
	sql = `select last_insert_id()`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var cutpoint_ref int

	rows.Next()
	rows.Scan(&cutpoint_ref)

	var students []student.StudentInfo
	students = cutpoint.Students

	for _, student := range students {
		updateStudent(db, cutpoint_ref, student.Ref)
	}
}

func update(db *sql.DB, cutpoint Cutpoint) {

	sql := `
		update cutpoint set 
			behavior_ref=?
		where ref=?`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		cutpoint.BehaviorRef,
		cutpoint.Ref)

	if err != nil {
		panic(err)
	}

	sql = `delete from cutpoint_students where cutpoint_ref=?`

	sttn, err = db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(cutpoint.Ref)

	if err != nil {
		panic(err)
	}

	var students []student.StudentInfo
	students = cutpoint.Students

	for _, student := range students {
		updateStudent(db, cutpoint.Ref, student.Ref)
	}
}

func updateStudent(db *sql.DB, cutpoint_ref int, student_ref int) {
	sql := `
		insert into cutpoint_students(
			cutpoint_ref, 
			student_ref
		) values(?,?)`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		cutpoint_ref,
		student_ref)

	if err != nil {
		panic(err)
	}
}

//func main() {
//
//	const urldb = "admin:35.103232@tcp(localhost:3306)/school"
//
//	Get(urldb)
//}


func main() {

	const urldb = "admin:35.103232@tcp(localhost:3306)/school"

	var	cutpoint = Cutpoint {
		BehaviorGroup: 2,
		BehaviorGroupName: "พฤติกรรมการเรียนการเข้าร่วมกิจกรรมโรงเรียน",
		BehaviorPoint: 10,
		BehaviorRef: 41,
		BehaviorTopicName: "หนีประชุม / ไม่เข้าร่วมกิจกรรมที่โรงเรียนจัด",
		CreateDate: "2021-02-25",
		Ref: 14,
		Students: []student.StudentInfo {
			{
				BirthDay: "1972-04-02",
				Code: "10001",
				Facebook: "patepath",
				FirstName: "Patipat",
				Gender: 0,
				GradeName: "อนุบาล 2",
				GradeNo: 2,
				IdCard: "3100904819071",
				LastName: "Punboonrat",
				LineId: "patipat.punboonrat",
				NickName: "amm",
				ParentFirstName1: "",
				ParentFirstName2: "",
				ParentFirstName3: "",
				ParentLastName1: "",
				ParentLastName2: "",
				ParentLastName3: "",
				ParentPhone1: "",
				ParentPhone2: "",
				ParentPhone3: "",
				ParentRef1: 0,
				ParentRef2: 0,
				ParentRef3: 0,
				ParentTitle1: 0,
				ParentTitle2: 0,
				ParentTitle3: 0,
				Phone: "0953586816",
				Ref: 1,
				RoomName: "1",
				RoomNo: 6,
				Title: 1,
			},
			{
				BirthDay: "2012-11-06",
				Code: "10003",
				Facebook: "",
				FirstName: "บวรนันท์",
				Gender: 0,
				GradeName: "อนุบาล 2",
				GradeNo: 2,
				IdCard: "1103704702223",
				LastName: "พันธุ์บุญรัตน์",
				LineId: "",
				NickName: "กัณฑ์",
				ParentFirstName1: "",
				ParentFirstName2: "",
				ParentFirstName3: "",
				ParentLastName1: "",
				ParentLastName2: "",
				ParentLastName3: "",
				ParentPhone1: "",
				ParentPhone2: "",
				ParentPhone3: "",
				ParentRef1: 0,
				ParentRef2: 0,
				ParentRef3: 0,
				ParentTitle1: 0,
				ParentTitle2: 0,
				ParentTitle3: 0,
				Phone: "0834761321",
				Ref: 3,
				RoomName: "2",
				RoomNo: 7,
				Title: 1,
			},
		},
	}

	Save(urldb, cutpoint)
}
