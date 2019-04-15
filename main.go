/**
 * @author Indra Octama
 * @Created Date 3 November 2016
 * Main Web With Golang
 */
package main

import (
	"net/http"
	"indragolangweb/apps/controllers"
	"indragolangweb/config"
)

var PathView string
func homePage(res http.ResponseWriter, req *http.Request) {

	PathView = config.PathView
	http.ServeFile(res, req, PathView+"/index.html")
}

func main() {
	config.GetPath()

	/**
	 * Web Display
	 */
	http.HandleFunc("/signup", controllers.SignupPage)
	http.HandleFunc("/login", controllers.LoginPage)
	http.HandleFunc("/user/list", controllers.UserList)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/emas",controllers.ResponseDummy)

	/**
	 * API
	 */
	http.HandleFunc("/api/userdata", controllers.UserData)

	/**
	 * Listening To Server
	 */
	http.ListenAndServe(":10099", nil)
}
