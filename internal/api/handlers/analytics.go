package handlers

import (
	"github.com/eminsonlu/salystic/internal/service"
	"github.com/eminsonlu/salystic/pkg/responses"
	"github.com/labstack/echo/v4"
)

type AnalyticsHandler struct {
	analyticsService *service.AnalyticsService
}

func NewAnalyticsHandler(analyticsService *service.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{
		analyticsService: analyticsService,
	}
}

func (h *AnalyticsHandler) GetGeneralAnalytics(c echo.Context) error {
	level := c.QueryParam("level")
	position := c.QueryParam("position")
	currency := c.QueryParam("currency")
	if currency == "" {
		currency = "TRY"
	}

	analytics, err := h.analyticsService.GetGeneralAnalytics(c.Request().Context(), level, position, currency)
	if err != nil {
		return responses.InternalServerError(c, "Failed to get analytics")
	}

	return responses.Success(c, analytics)
}

func (h *AnalyticsHandler) GetCareerAnalytics(c echo.Context) error {
	analytics, err := h.analyticsService.GetCareerAnalytics(c.Request().Context())
	if err != nil {
		return responses.InternalServerError(c, "Failed to get career analytics")
	}

	return responses.Success(c, analytics)
}

func (h *AnalyticsHandler) GetAvailablePositions(c echo.Context) error {
	positions, err := h.analyticsService.GetAvailablePositions(c.Request().Context())
	if err != nil {
		return responses.InternalServerError(c, "Failed to get available positions")
	}

	return responses.Success(c, positions)
}

func (h *AnalyticsHandler) GetAvailableLevels(c echo.Context) error {
	levels, err := h.analyticsService.GetAvailableLevels(c.Request().Context())
	if err != nil {
		return responses.InternalServerError(c, "Failed to get available levels")
	}

	return responses.Success(c, levels)
}