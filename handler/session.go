package handler

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/Amakuchisan/QuestionStore/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
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

	return c.Redirect(http.StatusTemporaryRedirect, authURL)
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

	if !isAuthorizedDomain(user.Email()) {
		accessToken := creds.Get("access_token").Str()
		err := revokeToken(accessToken)
		if err != nil {
			return err
		}

		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
	}

	// This function search whether login user is existed.
	_, err = u.userModel.FindByEmail(user.Email())

	if err != nil {
		usr := model.User{Name: user.Name(), Email: user.Email()}
		err = u.userModel.Create(&usr)
		if err != nil {
			return err
		}

	}

	data, err := u.userModel.FindByEmail(user.Email())
	if err != nil {
		return err
	}

	authCookieValue := objx.New(map[string]interface{}{
		"name": user.Name(),
		"id":   data.ID,
	}).MustBase64()

	cookie := &http.Cookie{
		Name:    "auth",
		Value:   authCookieValue,
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
	}
	c.SetCookie(cookie)

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func isAuthorizedDomain(email string) bool {
	return strings.HasSuffix(email, os.Getenv("AUTHORIZED_DOMAIN"))
}

func revokeToken(accessToken string) error {
	const googleRevokeURL = "https://accounts.google.com/o/oauth2/revoke"
	u, err := url.Parse(googleRevokeURL)
	if err != nil {
		return err
	}

	q := u.Query()
	q.Set("token", accessToken)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}

	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	return err
}
