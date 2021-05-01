package models

import (
	"net/http"
	"project-2-rizwijaya/db"
)

type Pelapor struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Nama     string `json:"pelapor"`
}

func SemuaPelapor() (Response, error) {
	var obj Pelapor
	var arrobj []Pelapor
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id_pelapor, username, nama_pelapor FROM pelapor"

	rows, err := con.Query(sqlStatement)

	if err != nil { //Jika terjadi eror
		return res, err
	}
	defer rows.Close()

	for rows.Next() { //Melakukan perulangan data
		err = rows.Scan(&obj.Id, &obj.Username, &obj.Nama)
		if err != nil { //Jika terjadi eror maka kembalikan ke controller
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}
