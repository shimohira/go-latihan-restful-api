package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"latihan-restful-api/helper"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/latihan_database_migration?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db

	// migrate create -ext sql -dir db/migrations create_table_first
	// migrate create -ext sql -dir db/migrations create_table_second
	// migrate create -ext sql -dir db/migrations create_table_third
	// migrate create -ext sql -dir db/migrations sample_dirty_state

	// migrate -database "mysql://root:my-secret-pw@tcp(localhost:3306)/latihan_database_migration" -path db/migrations up
	// migrate -database "mysql://root:my-secret-pw@tcp(localhost:3306)/latihan_database_migration" -path db/migrations down

	// migrate -database "mysql://root:my-secret-pw@tcp(localhost:3306)/latihan_database_migration" -path db/migrations version
	// migrate -database "mysql://root:my-secret-pw@tcp(localhost:3306)/latihan_database_migration" -path db/migrations force 20221126175431
}
