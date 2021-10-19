package report

import (
	//"encoding/json"

	//"net/http"
	//"os"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type ReportCheckin001 struct {
	Code      string
	FullName  string
	ClassRoom string
	Time_in   string
	Time_out  string
	Status    string
}

type ReportCheckin002 struct {
	ClassName     string
	Room          string
	StudentAmoung int
	Normal        int
	Absent        int
	Late1         int
	Late2         int
}

type Report1 struct {
	IdCard   string
	No       string
	Code     string
	Grade    string
	Title    string
	FullName string
	Room     string
	Point    int
}

type Report1Frm struct {
	No          int
	Description string
	Point       int
}

type Report004 struct {
	No           string
	Code         string
	Grade        string
	Room         string
	FullName     string
	ImgB64       string
	Created      string
	Behavior_ref int
}

type Report004Frm struct {
	Rptdate      string
	Code         string
	Behavior_ref int
	Created      string
	Point        int
}

type Report005 struct {
	Code     string
	Fullname string
	D01      string
	D02      string
	D03      string
	D04      string
	D05      string
	D06      string
	D07      string
	D08      string
	D09      string
	D10      string
	D11      string
	D12      string
	D13      string
	D14      string
	D15      string
	D16      string
	D17      string
	D18      string
	D19      string
	D20      string
	D21      string
	D22      string
	D23      string
	D24      string
	D25      string
	D26      string
	D27      string
	D28      string
	D29      string
	D30      string
	D31      string
	Total    string
}

func GetReport001(urldb string, yearedu int, room_ref string) []Report1 {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			ifnull(idcard, ''), 
			ifnull(no, ''),
			ifnull(code, ''),
			ifnull(C.name, 0) as grade,
			case
				when title=0 then 'N/A'
				when title=1 then 'ด.ช.'
				when title=2 then 'ด.ญ.'
				when title=3 then 'นาย'
				when title=4 then 'นางสาว'
			end,
			concat(first_name, ' ', last_name), 
			ifnull(B.name, 0) as room,
			ifnull((select sum(point) from  point_students where student_code=A.code), 0) as point
		from students as A
		join classroom as B on A.room = B.ref
		join grade as C on B.grade_ref = C.ref
		where room=?
		order by point desc;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(room_ref)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var reports []Report1
	var report Report1

	for rows.Next() {

		err = rows.Scan(
			&report.IdCard,
			&report.No,
			&report.Code,
			&report.Grade,
			&report.Title,
			&report.FullName,
			&report.Room,
			&report.Point,
		)

		if err != nil {
			panic(err)
		}

		reports = append(reports, report)

	}

	return reports
}

func GetReport001Frm(urldb string, yearedu int, code string) []Report1Frm {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select
			C.topicname as topicname,
			sum(A.point) as point
		from point_students as A
		join students as B on A.student_code=B.code
		join behavioral_point as C on A.behavior_ref=C.ref
		where yearedu=? and code=?
		group by A.student_code, behavior_ref;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(yearedu, code)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var results []Report1Frm
	var result Report1Frm
	i := 1

	for rows.Next() {

		err := rows.Scan(
			&result.Description,
			&result.Point,
		)

		if err != nil {
			panic(err)
		}

		result.No = i
		results = append(results, result)
		i = i + 1
	}

	return results
}

