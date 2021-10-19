package student

//package main

import (
	//"encoding/json"
	"fmt"
	//"net/http"
	"database/sql"
	"io/ioutil"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type StudentInfo struct {
	Ref              int
	Code             string
	IdCard           string
	Rfid             string
	Title            int
	FirstName        string
	LastName         string
	NickName         string
	Gender           int
	BirthDay         string
	GradeNo          int
	GradeName        string
	RoomNo           int
	RoomName         string
	CheckinProfile   string
	No               int
	Phone            string
	LineId           string
	Facebook         string
	ParentRef1       int
	ParentTitle1     int
	ParentFirstName1 string
	ParentLastName1  string
	ParentPhone1     string
	ParentRef2       int
	ParentTitle2     int
	ParentFirstName2 string
	ParentLastName2  string
	ParentPhone2     string
	ParentRef3       int
	ParentTitle3     int
	ParentFirstName3 string
	ParentLastName3  string
	ParentPhone3     string
	FaceImg          string
}

type StudentView struct {
	Code     string
	Idcard   string
	No       int
	FullName string
	Phone    string
}

type StudentImport struct {
	Room      int
	Code      string
	No        int
	FirstName string
	LastName  string
}

type ImportLot struct {
	Created  string
	Grade    int
	Students []StudentImport
}

func Import(urldb string, importlog ImportLot) {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		insert into importstudent (
			created,
			room_ref,
			code,
			no,
			firstname,
			lastname
		) values (?,?,?,?,?,?);
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

}

func GetFaceImg(code string) []byte {

	path := "imgs/library/"
	filename := code + ".b64"

	content, err := ioutil.ReadFile(path + filename)

	if err != nil {
		return []byte("")
	}

	return content
}

//func GetFaceImgByCode(urldb string, code string) []byte {
//
//	db, err := sql.Open("mysql", urldb)
//
//	if err != nil {
//		panic(err)
//	}
//
//	defer db.Close()
//
//	sql := `
//		select
//			year(birthday) as year,
//			idcard
//		from students
//		where code=?;
//	`
//
//	sttn, err := db.Prepare(sql)
//
//	if err != nil {
//		panic(err)
//	}
//
//	defer sttn.Close()
//
//	rows, err := sttn.Query(code)
//
//	if err != nil {
//		panic(err)
//	}
//
//	defer rows.Close()
//
//	var year string
//	var idcard string
//	var content []byte
//
//	if rows.Next() {
//
//		err := rows.Scan(&year, &idcard)
//
//		if err != nil {
//			panic(err)
//		}
//
//		content = GetFaceImg(year, idcard)
//	}
//
//	return content
//}

func Get(urldb string, param_grade string, param_room string, param_txtsearch string) []StudentInfo {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var studentInfos []StudentInfo

	cond := ""

	if param_grade != "0" {
		cond = "where grade=" + param_grade

		if param_room != "0" {
			cond = cond + " and room=" + param_room
		}

		if param_txtsearch != "" {
			cond = cond + " and (first_name like '%" + param_txtsearch + "%' or last_name like '%" + param_txtsearch + "%')"
		}

	} else if param_txtsearch != "" {
		cond = "where code like '" + param_txtsearch + "%' or first_name like '%" + param_txtsearch + "%' or last_name like '%" + param_txtsearch + "%'"
	}

	sql := `
		select 
			ifnull(students.ref, 0), 
			ifnull(code, ''), 
			ifnull(idcard, ''), 
			ifnull(rfid, ''),
			ifnull(title, 0),
			ifnull(first_name, ''), 
			ifnull(last_name, ''), 
			ifnull(nickname, ''), 
			ifnull(gender, 0), 
			ifnull(birthday, ''), 
			ifnull(grade, 0) as grade_ref,
			ifnull(grade.name, '') as grade_name,
			ifnull(room, 0) as room_ref,
			ifnull(classroom.name, '') as room_name,
			ifnull(checkinprofile, 0),
			if(no = '', 0, no),
			ifnull(phone, ''), 
			ifnull(line, ''), 
			ifnull(facebook, '') 
		from students 
		left join classroom on students.room = classroom.ref 
		left join grade on students.grade = grade.ref
		` + cond + ` order by grade_ref, room_ref, convert(no, integer) ;`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var studentInfo StudentInfo

	for rows.Next() {
		err := rows.Scan(
			&studentInfo.Ref,
			&studentInfo.Code,
			&studentInfo.IdCard,
			&studentInfo.Rfid,
			&studentInfo.Title,
			&studentInfo.FirstName,
			&studentInfo.LastName,
			&studentInfo.NickName,
			&studentInfo.Gender,
			&studentInfo.BirthDay,
			&studentInfo.GradeNo,
			&studentInfo.GradeName,
			&studentInfo.RoomNo,
			&studentInfo.RoomName,
			&studentInfo.CheckinProfile,
			&studentInfo.No,
			&studentInfo.Phone,
			&studentInfo.LineId,
			&studentInfo.Facebook,
		)

		if err != nil {
		}

		studentInfos = append(studentInfos, studentInfo)
	}

	return studentInfos
}

func GetByClassroom(urldb string, grade int, room int) []StudentView {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			ifnull(students.ref, 0), 
			ifnull(code, ''), 
			ifnull(idcard, ''), 
			ifnull(no, ''), 
			concat(ifnull(title, 0), ' ', ifnull(first_name, ''), ' ',  ifnull(last_name, '')), 
			ifnull(phone, '')
		from students 
		where room=?;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	var results []StudentView
	var result StudentView

	rows, err := sttn.Query(room)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
	_:
		err := rows.Scan(
			&result.Code,
			&result.Idcard,
			&result.No,
			&result.FullName,
			&result.Phone,
		)

		if err != nil {
			panic(err)
		}

		results = append(results, result)
	}

	return results
}

