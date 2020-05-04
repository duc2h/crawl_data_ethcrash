package database

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "hoangduc1998"
	dbName := "crawl_data_ethcrash"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB: connect success")
	}
	return db
}
