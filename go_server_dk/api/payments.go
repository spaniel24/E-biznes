package api

import (
	"github.com/labstack/echo/v4"
	"go_server_dk/databases"
	"go_server_dk/models"
	"net/http"
)

func RegisterPayment(c echo.Context) error {
	payment := new(models.Payment)
	c.Bind(payment)
	db := databases.GetDatabase()
	db.Create(&payment)
	return c.JSON(http.StatusOK, payment)
}
