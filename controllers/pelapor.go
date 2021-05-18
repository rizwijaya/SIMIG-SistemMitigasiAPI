package controllers

import (
	"net/http"
	"project-2-rizwijaya/models"

	"github.com/labstack/echo/v4"
)

// SemuaPelapor godoc
// @Summary Data Pelapor
// @Description Daftar warga yang pernah melaporkan bencana yang terjadi
// @Tags Laporkan Bencana
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.Daftar_Pelapor
// @Failure 422 "Unprocessable Entity"
// @Failure 400 "Bad Request/Token Salah"
// @Router /pelapor [get]
func SemuaPelapor(c echo.Context) error { //Daftar Pelapor
	result, err := models.SemuaPelapor()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()}) //Menampilakan interface error
	}

	return c.JSON(http.StatusOK, result)
}
