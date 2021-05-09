package controllers

import (
	"net/http"
	"project-2-rizwijaya/middlewares"

	"github.com/labstack/echo/v4"
)

func BerandaView(c echo.Context) error {
	flash := c.(*middlewares.CustomContext).GetFlash()
	auth := c.(*middlewares.CustomContext).Auth()
	return c.Render(http.StatusOK, "beranda.html", map[string]interface{}{
		"flash": flash,
		"auth":  auth,
	})
}
