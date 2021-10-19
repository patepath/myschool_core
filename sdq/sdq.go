package sdq

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Msg struct {
	Result bool
}

type Student struct {
	Code     string
	Fullname string
	Grade    string
	Room     string
	No       int
}

type SDQ struct {
	Code      string
	Type      int
	Topic1_1  int
	Topic1_2  int
	Topic1_3  int
	Topic1_4  int
	Topic1_5  int
	Topic1_6  int
	Topic1_7  int
	Topic1_8  int
	Topic1_9  int
	Topic1_10 int
	Topic1_11 int
	Topic1_12 int
	Topic1_13 int
	Topic1_14 int
	Topic1_15 int
	Topic1_16 int
	Topic1_17 int
	Topic1_18 int
	Topic1_19 int
	Topic1_20 int
	Topic1_21 int
	Topic1_22 int
	Topic1_23 int
	Topic1_24 int
	Topic1_25 int
	Topic2    int
	Topic2_1  int
	Topic2_2  int
	Topic2_3  int
	Topic2_4  int
	Topic2_5  int
	Topic2_6  int
	Topic2_7  int
	Topic3_0  int
	Topic3_1  int
	Topic3_2  int
	Topic3_3  int
	Topic3_4  int
	Topic3_5  int
	Topic3_6  int
	Topic3_7  int
	Topic3_8  int
	Topic3_9  int
	Topic3_10 int
	Topic3_11 int
	Topic3_12 int
	Topic3_13 int
	Topic3_14 int
	Topic3_15 int
	Topic3_16 int
	Topic3_17 int
	Topic3_18 int
	Topic3_19 int
	Topic3_20 int
	Topic3_21 int
	Topic3_22 int
	Topic3_23 int
	Topic3_24 int
	Topic3_25 int
	Topic3_26 int
	Topic3_27 int
	Topic3_28 int
	Topic3_29 int
	Topic3_30 int
	Topic3_31 int
	Topic3_32 int
	Topic3_33 int
	Topic3_34 int
	Topic3_35 int
	Topic3_36 int
	Topic3_37 int
	Topic3_38 int
	Topic3_39 int
	Topic3_40 int
	Topic3_41 int
	Topic3_42 int
	Topic3_43 int
	Topic3_44 int
	Topic3_45 int
	Topic3_46 int
	Topic3_47 int
	Topic3_48 int
	Topic3_49 int
	Topic3_50 int
	Topic3_51 int
	Topic3_52 int
}

func GetStudentByCode(urldb string, code string) Student {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select
			code,
			concat(first_name, ' ', last_name) as fullname,
			C.name as grade,
			B.name as room,
			no
		from students as A
		join classroom as B on B.ref = A.room
		join grade as C on C.ref = B.grade_ref
		where A.code=? limit 1;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	var student Student

	rows, err := sttn.Query(code)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(
			&student.Code,
			&student.Fullname,
			&student.Grade,
			&student.Room,
			&student.No,
		)

		if err != nil {
			panic(err)
		}
	}

	return student
}

