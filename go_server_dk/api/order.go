package api

import (
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"go_server_dk/databases"
	"go_server_dk/models"
	"net/http"
)

func CreateOrder(c echo.Context) error {
	userTokenFromFront := c.QueryParam("user_token")
	order := new(models.Order)
	db := databases.GetDatabase()

	var user = models.User{}
	db.First(&user, "user_key = ?", userTokenFromFront)

	if user.OauthKey == "" {
		return c.JSON(http.StatusForbidden, "403")
	}
	println(c.Request())
	c.Bind(order)
	order.UserId = user.Username
	db.Create(&order)
	return c.JSON(http.StatusOK, order)
}

const stripeApiKey = "sk_test_51LO42wLn4wYeRY5IoyHLhN0oOJDJkP0sDpDrt8v4V3JA6Iq6nenKqqyyWV4eZYNqSt4jgKW6JMAfF8apElSvj3Xg00gVWYM2cj"

func PayForOrder(c echo.Context) error {
	userTokenFromFront := c.QueryParam("user_token")
	var user = models.User{}
	db := databases.GetDatabase()
	db.First(&user, "user_key = ?", userTokenFromFront)
	if user.OauthKey == "" {
		return c.JSON(http.StatusForbidden, "403")
	}
	var oldOrder models.Order
	db.First(&oldOrder, "user_id = ?", user.Username)
	stripe.Key = stripeApiKey
	var price = oldOrder.Price * 100
	description := "Order made and paid by " + user.Username
	accepted := true
	_, err := charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(int64(price)),
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		Description:  stripe.String(description),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
		ReceiptEmail: stripe.String("daniel.koszyk024@gmail.com")})

	if err != nil {
		accepted = false
	}
	oldOrder.Status = "Paid"
	db.Save(&oldOrder)
	return c.JSON(http.StatusOK, accepted)
}
