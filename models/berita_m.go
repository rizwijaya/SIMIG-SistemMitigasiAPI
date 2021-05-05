package models

import (
	"net/http"
	"project-2-rizwijaya/db"
)

type Berita struct {
	Id          int    `json:"kode_berita"`
	Judul       string `json:"judul"`
	Isi_berita  string `json:"isi_berita"`
	Penulis     string `json:"penulis"`
	Tgl_ditulis string `json:"tanggal_ditulis"`
}

func SemuaBerita() (Response, error) { //Fungsi untuk menampilkan berita bencana
	var obj Berita
	var arrobj []Berita
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id_berita, judul, isi_berita, penulis, tanggal_ditulis FROM berita WHERE status != 0"

	rows, err := con.Query(sqlStatement)

	if err != nil { //Jika terjadi eror
		return res, err
	}
	defer rows.Close()

	for rows.Next() { //Melakukan perulangan data
		err = rows.Scan(&obj.Id, &obj.Judul, &obj.Isi_berita, &obj.Penulis, &obj.Tgl_ditulis)
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
