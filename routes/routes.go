package routes

import (
	"net/http"
	"project-2-rizwijaya/controllers"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/pelapor", controllers.FetchAllPelapor)

	return e
}
