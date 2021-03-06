package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/objx"
)

// MainPage -- top page
func MainPage(c echo.Context) error {
	auth, err := c.Cookie("auth")
	if err != nil {
		return err
	}
	userData := objx.MustFromBase64(auth.Value)

	return c.Render(http.StatusOK, "main", map[string]interface{}{
		"title": "TopPage",
		"name":  userData["name"],
	})
}
