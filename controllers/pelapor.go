package controllers

import (
	"net/http"
	"project-2-rizwijaya/models"

	"github.com/labstack/echo"
)

func SemuaPelapor(c echo.Context) error { //Menghandle response yang masuk
	result, err := models.SemuaPelapor()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()}) //Menampilakan interface error
	}

	return c.JSON(http.StatusOK, result)
}
