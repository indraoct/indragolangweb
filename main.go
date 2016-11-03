/**
 * @author Indra Octama
 * @Created Date 3 November 2016
 * Main Web With Golang
 */
package main

import (
	"net/http"
	_ "indragolangweb/apps/controllers"
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

	http.HandleFunc("/signup", controllers.SignupPage)
	http.HandleFunc("/login", controllers.LoginPage)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8181", nil)
}
