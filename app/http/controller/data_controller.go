package controller

import (
	"analyze-web/app/container"
	"analyze-web/app/http/response"
	"analyze-web/app/http/validator"
	"analyze-web/domain/usecases"
	"errors"
	"net/http"

	"analyze-web/pkg/logger/zap"

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

func (d *DataController) HomeHandler(c echo.Context) error {
	zap.Debug("runner", "test")
	zap.Error("ada", errors.New("sdfs"))
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
}

func (d *DataController) GetURLData(c echo.Context) error {
	urlStr := c.FormValue("url")
	if err := validator.ValidateURL(urlStr); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	result, err := d.DataServices.GetURLData(urlStr)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	res := map[string]any{
		"Result": result,
	}
	return response.Send(c, http.StatusOK, res)
}
