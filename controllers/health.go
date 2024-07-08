package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Health struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}

// HealthCheck godoc
// @Tags Health
// @Summary Check if service is active
// @Description Check health of the service
// @Produce  json
// @Success 200 {object} Health
// @Router / [get]
func HealthCheck(c echo.Context) error {
	response := &Health{
		Code:    http.StatusOK,
		Message: "Active!",
	}

	return c.JSON(http.StatusOK, response)
}
