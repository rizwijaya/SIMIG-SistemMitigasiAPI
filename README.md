##### Nama : Muhammad Rizqi Wijaya
##### NRP : 05311940000014
##### Departemen : Teknologi Informasi
##### Project 2 Pemprograman Integratif

---
# SIMIG - Sistem Informasi dan Mitigasi Bencana
---
 ## Daftar Isi
- [SIMIG - Sistem Informasi dan Mitigasi Bencana](#simig---sistem-informasi-dan-mitigasi-bencana)
- [Deskripsi Sistem Informasi dan Mitigasi Bencana](#deskripsi-sistem-informasi-dan-mitigasi-bencana)
- [Fitur yang tersedia](#fitur-yang-tersedia)
- [Struktur Project](#struktur-project)
- [Panduan Instalasi](#panduan-instalasi-installation-guide)
- [Dokumentasi](#dokumentasi)
- [Contoh Penggunaan](#contoh-penggunaan)
---
## Deskripsi Sistem Informasi dan Mitigasi Bencana
SIMIG - Sistem Informasi dan Mitigasi Bencana merupakan sistem API yang berisi informasi seputar bencana yang sedang terjadi, pernah terjadi, dan dilengkapi dengan fitur pelaporan bencana dimana pengguna dapat melakukan pelaporan bencana yang terjadi didaerah nya. Selain itu pada sistem ini juga dilengkapi dengan fitur berita yang akan memberikan informasi bencana tersebut secara lengkap yang mana informasi dan berita yang disampaikan dalam situs ini akan dilakukan verifikasi terlebih dahulu sebelum akhirnya dipublish, untuk memastikan bahwa berita serta informasi tersebut merupakan informasi yang benar sesuai dengan fakta yang ada dilapangan.

---
## Fitur yang tersedia
Terdapat beberapa fitur yang tersedia seperti sebagai berikut
+ Mitigasi Bencana
    + Data Bencana yang sedang terjadi
    + Perbarui Data Bencana
    + History Bencana
+ Laporkan Bencana
    + Lapor Bencana
    + Data Pelapor
+ Informasi Bencana
    + Informasi Bencana
    + Informasi Bencana dengan Kode Berita 
    + Perbarui Berita Bencana
    + Tulis Berita Bencana
    + Hapus Informasi Bencana
---
## Struktur Project
Struktur folder yang digunakan dalam project ini adalah sebagai berikut:
```
    ├── assets
    ├── cmd
    ├── config
    │   ├── config.go
    │   └── config.json
    ├── controllers
    │   ├── bencana.go
    │   ├── berita.go
    │   ├── login.go
    │   ├── pelapor.go
    │   └── users.go
    ├── db
    │   └── db.go
    ├── docs
    │   ├── docs.go
    │   ├── swagger.json
    │   └── swagger.yaml
    ├── dokumentasi
    ├── helpers
    │   ├── pwd_helper.go
    │   └── TemplateRegistry.go
    ├── middlewares
    │   ├── CustomContext.go
    │   └── middlewares.go
    ├── models
    │   ├── bencana_m.go
    │   ├── berita_m.go
    │   ├── documentation.go
    │   ├── login_m.go
    │   ├── pelapor_m.go
    │   └── response.go
    ├── routes
    │   └── routes.go
    ├── vendor
    ├── views
    │   ├── alerts.html
    │   ├── beranda.html
    │   ├── dashboard.html
    │   ├── documentation.html
    │   ├── footer.html
    │   ├── header.html
    │   ├── login.html
    │   └── register.html
    └── Readme.md
```
Pada struktur folder dari project ini menggunakan konsep MVC (Models, Views, dan Controller) yaitu dengan memisahkan antara Models yang digunakan untuk menghubungkan sistem dengan database, kemudian controllers digunakan untuk logika pada sistem dan views yang berisi mengenai tampilan depan dari website ini. Dimana selanjutnya terdapat folder middlewares yang berisi middlewares yang digunakan dalam project ini, docs berisi file documentasi dari project dan helpers berisi fungsi untuk melakukan hashing password serta renders template.

---
## Panduan Instalasi (Installation Guide)
##### Instalasi awal
Langkah awal sebelum menggunakan project adalah melakukan instalasi yaitu sebagai berikut.
1. Download dan install Golang.
    + [Download Golang](https://golang.org/doc/install#download)
    + [Documentasi Install Golang](https://golang.org/doc/install)
2. Buka folder $GOPATH /src/github.com/
3. Pada folder tersebut dapatkan source code dari repository dengan mengetikkan perintah berikut pada terminal.
    ```bash
    git clone https://github.com/2021-IT-Pemrograman-Integratif/project-2-rizwijaya.git
    ```
    Atau bisa menggunakan perintah berikut.
    ```bash
    go get -u https://github.com/2021-IT-Pemrograman-Integratif/project-2-rizwijaya
    ```
##### Download Dependensi
Download dependensi package yang digunakan dengan menggunakan depedency manager.
    ```bash
    dep ensure
    ```
Apabila tidak memiliki depedency manager bisa menggunakan perintah go.
1. Lacak dependensi pada project.
    ```bash
    go mod init
    ```
2. Selanjutnya Lakukan perintah berikut.
    ```bash
    go mod vendor
    ```
3. Build project menggunakan perintah berikut.
    ```bash
    go build
    ```
##### Konfigurasikan Database
1. Langkah pertama dalam konfigurasi database adalah buat database baru.
2. Selanjutnya import database dari project, pada project ini database terletak di folder database/mitigasi.sql.
3. Konfigurasikan database pada website, buka file config.json yang terletak pada folder config.
4. Pada file config.json lakukan konfigurasi sesuai dengan database dan host yang digunakan yaitu sebagai berikut.
    ```json
    {
        "DB_USERNAME" : "Username database yang digunakan",
        "DB_PASSWORD" : "Password database yang digunakan",
        "DB_PORT" : "Port dari database",
        "DB_HOST" : "Host dari database",
        "DB_NAME" : "Nama dari database"
    }
    ```
5. Simpan Konfigurasi tersebut.
##### Jalankan Web Server
Setelah selesai melakukan instalasi dan konfigurasi file serta database selanjutnya jalankan web server pada folder project menggunakan perintah.
```bash
go run main.go
```
Buka browser dan masukkan alamat sebagai berikut.
```
https://localhost:8080/
```
###### Note:
Apabila terdapat error depedensi atau package hilang saat menjalankan aplikasi website maka download dependensi yang hilang tersebut, berikut daftar dependensi package yang digunakan.
```
go get github.com/Masterminds/goutils
go get github.com/Masterminds/semver
go get github.com/Masterminds/sprig
go get github.com/alecthomas/template
go get github.com/cpuguy83/go-md2man/v2
go get github.com/dgrijalva/jwt-go
go get github.com/go-openapi/spec
go get github.com/go-openapi/swag
go get github.com/go-sql-driver/mysql
go get github.com/google/uuid
go get github.com/gookit/validate
go get github.com/gorilla/sessions
go get github.com/huandu/xstrings
go get github.com/imdario/mergo
go get github.com/labstack/echo-contrib
go get github.com/labstack/echo/v4
go get github.com/mailru/easyjson
go get github.com/mitchellh/copystructure
go get github.com/russross/blackfriday/v2
go get github.com/stretchr/testify
go get github.com/swaggo/echo-swagger
go get github.com/swaggo/swag
go get github.com/tkanos/gonfig
go get github.com/urfave/cli/v2
golang.org/x/crypto
golang.org/x/tools v0.1.1
```
Untuk mendownload dependensi dapat menggunakan perintah ```go get``` atau menggunakan dependency manager. Dengan contoh sebagai berikut.
```
go get <Nama Package>
```
Dependency Manager
```
dep ensure -add <Nama Package>
```
---
## Dokumentasi
+ ##### Mitigasi Bencana
    + [Data Bencana yang sedang terjadi (GET)](/dokumentasi/databencana.md)
    + [Perbarui Data Bencana (PUT)](/dokumentasi/perbaruibencana.md)
    + [History Bencana (GET)](/dokumentasi/historibencana.md)
+ ##### Laporkan Bencana
    + [Lapor Bencana (POST)](/dokumentasi/laporbencana.md)
    + [Data Pelapor (GET)](/dokumentasi/datapelapor.md)
+ ##### Informasi Bencana
    + [Informasi Bencana (GET)](/dokumentasi/informasibencana.md)
    + [Informasi Bencana dengan Kode Berita (GET)](/dokumentasi/informasibencanaid.md)
    + [Perbarui Berita Bencana (PUT)](/dokumentasi/perbaruiberita.md)
    + [Tulis Berita Bencana (POST)](/dokumentasi/tulisberita.md)
    + [Hapus Informasi Bencana (DELETE)](/dokumentasi/hapusberita.md)
---
## Contoh Penggunaan
Contoh Penggunaan API website secara realtime dapat diakses pada dokumentasi website atau akses pada url berikut [SIMIG - Sistem Informasi dan Mitigasi Bencana](#).
