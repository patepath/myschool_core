package province

import (
	//"encoding/json"
	//"fmt"
	//"net/http"
	//"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Province struct {
	Ref int
	OrderNo int
	Name string
}

func Get(urldb string) []Province {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `select ref, rank , name from province order by rank`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var provinces []Province
	var province Province

	for rows.Next() {
		err := rows.Scan(
			&province.Ref,
			&province.OrderNo,
			&province.Name,
		)

		if err != nil {
			panic(err)
		}

		provinces = append(provinces, province)

	}

	return provinces
}

func Save(urldb string, province Province) []Province {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	if province.Ref == 0 {
		insertProvince(db, province)

	} else {
		updateProvince(db, province)
	}

	return Get(urldb)
}

func insertProvince(db *sql.DB, province Province) {
	sql := `
		insert into  province(  
			rank,
			name
		) values (?,?)`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		province.OrderNo,
		province.Name,
	)

	if err != nil {
		panic(err)
	}
}

func updateProvince(db *sql.DB, province Province) {
	sql := `
		update province set 
			rank=?,
			name=?
		where ref=?`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		province.OrderNo,
		province.Name,
		province.Ref,
	)

	if err != nil {
		panic(err)
	}
}
