package db

import (
	"database/sql"
	"project-2-rizwijaya/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()
	connection := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err = sql.Open("mysql", connection)
	if err != nil {
		panic("connection Error..")
	}

	db.SetMaxIdleConns(5)
	err = db.Ping()
	if err != nil {
		panic("DNS Invalid")
	}
	defer db.Close()
}

func CreateCon() *sql.DB {
	return db
}
