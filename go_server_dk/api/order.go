package api

import (
	"github.com/labstack/echo/v4"
	"go_server_dk/databases"
	"go_server_dk/models"
	"net/http"
)

func CreateOrder(c echo.Context) error {
	order := new(models.Order)
	db := databases.GetDatabase()
	c.Bind(order)
	db.Create(&order)
	return c.JSON(http.StatusOK, order)
}

func PayForOrder(c echo.Context) error {
	orderId := c.Param("id")
	db := databases.GetDatabase()
	var oldOrder models.Order
	db.First(&oldOrder, "ID = ?", orderId)
	oldOrder.Status = "Paid"
	db.Save(&oldOrder)
	return c.JSON(http.StatusOK, oldOrder)
}
