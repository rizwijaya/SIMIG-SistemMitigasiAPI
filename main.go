package main

import (
	"project-2-rizwijaya/db"
	"project-2-rizwijaya/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