func GetByCode(urldb string, code string, t string) SDQ {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select
			code,
			type,
			ifnull(topic1_1, 0),  
			ifnull(topic1_2, 0),  
			ifnull(topic1_3, 0),  
			ifnull(topic1_4, 0),  
			ifnull(topic1_5, 0),  
			ifnull(topic1_6, 0),  
			ifnull(topic1_7, 0),  
			ifnull(topic1_8, 0),  
			ifnull(topic1_9, 0),  
			ifnull(topic1_10, 0),  
			ifnull(topic1_11, 0),  
			ifnull(topic1_12, 0),  
			ifnull(topic1_13, 0),  
			ifnull(topic1_14, 0),  
			ifnull(topic1_15, 0),  
			ifnull(topic1_16, 0),  
			ifnull(topic1_17, 0),  
			ifnull(topic1_18, 0),  
			ifnull(topic1_19, 0),  
			ifnull(topic1_20, 0),  
			ifnull(topic1_21, 0),  
			ifnull(topic1_22, 0),  
			ifnull(topic1_23, 0),  
			ifnull(topic1_24, 0),  
			ifnull(topic1_25, 0),  
			ifnull(topic2, 0),
			ifnull(topic2_1, 0),
			ifnull(topic2_2, 0),
			ifnull(topic2_3, 0),
			ifnull(topic2_4, 0),
			ifnull(topic2_5, 0),
			ifnull(topic2_6, 0),
			ifnull(topic2_7, 0),
			ifnull(topic3_1, 0),
			ifnull(topic3_2, 0),
			ifnull(topic3_3, 0),
			ifnull(topic3_4, 0),
			ifnull(topic3_5, 0),
			ifnull(topic3_6, 0),
			ifnull(topic3_7, 0),
			ifnull(topic3_8, 0),
			ifnull(topic3_9, 0),
			ifnull(topic3_10, 0),
			ifnull(topic3_11, 0),
			ifnull(topic3_12, 0),
			ifnull(topic3_13, 0),
			ifnull(topic3_14, 0),
			ifnull(topic3_15, 0),
			ifnull(topic3_16, 0),
			ifnull(topic3_17, 0),
			ifnull(topic3_18, 0),
			ifnull(topic3_19, 0),
			ifnull(topic3_20, 0),
			ifnull(topic3_21, 0),
			ifnull(topic3_22, 0),
			ifnull(topic3_23, 0),
			ifnull(topic3_24, 0),
			ifnull(topic3_25, 0),
			ifnull(topic3_26, 0),
			ifnull(topic3_27, 0),
			ifnull(topic3_28, 0),
			ifnull(topic3_29, 0),
			ifnull(topic3_30, 0),
			ifnull(topic3_31, 0),
			ifnull(topic3_32, 0),
			ifnull(topic3_33, 0),
			ifnull(topic3_34, 0),
			ifnull(topic3_35, 0),
			ifnull(topic3_36, 0),
			ifnull(topic3_37, 0),
			ifnull(topic3_38, 0),
			ifnull(topic3_39, 0),
			ifnull(topic3_40, 0),
			ifnull(topic3_41, 0),
			ifnull(topic3_42, 0),
			ifnull(topic3_43, 0),
			ifnull(topic3_44, 0),
			ifnull(topic3_45, 0),
			ifnull(topic3_46, 0),
			ifnull(topic3_47, 0),
			ifnull(topic3_48, 0),
			ifnull(topic3_49, 0),
			ifnull(topic3_50, 0),
			ifnull(topic3_51, 0),
			ifnull(topic3_52, 0)
		from sdq
		where code=? and type=?
		limit 1;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	row, err := sttn.Query(code, t)

	if err != nil {
		panic(err)
	}

	defer row.Close()

	var sdq SDQ

	if row.Next() {

		err := row.Scan(
			&sdq.Code,
			&sdq.Type,
			&sdq.Topic1_1,
			&sdq.Topic1_2,
			&sdq.Topic1_3,
			&sdq.Topic1_4,
			&sdq.Topic1_5,
			&sdq.Topic1_6,
			&sdq.Topic1_7,
			&sdq.Topic1_8,
			&sdq.Topic1_9,
			&sdq.Topic1_10,
			&sdq.Topic1_11,
			&sdq.Topic1_12,
			&sdq.Topic1_13,
			&sdq.Topic1_14,
			&sdq.Topic1_15,
			&sdq.Topic1_16,
			&sdq.Topic1_17,
			&sdq.Topic1_18,
			&sdq.Topic1_19,
			&sdq.Topic1_20,
			&sdq.Topic1_21,
			&sdq.Topic1_22,
			&sdq.Topic1_23,
			&sdq.Topic1_24,
			&sdq.Topic1_25,
			&sdq.Topic2,
			&sdq.Topic2_1,
			&sdq.Topic2_2,
			&sdq.Topic2_3,
			&sdq.Topic2_4,
			&sdq.Topic2_5,
			&sdq.Topic2_6,
			&sdq.Topic2_7,
			&sdq.Topic3_1,
			&sdq.Topic3_2,
			&sdq.Topic3_3,
			&sdq.Topic3_4,
			&sdq.Topic3_5,
			&sdq.Topic3_6,
			&sdq.Topic3_7,
			&sdq.Topic3_8,
			&sdq.Topic3_9,
			&sdq.Topic3_10,
			&sdq.Topic3_11,
			&sdq.Topic3_12,
			&sdq.Topic3_13,
			&sdq.Topic3_14,
			&sdq.Topic3_15,
			&sdq.Topic3_16,
			&sdq.Topic3_17,
			&sdq.Topic3_18,
			&sdq.Topic3_19,
			&sdq.Topic3_20,
			&sdq.Topic3_21,
			&sdq.Topic3_22,
			&sdq.Topic3_23,
			&sdq.Topic3_24,
			&sdq.Topic3_25,
			&sdq.Topic3_26,
			&sdq.Topic3_27,
			&sdq.Topic3_28,
			&sdq.Topic3_29,
			&sdq.Topic3_30,
			&sdq.Topic3_31,
			&sdq.Topic3_32,
			&sdq.Topic3_33,
			&sdq.Topic3_34,
			&sdq.Topic3_35,
			&sdq.Topic3_36,
			&sdq.Topic3_37,
			&sdq.Topic3_38,
			&sdq.Topic3_39,
			&sdq.Topic3_40,
			&sdq.Topic3_41,
			&sdq.Topic3_42,
			&sdq.Topic3_43,
			&sdq.Topic3_44,
			&sdq.Topic3_45,
			&sdq.Topic3_46,
			&sdq.Topic3_47,
			&sdq.Topic3_48,
			&sdq.Topic3_49,
			&sdq.Topic3_50,
			&sdq.Topic3_51,
			&sdq.Topic3_52,
		)

		if err != nil {
			panic(err)
		}

	}

	return sdq
}

