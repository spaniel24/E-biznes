package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go_server_dk/api"
)

func initRouting() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/products", api.GetProducts)
	e.GET("/products/:id", api.GetProduct)
	e.POST("/products", api.AddProduct)
	e.PUT("/products/:id", api.UpdateProduct)
	e.DELETE("/products", api.DeleteProduct)

	e.GET("/categories", api.GetCategories)
	e.GET("/categories/:id", api.GetCategory)
	e.POST("/categories", api.AddCategory)
	e.PUT("/categories/:id", api.UpdateCategory)
	e.DELETE("/categories", api.DeleteCategory)

	e.POST("/payments", api.RegisterPayment)

	e.POST("/register", api.RegisterUser)
	e.GET("/login", api.Login)
	e.GET("/logout", api.Logout)

	e.POST("/order", api.CreateOrder)
	e.GET("/order/:id", api.PayForOrder)
	return e
}
