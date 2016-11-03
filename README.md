# indragolangweb

Bagaimana menggunakan coding indragolangweb:

1. Pastikan komputer anda sudah terinstall Golang
2. Settingan GOPATH dan 3 folder utama : pkg, src, bin sudah tersedia
3. Pastikan kodingan ini berada di folder src
4. Carilah Package mysql,dan crypt, berikut ini perintahnya :

    go get github.com/go-sql-driver/mysql
    go get golang.org/x/crypto/bcrypt

5. Jalankan perintah : go build indragolangweb/
6. Jika anda berhasil, maka anda akan menemukan halaman web dengan cara membuka browser ketik : http://localhost:8181
    dengan tampilan Homepage. Ada 2 link : Login dan signup.
7. Jangan lupa dump database golang.sql ke dalam engine mysql anda.
8. Selamat mencoba dan selamat belajar!

