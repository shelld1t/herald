package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shelld1t/core/httpServer"
	"github.com/shelld1t/core/model"
)

type Data struct {
	value string
}

func NewDataController() *Data {
	return &Data{}
}

func (h *Data) DataEndpoints() []*httpServer.Endpoint {
	return []*httpServer.Endpoint{
		{
			Path:   "/data/",
			Method: http.MethodGet,
			Handle: h.GetData,
		},
	}
}

func (h *Data) GetData(ectx echo.Context) model.Response {
	return model.Ok("ok")
}