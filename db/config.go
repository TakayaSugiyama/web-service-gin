package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func configTestCliant() string {
	if value, ok := os.LookupEnv("TEST_CLIENT"); ok && value == "CI" {
		return "ginuser:password@tcp(127.0.0.1:3306)/gin-service?parseTime=true"
	}
	if value, ok := os.LookupEnv("DSN"); ok {
		return value
	}
	return "ginuser:password@tcp(gin_service_db:3306)/gin-service?parseTime=true"
}

func ConnectDB() *sql.DB {
	var dnsNmae string = configTestCliant()
	db, err := sql.Open("mysql", dnsNmae)
	if err != nil {
		panic(err.Error())
	}

	return db
}