func GetByRef(urldb string, ref string) StudentInfo {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			ifnull(students.ref, 0), 
			ifnull(code, ''), 
			ifnull(idcard, ''), 
			ifnull(rfid, ''),
			ifnull(title, 0),
			ifnull(first_name, ''), 
			ifnull(last_name, ''), 
			ifnull(nickname, ''), 
			ifnull(gender, 0), 
			ifnull(birthday, ''), 
			ifnull(grade, 0) as grade_ref,
			ifnull(grade.name, '') as grade_name,
			ifnull(room, 0) as room_ref,
			ifnull(classroom.name, '') as room_name,
			ifnull(checkinprofile, 0),
			if(no = '', 0, no),
			ifnull(phone, ''), 
			ifnull(line, ''), 
			ifnull(facebook, '') 
		from students 
		left join classroom on students.room = classroom.ref 
		left join grade on students.grade = grade.ref
		where students.ref=?
		limit 1; 
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	var studentInfo StudentInfo

	row, err := sttn.Query(ref)

	if err != nil {
		panic(err)
	}

	defer row.Close()

	if row.Next() {

		err := row.Scan(
			&studentInfo.Ref,
			&studentInfo.Code,
			&studentInfo.IdCard,
			&studentInfo.Rfid,
			&studentInfo.Title,
			&studentInfo.FirstName,
			&studentInfo.LastName,
			&studentInfo.NickName,
			&studentInfo.Gender,
			&studentInfo.BirthDay,
			&studentInfo.GradeNo,
			&studentInfo.GradeName,
			&studentInfo.RoomNo,
			&studentInfo.RoomName,
			&studentInfo.CheckinProfile,
			&studentInfo.No,
			&studentInfo.Phone,
			&studentInfo.LineId,
			&studentInfo.Facebook,
		)

		if err != nil {
			panic(err)
		}
	}

	return studentInfo
}

