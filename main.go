package main

import (
	"project-2-rizwijaya/db"
	"project-2-rizwijaya/routes"
)

// @title Sistem Informasi dan Mitigasi Bencana
// @version 1.0
//Format token: Bearer <token>
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /
func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
