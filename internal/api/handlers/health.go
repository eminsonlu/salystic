package handlers

import (
	"github.com/eminsonlu/salystic/pkg/database"
	"github.com/eminsonlu/salystic/pkg/responses"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	db *database.MongoDB
}

func NewHealthHandler(db *database.MongoDB) *HealthHandler {
	return &HealthHandler{
		db: db,
	}
}

func (h *HealthHandler) Health(c echo.Context) error {
	dbStatus := "connected"
	if err := h.db.Health(); err != nil {
		dbStatus = "disconnected"
	}

	healthData := responses.HealthResponse{
		Status:   "ok",
		Database: dbStatus,
		Version:  "1.0.0",
	}

	return responses.Success(c, healthData)
}