package api

import (
	"github.com/labstack/echo/v4"
	"go_server_dk/databases"
	"go_server_dk/models"
	"net/http"
)

func RegisterUser(c echo.Context) error {
	userData := new(models.User)
	c.Bind(userData)
	db := databases.GetDatabase()
	db.Create(&userData)
	return c.JSON(http.StatusOK, userData)
}

func Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}

func Logout(c echo.Context) error {
	username := c.Param("username")
	db := databases.GetDatabase()
	var user models.User
	db.Find(&user, "Username = ?", username)
	user.OauthKey = ""
	user.UserKey = ""
	db.Save(&user)
	return c.JSON(http.StatusOK, "ok")
}
