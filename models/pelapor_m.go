package models

import (
	"net/http"
	"project-2-rizwijaya/db"
)

type Pelapor struct {
	Id    int    `json:"id"`
	Nama  string `json:"pelapor"`
	Telp  string `json:"telepon"`
	Email string `json:"email"`
	Date  string `json:"tanggal"`
}

func FetchAllPelapor() (Response, error) {
	var obj Pelapor
	var arrobj []Pelapor
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM pelapor"

	rows, err := con.Query(sqlStatement)

	if err != nil { //Jika terjadi eror
		return res, err
	}
	defer rows.Close()

	for rows.Next() { //Melakukan perulangan data
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Telp, &obj.Email, &obj.Date)
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
