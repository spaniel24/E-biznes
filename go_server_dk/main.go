package main

import (
	"github.com/labstack/echo/v4/middleware"
	"go_server_dk/databases"
	"os"
)

func main() {
	databases.InitDatabase()
	println("siema 1" + os.Getenv("TEST_VAR"))
	println("siema 2" + os.Getenv("TEST_VAR_TEST_VAR"))
	println("siema 3" + os.Getenv("TEST_test"))
	println("siema 4" + os.Getenv("APPSETTING_test"))
	println("siema 5" + os.Getenv("APPSETTING_TEST_VAR"))
	println("siema 6" + os.Getenv("test"))
	e := initRouting()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://shopworking.azurewebsites.net"},
	}))
	e.Logger.Fatal(e.Start(":8080"))
}
