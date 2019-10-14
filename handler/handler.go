package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

// MainPage -- top page
func MainPage(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "Dolly!",
	})
}
