package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go_server_dk/api"
)

var productsRoute = "/products"
var categoriesRoute = "/categories"

func initRouting() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET(productsRoute, api.GetProducts)
	e.GET("/products/:id", api.GetProduct)
	e.POST(productsRoute, api.AddProduct)
	e.PUT("/products/:id", api.UpdateProduct)
	e.DELETE(productsRoute, api.DeleteProduct)

	e.GET(categoriesRoute, api.GetCategories)
	e.GET("/categories/:id", api.GetCategory)
	e.POST(categoriesRoute, api.AddCategory)
	e.PUT("/categories/:id", api.UpdateCategory)
	e.DELETE(categoriesRoute, api.DeleteCategory)

	e.POST("/payments", api.PayForOrder)

	e.POST("/register", api.RegisterUser)
	e.GET("/login", api.Login)
	e.GET("/logout", api.Logout)

	e.POST("/order", api.CreateOrder)

	e.GET("/oauth/login/:client", api.OauthLoginUrl)
	e.GET("/github/callback", api.OauthCallbackGithub)
	e.GET("/gogcall", api.OauthCallbackGoogle)
	e.GET("/facebook", api.OauthCallbackFacebook)
	e.GET("/linkedin", api.OauthCallbackLinkedin)
	return e
}
