package models

import (
	"database/sql"
	"errors"
	"fmt"
	"project-2-rizwijaya/db"
	"project-2-rizwijaya/helpers"

	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID       string `json:"id,omitempty" form:"id,omitempty"`
	Nama     string `json:"nama" form:"nama" validate:"required|minLen:5|maxLen:150"`
	Username string `json:"username" form:"username" validate:"required|minLen:5|maxLen:50"`
	Email    string `json:"email" form:"email" validate:"required|email"`
	Password string `json:"password" form:"password" validate:"required|minLen:6"`
}

func (f User) ConfigValidation(v *validate.Validation) {
	v.StopOnError = false
	v.WithScenes(validate.SValues{
		"login":    []string{"Username", "Password"},
		"register": []string{"Nama", "Username", "Email", "Password"},
	})
}

func (f User) Messages() map[string]string {
	return validate.MS{
		"required":        "Bidang {field} harus diisi.",
		"Email.email":     "Format Email tidak valid.",
		"Nama.minLen":     "Minimal panjang Nama adalah 5",
		"Nama.maxLen":     "Maksimal panjang Nama adalah 150",
		"Username.minLen": "Minimal panjang Username adalah 5",
		"Username.maxLen": "Maksimal panjang Username adalah 50",
		"Password.minLen": "Minimal panjang Password adalah 6",
	}
}

func CheckLogin(usernameForm string) (User, error) {
	var user User
	var id, nama, username, email, password string

	con := db.CreateCon()

	sqlStatement := "SELECT id_user, nama_user, email, username, password FROM users WHERE username = ?;"

	err := con.QueryRow(sqlStatement, usernameForm).Scan(&id, &nama, &email, &username, &password)

	if err == sql.ErrNoRows {
		//fmt.Print("Username not found")
		return user, errors.New("Username/Password salah.")
	}

	user = User{ID: id, Nama: nama, Email: email, Username: username, Password: password}
	return user, nil
}

func SaveUser(c echo.Context) sql.Result {
	con := db.CreateCon()
	user := new(User)
	if err := c.Bind(user); err != nil {
		fmt.Println(err.Error())
	}

	hash, _ := helpers.HashPassword(user.Password)
	user.Password = hash
	//Tambahkan User baru ke database
	sqlStatement := "INSERT users (nama_user, email, username, password) VALUES (?, ?, ?, ?)"

	stmt, _ := con.Prepare(sqlStatement)
	defer stmt.Close()

	result, _ := stmt.Exec(user.Nama, user.Email, user.Username, user.Password)
	return result
}

func CheckUser(usernameForm string) (bool, error) {
	var id, nama, username, email, password string

	con := db.CreateCon()

	sqlStatement := "SELECT id_user, nama_user, email, username, password FROM users WHERE username = ?;"

	err := con.QueryRow(sqlStatement, usernameForm).Scan(&id, &nama, &email, &username, &password)

	if err == sql.ErrNoRows { //Jika tidak terdapat username
		return true, err
	}

	return false, nil
}

func CheckEmail(emailForm string) (bool, error) {
	var id, nama, username, email, password string

	con := db.CreateCon()

	sqlStatement := "SELECT id_user, nama_user, email, username, password FROM users WHERE email = ?;"

	err := con.QueryRow(sqlStatement, emailForm).Scan(&id, &nama, &email, &username, &password)

	if err == sql.ErrNoRows { //Jika tidak terdapat email
		return true, err
	}

	return false, nil
}
