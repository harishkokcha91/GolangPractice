package dbutils

import (
	"database/sql"
)

func DbConn() *sql.DB {

	db, err := sql.Open("mysql", "sammy:Password@123@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	return db
}
