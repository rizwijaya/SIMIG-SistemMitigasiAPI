package controllers

import (
	"net/http"
	"project-2-rizwijaya/models"
	"strconv"

	"github.com/labstack/echo"
)

func SemuaBerita(c echo.Context) error { //Menghandle response yang masuk
	result, err := models.SemuaBerita()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()}) //Menampilakan interface error
	}

	return c.JSON(http.StatusOK, result)
}

//Fungsi untuk menulis berita
func TulisBerita(c echo.Context) error {
	judul := c.FormValue("judul")
	isi_berita := c.FormValue("isi_berita")
	penulis := c.FormValue("penulis")
	tanggal_ditulis := c.FormValue("tanggal_ditulis")

	result, err := models.TulisBerita(judul, isi_berita, penulis, tanggal_ditulis)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

//Fungsi untuk memperbarui berita
func UpdateBerita(c echo.Context) error {
	id := c.FormValue("id")
	judul := c.FormValue("judul")
	isi_berita := c.FormValue("isi_berita")
	penulis := c.FormValue("penulis")
	tanggal_ditulis := c.FormValue("tanggal_ditulis")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateBerita(conv_id, judul, isi_berita, penulis, tanggal_ditulis)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

//Fungsi hapus berita
func HapusBerita(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.HapusBerita(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
