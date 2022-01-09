package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shelld1t/core/httpServer"
	"github.com/shelld1t/core/model"
)

type Health struct {
}

func NewHealthController() *Health {
	return &Health{}
}

func (h *Health) HealthEndpoints() []*httpServer.Endpoint {
	return []*httpServer.Endpoint{
		{
			Path:   "/health",
			Method: http.MethodGet,
			Handle: h.Ping,
		},
	}
}

func (h *Health) Ping(ectx echo.Context) model.Response {
	return model.Ok("ok")
}
