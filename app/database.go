package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"latihan-restful-api/helper"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/belajar_golang_restful_api?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
