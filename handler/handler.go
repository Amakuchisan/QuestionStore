package handler

import (
	"github.com/Amakuchisan/QuestionBox/database"
	"github.com/Amakuchisan/QuestionBox/model"
	"github.com/labstack/echo"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	"net/http"
)

// MainPage -- top page
func MainPage(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

// UsersPage -- user page
func UsersPage(c echo.Context) error {
	users := []model.User{}
	err := database.DB.Select(&users, "SELECT * from user")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

// LoginHandler -- Login to each provider
func LoginHandler(c echo.Context) error {
	provider, err := gomniauth.Provider(c.Param("provider"))
	if err != nil {
		return err
	}
	state := gomniauth.NewState("after", "success")
	authURL, err := provider.GetBeginAuthURL(state, nil)

	if err != nil {
		return err
	}
	return c.Redirect(http.StatusMovedPermanently, authURL)

}

// CallbackHandler -- Provider called this handler after login
func CallbackHandler(c echo.Context) error {
	provider, err := gomniauth.Provider(c.Param("provider"))
	if err != nil {
		return err
	}

	omap, err := objx.FromURLQuery(c.QueryString())
	if err != nil {
		return err
	}

	creds, err := provider.CompleteAuth(omap)
	if err != nil {
		return err
	}

	user, err := provider.GetUser(creds)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
