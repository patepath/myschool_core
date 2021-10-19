package behavioral

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type BehavioralPoint struct {
	GroupNo   int
	GroupName string
	Ref       int
	TopicNo   int
	TopicName string
	Point     int
}

type Opt struct {
	Value int
	Name  string
}

func Get(urldb string) []BehavioralPoint {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			ref,
			groupno,
			groupname,
			topicno,
			topicname,
			point
		from behavioral_point 
		where groupno <> 0
		order by groupno, topicno`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var behaviorals []BehavioralPoint
	var behavioral BehavioralPoint

	for rows.Next() {
		err := rows.Scan(
			&behavioral.Ref,
			&behavioral.GroupNo,
			&behavioral.GroupName,
			&behavioral.TopicNo,
			&behavioral.TopicName,
			&behavioral.Point,
		)

		if err != nil {
		}

		behaviorals = append(behaviorals, behavioral)
	}

	return behaviorals
}

func GetByRef(urldb string, ref string) BehavioralPoint {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			ref,
			groupno,
			groupname,
			topicno,
			topicname,
			point
		from behavioral_point 
		where ref=?;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(ref)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var behavioral BehavioralPoint

	if rows.Next() {
		err := rows.Scan(
			&behavioral.Ref,
			&behavioral.GroupNo,
			&behavioral.GroupName,
			&behavioral.TopicNo,
			&behavioral.TopicName,
			&behavioral.Point,
		)

		if err != nil {
			panic(err)
		}

	}

	return behavioral
}

func GetByGroup(urldb string, groupno string) []BehavioralPoint {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			ref,
			groupno,
			groupname,
			topicno,
			topicname,
			point
		from behavioral_point 
		where groupno=?
		order by groupno, topicno`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(groupno)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var behaviorals []BehavioralPoint
	var behavioral BehavioralPoint

	for rows.Next() {
		err := rows.Scan(
			&behavioral.Ref,
			&behavioral.GroupNo,
			&behavioral.GroupName,
			&behavioral.TopicNo,
			&behavioral.TopicName,
			&behavioral.Point,
		)

		if err != nil {
		}

		behaviorals = append(behaviorals, behavioral)
	}

	return behaviorals

}

func GetGroup(urldb string) []Opt {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select groupno, groupname 
		from behavioral_point
		where groupno <> 0
		group by groupno

	`
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var groups []Opt
	var group Opt

	for rows.Next() {
		err := rows.Scan(
			&group.Value,
			&group.Name,
		)

		if err == nil {
		}

		groups = append(groups, group)
	}

	return groups
}

func GetTopic(urldb string, groupno string) []Opt {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	cond := "where groupno=" + groupno

	sql := `
		select ref, topicname 
		from behavioral_point 
	` + cond

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var topics []Opt
	var topic Opt

	for rows.Next() {
		err := rows.Scan(
			&topic.Value,
			&topic.Name,
		)

		if err == nil {
		}

		topics = append(topics, topic)
	}

	return topics
}

func Save(urldb string, behavioral BehavioralPoint) []BehavioralPoint {

	if behavioral.GroupNo == 0 {
		behavioral.GroupNo = getNewGroupNo(urldb)
	}

	if behavioral.Ref == 0 {
		insert(urldb, behavioral)
	} else {
		update(urldb, behavioral)
	}

	return GetByGroup(urldb, strconv.Itoa(behavioral.GroupNo))
}

func getNewGroupNo(urldb string) int {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select groupno
		from behavioral_point 
		order by groupno desc
		limit 1;
	`

	var groupno int

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&groupno)

		if err != nil {
			panic(err)
		}

		groupno = groupno + 1

	} else {
		groupno = 1
	}

	return groupno
}

func getNewTopicNo(urldb string, groupno int) int {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select topicno 
		from behavioral_point
		where groupno=` + strconv.Itoa(groupno) + `
		order by topicno desc
		limit 1;
	`
	var topicno int

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&topicno)

		if err != nil {
			panic(err)
		}

		topicno = topicno + 1

	} else {
		topicno = 1
	}

	return topicno
}

func insert(urldb string, behavioral BehavioralPoint) {

	behavioral.TopicNo = getNewTopicNo(urldb, behavioral.GroupNo)

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		insert into  behavioral_point (
			groupno,
			groupname,
			topicno,
			topicname,
			point
		) values(?,?,?,?,?)
	`
	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	_, err = sttn.Exec(
		behavioral.GroupNo,
		behavioral.GroupName,
		behavioral.TopicNo,
		behavioral.TopicName,
		behavioral.Point,
	)

	if err != nil {
		panic(err)
	}
}

func update(urldb string, behavioral BehavioralPoint) {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		update behavioral_point set 
			groupno=?,
			groupname=?,
			topicno=?,
			topicname=?,
			point=?
		where ref=?
	`
	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	_, err = sttn.Exec(
		behavioral.GroupNo,
		behavioral.GroupName,
		behavioral.TopicNo,
		behavioral.TopicName,
		behavioral.Point,
		behavioral.Ref,
	)

	if err != nil {
		panic(err)
	}
}

func Remove(urldb string, behavior BehavioralPoint) []BehavioralPoint {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `delete from behavioral_point where ref=?`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(behavior.Ref)

	if err != nil {
		panic(err)
	}

	return GetByGroup(urldb, strconv.Itoa(behavior.GroupNo))
}
