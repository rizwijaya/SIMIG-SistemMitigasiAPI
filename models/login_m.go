package models

import (
	"database/sql"
	"fmt"
	"project-2-rizwijaya/db"
	"project-2-rizwijaya/helpers"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT id_user, username, password FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(&obj.Id, &obj.Username, &pwd)

	if err == sql.ErrNoRows {
		fmt.Print("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	//Jika tidak ada error, mengecek password
	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash dan password doesn't match.")
		return false, err
	}

	return true, nil
}
