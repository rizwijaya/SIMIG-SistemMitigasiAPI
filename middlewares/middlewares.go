package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

func IsNotLogged(next echo.HandlerFunc) echo.HandlerFunc { //Apabila user sudah login maka redirect ke dashboard
	return func(c echo.Context) error {
		auth := c.(*CustomContext).Auth()
		if auth.Authenticated == false {
			return next(c)
		}
		return c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func IsLogged(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.(*CustomContext).Auth()
		if auth.Authenticated == true {
			return next(c)
		}
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
}
