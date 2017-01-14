/**
 * Updated Date : 14 January 2017
 * Make Function close database
 */
package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func ConnectDB(){

	var err error
	Db, err =  sql.Open("mysql", "root:blink182@/golang")
	if err != nil {
		panic(err.Error())
	}
	if err = Db.Ping(); err != nil {
		panic(err.Error())
	}

}

func CloseDB(){

	Db.Close()

}
