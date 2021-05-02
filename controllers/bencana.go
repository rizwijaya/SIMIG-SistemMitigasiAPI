package controllers

import (
	"net/http"
	"project-2-rizwijaya/models"

	"github.com/labstack/echo"
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