package department

import (
	//"encoding/json"
	//"fmt"
	//"net/http"
	//"strconv"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Department struct {
	Ref  int
	Code string
	Name string
}

func Get(urldb string) []Department {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select ref, code, name
		from department order by name;
	`
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var depts []Department
	var dept Department

	for rows.Next() {
		err := rows.Scan(&dept.Ref, &dept.Code, &dept.Name)

		if err != nil {
			panic(err)
		}

		depts = append(depts, dept)
	}

	return depts
}

func insert(urldb string, dept Department) {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		insert into department(code, name) values(?,?)
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(dept.Code, dept.Name)

	if err != nil {
		panic(err)
	}

}

func update(urldb string, dept Department) {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		update department set code=?, name=? where ref=?;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(dept.Code, dept.Name, dept.Ref)

	if err != nil {
		panic(err)
	}
}

func Save(urldb string, dept Department) []Department {

	if dept.Ref == 0 {
		insert(urldb, dept)

	} else {
		update(urldb, dept)
	}

	return Get(urldb)
}

func Delete(urldb string, ref string) []Department {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		delete from  department  where ref=?;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(ref)

	if err != nil {
		panic(err)
	}

	return Get(urldb)
}