func GetReportCheckin001(urldb string, rptdate string, room_ref string) []ReportCheckin001 {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var sql string

	if room_ref != "0" {

		sql = `
			select
				A.code,
				concat(A.first_name, ' ', A.last_name) as fullname,
				concat(C.name, '/' ,D.name) as classroom,
				ifnull(B.time_in, '-') as time_in,
				ifnull(
					(
						select time(created) 
						from checkinout 
						where date(created)=B.date 
						order by time(created) desc 
						limit 1
					), '-') as time_out,
				case
					when time_in <= E.time1 then 'ปกติ' 
					when time_in > E.time1 and time_in <= E.time2 then 'สาย' 
					when time_in > E.time2 and time_in <= E.time3 then 'สาย'
					when time_in > E.time3 then 'สาย' 
					else 'ขาดเรียน' 
				end as point
				
			from students as A

			left join (
				select idcard, date(created) as date, time(created) as time_in 
				from checkinout 
				where date(created)=?
				group by idcard
			) as B on B.idcard = A.code

			left join grade as C on C.ref = A.grade
			
			left join classroom as D on D.ref = A.room
			
			left join (select date_in, time1, time2, time3 from checkinprofile where grp=0) as E on E.date_in = B.date 

			where room=? 

			group by A.code;
			`

	} else {
		sql = `
				select
					A.code,
					concat(A.first_name, ' ', A.last_name) as fullname,
					concat(C.name, '/' ,D.name) as classroom,
					ifnull(B.time_in, '-') as time_in,
					ifnull(
						(
							select time(created)
							from checkinout
							where date(created)=B.date
							order by time(created) desc
							limit 1
						), '-') as time_out,
					case
						when time_in <= E.time1 then 'ปกติ' 
						when time_in > E.time1 and time_in <= E.time2 then 'สาย'
						when time_in > E.time2 and time_in <= E.time3 then 'สาย'
						when time_in > E.time3 then 'สาย'
						else 'ขาดเรียน' 
					end as point
	
				from students as A
	
				left join (
					select idcard, date(created) as date, time(created) as time_in
					from checkinout
					where date(created)=?
					group by idcard
				) as B on B.idcard = A.code
	
				left join grade as C on C.ref = A.grade
	
				left join classroom as D on D.ref = A.room
	
				left join (select date_in, time1, time2, time3 from checkinprofile where grp=0) as E on E.date_in = B.date
	
				order by A.grade, A.room, A.no;
			`
	}

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	var reports []ReportCheckin001
	var report ReportCheckin001

	if room_ref != "0" {
		rows, err := sttn.Query(rptdate, room_ref)

		if err != nil {
			panic(err)
		}

		defer rows.Close()

		for rows.Next() {

			err := rows.Scan(
				&report.Code,
				&report.FullName,
				&report.ClassRoom,
				&report.Time_in,
				&report.Time_out,
				&report.Status,
			)

			if err != nil {
				panic(err)
			}

			reports = append(reports, report)
		}

	} else {
		rows, err := sttn.Query(rptdate)

		if err != nil {
			panic(err)
		}

		defer rows.Close()

		for rows.Next() {

			err := rows.Scan(
				&report.Code,
				&report.FullName,
				&report.ClassRoom,
				&report.Time_in,
				&report.Time_out,
				&report.Status,
			)

			if err != nil {
				panic(err)
			}

			reports = append(reports, report)
		}
	}

	return reports
}

func GetNormal(urldb string, yearedu int, group int, room_ref int, rptdate string) int {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select ifnull(sum(D.count), 0) as normal

		from (
			select  
				case
				when time(B.created) <= C.time1 then 1
				else 0
				end as count

			from (
				select A.code   
				from students as A
				left join checkinout as B on B.idcard = A.code
				where room=? 
				group by A.code) as A
				left join (select * from checkinout where date(created)=?
			) as B on B.idcard=A.code

			left join (
				select * 
				from checkinprofile 
				where yearedu=? and grp=?
			) as C on C.date_in = date(B.created)
			
			group by A.code

		) as D;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(room_ref, rptdate, yearedu, group)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var count int

	if rows.Next() {
		err = rows.Scan(&count)

		if err != nil {
			panic(err)
		}
	}

	return count
}

func GetAbsent(urldb string, room_ref int, rptdate string) int {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select count(*)
		from (
			select A.code, B.created  
			from students as A
			left join (select * from checkinout where date(created)=?) as B on B.idcard = A.code
			where room=? and B.created is null
			group by code
		) as B;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(rptdate, room_ref)

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

	return count
}

func GetLate1(urldb string, room_ref int, rptdate string) int {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select ifnull(sum(point), 0)
		from (
			select 
				case
					when time_in > C.time1 and time_in <= time2 then 1
					else 0
				end as point
			from students as A
			left join (select idcard, created, time(created) as time_in from checkinout where date(created)=?) as B on B.idcard = A.code
			left join checkinprofile as C on C.date_in=date(B.created)
			where room=? and B.time_in is not null
			group by code
		) as D;
	`
	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(rptdate, room_ref)

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

	return count
}

func GetLate2(urldb string, room_ref int, rptdate string) int {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select ifnull(sum(point), 0)
		from (
			select 
				case
					when time_in > C.time2 then 1
					else 0
				end as point
			from students as A
			left join (select idcard, created, time(created) as time_in from checkinout where date(created)=?) as B on B.idcard = A.code
			left join checkinprofile as C on C.date_in=date(B.created)
			where room=? and B.time_in is not null
			group by code
		) as D;
	`
	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(rptdate, room_ref)

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

	return count
}

