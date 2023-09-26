package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthzHandler struct {
}

func (h *HealthzHandler) Healthz(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
