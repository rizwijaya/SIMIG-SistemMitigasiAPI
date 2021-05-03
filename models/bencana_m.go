package models

import (
	"net/http"
	"project-2-rizwijaya/db"
)

type Bencana struct {
	Id          int    `json:"kode_laporan"`
	Username    string `json:"pelapor"`
	Bencana     string `json:"bencana"`
	Lokasi      string `json:"lokasi"`
	Tgl_laporan string `json:"tanggal_laporan"`
	Tgl_terjadi string `json:"tanggal_terjadi"`
	Tgl_selesai string `json:"tanggal_selesai"`
}

func SemuaBencana() (Response, error) {
	var obj Bencana
	var arrobj []Bencana
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT t2.id_bencana, t1.username, t2.bencana, t2.lokasi, t1.tgl_laporan, t2.tgl_terjadi FROM pelapor t1 INNER JOIN bencana t2 ON t1.id_pelapor=t2.id_pelapor WHERE t2.tgl_selesai IS NULL AND t2.status = 1"

	rows, err := con.Query(sqlStatement)

	if err != nil { //Jika terjadi eror
		return res, err
	}
	defer rows.Close()

	for rows.Next() { //Melakukan perulangan data
		err = rows.Scan(&obj.Id, &obj.Username, &obj.Bencana, &obj.Lokasi, &obj.Tgl_laporan, &obj.Tgl_terjadi)
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

//Fungsi Laporkan bila terjadi bencana
func LaporBencana(username string, nama string, telp string, email string, bencana string, lokasi string, tgl_terjadi string) (Response, error) {
	var res Response

	con := db.CreateCon()

	//Tambahkan data Pelapor ke database
	sqlStatement := "INSERT pelapor (username, nama_pelapor, no_telp, alamat_email) VALUES (?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result2, err := stmt.Exec(username, nama, telp, email)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result2.LastInsertId() //Dapatkan id pelapor
	if err != nil {
		return res, err
	}

	//Tambahkan data bencana ke database
	sqlbencana := "INSERT bencana (id_pelapor, bencana, lokasi, tgl_terjadi) VALUES (?, ?, ?, ?)"

	stmt2, err := con.Prepare(sqlbencana)
	if err != nil {
		return res, err
	}

	result, err := stmt2.Exec(lastInsertedId, bencana, lokasi, tgl_terjadi)
	if err != nil {
		return res, err
	}

	lastInsertedId2, err := result.LastInsertId() //kode laporan bencana
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"Kode Laporan Bencana": lastInsertedId2,
	}

	return res, nil
}

//Menampilkan histori bencana yang pernah terjadi
func HistoriBencana() (Response, error) {
	var obj Bencana
	var arrobj []Bencana
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT t2.id_bencana, t1.username, t2.bencana, t2.lokasi, t1.tgl_laporan, t2.tgl_terjadi, t2.tgl_selesai FROM pelapor t1 INNER JOIN bencana t2 ON t1.id_pelapor=t2.id_pelapor WHERE t2.tgl_selesai IS NOT NULL AND t2.status = 1"

	rows, err := con.Query(sqlStatement)

	if err != nil { //Jika terjadi eror
		return res, err
	}
	defer rows.Close()

	for rows.Next() { //Melakukan perulangan data
		err = rows.Scan(&obj.Id, &obj.Username, &obj.Bencana, &obj.Lokasi, &obj.Tgl_laporan, &obj.Tgl_terjadi, &obj.Tgl_selesai)
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
