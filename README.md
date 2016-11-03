# indragolangweb

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

Bagaimana menggunakan coding indragolangweb:

1. Pastikan komputer anda sudah terinstall Golang
2. Settingan GOPATH dan 3 folder utama : pkg, src, bin sudah tersedia
3. Pastikan kodingan ini berada di folder src
4. Carilah File config/path.go , silahkan ubah path sesuai dengan absolute path komputer anda
5. Carilah Package mysql,dan crypt, berikut ini perintahnya :

    go get github.com/go-sql-driver/mysql
    go get golang.org/x/crypto/bcrypt

6. Jalankan perintah : go build indragolangweb/
7. Jika anda berhasil, maka anda akan menemukan halaman web dengan cara membuka browser ketik : http://localhost:8181
    dengan tampilan Homepage. Ada 2 link : Login dan signup.
8. Jangan lupa dump database golang.sql ke dalam engine mysql anda.
9. Selamat mencoba dan selamat belajar!

