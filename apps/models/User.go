package models

import (
	"indragolangweb/config"
)

type UserData struct {
	Err error
	DatabaseUser  string
	DatabasePassword string
}

/**
* Get User
 */
func GetUser(username string) error{

	var user string
	config.ConnectDB()
	var err = config.Db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)
	config.CloseDB()

	return err

}

/**
* Get User
 */
func GetUserLogin(username,databaseUsername,databasePassword string) UserData{

	config.ConnectDB()
	var err = config.Db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)
	config.CloseDB()
	u := UserData{err, databaseUsername,databasePassword}

	return u

}

/**
 *Insert User
 */
func InsertUser(username string,hashedPassword []byte) error{

	config.ConnectDB()
	var _, err = config.Db.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, hashedPassword)
	config.CloseDB()

	return err
}



