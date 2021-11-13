package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	//"net/http/cgi"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	//"strconv"
	"schoolcore/behavioral"
	"schoolcore/behaviorpoint"
	"schoolcore/cam"
	"schoolcore/checkinout"
	"schoolcore/checkinprofile"
	"schoolcore/checkinsubject"
	"schoolcore/cutpoint"
	"schoolcore/department"
	"schoolcore/employee"
	"schoolcore/grade"
	"schoolcore/linereg"
	"schoolcore/parent"
	"schoolcore/province"
	"schoolcore/register"
	"schoolcore/report"
	"schoolcore/room"
	"schoolcore/sdq"
	"schoolcore/student"
	"schoolcore/subject"
	"schoolcore/subjectgroup"
	"schoolcore/subjectstudent"
	"schoolcore/teacher"
	"schoolcore/title"
	"schoolcore/user"
	//"log"
	//"schoolcore/database"
)

type Message struct {
	Service string
	Action  int // 1=add, 2=edit, 3=delete
	Data    string
}

type Picture struct {
	Camno       int
	Datetime    string
	Idcard      string
	Temperature float32
	Faceimage   string
}

const urldb = "root:35.103232@tcp(localhost:3306)/school"
const urlregdb = "root:35.103232@tcp(localhost:3306)/register"

//const urldb = "admin:35.103232@tcp(localhost:3306)/school"

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	(*w).Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	(*w).Header().Set("content-type", "application/json; charset=utf-8")
}

