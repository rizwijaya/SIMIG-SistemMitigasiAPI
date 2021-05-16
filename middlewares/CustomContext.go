package middlewares

import (
	"encoding/json"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
}

type Auth struct {
	Nama          string `json:"nama"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Token         string `json:"token"`
	Authenticated bool   `json:"authenticated"`
}

func (c CustomContext) Auth() Auth {
	var auth Auth
	session, _ := session.Get("session", c)
	if session.Values["authenticated"] == nil {
		auth.Nama = ""
		auth.Username = ""
		auth.Email = ""
		auth.Token = ""
		auth.Authenticated = false
		return auth
	}
	auth.Nama = session.Values["nama"].(string)
	auth.Username = session.Values["username"].(string)
	auth.Email = session.Values["email"].(string)
	auth.Token = session.Values["token"].(string)
	auth.Authenticated = session.Values["authenticated"].(bool)
	return auth
}

func (c CustomContext) SetFlash(t string, msg interface{}) bool {
	session, _ := session.Get("flash", c)
	session.Options = &sessions.Options{
		MaxAge: 1,
	}
	flashes := map[string]interface{}{"type": t, "msg": msg}
	flashesJSON, _ := json.Marshal(flashes)
	flashesString := string(flashesJSON)
	session.AddFlash(flashesString)
	session.Save(c.Request(), c.Response())
	return true
}

func (c CustomContext) GetFlash() map[string]interface{} {
	session, _ := session.Get("flash", c)
	if flash := session.Flashes(); len(flash) > 0 {
		flashes := make(map[string]interface{})
		json.Unmarshal([]byte(flash[0].(string)), &flashes)
		return flashes
	}
	return map[string]interface{}{}
}
