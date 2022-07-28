package main

import (
	"go_server_dk/databases"
	"os"
)

func main() {
	databases.InitDatabase()
	zmienna := os.Getenv("TEST_VAR")
	println(zmienna)
	e := initRouting()
	e.Logger.Fatal(e.Start(":8080"))
}
