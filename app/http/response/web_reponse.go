package response

import (
	"github.com/labstack/echo/v4"
)

func Send(c echo.Context, status int, result map[string]any) error {
	return c.Render(status, "index.html", map[string]any{
		"result": result,
	})
}
