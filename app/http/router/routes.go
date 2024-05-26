package router

import (
	"analyze-web/app/config"
	"analyze-web/app/container"
	"analyze-web/app/http/controller"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(cfg *config.Config, ctr *container.Container) *echo.Echo {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, " I'm breathing..")
	})

	e.Renderer = NewWebRenderer(cfg.Service.WebPath)
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/home")
	})
	dc := controller.NewDataController(ctr)
	e.GET("/home", dc.HomeHandler)
	e.POST("/analyze", dc.GetURLData)
	go func() {
		openBrowser(fmt.Sprintf("%s:%d", cfg.Service.BaseURL, cfg.Service.Port))
	}()
	return e
}
