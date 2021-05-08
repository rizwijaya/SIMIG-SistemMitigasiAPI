package controllers

import (
	"net/http"
	"project-2-rizwijaya/helpers"
	"project-2-rizwijaya/models"

	"github.com/labstack/echo"
)

func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	return c.String(http.StatusOK, "Berhasil login")

}

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}
