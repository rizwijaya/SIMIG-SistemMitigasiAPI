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

	//Fitur Mitigasi Bencana
	e.GET("/pelapor", controllers.SemuaPelapor)  //Dapatkan data pelapor bencana
	e.GET("/bencana", controllers.SemuaBencana)  //Dapatkan data bencana yang pernah terjadi
	e.POST("/bencana", controllers.LaporBencana) //Laporkan bila terjadi bencana

	return e
}
