package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //nameless import karena tidak menggunakan data apa-apa dari mysql ini
)
func InitDb() *sql.DB {
	dsn := "root@tcp(phpmyadmin.test)/golang-todos"
	db, err :=sql.Open("mysql", dsn) //data driver mysql
	if err != nil {
		panic(err)
	}

	return db
}