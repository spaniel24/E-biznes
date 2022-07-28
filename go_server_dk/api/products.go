package api

import (
	"github.com/labstack/echo/v4"
	"go_server_dk/databases"
	"go_server_dk/models"
	"net/http"
)

func GetProducts(c echo.Context) error {

	db := databases.GetDatabase()
	var products []models.Product
	category := c.QueryParam("category")
	if category != "" {
		db.Find(&products, "Category = ?", category)
	} else {
		db.Find(&products)
	}
	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {

	id := c.Param("id")
	db := databases.GetDatabase()
	var product models.Product
	db.First(&product, idQuery, id)
	return c.JSON(http.StatusOK, product)
}

func AddProduct(c echo.Context) error {

	product := new(models.Product)
	c.Bind(product)
	db := databases.GetDatabase()
	db.Create(&product)
	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	product := new(models.Product)
	c.Bind(product)
	db := databases.GetDatabase()
	var oldProduct models.Product
	db.First(&oldProduct, idQuery, id)
	oldProduct.Name = product.Name
	oldProduct.Description = product.Description
	oldProduct.Category = product.Category
	oldProduct.Price = product.Price
	oldProduct.ImageUrl = product.ImageUrl
	db.Save(&oldProduct)
	return c.JSON(http.StatusOK, oldProduct)
}

func DeleteProduct(c echo.Context) error {

	id := c.Param("id")
	db := databases.GetDatabase()
	var product models.Product
	db.Delete(&product, idQuery, id)
	return c.JSON(http.StatusOK, product)
}
