package main

import (
	"MINI_PROJECT_ALTERRA/routes"
)

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
