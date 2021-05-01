package models

import (
	"net/http"
	"project-2-rizwijaya/db"
)

type Bencana struct {
	Id          int    `json:"id_laporan"`
	Username    string `json:"username"`
	Bencana     string `json:"bencana"`
	Lokasi      string `json:"lokasi"`
	Tgl_terjadi string `json:"tanggal_terjadi"`
}

func SemuaBencana() (Response, error) {
	var obj Bencana
	var arrobj []Bencana
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT t2.id_bencana, t1.username, t2.bencana, t2.lokasi, t2.tgl_terjadi FROM pelapor t1 INNER JOIN bencana t2 ON t1.id_pelapor=t2.id_pelapor WHERE t2.tgl_selesai IS NULL"

	rows, err := con.Query(sqlStatement)

	if err != nil { //Jika terjadi eror
		return res, err
	}
	defer rows.Close()

	for rows.Next() { //Melakukan perulangan data
		err = rows.Scan(&obj.Id, &obj.Username, &obj.Bencana, &obj.Lokasi, &obj.Tgl_terjadi)
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
