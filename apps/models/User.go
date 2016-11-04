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
	var err = config.Debe.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

	return err

}

/**
* Get User
 */
func GetUserLogin(username,databaseUsername,databasePassword string) UserData{

	config.ConnectDB()
	var err = config.Debe.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)

	u := UserData{err, databaseUsername,databasePassword}

	return u

}

/**
 *Insert User
 */
func InsertUser(username string,hashedPassword []byte) error{

	config.ConnectDB()
	var _, err = config.Debe.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, hashedPassword)

	return err
}



