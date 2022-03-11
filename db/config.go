package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "ginuser:password@tcp(gin_service_db:3306)/gin-service?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	return db
}
