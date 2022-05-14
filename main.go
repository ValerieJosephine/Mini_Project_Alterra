package main

import (
	db "MINI_PROJECT_ALTERRA/database"
	"MINI_PROJECT_ALTERRA/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
