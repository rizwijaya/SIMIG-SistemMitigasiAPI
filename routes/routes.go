package routes

import (
	"net/http"
	"project-2-rizwijaya/controllers"
	"project-2-rizwijaya/middleware"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	//Fitur Mitigasi Bencana
	e.GET("/pelapor", controllers.SemuaPelapor, middleware.IsAuthenticated)          //Dapatkan data pelapor bencana
	e.GET("/bencana", controllers.SemuaBencana)                                      //Dapatkan data bencana yang pernah terjadi
	e.POST("/bencana", controllers.LaporBencana, middleware.IsAuthenticated)         //Laporkan bila terjadi bencana
	e.GET("/historibencana", controllers.HistoriBencana, middleware.IsAuthenticated) //Menampilkan Histori bencana yang pernah terjadi
	e.PUT("/bencana", controllers.UpdateBencana, middleware.IsAuthenticated)         //Update bencana pada Admin - jika selesai

	//Fitur berita bencana
	e.GET("/berita", controllers.SemuaBerita)                                //Daftar berita bencana terbaru
	e.POST("/berita", controllers.TulisBerita, middleware.IsAuthenticated)   //Tambah berita bencana
	e.PUT("/berita", controllers.UpdateBerita, middleware.IsAuthenticated)   //Edit berita bencana
	e.DELETE("/berita", controllers.HapusBerita, middleware.IsAuthenticated) //Hapus berita bencana

	//e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)
	return e
}
