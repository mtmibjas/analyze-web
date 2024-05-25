package controller

import (
	"analyze-web/app/container"
	"analyze-web/domain/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DataController struct {
	Adapters     *container.Container
	DataServices *usecases.Service
}

func NewDataController(ctr *container.Container) *DataController {
	return &DataController{
		Adapters:     ctr,
		DataServices: usecases.NewDataService(ctr),
	}
}

func (d *DataController) GetURLData(c echo.Context) error {
	urlStr := c.FormValue("url")

	result, err := d.DataServices.GetURLData(urlStr)
	if err != nil {

	}

	return c.Render(http.StatusOK, "index.html", map[string]any{
		"result": result,
	})
}
