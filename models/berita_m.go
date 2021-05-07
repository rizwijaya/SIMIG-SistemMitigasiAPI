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

//Fungsi Tulis berita
func TulisBerita(judul string, isi_berita string, penulis string, tgl_ditulis string) (Response, error) {
	var res Response

	con := db.CreateCon()

	//Tambahkan Berita ke database berita
	sqlStatement := "INSERT berita (judul, isi_berita, penulis, tanggal_ditulis) VALUES (?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result2, err := stmt.Exec(judul, isi_berita, penulis, tgl_ditulis)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result2.LastInsertId() //Dapatkan id berita
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"Kode Berita": lastInsertedId,
	}

	return res, nil
}

//Update untuk memperbarui berita
func UpdateBerita(id int, judul string, isi_berita string, penulis string, tanggal_ditulis string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "Update berita SET judul = ?, isi_berita = ?, penulis = ?, tanggal_ditulis = ?, status = 0 WHERE id_berita = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(judul, isi_berita, penulis, tanggal_ditulis, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_Affected": rowsAffected,
	}

	return res, nil
}

//Fungsi untuk hapus berita
func HapusBerita(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM berita where id_berita =?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_Affected": rowsAffected,
	}

	return res, nil
}
