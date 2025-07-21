package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/eminsonlu/salystic/internal/service"
	"github.com/eminsonlu/salystic/pkg/responses"
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
	analytics, err := h.analyticsService.GetGeneralAnalytics(c.Request().Context())
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