func Save(urldb string, sdq SDQ) Msg {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		select code
		from sdq
		where code=? and type=?;
	`

	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	row, err := sttn.Query(sdq.Code, sdq.Type)

	if err != nil {
		panic(err)
	}

	defer row.Close()

	if row.Next() {
		return update(urldb, sdq)
	}

	return insert(urldb, sdq)
}

func insert(urldb string, sdq SDQ) Msg {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		insert into sdq (
			code,
			type,
			topic1_1,  
			topic1_2,  
			topic1_3,  
			topic1_4,  
			topic1_5,  
			topic1_6,  
			topic1_7,  
			topic1_8,  
			topic1_9,  
			topic1_10,  
			topic1_11,  
			topic1_12,  
			topic1_13,  
			topic1_14,  
			topic1_15,  
			topic1_16,  
			topic1_17,  
			topic1_18,  
			topic1_19,  
			topic1_20,  
			topic1_21,  
			topic1_22,  
			topic1_23,  
			topic1_24,  
			topic1_25,  
			topic2,
			topic2_1,
			topic2_2,
			topic2_3,
			topic2_4,
			topic2_5,
			topic2_6,
			topic2_7,
			topic3_1,
			topic3_2,
			topic3_3,
			topic3_4,
			topic3_5,
			topic3_6,
			topic3_7,
			topic3_8,
			topic3_9,
			topic3_10,
			topic3_11,
			topic3_12,
			topic3_13,
			topic3_14,
			topic3_15,
			topic3_16,
			topic3_17,
			topic3_18,
			topic3_19,
			topic3_20,
			topic3_21,
			topic3_22,
			topic3_23,
			topic3_24,
			topic3_25,
			topic3_26,
			topic3_27,
			topic3_28,
			topic3_29,
			topic3_30,
			topic3_31,
			topic3_32,
			topic3_33,
			topic3_34,
			topic3_35,
			topic3_36,
			topic3_37,
			topic3_38,
			topic3_39,
			topic3_40,
			topic3_41,
			topic3_42,
			topic3_43,
			topic3_44,
			topic3_45,
			topic3_46,
			topic3_47,
			topic3_48,
			topic3_49,
			topic3_50,
			topic3_51,
			topic3_52
		) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);
	`
	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		sdq.Code,
		sdq.Type,
		sdq.Topic1_1,
		sdq.Topic1_2,
		sdq.Topic1_3,
		sdq.Topic1_4,
		sdq.Topic1_5,
		sdq.Topic1_6,
		sdq.Topic1_7,
		sdq.Topic1_8,
		sdq.Topic1_9,
		sdq.Topic1_10,
		sdq.Topic1_11,
		sdq.Topic1_12,
		sdq.Topic1_13,
		sdq.Topic1_14,
		sdq.Topic1_15,
		sdq.Topic1_16,
		sdq.Topic1_17,
		sdq.Topic1_18,
		sdq.Topic1_19,
		sdq.Topic1_20,
		sdq.Topic1_21,
		sdq.Topic1_22,
		sdq.Topic1_23,
		sdq.Topic1_24,
		sdq.Topic1_25,
		sdq.Topic2,
		sdq.Topic2_1,
		sdq.Topic2_2,
		sdq.Topic2_3,
		sdq.Topic2_4,
		sdq.Topic2_5,
		sdq.Topic2_6,
		sdq.Topic2_7,
		sdq.Topic3_1,
		sdq.Topic3_2,
		sdq.Topic3_3,
		sdq.Topic3_4,
		sdq.Topic3_5,
		sdq.Topic3_6,
		sdq.Topic3_7,
		sdq.Topic3_8,
		sdq.Topic3_9,
		sdq.Topic3_10,
		sdq.Topic3_11,
		sdq.Topic3_12,
		sdq.Topic3_13,
		sdq.Topic3_14,
		sdq.Topic3_15,
		sdq.Topic3_16,
		sdq.Topic3_17,
		sdq.Topic3_18,
		sdq.Topic3_19,
		sdq.Topic3_20,
		sdq.Topic3_21,
		sdq.Topic3_22,
		sdq.Topic3_23,
		sdq.Topic3_24,
		sdq.Topic3_25,
		sdq.Topic3_26,
		sdq.Topic3_27,
		sdq.Topic3_28,
		sdq.Topic3_29,
		sdq.Topic3_30,
		sdq.Topic3_31,
		sdq.Topic3_32,
		sdq.Topic3_33,
		sdq.Topic3_34,
		sdq.Topic3_35,
		sdq.Topic3_36,
		sdq.Topic3_37,
		sdq.Topic3_38,
		sdq.Topic3_39,
		sdq.Topic3_40,
		sdq.Topic3_41,
		sdq.Topic3_42,
		sdq.Topic3_43,
		sdq.Topic3_44,
		sdq.Topic3_45,
		sdq.Topic3_46,
		sdq.Topic3_47,
		sdq.Topic3_48,
		sdq.Topic3_49,
		sdq.Topic3_50,
		sdq.Topic3_51,
		sdq.Topic3_52,
	)

	if err != nil {
		panic(err)
	}

	var msg Msg
	msg.Result = true

	return msg
}

