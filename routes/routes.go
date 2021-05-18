package routes

import (
	"html/template"
	"project-2-rizwijaya/controllers"
	_ "project-2-rizwijaya/docs"
	. "project-2-rizwijaya/helpers"
	"project-2-rizwijaya/middlewares"

	"github.com/Masterminds/sprig"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init() *echo.Echo {
	e := echo.New()

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World!")
	// })
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))
	//catat log website diterminal
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.Static("assets"))                     //set folder assets html
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{ //mengijinkan method dengan CORS
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret")))) //Buat cookie baru
	//Template parser
	templates := make(map[string]*template.Template)
	templates["beranda.html"] = template.Must(template.New("base").Funcs(sprig.FuncMap()).ParseFiles("views/header.html", "views/beranda.html", "views/footer.html", "views/alert.html"))
	templates["login.html"] = template.Must(template.New("base").Funcs(sprig.FuncMap()).ParseFiles("views/login.html", "views/alert.html"))
	templates["dashboard.html"] = template.Must(template.New("base").Funcs(sprig.FuncMap()).ParseFiles("views/header.html", "views/dashboard.html", "views/footer.html", "views/alert.html"))
	templates["documentation.html"] = template.Must(template.New("base").Funcs(sprig.FuncMap()).ParseFiles("views/header.html", "views/documentation.html", "views/footer.html", "views/alert.html"))
	templates["register.html"] = template.Must(template.New("base").Funcs(sprig.FuncMap()).ParseFiles("views/header.html", "views/register.html", "views/footer.html", "views/alert.html"))
	e.Renderer = &TemplateRegistry{ //Lakukan render template
		Templates: templates,
	}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc { //Menampilkan pesan dan error
		return func(c echo.Context) error {
			cc := &middlewares.CustomContext{c}
			return next(cc)
		}
	})

	e.GET("/documentation/*", echoSwagger.WrapHandler)
	//Routing URL pada website
	e.GET("/", controllers.BerandaView)
	e.GET("/dashboard", controllers.DashboardView, middlewares.IsLogged)
	e.GET("/documentation", controllers.DocumentationView)
	//Login dan Logout
	e.GET("/login", controllers.LoginView, middlewares.IsNotLogged)
	e.POST("/login", controllers.LoginUser, controllers.CheckLogin)
	e.DELETE("/logout", controllers.LogoutUser, controllers.Logouting)
	//Register user
	e.GET("/register", controllers.RegisterView)
	e.POST("/register", controllers.RegisterUser, controllers.Registering)

	//Fitur Mitigasi Bencana
	e.GET("/pelapor", controllers.SemuaPelapor, middlewares.IsAuthenticated)          //Dapatkan data pelapor bencana
	e.GET("/bencana", controllers.SemuaBencana)                                       //Dapatkan data bencana yang pernah terjadi
	e.POST("/bencana", controllers.LaporBencana, middlewares.IsAuthenticated)         //Laporkan bila terjadi bencana
	e.GET("/historibencana", controllers.HistoriBencana, middlewares.IsAuthenticated) //Menampilkan Histori bencana yang pernah terjadi
	e.PUT("/bencana", controllers.UpdateBencana, middlewares.IsAuthenticated)         //Update bencana pada Admin - jika selesai

	//Fitur berita bencana
	e.GET("/berita", controllers.SemuaBerita)                                 //Daftar berita bencana terbaru
	e.POST("/berita", controllers.TulisBerita, middlewares.IsAuthenticated)   //Tambah berita bencana
	e.PUT("/berita", controllers.UpdateBerita, middlewares.IsAuthenticated)   //Edit berita bencana
	e.DELETE("/berita", controllers.HapusBerita, middlewares.IsAuthenticated) //Hapus berita bencana

	//e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	return e
}
