package main

import "go_server_dk/databases"

func main() {
	databases.InitDatabase()
	e := initRouting()
	e.Logger.Fatal(e.Start(":8080"))
}
