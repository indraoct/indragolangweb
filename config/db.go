package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Debe *sql.DB

func ConnectDB(){

	var err error
	Debe, err =  sql.Open("mysql", "root@/golang")
	if err != nil {
		panic(err.Error())
	}
	if err = Debe.Ping(); err != nil {
		panic(err.Error())
	}

}