func CGIHandle(res http.ResponseWriter, req *http.Request) {

	enableCors(&res)

	if req.Method == "OPTIONS" {
		res.WriteHeader(http.StatusOK)
		return
	}

	parms_service, _ := req.URL.Query()["service"]
	parms_action, _ := req.URL.Query()["action"]

	var msgOut []byte

	switch service := parms_service[0]; service {

	case "subjectstudent":

		switch action := parms_action[0]; action {

		case "get":
			var yearedu = req.URL.Query()["yearedu"][0]
			var grade_ref = req.URL.Query()["grade_ref"][0]
			var subject_ref = req.URL.Query()["subject_ref"][0]

			result := subjectstudent.Get(urldb, yearedu, grade_ref, subject_ref)
			msgOut, _ = json.Marshal(result)

		case "get_by_room":
			var yearedu = req.URL.Query()["yearedu"][0]
			var room_ref = req.URL.Query()["room_ref"][0]
			var subject_ref = req.URL.Query()["subject_ref"][0]

			result := subjectstudent.GetByRoom(urldb, yearedu, room_ref, subject_ref)
			msgOut, _ = json.Marshal(result)

		case "insert_all_grade":
			var yearedu = req.URL.Query()["yearedu"][0]
			var grade_ref = req.URL.Query()["grade_ref"][0]
			var subject_ref = req.URL.Query()["subject_ref"][0]

			result := subjectstudent.InsertAllGrade(urldb, yearedu, grade_ref, subject_ref)
			msgOut, _ = json.Marshal(result)

		case "insert_by_room":
			var yearedu = req.URL.Query()["yearedu"][0]
			var room_ref = req.URL.Query()["room_ref"][0]
			var subject_ref = req.URL.Query()["subject_ref"][0]

			result := subjectstudent.InsertByRoom(urldb, yearedu, room_ref, subject_ref)
			msgOut, _ = json.Marshal(result)

		case "insert_by_code":
			var yearedu = req.URL.Query()["yearedu"][0]
			var student_code = req.URL.Query()["student_code"][0]
			var subject_ref = req.URL.Query()["subject_ref"][0]

			result := subjectstudent.InsertByCode(urldb, yearedu, subject_ref, student_code)
			msgOut, _ = json.Marshal(result)

		case "delete_all_grade":
			var yearedu = req.URL.Query()["yearedu"][0]
			var grade_ref = req.URL.Query()["grade_ref"][0]
			var subject_ref = req.URL.Query()["subject_ref"][0]

			result := subjectstudent.DeleteAllGrade(urldb, yearedu, grade_ref, subject_ref)
			msgOut, _ = json.Marshal(result)

		case "delete_by_room":
			var yearedu = req.URL.Query()["yearedu"][0]
			var room_ref = req.URL.Query()["room_ref"][0]
			var subject_ref = req.URL.Query()["subject_ref"][0]

			result := subjectstudent.DeleteByRoom(urldb, yearedu, room_ref, subject_ref)
			msgOut, _ = json.Marshal(result)

		case "delete_by_code":
			var yearedu = req.URL.Query()["yearedu"][0]
			var student_code = req.URL.Query()["student_code"][0]
			var subject_ref = req.URL.Query()["subject_ref"][0]

			result := subjectstudent.DeleteByCode(urldb, yearedu, student_code, subject_ref)
			msgOut, _ = json.Marshal(result)
		}

	case "checkinsubject":

		switch action := parms_action[0]; action {

		case "get":
			var created = req.URL.Query()["created"][0]
			var room_ref = req.URL.Query()["room_ref"][0]

			result := checkinsubject.Get(urldb, created, room_ref)
			msgOut, _ = json.Marshal(result)

		case "get_by_key":
			var created = req.URL.Query()["created"][0]
			var period = req.URL.Query()["period"][0]
			var room_ref = req.URL.Query()["room_ref"][0]

			result := checkinsubject.GetByKey(urldb, created, period, room_ref)
			msgOut, _ = json.Marshal(result)

		case "get_status_normal":
			var created = req.URL.Query()["created"][0]
			var room_ref = req.URL.Query()["room_ref"][0]

			result := checkinsubject.GetNormal(urldb, created, room_ref)
			msgOut, _ = json.Marshal(result)

		case "get_status_absent":
			var created = req.URL.Query()["created"][0]
			var room_ref = req.URL.Query()["room_ref"][0]

			result := checkinsubject.GetAbsent(urldb, created, room_ref)
			msgOut, _ = json.Marshal(result)

		case "check_duplicate":
			var created = req.URL.Query()["created"][0]
			var period = req.URL.Query()["period"][0]
			var room_ref = req.URL.Query()["room_ref"][0]

			result := checkinsubject.CheckDuplicate(urldb, created, period, room_ref)
			msgOut, _ = json.Marshal(result)

		case "save":
			var payload checkinsubject.CheckinSubject

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &payload)
			result := checkinsubject.Save(urldb, payload)

			msgOut, _ = json.Marshal(result)

		case "change":
			var created = req.URL.Query()["created"][0]
			var period = req.URL.Query()["period"][0]
			var room_ref = req.URL.Query()["room_ref"][0]
			var payload checkinsubject.CheckinSubject

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &payload)
			result := checkinsubject.Change(urldb, created, period, room_ref, payload)

			msgOut, _ = json.Marshal(result)

		case "update":
			var payload checkinsubject.CheckinSubject

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &payload)
			result := checkinsubject.Update(urldb, payload)

			msgOut, _ = json.Marshal(result)

		case "delete":
			var created = req.URL.Query()["created"][0]
			var period = req.URL.Query()["period"][0]
			var room_ref = req.URL.Query()["room_ref"][0]

			result := checkinsubject.Delete(urldb, created, period, room_ref)
			msgOut, _ = json.Marshal(result)
		}

	case "subject":

		switch action := parms_action[0]; action {

		case "get":
			result := subject.Get(urldb)
			msgOut, _ = json.Marshal(result)

		case "get_by_group":
			var group_ref = req.URL.Query()["group_ref"][0]

			result := subject.GetByGroup(urldb, group_ref)
			msgOut, _ = json.Marshal(result)

		case "get_by_grade_group":
			var grade_ref = req.URL.Query()["grade_ref"][0]
			var group_ref = req.URL.Query()["group_ref"][0]

			result := subject.GetByGradeGroup(urldb, grade_ref, group_ref)
			msgOut, _ = json.Marshal(result)

		case "save":
			var payload subject.Subject

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &payload)
			result := subject.Save(urldb, payload)
			msgOut, _ = json.Marshal(result)

		case "del":
			var payload subject.Subject

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &payload)
			result := subject.Del(urldb, payload)
			msgOut, _ = json.Marshal(result)
		}

	case "subjectgroup":

		switch action := parms_action[0]; action {

		case "get":
			result := subjectgroup.Get(urldb)
			msgOut, _ = json.Marshal(result)

		case "save":
			var payload subjectgroup.SubjectGroup

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &payload)
			result := subjectgroup.Save(urldb, payload)
			msgOut, _ = json.Marshal(result)

		case "del":
			var payload subjectgroup.SubjectGroup

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &payload)
			result := subjectgroup.Del(urldb, payload)
			msgOut, _ = json.Marshal(result)

		}

	case "line":

		switch action := parms_action[0]; action {

		case "register_student":
			var payload linereg.StudentReg
			body, _ := ioutil.ReadAll(req.Body)

			if string(body) != "" {
				json.Unmarshal([]byte(string(body)), &payload)
				msgOut, _ = json.Marshal(linereg.RegisterStudent(urldb, payload))
			}

		}

	case "register":

		switch action := parms_action[0]; action {

		case "get":
			result := register.Get(urlregdb)
			msgOut, _ = json.Marshal(result)
		}

	case "user":

		switch action := parms_action[0]; action {

		case "login":
			var payload user.Payload

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &payload)
			result := user.Login(urldb, payload)
			msgOut, _ = json.Marshal(result)

		case "get_all":
			result := user.GetAll(urldb)
			msgOut, _ = json.Marshal(result)

		case "get_by_ref":
			var ref = req.URL.Query()["ref"][0]

			result := user.GetByRef(urldb, ref)
			msgOut, _ = json.Marshal(result)

		case "save":
			var userInfo user.User

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &userInfo)

			if userInfo.Name != "" {
				result := user.Save(urldb, userInfo)
				msgOut, _ = json.Marshal(result)
			}

		case "delete":
			var ref = req.URL.Query()["ref"][0]

			result := user.Remove(urldb, ref)
			msgOut, _ = json.Marshal(result)
		}

	case "cam":

		body, err := ioutil.ReadAll(req.Body)

		if err != nil {
			panic(err)
		}

		switch action := parms_action[0]; action {

		case "save":
			var pic cam.Picture
			json.Unmarshal([]byte(string(body)), &pic)

			if pic.Camno != 0 {
				result := cam.Save(urldb, pic)
				msgOut, _ = json.Marshal(result)
			}
		}

	case "faceimg":
		parms_datetime, _ := req.URL.Query()["datetime"]
		parms_code, _ := req.URL.Query()["code"]

		dt := strings.Fields(parms_datetime[0])

		year := strings.Split(dt[0], "-")[0]
		month := strings.Split(dt[0], "-")[1]
		day := strings.Split(dt[0], "-")[2]
		path := "imgs/checkinout/" + year + "/" + month + "/" + day + "/"
		filename := dt[0] + "_" + dt[1] + "_" + parms_code[0] + ".b64"

		content, err := ioutil.ReadFile(path + filename)

		if err != nil {
			panic(err)
		}

		msgOut = content

	case "teacher":

		switch action := parms_action[0]; action {

		case "get":
			result := teacher.Get(urldb)
			msgOut, _ = json.Marshal(result)

		case "get_by_ref":
			var ref = req.URL.Query()["ref"][0]

			if ref != "" {
				result := teacher.GetByRef(urldb, ref)
				msgOut, _ = json.Marshal(result)
			}

		case "get_by_code":
			var code = req.URL.Query()["code"][0]

			result := teacher.GetByCode(urldb, code)
			msgOut, _ = json.Marshal(result)

		case "get_by_grade":
			var grade = req.URL.Query()["grade"][0]

			if grade != "" {
				result := teacher.GetByGrade(urldb, grade)
				msgOut, _ = json.Marshal(result)
			}

		case "save":
			var teacherInfo teacher.Teacher

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &teacherInfo)

			if teacherInfo.Code != "" {
				result := teacher.Save(urldb, teacherInfo)
				msgOut, _ = json.Marshal(result)
			}

		case "delete":
			var ref = req.URL.Query()["ref"][0]
			var grade_ref = req.URL.Query()["grade_ref"][0]

			if ref != "" {
				result := teacher.Remove(urldb, ref, grade_ref)
				msgOut, _ = json.Marshal(result)
			}

		}

	case "student":

		switch action := parms_action[0]; action {

		case "get":
			grade, _ := req.URL.Query()["grade"]
			room, _ := req.URL.Query()["room"]
			txtsearch, _ := req.URL.Query()["txtsearch"]

			studentInfos := student.Get(urldb, grade[0], room[0], txtsearch[0])
			msgOut, _ = json.Marshal(studentInfos)

		case "2":
			var studentInfo student.StudentInfo

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &studentInfo)

			if studentInfo.Code != "" {
				studentInfos := student.Save(urldb, studentInfo)
				msgOut, _ = json.Marshal(studentInfos)
			}

		case "3":
			code, _ := req.URL.Query()["code"]

			studentInfos := student.GetByCode(urldb, code[0])
			msgOut, _ = json.Marshal(studentInfos)

		case "4":
			parentref, _ := req.URL.Query()["parentref"]

			studentInfos := student.GetByParent(urldb, parentref[0])
			msgOut, _ = json.Marshal(studentInfos)

		case "5":
			uid, _ := req.URL.Query()["uid"]

			studentInfo := student.GetByUID(urldb, uid[0])
			msgOut, _ = json.Marshal(studentInfo)

		case "get_faceimg":
			code, _ := req.URL.Query()["code"]

			msgOut = student.GetFaceImg(code[0])

		case "7":
			code := req.URL.Query()["code"][0]

			msgOut = student.GetFaceImg(code)

		case "import":
			var importlot student.ImportLot

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &importlot)

		case "get_by_ref":
			ref, _ := req.URL.Query()["ref"]

			result := student.GetByRef(urldb, ref[0])
			msgOut, _ = json.Marshal(result)

		case "delete":
			ref := req.URL.Query()["ref"][0]

			result := student.Remove(urldb, ref)
			msgOut, _ = json.Marshal(result)

		}

	case "parent":

		switch action := parms_action[0]; action {

		case "1":
			result := parent.Get(urldb)
			msgOut, _ = json.Marshal(result)

		case "2":
			var parentInfo parent.Parent

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &parentInfo)

			if parentInfo.FirstName != "" {
				result := parent.Save(urldb, parentInfo)
				msgOut, _ = json.Marshal(result)
			}

		case "3":
			studentref, _ := req.URL.Query()["studentref"]
			result := parent.GetByStudent(urldb, studentref[0])
			msgOut, _ = json.Marshal(result)

		case "get_by_ref":
			ref := req.URL.Query()["ref"][0]

			result := parent.GetByRef(urldb, ref)
			msgOut, _ = json.Marshal(result)

		case "delete":
			ref := req.URL.Query()["ref"][0]

			result := parent.Remove(urldb, ref)
			msgOut, _ = json.Marshal(result)
		}

	case "checkinout":

		switch action := parms_action[0]; action {

		case "1":
			result := checkinout.Get(urldb)
			msgOut, _ = json.Marshal(result)

		case "get_by_ref":
			ref := req.URL.Query()["ref"][0]

			result := checkinout.GetByRef(urldb, ref)
			msgOut, _ = json.Marshal(result)

		case "get_by_code":
			code := req.URL.Query()["code"][0]

			result := checkinout.GetByCode(urldb, code)
			msgOut, _ = json.Marshal(result)

		}

	case "checkinprofile":

		switch action := parms_action[0]; action {

		case "get":
			var result []checkinprofile.CheckinProfile

			result = checkinprofile.Get(urldb)
			msgOut, _ = json.Marshal(result)

		case "save":
			var checkinprofileInfo checkinprofile.CheckinProfile

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &checkinprofileInfo)

			if checkinprofileInfo.DateIn != "" {
				result := checkinprofile.Save(urldb, checkinprofileInfo)
				msgOut, _ = json.Marshal(result)
			}

		case "get_group":
			result := checkinprofile.GetGroups(urldb)
			msgOut, _ = json.Marshal(result)

		case "delete":
			var yearedu, _ = strconv.Atoi(req.URL.Query()["yearedu"][0])
			var group, _ = strconv.Atoi(req.URL.Query()["group"][0])
			var datein = req.URL.Query()["datein"][0]

			result := checkinprofile.Remove(urldb, yearedu, group, datein)
			msgOut, _ = json.Marshal(result)
		}

	case "behavioral":

		switch action := parms_action[0]; action {

		case "1":
			result := behavioral.Get(urldb)
			msgOut, _ = json.Marshal(result)

		case "2":
			var behavioralInfo behavioral.BehavioralPoint

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &behavioralInfo)

			if behavioralInfo.GroupName != "" {
				result := behavioral.Save(urldb, behavioralInfo)
				msgOut, _ = json.Marshal(result)
			}

		case "3":
			result := behavioral.GetGroup(urldb)
			out, err := json.Marshal(result)

			if err != nil {
				panic(err)
			}

			msgOut = out

		case "4":
			groupno := req.URL.Query()["groupno"][0]

			result := behavioral.GetTopic(urldb, groupno)
			msgOut, _ = json.Marshal(result)

		case "5":
			groupno := req.URL.Query()["groupno"][0]

			var result []behavioral.BehavioralPoint

			if groupno == "0" {
				result = behavioral.Get(urldb)

			} else {
				result = behavioral.GetByGroup(urldb, groupno)
			}

			msgOut, _ = json.Marshal(result)

		case "6":
			var behavioralInfo behavioral.BehavioralPoint

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &behavioralInfo)

			if behavioralInfo.GroupName != "" {
				result := behavioral.Remove(urldb, behavioralInfo)
				msgOut, _ = json.Marshal(result)
			}

		case "get_by_ref":
			var ref = req.URL.Query()["ref"][0]

			result := behavioral.GetByRef(urldb, ref)
			msgOut, _ = json.Marshal(result)
		}

	case "behaviorpoint":

		switch action := parms_action[0]; action {

		case "2":
			var result []behaviorpoint.BehaviorPoint

			result = behaviorpoint.Get(urldb)
			msgOut, _ = json.Marshal(result)
		}

	case "province":

		switch action := parms_action[0]; action {

		case "1":
			result := province.Get(urldb)
			msgOut, _ = json.Marshal(result)

		case "2":
			var provinceInfo province.Province

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &provinceInfo)
			result := province.Save(urldb, provinceInfo)
			msgOut, _ = json.Marshal(result)
		}

	case "title":

		switch action := parms_action[0]; action {

		case "1":
			result := title.Get(urldb)
			msgOut, _ = json.Marshal(result)

		case "2":
			var data title.Title

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &data)

			if data.Name != "" {
				result := title.Save(urldb, data)
				msgOut, _ = json.Marshal(result)
			}
		}

	case "grade":

		switch action := parms_action[0]; action {

		case "1":
			result := grade.Get(urldb)
			msgOut, _ = json.Marshal(result)

		case "2":
			var gradeInfo grade.Grade

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &gradeInfo)

			if gradeInfo.Name != "" {
				result := grade.Save(urldb, gradeInfo)
				msgOut, _ = json.Marshal(result)
			}

		case "3":
			var gradeInfo grade.Grade

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &gradeInfo)

			if gradeInfo.Name != "" {
				result := grade.Remove(urldb, gradeInfo)
				msgOut, _ = json.Marshal(result)
			}

		case "get_by_ref":

			ref := req.URL.Query()["ref"][0]
			result := grade.GetByRef(urldb, ref)
			msgOut, _ = json.Marshal(result)
		}

	case "room":

		switch action := parms_action[0]; action {

		case "1":
			result := room.Get(urldb)
			msgOut, _ = json.Marshal(result)

		case "2":
			var roomInfo room.Room

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &roomInfo)

			if roomInfo.GradeRef != 0 {
				result := room.Save(urldb, roomInfo)
				msgOut, _ = json.Marshal(result)
			}

		case "3":
			gradeRef := req.URL.Query()["grade_ref"][0]
			result := room.GetByGrade(urldb, gradeRef)
			msgOut, _ = json.Marshal(result)

		case "4":
			var roomInfo room.Room

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &roomInfo)

			if roomInfo.Ref != 0 {
				result := room.Remove(urldb, roomInfo)
				msgOut, _ = json.Marshal(result)
			}

		case "get_by_ref":
			ref := req.URL.Query()["ref"][0]
			result := room.GetByRef(urldb, ref)
			msgOut, _ = json.Marshal(result)

		}

	case "cutpoint":

		switch action := parms_action[0]; action {

		case "1":
			result := cutpoint.Get(urldb)
			msgOut, _ = json.Marshal(result)

		case "2":
			var data cutpoint.Cutpoint

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &data)
			result := cutpoint.Save(urldb, data)
			msgOut, _ = json.Marshal(result)

		case "3":
			cutpoint_ref := req.URL.Query()["cutpoint_ref"][0]
			result := cutpoint.GetStudentsByRef(urldb, cutpoint_ref)
			msgOut, _ = json.Marshal(result)

		case "4":
			student_ref := req.URL.Query()["student_ref"][0]
			result := cutpoint.GetByStudent(urldb, student_ref)
			msgOut, _ = json.Marshal(result)

		case "5":
			room_ref := req.URL.Query()["room_ref"][0]

			result := cutpoint.GetReport1(urldb, room_ref)
			msgOut, _ = json.Marshal(result)

		case "6":
			code := req.URL.Query()["code"][0]

			result := cutpoint.GetStudentBehaviorPoint(urldb, 2564, code)
			msgOut, _ = json.Marshal(result)
		}

	case "report":

		switch action := parms_action[0]; action {

		case "report001":
			room_ref := req.URL.Query()["room_ref"][0]

			result := report.GetReport001(urldb, 2564, room_ref)
			msgOut, _ = json.Marshal(result)

		case "report001_frm":
			code := req.URL.Query()["code"][0]

			result := report.GetReport001Frm(urldb, 2564, code)
			msgOut, _ = json.Marshal(result)

		case "checkin001":
			rptdate := req.URL.Query()["rptdate"][0]
			room_ref := req.URL.Query()["room_ref"][0]

			result := report.GetReportCheckin001(urldb, rptdate, room_ref)
			msgOut, _ = json.Marshal(result)

		case "checkin002":
			rptdate := req.URL.Query()["rptdate"][0]

			result := report.GetReportCheckin002(urldb, 2564, 0, rptdate)
			msgOut, _ = json.Marshal(result)

		case "report004":
			rptdate := req.URL.Query()["rptdate"][0]
			room_ref := req.URL.Query()["room_ref"][0]

			result := report.GetReport004(urldb, rptdate, room_ref)
			msgOut, _ = json.Marshal(result)

		case "report004_frm":
			var data report.Report004Frm

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &data)

			if data.Rptdate != "" {
				report.UpdateReport004(urldb, data)
				msgOut = []byte("success")
			}

		case "report005":
			monthyr := req.URL.Query()["monthyr"][0]
			room_ref := req.URL.Query()["room_ref"][0]

			result := report.GetReport005(urldb, monthyr, room_ref)
			msgOut, _ = json.Marshal(result)
		}

	case "department":

		switch action := parms_action[0]; action {

		case "get":
			result := department.Get(urldb)
			msgOut, _ = json.Marshal(result)

		case "save":
			var data department.Department

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &data)

			if data.Name != "" {
				result := department.Save(urldb, data)
				msgOut, _ = json.Marshal(result)
			}

		case "delete":
			ref := req.URL.Query()["ref"][0]

			if ref != "" {
				result := department.Delete(urldb, ref)
				msgOut, _ = json.Marshal(result)
			}

		}

	case "employee":

		switch action := parms_action[0]; action {

		case "get":
			result := employee.Get(urldb)
			msgOut, _ = json.Marshal(result)

		case "save":
			var data employee.Employee

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &data)

			if data.FirstName != "" {
				result := employee.Save(urldb, data)
				msgOut, _ = json.Marshal(result)
			}
		}

	case "sdq":

		switch action := parms_action[0]; action {

		case "get_student_by_code":
			student_code := req.URL.Query()["code"][0]

			if student_code != "" {
				result := sdq.GetStudentByCode(urldb, student_code)
				msgOut, _ = json.Marshal(result)
			}

		case "get_by_code":
			code := req.URL.Query()["code"][0]
			t := req.URL.Query()["type"][0]

			if code != "" {
				result := sdq.GetByCode(urldb, code, t)
				msgOut, _ = json.Marshal(result)
			}

		case "save":
			var data sdq.SDQ

			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal([]byte(string(body)), &data)

			if data.Code != "" {
				result := sdq.Save(urldb, data)
				msgOut, _ = json.Marshal(result)
			}
		}

	}

	fmt.Fprintf(res, "%s\n", msgOut)
}

func handleCam(msg Message) {

	var pic Picture

	camno := pic.Camno
	datetime := strings.ReplaceAll(strings.ReplaceAll(pic.Datetime, "T", " "), "+07:00", "")
	idcard := pic.Idcard
	temperature := pic.Temperature
	faceimg := pic.Faceimage

	facefile := []byte(faceimg)
	ioutil.WriteFile("imgs/faces.b64", facefile, 644)

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sttn, err := db.Prepare(`
		insert into checkinout (
			camno, 
			created, 
			temperature )
		values (?,?,?,?)
	`)

	if err != nil {
		panic(err)
	}

	sttn.Exec(camno, datetime, idcard, temperature)
}

func main() {
	//err := cgi.Serve(http.HandlerFunc(CGIHandle))

	//if err != nil {
	//	panic(err)
	//}

	fmt.Println("Start server port: 8080")

	http.HandleFunc("/schoolcore", CGIHandle)
	http.ListenAndServe(":8080", nil)
}
