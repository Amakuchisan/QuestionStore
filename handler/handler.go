package handler

import (
	"github.com/labstack/echo"
	"github.com/stretchr/objx"
	"net/http"
)

// MainPage -- top page
func MainPage(c echo.Context) error {
	auth, err := c.Cookie("auth")
	if err != nil {
		return c.Redirect(http.StatusMovedPermanently, "/auth/login/google")
	}
	userData := objx.MustFromBase64(auth.Value)

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "Dolly!",
	})
}
