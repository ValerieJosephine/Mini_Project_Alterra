package main

import (
	"MINI_PROJECT_ALTERRA/db"
	"MINI_PROJECT_ALTERRA/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
