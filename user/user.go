package user

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Payload struct {
	A string
	B string
}

type User struct {
	Ref      int
	IdCard   string
	Name     string
	Password string
	Fullname string
	Role     string
}

func Login(urldb string, payload Payload) User {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select name, full_name, role
		from user
		where md5(name)=? and password=?
		limit 1;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	var user User

	rows, err := sttn.Query(payload.A, payload.B)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {

		err := rows.Scan(
			&user.Name,
			&user.Fullname,
			&user.Role,
		)

		if err != nil {
			panic(err)
		}

	}

	return user
}

func GetAll(urldb string) []User {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select
			ref, 
			ifnull(idcard, ''), 
			name, 
			full_name, 
			role
		from user
		order by name;
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var users []User
	var user User

	for rows.Next() {

		err := rows.Scan(
			&user.Ref,
			&user.IdCard,
			&user.Name,
			&user.Fullname,
			&user.Role,
		)

		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	return users
}

func GetByRef(urldb string, ref string) User {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select
			ref, 
			ifnull(idcard, ''), 
			name, 
			full_name, 
			role
		from user
		where ref=` + ref + `
		limit 1;
	`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var user User

	if rows.Next() {

		err := rows.Scan(
			&user.Ref,
			&user.IdCard,
			&user.Name,
			&user.Fullname,
			&user.Role,
		)

		if err != nil {
			panic(err)
		}

	}

	return user
}

func Save(urldb string, user User) []User {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(nil)
	}

	defer db.Close()

	if user.Ref == 0 {

		sql := `
			insert into 
			user (idcard, name, full_name, password, role) 
			values (?,?,?,md5(?),?);
		`
		sttn, err := db.Prepare(sql)
		if err != nil {
			panic(err)
		}

		defer sttn.Close()

		_, err = sttn.Exec(user.IdCard, user.Name, user.Fullname, user.Password, user.Role)

		if err != nil {
			panic(err)
		}

	} else {

		if user.Password == "" {

			sql := `
				update user set
					idcard=?, 
					name=?,
					full_name=?,
					role=?
				where ref=?;
			`
			sttn, err := db.Prepare(sql)
			if err != nil {
				panic(err)
			}

			defer sttn.Close()

			_, err = sttn.Exec(user.IdCard, user.Name, user.Fullname, user.Role, user.Ref)

			if err != nil {
				panic(err)
			}

		} else {

			sql := `
				update user set
					idcard=?,
					name=?,
					full_name=?,
					role=?,
					password=md5(?)
				where ref=?;
			`
			sttn, err := db.Prepare(sql)
			if err != nil {
				panic(err)
			}

			defer sttn.Close()

			_, err = sttn.Exec(user.IdCard, user.Name, user.Fullname, user.Role, user.Password, user.Ref)

			if err != nil {
				panic(err)
			}
		}

	}

	return GetAll(urldb)
}

func Remove(urldb string, ref string) []User {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(nil)
	}

	defer db.Close()

	sql := `
		delete from user where ref=?;
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

	return GetAll(urldb)
}
