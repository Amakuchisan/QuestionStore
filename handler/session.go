package handler

import (
	"github.com/Amakuchisan/QuestionBox/model"
	"github.com/labstack/echo"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	"net/http"
	"time"
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
func (u *userHandler) CallbackHandler(c echo.Context) error {
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

	usr, err := u.userModel.FindByEmail(user.Email())
	if err != nil && usr == nil {
		usr := model.User{Name: user.Name(), Email: user.Email()}
		err = u.userModel.Create(&usr)
		if err != nil {
			return err
		}
	}

	authCookieValue := objx.New(map[string]interface{}{
		"name": user.Name(),
	}).MustBase64()

	cookie := &http.Cookie{
		Name:    "auth",
		Value:   authCookieValue,
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
	}
	c.SetCookie(cookie)

	return c.Redirect(http.StatusMovedPermanently, "/")
}
