package main

import (
	"github.com/labstack/echo/v4/middleware"
	"go_server_dk/databases"
	"os"
)

func main() {
	databases.InitDatabase()
	zmienna := os.Getenv("TEST_VAR")
	println(zmienna)
	e := initRouting()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://shopworking.azurewebsites.net"},
	}))
	e.Logger.Fatal(e.Start(":8080"))
}
