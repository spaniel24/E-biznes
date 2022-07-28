package main

import (
	"github.com/labstack/echo/v4/middleware"
	"go_server_dk/databases"
)

func main() {
	databases.InitDatabase()
	e := initRouting()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://shopworking.azurewebsites.net"},
	}))
	e.Logger.Fatal(e.Start(":8080"))
}
