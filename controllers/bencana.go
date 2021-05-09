package controllers

import (
	"net/http"
	"project-2-rizwijaya/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SemuaBencana(c echo.Context) error { //Menghandle response yang masuk
	result, err := models.SemuaBencana()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()}) //Menampilakan interface error
	}

	return c.JSON(http.StatusOK, result)
}

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

//Fungsi untuk menampilkan histori bencana alam
func HistoriBencana(c echo.Context) error { //Menghandle response yang masuk
	result, err := models.HistoriBencana()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()}) //Menampilakan interface error
	}

	return c.JSON(http.StatusOK, result)
}

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
