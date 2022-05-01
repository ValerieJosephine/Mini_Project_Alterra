package main

func main() {
	e := routes.routes()

	e.Logger.Fatal(e.start(":8000"))
}
