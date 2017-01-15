package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"indragolangweb/config"
	"indragolangweb/apps/models"
	"database/sql"
)

var err error


func SignupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		config.GetPath()
		http.ServeFile(res, req, config.PathView+"/signup.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")


	err = models.GetUser(username)

	switch {
		case err == sql.ErrNoRows:
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				http.Error(res, "Server error, unable to create your account.", 500)
				return
			}

			err = models.InsertUser(username,hashedPassword)

			if err != nil {
				http.Error(res, "Server error, unable to create your account.", 500)
				return
			}

			res.Write([]byte("User created!"))
			return
		case err != nil:
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		default:
			http.Redirect(res, req, "/", 301)
	}
}

func LoginPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		config.GetPath()
		http.ServeFile(res, req, config.PathView+"/login.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var databaseUsername string
	var databasePassword string



	var userData models.UserData = models.GetUserLogin(username,databaseUsername,databasePassword)


	if userData.Err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.DatabasePassword), []byte(password))
	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	res.Write([]byte("Hello " + userData.DatabaseUser))

}

func UserList(res http.ResponseWriter, req *http.Request){

	config.GetPath()
	http.ServeFile(res, req, config.PathView+"/list.html")
	return

}

/**
 * API Get Data User
 */

func UserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var data,_  = models.GetuserList()
		w.Write([]byte(data))
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}