func GetReportCheckin002(urldb string, yearedu int, group int, rptdate string) []ReportCheckin002 {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			ifnull(B.name, ""),
			ifnull(A.ref, ""),
			A.name,
			(select count(*) from students where room=A.ref) as student_amoung
		from classroom as A
		left join grade as B on B.ref = A.grade_ref;
	`
	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query()

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var reports []ReportCheckin002
	var report ReportCheckin002
	var room_ref int

	for rows.Next() {

		err := rows.Scan(
			&report.ClassName,
			&room_ref,
			&report.Room,
			&report.StudentAmoung,
		)

		if err != nil {
			panic(err)
		}

		report.Normal = GetNormal(urldb, yearedu, group, room_ref, rptdate)
		report.Absent = GetAbsent(urldb, room_ref, rptdate)
		report.Late1 = GetLate1(urldb, room_ref, rptdate)
		report.Late2 = GetLate2(urldb, room_ref, rptdate)

		reports = append(reports, report)
	}

	return reports
}

func GetReport004(urldb string, rptdate string, room_ref string) []Report004 {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			no, 
			code, 
			grade,
			room,
			concat(first_name, ' ', last_name) as fullname, 
			ifnull(B.created, '') as created,
			ifnull(B.behavior_ref, 52) as behavior_ref
		from students as A
		left join (
			select created, behavior_ref, student_code 
			from point_students 
			where date(created)=? 
		) as B on A.code = B.student_code
		where room=? 
		order by cast(no as int) ;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(rptdate, room_ref)

	defer rows.Close()

	var results []Report004
	var result Report004

	for rows.Next() {

		err := rows.Scan(
			&result.No,
			&result.Code,
			&result.Grade,
			&result.Room,
			&result.FullName,
			&result.Created,
			&result.Behavior_ref,
		)

		if err != nil {
			panic(err)
		}

		result.ImgB64 = ""
		results = append(results, result)
	}

	return results
}

func UpdateReport004(urldb string, data Report004Frm) {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select student_code
		from point_students
		where date(created)=? and student_code=? 
		order by created limit 1
		;
	`
	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(data.Rptdate, data.Code)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if !rows.Next() {

		sql := `
			insert into point_students (yearedu, created, behavior_ref, student_code, point)
			values (2563,?,?,?,?)
		`
		sttn, err := db.Prepare(sql)

		if err != nil {
			panic(err)
		}

		defer sttn.Close()

		_, err = sttn.Exec(data.Rptdate, data.Behavior_ref, data.Code, data.Point)

		if err != nil {
			panic(err)
		}

	} else {

		sql := `
			update point_students set
				yearedu=2563,
				behavior_ref=?,
				point=?
			where created=? and student_code=?;
		`
		sttn, err := db.Prepare(sql)

		if err != nil {
			panic(err)
		}

		defer sttn.Close()

		_, err = sttn.Exec(data.Behavior_ref, data.Point, data.Created, data.Code)

		if err != nil {
			panic(err)
		}

	}

	switch data.Behavior_ref {

	case 50:
		fmt.Println("late1")

	case 51:
		fmt.Println("late2")

	case 52:
		fmt.Println("off")

	case 102:
		fmt.Println("normal")

	case 103:
		fmt.Println("event")

	}
}

func changToNormal(urldb string, data Report004Frm) {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select ref 
		from checkinout
		where date(created)=? adn idcard=?
		order by created limit 1
		;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(data.Created, data.Code)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {

		created := data.Rptdate + " 8:00:00"

		fmt.Println(created)

		sql = `
			update point_students set
				created=?,
				hehavior_ref=?,
				point=?
			where created=? and student_code=?
			;
		`

		sttn, err = db.Prepare(sql)

		if err != nil {
			panic(err)
		}

		_, err = sttn.Exec(created, data.Behavior_ref, data.Point, data.Created, data.Code)

		if err != nil {
			panic(err)
		}

	}

}

func getCheckin(urldb string, code string, date string) string {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select created
		from checkinout 
		where date(created)=? and idcard=?
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(date, code)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var result string

	result = ""

	if rows.Next() {
		result = "/"
	}

	return result
}

