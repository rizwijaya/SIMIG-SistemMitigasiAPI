package controllers

import (
	"net/http"
	"project-2-rizwijaya/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary Data bencana yang sedang terjadi
// @Description Menampilkan daftar bencana yang sedang terjadi diberbagai wilayah di indonesia.
// @Description Jika request berhasil, maka akan menampilkan data daftar bencana yang sedang terjadi.
// @Tags Mitigasi Bencana
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Bencana
// @Failure 422 "Unprocessable Entity"
// @Failure 400 "Bad Request"
// @Router /bencana [get]
func SemuaBencana(c echo.Context) error { //Menampilkan data bencana yang sedang terjadi
	result, err := models.SemuaBencana()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()}) //Menampilakan interface error
	}

	return c.JSON(http.StatusOK, result)
}

// @Summary Lapor Bencana
// @Description Lapor Bencana, pengguna dapat melaporkan bencana yang sedang terjadi dilokasinya.
// @Description Apabila laporan bencana berhasil diterima maka akan mendapatkan kode laporan.
// @Tags Laporkan Bencana
// @Produce  json
// @Security ApiKeyAuth
// @Param username query string true "Username"
// @Param nama query string true "Nama"
// @Param telp query string true "Telp"
// @Param email query string true "Email"
// @Param bencana query string true "Bencana"
// @Param lokasi query string true "Lokasi"
// @Param tgl_terjadi query string true "tgl_terjadi"
// @Success 200 {object} models.Lapor_Bencana
// @Failure 422 "Unprocessable Entity"
// @Failure 400 "Bad Request/Token Salah"
// @Router /bencana [post]
//Fungsi Laporkan bila terjadi bencana
func LaporBencana(c echo.Context) error {
	username := c.FormValue("username")
	nama := c.FormValue("nama")
	telp := c.FormValue("telp")
	email := c.FormValue("email")
	bencana := c.FormValue("bencana")
	lokasi := c.FormValue("lokasi")
	tgl_terjadi := c.FormValue("tgl_terjadi")

	result, err := models.LaporBencana(username, nama, telp, email, bencana, lokasi, tgl_terjadi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// @Summary History Bencana
// @Description Pada history bencana ini akan menampilkan bencana yang pernah terjadi sebelumnya diberbagai wilayah di indonesia.
// @Tags Mitigasi Bencana
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.Bencana
// @Failure 422 "Unprocessable Entity"
// @Failure 400 "Bad Request/Token Salah"
// @Router /historibencana [get]
//Fungsi untuk menampilkan histori bencana alam
func HistoriBencana(c echo.Context) error { //Menghandle response yang masuk
	result, err := models.HistoriBencana()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()}) //Menampilakan interface error
	}

	return c.JSON(http.StatusOK, result)
}

// @Summary Perbarui Data Bencana
// @Description Melalui perbarui data bencana, pengguna dapat memperbarui tanggal bencana jika sudah berhasil diatasi.
// @Description Dalam memperbarui data tersebut menggunakan kode laporan bencana yang sebelumnya dilaporkan.
// @Description Jika data berhasil diperbarui maka akan mengembalikan row affected 1 jika gagal row affected 0.
// @Tags Mitigasi Bencana
// @Produce json
// @Security ApiKeyAuth
// @Param id query int true "Id"
// @Param tgl_selesai query string true "tgl_selesai"
// @Success 200 {object} models.Perbarui_Bencana
// @Failure 422 "Unprocessable Entity"
// @Failure 400 "Bad Request/Token Salah"
// @Router /bencana [put]
func UpdateBencana(c echo.Context) error {
	id := c.FormValue("id")
	tgl_selesai := c.FormValue("tgl_selesai")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateBencana(conv_id, tgl_selesai)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
