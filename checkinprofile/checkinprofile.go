package checkinprofile

import (
	"fmt"
	//"os"
	//"strings"
	//"io/ioutil"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type CheckinProfile struct {
	Group   int
	DateIn  string
	YearEdu int
	Name    string
	Time1   string
	Time2   string
	Time3   string
	Point1  int
	Point2  int
	Point3  int
}

func Get(urldb string) []CheckinProfile {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			grp, 
			date_in, 
			yearedu, 
			ifnull(name, '') as name, 
			time1, 
			time2, 
			time3, 
			point1, 
			point2, 
			point3
		from checkinprofile
		order by grp, date_in
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var checkinprofiles []CheckinProfile
	var checkinprofile CheckinProfile

	for rows.Next() {
		err := rows.Scan(
			&checkinprofile.Group,
			&checkinprofile.DateIn,
			&checkinprofile.YearEdu,
			&checkinprofile.Name,
			&checkinprofile.Time1,
			&checkinprofile.Time2,
			&checkinprofile.Time3,
			&checkinprofile.Point1,
			&checkinprofile.Point2,
			&checkinprofile.Point3,
		)

		if err != nil {
			panic(err)
		}

		checkinprofiles = append(checkinprofiles, checkinprofile)
	}

	return checkinprofiles
}

func GetGroups(urldb string) []string {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select ifnull( name, '')
		from checkinprofile
		group by name;
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var names []string
	var name string

	for rows.Next() {

		err := rows.Scan(&name)

		if err != nil {
			panic(err)
		}

		names = append(names, name)
	}

	return names
}

func Save(urldb string, checkinprofile CheckinProfile) []CheckinProfile {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
	select * 
	from checkinprofile 
	where yearedu=? and grp=? and date_in=?
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(
		checkinprofile.YearEdu,
		checkinprofile.Group,
		checkinprofile.DateIn,
	)

	if err != nil {
	}

	defer rows.Close()

	if rows.Next() {
		sql := `
		update checkinprofile set
			name=?,
			time1=?,
			time2=?,
			time3=?,
			point1=?,
			point2=?,
			point3=?
		where yearedu=? and grp=? and date_in=?;
		`

		sttn, err := db.Prepare(sql)

		if err != nil {
			panic(err)
		}

		defer sttn.Close()

		_, err = sttn.Exec(
			checkinprofile.Name,
			checkinprofile.Time1,
			checkinprofile.Time2,
			checkinprofile.Time3,
			checkinprofile.Point1,
			checkinprofile.Point2,
			checkinprofile.Point3,
			checkinprofile.YearEdu,
			checkinprofile.Group,
			checkinprofile.DateIn,
		)

		if err != nil {
			fmt.Println("update: ", err)
		}

	} else {
		sql := `
		insert into checkinprofile (
			yearedu,
			grp,
			name,
			date_in,
			time1,
			time2,
			time3,
			point1,
			point2,
			point3
		) values(?,?,?,?,?,?,?,?,?,?)
		`
		sttn, err := db.Prepare(sql)

		if err != nil {
			panic(err)
		}

		defer sttn.Close()

		_, err = sttn.Exec(
			checkinprofile.YearEdu,
			checkinprofile.Group,
			checkinprofile.Name,
			checkinprofile.DateIn,
			checkinprofile.Time1,
			checkinprofile.Time2,
			checkinprofile.Time3,
			checkinprofile.Point1,
			checkinprofile.Point2,
			checkinprofile.Point3,
		)

		if err != nil {
			fmt.Println("insert: ", err)
		}

	}

	return Get(urldb)
}

func Remove(urldb string, yearedu int, group int, datein string) []CheckinProfile {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		delete from checkinprofile 
		where yearedu=? and grp=? and date_in=?;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(yearedu, group, datein)

	if err != nil {
		panic(err)
	}

	return Get(urldb)
}