func GetReport005(urldb string, monthyr string, room string) []Report005 {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select
			A.code,
			concat(A.first_name, ' ', A.last_name) as fullname
		from students as A
		where A.room=?
		order by cast(no as int) 
		;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(room)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var reports []Report005
	var report Report005
	var total int

	for rows.Next() {

		err := rows.Scan(
			&report.Code,
			&report.Fullname,
		)

		if err != nil {
			panic(err)
		}

		total = 0

		report.D01 = getCheckin(urldb, report.Code, monthyr+"-01")

		if report.D01 != "" {
			total = total + 1
		}

		report.D02 = getCheckin(urldb, report.Code, monthyr+"-02")

		if report.D02 != "" {
			total = total + 1
		}

		report.D03 = getCheckin(urldb, report.Code, monthyr+"-03")

		if report.D03 != "" {
			total = total + 1
		}

		report.D04 = getCheckin(urldb, report.Code, monthyr+"-04")

		if report.D04 != "" {
			total = total + 1
		}

		report.D05 = getCheckin(urldb, report.Code, monthyr+"-05")

		if report.D05 != "" {
			total = total + 1
		}

		report.D06 = getCheckin(urldb, report.Code, monthyr+"-06")

		if report.D06 != "" {
			total = total + 1
		}

		report.D07 = getCheckin(urldb, report.Code, monthyr+"-07")

		if report.D07 != "" {
			total = total + 1
		}

		report.D08 = getCheckin(urldb, report.Code, monthyr+"-08")

		if report.D08 != "" {
			total = total + 1
		}

		report.D09 = getCheckin(urldb, report.Code, monthyr+"-09")

		if report.D09 != "" {
			total = total + 1
		}

		report.D10 = getCheckin(urldb, report.Code, monthyr+"-10")

		if report.D10 != "" {
			total = total + 1
		}

		report.D11 = getCheckin(urldb, report.Code, monthyr+"-11")

		if report.D11 != "" {
			total = total + 1
		}

		report.D12 = getCheckin(urldb, report.Code, monthyr+"-12")

		if report.D12 != "" {
			total = total + 1
		}

		report.D13 = getCheckin(urldb, report.Code, monthyr+"-13")

		if report.D13 != "" {
			total = total + 1
		}

		report.D14 = getCheckin(urldb, report.Code, monthyr+"-14")

		if report.D14 != "" {
			total = total + 1
		}

		report.D15 = getCheckin(urldb, report.Code, monthyr+"-15")

		if report.D15 != "" {
			total = total + 1
		}

		report.D16 = getCheckin(urldb, report.Code, monthyr+"-16")

		if report.D16 != "" {
			total = total + 1
		}

		report.D17 = getCheckin(urldb, report.Code, monthyr+"-17")

		if report.D17 != "" {
			total = total + 1
		}

		report.D18 = getCheckin(urldb, report.Code, monthyr+"-18")

		if report.D18 != "" {
			total = total + 1
		}

		report.D19 = getCheckin(urldb, report.Code, monthyr+"-19")

		if report.D19 != "" {
			total = total + 1
		}

		report.D20 = getCheckin(urldb, report.Code, monthyr+"-20")

		if report.D20 != "" {
			total = total + 1
		}

		report.D21 = getCheckin(urldb, report.Code, monthyr+"-21")

		if report.D21 != "" {
			total = total + 1
		}

		report.D22 = getCheckin(urldb, report.Code, monthyr+"-22")

		if report.D22 != "" {
			total = total + 1
		}

		report.D23 = getCheckin(urldb, report.Code, monthyr+"-23")

		if report.D23 != "" {
			total = total + 1
		}

		report.D24 = getCheckin(urldb, report.Code, monthyr+"-24")

		if report.D24 != "" {
			total = total + 1
		}

		report.D25 = getCheckin(urldb, report.Code, monthyr+"-25")

		if report.D25 != "" {
			total = total + 1
		}

		report.D26 = getCheckin(urldb, report.Code, monthyr+"-26")

		if report.D26 != "" {
			total = total + 1
		}

		report.D27 = getCheckin(urldb, report.Code, monthyr+"-27")

		if report.D27 != "" {
			total = total + 1
		}

		report.D28 = getCheckin(urldb, report.Code, monthyr+"-28")

		if report.D28 != "" {
			total = total + 1
		}

		report.D29 = getCheckin(urldb, report.Code, monthyr+"-29")

		if report.D29 != "" {
			total = total + 1
		}

		report.D30 = getCheckin(urldb, report.Code, monthyr+"-30")

		if report.D30 != "" {
			total = total + 1
		}

		report.D31 = getCheckin(urldb, report.Code, monthyr+"-31")

		if report.D31 != "" {
			total = total + 1
		}

		report.Total = strconv.Itoa(total)

		reports = append(reports, report)
	}

	return reports
}
