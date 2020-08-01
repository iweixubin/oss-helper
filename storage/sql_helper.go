package storage

import (
	"fmt"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var _db *sql.DB

func init() {
	return
	dataSource := "root:123456@tcp(192.168.86.101:3306)/storage_service?charset=ascii&parseTime=False&loc=Local&multiStatements=true"
	_db, err := sql.Open("mysql", dataSource)
	if err != nil {
		println(err)
	}

	// 测试一下连接是否有问题
	err = _db.Ping()
	if err != nil {
		panic(err)
	}

}

func GetDb() *sql.DB {
	return _db
}

func tDb() {
	dataSource := "root:123456@tcp(192.168.86.101:3306)/storage_service?charset=ascii&parseTime=False&loc=Local&multiStatements=true"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		println(err)
	}

	db.SetConnMaxLifetime(time.Second * 12)

	db.Ping()

	//defer db.Close()

	rows, err := db.Query("show charset")
	if err != nil {
		println(err)
	}

	defer rows.Close()

	type charset struct {
		Charset     string
		Description string
		Default     string
		Maxlen      int
	}

	for rows.Next() {
		c := charset{}
		if e := rows.Scan(&c.Charset, &c.Description, &c.Default, &c.Maxlen); e != nil {
			println(e.Error())
		} else {
			fmt.Printf("%v ---- %v ---- %v ---- %v\n", c.Charset, c.Description, c.Default, c.Maxlen)
		}

	}

	db.Ping()

	_ = db
	_ = err
}
