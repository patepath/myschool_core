package behaviorpoint

import (
	//"fmt"
	//"os"
	//"strings"
	//"io/ioutil"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


type BehaviorPoint struct {
	Code	string
	Title	string
	FullName	string
	Grade	string
	Room	string
	Point	int
}

func Get(urldb string) []BehaviorPoint {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
	select grp, date_in, time1, time2, time3, point1, point2, point3 
	from checkinprofile
	order by grp, date_in
	`
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()


	return nil
}
