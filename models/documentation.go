package models

//Format token: Bearer <token>
type Lapor_Bencana struct {
	Id int `json:"Kode Laporan Bencana"`
}

type Perbarui_Bencana struct {
	Id int `json:"rows_Affected"`
}

type Daftar_Pelapor struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Nama     string `json:"pelapor"`
}

type Tulis_Berita struct {
	Id int `json:"Kode Berita"`
}
