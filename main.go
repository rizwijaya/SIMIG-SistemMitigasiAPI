package main

import "project-2-rizwijaya/routes"

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