func GetByCode(urldb string, code string) StudentInfo {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			ifnull(students.ref, 0), 
			ifnull(code, ''), 
			ifnull(idcard, ''), 
			ifnull(rfid, ''),
			ifnull(title, 0),
			ifnull(first_name, ''), 
			ifnull(last_name, ''), 
			ifnull(nickname, ''), 
			ifnull(gender, 0), 
			ifnull(birthday, ''), 
			ifnull(grade, 0) as grade_ref,
			ifnull(grade.name, '') as grade_name,
			ifnull(room, 0) as room_ref,
			ifnull(classroom.name, '') as room_name,
			ifnull(checkinprofile, 0),
			if(no = '', 0, no),
			ifnull(phone, ''), 
			ifnull(line, ''), 
			ifnull(facebook, '') 
		from students 
		left join classroom on students.room = classroom.ref 
		left join grade on students.grade = grade.ref
		where code='` + code + `'
		limit 1;`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var studentInfo StudentInfo

	if rows.Next() {
		err := rows.Scan(
			&studentInfo.Ref,
			&studentInfo.Code,
			&studentInfo.IdCard,
			&studentInfo.Rfid,
			&studentInfo.Title,
			&studentInfo.FirstName,
			&studentInfo.LastName,
			&studentInfo.NickName,
			&studentInfo.Gender,
			&studentInfo.BirthDay,
			&studentInfo.GradeNo,
			&studentInfo.GradeName,
			&studentInfo.RoomNo,
			&studentInfo.RoomName,
			&studentInfo.CheckinProfile,
			&studentInfo.No,
			&studentInfo.Phone,
			&studentInfo.LineId,
			&studentInfo.Facebook,
		)

		if err != nil {
			panic(err)
		}
	}

	return studentInfo
}

func GetByParent(urldb string, strparentref string) []StudentInfo {

	parentref, err := strconv.Atoi(strparentref)

	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			students.ref,
			if(isnull(code), '', code),
			if(isnull(idcard), '',idcard),
			if(isnull(rfid), '',rfid),
			if(isnull(title), 0, title),
			if(isnull(first_name), '', first_name),
			if(isnull(last_name), '', last_name),
			if(isnull(nickname), '', nickname),
			if(isnull(gender), 0, gender),
			if(isnull(birthday), '', birthday),
			if(isnull(grade), 0, grade),
			if(isnull(grade.name), '', grade.name) as grade_name,
			if(isnull(room), 0, room),
			if(isnull(classroom.name), '', classroom.name) as room_name,
			checkinprofile,
			if(no = '', 0, no),
			if(isnull(phone), '', phone),
			if(isnull(line), '', line),
			if(isnull(facebook), '', facebook)
		from student_parents 
		left join students on student_parents.student_ref = students.ref 
		left join grade on students.grade = grade.ref
		left join classroom on students.room = classroom.ref
		where student_parents.parent_ref=?
		;`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	rows, err := sttn.Query(parentref)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	var studentInfos []StudentInfo
	var studentInfo StudentInfo

	for rows.Next() {
		err := rows.Scan(
			&studentInfo.Ref,
			&studentInfo.Code,
			&studentInfo.IdCard,
			&studentInfo.Rfid,
			&studentInfo.Title,
			&studentInfo.FirstName,
			&studentInfo.LastName,
			&studentInfo.NickName,
			&studentInfo.Gender,
			&studentInfo.BirthDay,
			&studentInfo.GradeNo,
			&studentInfo.GradeName,
			&studentInfo.RoomNo,
			&studentInfo.RoomName,
			&studentInfo.CheckinProfile,
			&studentInfo.No,
			&studentInfo.Phone,
			&studentInfo.LineId,
			&studentInfo.Facebook,
		)

		if err != nil {
			panic(err)
		}

		studentInfos = append(studentInfos, studentInfo)
	}

	return studentInfos
}

func SavePicImg(student StudentInfo) {

	path := "imgs/library/"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0777)
	}

	filename := student.Code + ".b64"

	err := ioutil.WriteFile(path+filename, []byte(student.FaceImg), 0777)

	if err != nil {
		panic(err)
	}
}

func Save(urldb string, student StudentInfo) []StudentInfo {

	if student.Ref == 0 {
		insert(urldb, student)

	} else {
		update(urldb, student)
	}

	SavePicImg(student)

	return Get(urldb, "0", "0", "")
}

