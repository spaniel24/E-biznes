package main

import (
	"github.com/labstack/echo/v4/middleware"
	"go_server_dk/databases"
	"os"
)

func main() {
	databases.InitDatabase()
	println(os.Getenv("TEST_VAR"))
	println(os.Getenv("TEST_VAR_TEST_VAR"))
	println(os.Getenv("TEST_test"))
	println(os.Getenv("APPSETTING_test"))
	println(os.Getenv("APPSETTING_TEST_VAR"))
	println(os.Getenv("test"))
	e := initRouting()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://shopworking.azurewebsites.net"},
	}))
	e.Logger.Fatal(e.Start(":8080"))
}
