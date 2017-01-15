package models

import (
	"indragolangweb/config"
	"fmt"
	"encoding/json"
)

type UserData struct {
	Err error
	DatabaseUser  string
	DatabasePassword string
}

type UserDataList struct{
	id int
	username string
	password string

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
 * Get User List
 */

func GetuserList() (string, error) {
	config.ConnectDB()
	rows, err := config.Db.Query("select id, username, password from users")
	if err != nil {
		return "", err
	}
	defer rows.Close()
	config.CloseDB()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	fmt.Println(string(jsonData))
	return string(jsonData), nil
}

	/**
	* Get User Login
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



