package dal

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"oss-helper/internal/config"
)

// dal 就是三层架构中 数据访问层 的意思

var db *sql.DB

func InitDb(cfg config.Config) *sql.DB {
	var err error
	db, err = sql.Open(cfg.Database.Driver, cfg.Database.Source)
	if err != nil {
		println(err)
	}

	if cfg.Database.MaxOpen > 0 {
		db.SetMaxOpenConns(cfg.Database.MaxOpen)
	}

	if cfg.Database.MaxLifetime > 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(cfg.Database.MaxLifetime))
	}

	if cfg.Database.MaxIdle > 0 {
		db.SetMaxIdleConns(cfg.Database.MaxIdle)
	}

	// 测试一下连接是否有问题
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func createdAt() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
