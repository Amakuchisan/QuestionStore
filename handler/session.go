package handler

import (
	"github.com/labstack/echo"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	"net/http"
)

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
