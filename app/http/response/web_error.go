package response

import (
	"github.com/labstack/echo/v4"
)

func Error(c echo.Context, status int, err error) error {
	return c.Render(status, "index.html", map[string]any{
		"Error": err.Error(),
	})
}