func update(urldb string, sdq SDQ) Msg {

	db, err := sql.Open("mysql", urldb)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql := `
		update sdq set
			topic1_1=?,  
			topic1_2=?,  
			topic1_3=?,  
			topic1_4=?,  
			topic1_5=?,  
			topic1_6=?,  
			topic1_7=?,  
			topic1_8=?,  
			topic1_9=?,  
			topic1_10=?,  
			topic1_11=?,  
			topic1_12=?,  
			topic1_13=?,  
			topic1_14=?,  
			topic1_15=?,  
			topic1_16=?,  
			topic1_17=?,  
			topic1_18=?,  
			topic1_19=?,  
			topic1_20=?,  
			topic1_21=?,  
			topic1_22=?,  
			topic1_23=?,  
			topic1_24=?,  
			topic1_25=?,  
			topic2=?,
			topic2_1=?,
			topic2_2=?,
			topic2_3=?,
			topic2_4=?,
			topic2_5=?,
			topic2_6=?,
			topic2_7=?,
			topic3_1=?,
			topic3_2=?,
			topic3_3=?,
			topic3_4=?,
			topic3_5=?,
			topic3_6=?,
			topic3_7=?,
			topic3_8=?,
			topic3_9=?,
			topic3_10=?,
			topic3_11=?,
			topic3_12=?,
			topic3_13=?,
			topic3_14=?,
			topic3_15=?,
			topic3_16=?,
			topic3_17=?,
			topic3_18=?,
			topic3_19=?,
			topic3_20=?,
			topic3_21=?,
			topic3_22=?,
			topic3_23=?,
			topic3_24=?,
			topic3_25=?,
			topic3_26=?,
			topic3_27=?,
			topic3_28=?,
			topic3_29=?,
			topic3_30=?,
			topic3_31=?,
			topic3_32=?,
			topic3_33=?,
			topic3_34=?,
			topic3_35=?,
			topic3_36=?,
			topic3_37=?,
			topic3_38=?,
			topic3_39=?,
			topic3_40=?,
			topic3_41=?,
			topic3_42=?,
			topic3_43=?,
			topic3_44=?,
			topic3_45=?,
			topic3_46=?,
			topic3_47=?,
			topic3_48=?,
			topic3_49=?,
			topic3_50=?,
			topic3_51=?,
			topic3_52=?
		where code=? and type=?;
	`
	sttn, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer sttn.Close()

	_, err = sttn.Exec(
		sdq.Topic1_1,
		sdq.Topic1_2,
		sdq.Topic1_3,
		sdq.Topic1_4,
		sdq.Topic1_5,
		sdq.Topic1_6,
		sdq.Topic1_7,
		sdq.Topic1_8,
		sdq.Topic1_9,
		sdq.Topic1_10,
		sdq.Topic1_11,
		sdq.Topic1_12,
		sdq.Topic1_13,
		sdq.Topic1_14,
		sdq.Topic1_15,
		sdq.Topic1_16,
		sdq.Topic1_17,
		sdq.Topic1_18,
		sdq.Topic1_19,
		sdq.Topic1_20,
		sdq.Topic1_21,
		sdq.Topic1_22,
		sdq.Topic1_23,
		sdq.Topic1_24,
		sdq.Topic1_25,
		sdq.Topic2,
		sdq.Topic2_1,
		sdq.Topic2_2,
		sdq.Topic2_3,
		sdq.Topic2_4,
		sdq.Topic2_5,
		sdq.Topic2_6,
		sdq.Topic2_7,
		sdq.Topic3_1,
		sdq.Topic3_2,
		sdq.Topic3_3,
		sdq.Topic3_4,
		sdq.Topic3_5,
		sdq.Topic3_6,
		sdq.Topic3_7,
		sdq.Topic3_8,
		sdq.Topic3_9,
		sdq.Topic3_10,
		sdq.Topic3_11,
		sdq.Topic3_12,
		sdq.Topic3_13,
		sdq.Topic3_14,
		sdq.Topic3_15,
		sdq.Topic3_16,
		sdq.Topic3_17,
		sdq.Topic3_18,
		sdq.Topic3_19,
		sdq.Topic3_20,
		sdq.Topic3_21,
		sdq.Topic3_22,
		sdq.Topic3_23,
		sdq.Topic3_24,
		sdq.Topic3_25,
		sdq.Topic3_26,
		sdq.Topic3_27,
		sdq.Topic3_28,
		sdq.Topic3_29,
		sdq.Topic3_30,
		sdq.Topic3_31,
		sdq.Topic3_32,
		sdq.Topic3_33,
		sdq.Topic3_34,
		sdq.Topic3_35,
		sdq.Topic3_36,
		sdq.Topic3_37,
		sdq.Topic3_38,
		sdq.Topic3_39,
		sdq.Topic3_40,
		sdq.Topic3_41,
		sdq.Topic3_42,
		sdq.Topic3_43,
		sdq.Topic3_44,
		sdq.Topic3_45,
		sdq.Topic3_46,
		sdq.Topic3_47,
		sdq.Topic3_48,
		sdq.Topic3_49,
		sdq.Topic3_50,
		sdq.Topic3_51,
		sdq.Topic3_52,
		sdq.Code,
		sdq.Type,
	)

	if err != nil {
		panic(err)
	}

	var msg Msg
	msg.Result = true

	return msg
}
