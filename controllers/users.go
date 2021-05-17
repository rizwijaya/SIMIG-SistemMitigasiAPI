package controllers

import (
	"fmt"
	"net/http"
	"project-2-rizwijaya/middlewares"
	"project-2-rizwijaya/models"

	"github.com/labstack/echo/v4"
)

func LoginView(c echo.Context) error {
	flash := c.(*middlewares.CustomContext).GetFlash()
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"flash": flash,
	})
}

func BerandaView(c echo.Context) error {
	flash := c.(*middlewares.CustomContext).GetFlash()
	auth := c.(*middlewares.CustomContext).Auth()
	return c.Render(http.StatusOK, "beranda.html", map[string]interface{}{
		"flash": flash,
		"auth":  auth,
	})
}

func DashboardView(c echo.Context) error {
	auth := c.(*middlewares.CustomContext).Auth()
	return c.Render(http.StatusOK, "dashboard.html", map[string]interface{}{
		"auth": auth,
	})
}

func RegisterView(c echo.Context) error {
	flash := c.(*middlewares.CustomContext).GetFlash()
	return c.Render(http.StatusOK, "register.html", map[string]interface{}{
		"flash": flash,
	})
}

func RegisterUser(c echo.Context) error {
	result := models.SaveUser(c)
	fmt.Println(result.LastInsertId())
	c.(*middlewares.CustomContext).SetFlash("done", "Akun berhasil dibuat, silahkan login!")
	return c.Redirect(http.StatusMovedPermanently, "/login")
}
