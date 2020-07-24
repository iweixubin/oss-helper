package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func tDb() {
	dataSource := "root:123456@tcp(192.168.86.101:3306)/storage_service?charset=ascii&parseTime=False&loc=Local&multiStatements=true"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		println(err)
	}

	defer db.Close()

	rows, err := db.Query("show charset")
	if err != nil {
		println(err)
	}

	type charset struct {
		Charset     string
		Description string
	}

	for rows.Next() {
		c := charset{}
		rows.Scan(&c.Charset, &c.Description)
	}

	_ = db
	_ = err
}
