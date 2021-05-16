package controllers

import (
	"log"
	"net/http"
	. "project-2-rizwijaya/middlewares"
	"project-2-rizwijaya/models"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gookit/validate"
	"github.com/labstack/echo-contrib/session"
	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

func LoginUser(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/dashboard")
}

func CheckLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := session.Get("session", c) //Get sessions

		data := new(models.User)
		if err2 := c.Bind(data); err2 != nil {
			log.Fatal(err2)
		}
		v := validate.Struct(data) //Lakukan validasi input
		if !v.AtScene("login").Validate() {
			c.(*CustomContext).SetFlash("error", v.Errors)
			//c.(*CustomContext).SetFlash("error", "Email are incorrects format.")
			return c.Redirect(http.StatusMovedPermanently, "/login")
		}

		username := c.FormValue("username")
		password := c.FormValue("password")

		user, err := models.CheckLogin(username) //Mengecek apakah username terdapat pada db.
		//user, err := GetSingleUser(email)
		if err != nil {
			c.(*CustomContext).SetFlash("error", err.Error())
			return c.Redirect(http.StatusMovedPermanently, "/login")
		}
		//Melakukan pengecekan password apakah sesuai
		if err1 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err1 != nil {
			c.(*CustomContext).SetFlash("error", "Username/Password salah.")
			return c.Redirect(http.StatusMovedPermanently, "/login")
		}

		//Jika berhasil login, maka buat Token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = username
		claims["roles"] = "application"
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil { //Jika terdapat eror maka cetak erornya
			c.(*CustomContext).SetFlash("error", "Terdapat Kesalahan, silahkan hubungi Administrator.")
			return c.Redirect(http.StatusMovedPermanently, "/login")
		}

		//Jika tidak ada eror maka buat sessions
		session.Values["authenticated"] = true
		session.Values["nama"] = user.Nama
		session.Values["username"] = user.Username
		session.Values["email"] = user.Email
		session.Values["token"] = t
		session.Save(c.Request(), c.Response())

		return next(c)
	}
}

func LogoutUser(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/login")
}

func Logouting(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.(*CustomContext).Auth()
		if auth.Authenticated == true {
			session, _ := session.Get("session", c)
			session.Values["authenticated"] = false
			session.Values["nama"] = ""
			session.Values["username"] = ""
			session.Values["email"] = ""
			session.Values["token"] = ""
			session.Save(c.Request(), c.Response())
			c.(*CustomContext).SetFlash("done", "Berhasil untuk keluar.")
			return next(c)
		}
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
}
