package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/zuu-development/fullstack-examination-2024/internal/api/dto"
)

// HealthHandler is the request handler for the health endpoint.
type HealthHandler interface {
	Healthz(c echo.Context) error
}

type healthHandler struct{}

// NewHealth returns a new instance of the health handler.
func NewHealth() HealthHandler {
	return &healthHandler{}
}

//	@Summary	Health check
//	@Tags		health
//	@Produce	json
//	@Success	200	{object}	dto.ResponseData
//	@Router		/healthz [get]
func (t *healthHandler) Healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, dto.ResponseData{Data: time.Now()})
}
