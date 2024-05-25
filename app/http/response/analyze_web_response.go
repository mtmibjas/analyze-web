package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Send(c echo.Context, status int, err error, result map[string]any) error {
	return c.Render(http.StatusOK, "index.html", map[string]any{
		"result": result,
	})
}
