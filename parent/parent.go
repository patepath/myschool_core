package parent

//package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Parent struct {
	Ref          int
	IdCard       string
	Title        int
	FirstName    string
	LastName     string
	Gender       int
	Addr1        string
	Addr2        string
	Addr3        string
	ProvinceCode int
	ProvinceName string
	ZipCode      string
	Phone        string
	LineId       string
	Facebook     string
	LineUid      string
	//	OccupationCode: number;
	//	OccupationName: string;
	//	Position: string;
	//	CompanyName: string;
	//	CompanyAddr1: string;
	//	CompanyAddr2: string;
	//	CompanyAddr3: string;
	//	CompanyProvinceCode: number;
	//	CompanyProvinceName: string;
	//	CompanyZipCode: string;
	//	CompanyPhone: string;
	//	Salary: number
}

func Get(urldb string) []Parent {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
	select 
		ref, 
		if(isnull(idcard), '', idcard),
		if(isnull(title), 0, title),
		if(isnull(first_name), '', first_name),
		if(isnull(last_name), '', last_name),
		if(isnull(gender), 0, gender),
		if(isnull(addr1), '', addr1),
		if(isnull(addr2), '', addr2),
		if(isnull(addr3), '', addr3),
		if(isnull(province_code), 0, province_code),
		if(isnull(province_name), '', province_name),
		if(isnull(zipcode), '', zipcode),
		if(isnull(phone), '', phone),
		if(isnull(line), '', line),
		if(isnull(facebook), '', facebook),   
		if(isnull(lineuid), '', lineuid)   
	from parents
	order by first_name; 
	`
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var parents []Parent
	var parent Parent

	for rows.Next() {
		err := rows.Scan(
			&parent.Ref,
			&parent.IdCard,
			&parent.Title,
			&parent.FirstName,
			&parent.LastName,
			&parent.Gender,
			&parent.Addr1,
			&parent.Addr2,
			&parent.Addr3,
			&parent.ProvinceCode,
			&parent.ProvinceName,
			&parent.ZipCode,
			&parent.Phone,
			&parent.LineId,
			&parent.Facebook,
			&parent.LineUid,
		)

		if err != nil {
			panic(err)
		}

		parents = append(parents, parent)
	}

	return parents
}

func GetByRef(urldb string, ref string) Parent {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			ref, 
			if(isnull(idcard), '', idcard),
			if(isnull(title), 0, title),
			if(isnull(first_name), '', first_name),
			if(isnull(last_name), '', last_name),
			if(isnull(gender), 0, gender),
			if(isnull(addr1), '', addr1),
			if(isnull(addr2), '', addr2),
			if(isnull(addr3), '', addr3),
			if(isnull(province_code), 0, province_code),
			if(isnull(province_name), '', province_name),
			if(isnull(zipcode), '', zipcode),
			if(isnull(phone), '', phone),
			if(isnull(line), '', line),
			if(isnull(facebook), '', facebook),   
			if(isnull(lineuid), '', lineuid)   
		from parents
		where ref=` + ref

	row, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer row.Close()

	var parent Parent

	if row.Next() {
		err := row.Scan(
			&parent.Ref,
			&parent.IdCard,
			&parent.Title,
			&parent.FirstName,
			&parent.LastName,
			&parent.Gender,
			&parent.Addr1,
			&parent.Addr2,
			&parent.Addr3,
			&parent.ProvinceCode,
			&parent.ProvinceName,
			&parent.ZipCode,
			&parent.Phone,
			&parent.LineId,
			&parent.Facebook,
			&parent.LineUid,
		)

		if err != nil {
			panic(err)
		}

	}

	return parent
}

func GetByStudent(urldb string, studentref string) []Parent {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select 
			ref,  
			title,
			first_name, 
			last_name, 
			phone 
		from student_parents 
		left join parents on student_parents.parent_ref = parents.ref 
		where ref is not null and student_ref=?;`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	rows, err := sttn.Query(studentref)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var parents []Parent
	var parent Parent

	for rows.Next() {
		err := rows.Scan(
			&parent.Ref,
			&parent.Title,
			&parent.FirstName,
			&parent.LastName,
			&parent.Phone,
		)

		if err != nil {
			panic(err)
		}

		parents = append(parents, parent)
	}

	return parents
}

func Save(urldb string, parent Parent) []Parent {

	if parent.Ref == 0 {
		insert(urldb, parent)
	} else {
		update(urldb, parent)
	}

	return Get(urldb)
}

func insert(urldb string, parent Parent) {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		insert into parents ( 
			idcard,
			title,
			first_name,
			last_name,
			gender,
			addr1,
			addr2,
			addr3,
			province_code,
			province_name,
			zipcode,
			phone,
			line,
			facebook,
			lineuid
		) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)	
	`
	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	_, err = sttn.Exec(
		parent.IdCard,
		parent.Title,
		parent.FirstName,
		parent.LastName,
		parent.Gender,
		parent.Addr1,
		parent.Addr2,
		parent.Addr3,
		parent.ProvinceCode,
		parent.ProvinceName,
		parent.ZipCode,
		parent.Phone,
		parent.LineId,
		parent.Facebook,
		parent.LineUid,
	)

	if err != nil {
		panic(err)
	}

}

func update(urldb string, parent Parent) {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		update parents set 
			idcard=?,
			title=?,
			first_name=?,
			last_name=?,
			gender=?,
			addr1=?,
			addr2=?,
			addr3=?,
			province_code=?,
			province_name=?,
			zipcode=?,
			phone=?,
			line=?,
			facebook=?,
			lineuid=?
		where ref=?`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	_, err = sttn.Exec(
		parent.IdCard,
		parent.Title,
		parent.FirstName,
		parent.LastName,
		parent.Gender,
		parent.Addr1,
		parent.Addr2,
		parent.Addr3,
		parent.ProvinceCode,
		parent.ProvinceName,
		parent.ZipCode,
		parent.Phone,
		parent.LineId,
		parent.Facebook,
		parent.LineUid,
		parent.Ref,
	)

	if err != nil {
		panic(err)
	}

}

func Remove(urldb string, ref string) []Parent {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		delete from parents
		where ref=` + ref

	_, err = db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return Get(urldb)
}

func main() {
	const urldb = "admin:35.103232@tcp(localhost:3306)/school"
	//GetByStudent(urldb, "1")

	var parent = Parent{
		Addr1:        "",
		Addr2:        "",
		Addr3:        "",
		Facebook:     "",
		FirstName:    "cccc",
		Gender:       0,
		IdCard:       "1452798451456",
		LastName:     "cccc",
		LineId:       "nattaya.pun",
		LineUid:      "",
		Phone:        "3333",
		ProvinceCode: 0,
		ProvinceName: "",
		Ref:          1,
		Title:        1,
		ZipCode:      "",
	}

	Save(urldb, parent)
}
