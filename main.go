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
)

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "/Users/mmdc/go-work/src/indragolangweb/apps/views/index.html")
}

func main() {


	http.HandleFunc("/signup", controllers.SignupPage)
	http.HandleFunc("/login", controllers.LoginPage)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8181", nil)
}
