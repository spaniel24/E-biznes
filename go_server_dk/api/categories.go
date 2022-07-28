package api

import (
	"github.com/labstack/echo/v4"
	"go_server_dk/databases"
	"go_server_dk/models"
	"net/http"
)

var idQuery = "ID = ?"

func GetCategories(c echo.Context) error {

	db := databases.GetDatabase()
	var categories []models.Category
	db.Find(&categories)
	return c.JSON(http.StatusOK, categories)
}

func GetCategory(c echo.Context) error {

	id := c.Param("id")
	db := databases.GetDatabase()
	var category models.Category
	db.First(&category, idQuery, id)
	return c.JSON(http.StatusOK, category)
}

func AddCategory(c echo.Context) error {

	category := new(models.Category)
	c.Bind(category)
	db := databases.GetDatabase()
	db.Create(&category)
	return c.JSON(http.StatusOK, category)
}

func UpdateCategory(c echo.Context) error {
	id := c.Param("id")
	category := new(models.Category)
	c.Bind(category)
	db := databases.GetDatabase()
	var oldCategory models.Category
	db.First(&oldCategory, idQuery, id)
	oldCategory.Name = category.Name
	db.Save(&oldCategory)
	return c.JSON(http.StatusOK, oldCategory)
}

func DeleteCategory(c echo.Context) error {

	id := c.Param("id")
	db := databases.GetDatabase()
	var category models.Category
	db.Delete(&category, idQuery, id)
	return c.JSON(http.StatusOK, category)
}
