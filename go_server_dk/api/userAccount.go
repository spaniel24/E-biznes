package api

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go_server_dk/databases"
	"go_server_dk/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"os"
)

func OauthConfigGoogle() *oauth2.Config {

	//Provide default configuration for oauth provider
	oauthConfig := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		RedirectURL:  "https://shopworkingbackend.azurewebsites.net/gogcall",
	}

	return oauthConfig
}

func OauthConfigGithub() *oauth2.Config {

	//Provide default configuration for oauth provider
	oauthConfig := &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Endpoint:     github.Endpoint,
		Scopes:       []string{"user:email", "read:user"},
		RedirectURL:  "https://shopworkingbackend.azurewebsites.net/github/callback",
	}

	return oauthConfig
}

func OauthConfigFacebook() *oauth2.Config {

	//Provide default configuration for oauth provider
	//oauthConfig := &oauth2.Config{
	//	ClientID:     "1686548605061476",
	//	ClientSecret: "683a5341676d883707d7feec31feef5b",
	//	Endpoint:     facebook.Endpoint,
	//	Scopes:       []string{"user:email", "read:user"},
	//	RedirectURL:  "http://localhost:8080/github/callback",
	//}

	var FACEBOOK = &oauth2.Config{
		ClientID:     os.Getenv("FACEBOOK_APP_ID"),
		ClientSecret: os.Getenv("FACEBOOK_APP_SECRET"),
		Endpoint:     facebook.Endpoint,
		RedirectURL:  "https://shopworkingbackend.azurewebsites.net/facebook",
	}

	return FACEBOOK
}

func OauthLoginUrl(c echo.Context) error {
	client := c.Param("client")

	if client == "google" {
		oauthUrl := OauthConfigGoogle().AuthCodeURL("state")
		return c.JSON(http.StatusOK, oauthUrl)
	}
	if client == "github" {
		oauthUrl := OauthConfigGithub().AuthCodeURL("state")
		return c.JSON(http.StatusOK, oauthUrl)
	}
	if client == "facebook" {
		oauthUrl := OauthConfigFacebook().AuthCodeURL("state")
		return c.JSON(http.StatusOK, oauthUrl)
	}

	return c.JSON(http.StatusOK, "ni mom")
}

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

func UserInDatabase(Username string) bool {

	//Obtain current database connection and fetch user by username
	db := databases.GetDatabase()
	user := models.User{}
	db.Where("username = ?", Username).Find(&user)

	//User exists if object returned from DB does not contain empty fields
	return user.Username != ""
}

//var (
//	key   = []byte("super-secret-key")
//	store = sessions.NewCookieStore(key)
//)

func OauthCallbackGithub(c echo.Context) error {

	//Request access token from provider
	oauthToken, err := OauthConfigGithub().Exchange(context.Background(), c.QueryParam("code"))

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	userRequest, err := http.NewRequest("GET", "https://api.github.com/user", nil)

	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	userRequest.Header.Add("Accept", "application/vnd.github.v3+json")
	userRequest.Header.Add("Authorization", "token "+oauthToken.AccessToken)

	//Perform user data request
	userResponse, err := http.DefaultClient.Do(userRequest)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	defer userResponse.Body.Close()

	userData, err := ioutil.ReadAll(userResponse.Body)
	userDataString := string(userData)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	//Create temporary struct to hold user data returnerd from request
	userDataStruct := struct {
		ID    int
		Login string
	}{}

	//Convert user data json to temporary struct
	json.Unmarshal([]byte(userDataString), &userDataStruct)

	//session, _ := store.Get(r, "cookie-name")
	//session.Values["authenticated"] = true
	//session.Save(r, w)

	//Create new internal user token to save or refresh
	userToken := uuid.New()

	//Obtain current database connection
	db := databases.GetDatabase()

	if !UserInDatabase(userDataStruct.Login) {

		internalUser := models.User{}
		internalUser.Username = userDataStruct.Login
		internalUser.OauthId = userDataStruct.ID
		internalUser.OauthKey = oauthToken.AccessToken
		internalUser.UserKey = userToken.String()
		//internalUser.Cart = newCart

		db.Create(&internalUser)
	} else {

		//If user exists refresh access token
		user := models.User{}
		db.Where("Username = ?", userDataStruct.Login).Find(&user)

		user.OauthKey = oauthToken.AccessToken
		user.UserKey = userToken.String()

		db.Save(&user)
	}

	//Redirect the user to the home page with acces token as query param
	return c.Redirect(http.StatusFound, "https://shopworking.azurewebsites.net?user_token="+userToken.String())
}

func OauthCallbackFacebook(c echo.Context) error {
	oauthToken, err := OauthConfigFacebook().Exchange(context.Background(), c.QueryParam("code"))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	userRequest, err := http.NewRequest("GET", "https://graph.facebook.com/me?access_token="+oauthToken.AccessToken, nil)

	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	userResponse, err := http.DefaultClient.Do(userRequest)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	defer userResponse.Body.Close()
	userData, err := ioutil.ReadAll(userResponse.Body)
	userDataString := string(userData)

	println(userDataString)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	userDataStruct := struct {
		Name string
	}{}

	//Convert user data json to temporary struct
	json.Unmarshal([]byte(userDataString), &userDataStruct)
	userToken := uuid.New()
	db := databases.GetDatabase()
	println(userDataStruct.Name)
	if !UserInDatabase(userDataStruct.Name) {

		internalUser := models.User{}
		internalUser.Username = userDataStruct.Name
		internalUser.OauthKey = oauthToken.AccessToken
		internalUser.UserKey = userToken.String()

		db.Create(&internalUser)
	} else {

		//If user exists refresh access token
		user := models.User{}
		db.Where("username = ?", userDataStruct.Name).Find(&user)

		user.OauthKey = oauthToken.AccessToken
		user.UserKey = userToken.String()

		db.Save(&user)
	}

	return c.Redirect(http.StatusFound, "https://shopworking.azurewebsites.net?user_token="+userToken.String())
}

func OauthCallbackGoogle(c echo.Context) error {
	oauthToken, err := OauthConfigGoogle().Exchange(context.Background(), c.QueryParam("code"))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	userRequest, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo?access_token="+oauthToken.AccessToken, nil)

	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	userResponse, err := http.DefaultClient.Do(userRequest)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	defer userResponse.Body.Close()
	userData, err := ioutil.ReadAll(userResponse.Body)
	userDataString := string(userData)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	userDataStruct := struct {
		Email string
	}{}

	//Convert user data json to temporary struct
	json.Unmarshal([]byte(userDataString), &userDataStruct)
	userToken := uuid.New()
	db := databases.GetDatabase()

	if !UserInDatabase(userDataStruct.Email) {

		internalUser := models.User{}
		internalUser.Username = userDataStruct.Email
		internalUser.OauthKey = oauthToken.AccessToken
		internalUser.UserKey = userToken.String()

		db.Create(&internalUser)
	} else {

		//If user exists refresh access token
		user := models.User{}
		db.Where("username = ?", userDataStruct.Email).Find(&user)

		user.OauthKey = oauthToken.AccessToken
		user.UserKey = userToken.String()

		db.Save(&user)
	}

	return c.Redirect(http.StatusFound, "https://shopworking.azurewebsites.net?user_token="+userToken.String())
}