func insert(urldb string, student StudentInfo) {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var sql string

	sql = `
		insert into students (
			code,			
			idcard,
			title,
			first_name,
			last_name,
			nickname,
			gender,
			grade,
			room,
			checkinprofile,
			no,
			phone,
			line,
			facebook,
			rfid
		) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	_, err = sttn.Exec(
		student.Code,
		student.IdCard,
		student.Title,
		student.FirstName,
		student.LastName,
		student.NickName,
		student.Gender,
		student.GradeNo,
		student.RoomNo,
		student.CheckinProfile,
		student.No,
		student.Phone,
		student.LineId,
		student.Facebook,
		student.Rfid,
	)

	if err != nil {
		panic(err)
	}
}

func update(urldb string, student StudentInfo) {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var sql string

	sql = `
		update students set
			code=?,			
			idcard=?,
			rfid=?,
			title=?,
			first_name=?,
			last_name=?,
			nickname=?,
			gender=?,
			grade=?,
			room=?,
			checkinprofile=?,
			no=?,
			phone=?,
			line=?,
			facebook=?
		where ref=?			
	`
	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	_, err = sttn.Exec(
		student.Code,
		student.IdCard,
		student.Rfid,
		student.Title,
		student.FirstName,
		student.LastName,
		student.NickName,
		student.Gender,
		student.GradeNo,
		student.RoomNo,
		student.CheckinProfile,
		student.No,
		student.Phone,
		student.LineId,
		student.Facebook,
		student.Ref,
	)

	if err != nil {
		panic(err)
	}

	// update parent info

	sql = "select count(*) from student_parents where student_ref=?"

	sttn, err = db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(student.Ref)

	if err != nil {
		panic(err)
	}

	var count int
	var parentref1 int
	var parentref2 int
	var parentref3 int

	// check existing parents
	rows.Next()
	rows.Scan(&count)

	if count == 0 {
		parentref1 = insertParent(db, student.ParentTitle1, student.ParentFirstName1, student.ParentLastName1, student.ParentPhone1)
		parentref2 = insertParent(db, student.ParentTitle2, student.ParentFirstName2, student.ParentLastName2, student.ParentPhone2)
		parentref3 = insertParent(db, student.ParentTitle3, student.ParentFirstName3, student.ParentLastName3, student.ParentPhone3)

		mapParent(db, student.Ref, parentref1)
		mapParent(db, student.Ref, parentref2)
		mapParent(db, student.Ref, parentref3)

	} else {
		updateParent(db, student.ParentRef1, student.ParentTitle1, student.ParentFirstName1, student.ParentLastName1, student.ParentPhone1)
		updateParent(db, student.ParentRef2, student.ParentTitle2, student.ParentFirstName2, student.ParentLastName2, student.ParentPhone2)
		updateParent(db, student.ParentRef3, student.ParentTitle3, student.ParentFirstName3, student.ParentLastName3, student.ParentPhone3)

	}

}

func insertParent(db *sql.DB, title int, firstname string, lastname string, phone string) int {

	sql := `insert into parents (title, first_name, last_name, phone) values(?,?,?,?)`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(title, firstname, lastname, phone)

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

	var ref int

	rows.Next()
	rows.Scan(&ref)

	return ref
}

func updateParent(db *sql.DB, ref int, title int, firstname string, lastname string, phone string) {
	sql := `
		update parents set 
			title=?, 
			first_name=?, 
			last_name=?, 
			phone=?
		where ref=?`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(title, firstname, lastname, phone, ref)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()
}

func mapParent(db *sql.DB, studentref int, parentref int) {

	sql := `insert into student_parents (student_ref, parent_ref) values(?,?)`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(studentref, parentref)

	if err != nil {
		panic(err)
	}
}

func GetByUID(urldb string, uid string) StudentInfo {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			students.ref, 
			code, 
			idcard, 
			rfid,
			title,
			first_name, 
			last_name, 
			nickname, 
			gender, 
			birthday, 
			grade as grade_ref,
			grade.name as grade_name,
			room as room_ref,
			classroom.name as room_name,
			checkinprofile,
			if(no = '', 0,  no),
			phone, 
			line, 
			facebook 
		from students 
		left join classroom on students.room = classroom.ref 
		left join grade on students.grade = grade.ref
		where rfid='` + uid + `' limit 1`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var studentInfo StudentInfo

	if rows.Next() {
		err := rows.Scan(
			&studentInfo.Ref,
			&studentInfo.Code,
			&studentInfo.IdCard,
			&studentInfo.Rfid,
			&studentInfo.Title,
			&studentInfo.FirstName,
			&studentInfo.LastName,
			&studentInfo.NickName,
			&studentInfo.Gender,
			&studentInfo.BirthDay,
			&studentInfo.GradeNo,
			&studentInfo.GradeName,
			&studentInfo.RoomNo,
			&studentInfo.RoomName,
			&studentInfo.CheckinProfile,
			&studentInfo.No,
			&studentInfo.Phone,
			&studentInfo.LineId,
			&studentInfo.Facebook,
		)

		if err != nil {
		}

	}

	return studentInfo
}

func Remove(urldb string, ref string) []StudentInfo {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var sql string

	sql = `
		delete from students where ref=` + ref + `; 
	`

	_, err = db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return Get(urldb, "0", "0", "")
}

func main() {
	const urldb = "admin:35.103232@tcp(localhost:3306)/school"

	fmt.Print(GetByParent(urldb, "41"))
}
