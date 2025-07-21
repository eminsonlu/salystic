package handlers

import (
	"github.com/eminsonlu/salystic/internal/repo"
	"github.com/eminsonlu/salystic/pkg/responses"

	"github.com/labstack/echo/v4"
)

type ConstantsHandler struct {
	constantsRepo repo.ConstantsRepository
}

func NewConstantsHandler(constantsRepo repo.ConstantsRepository) *ConstantsHandler {
	return &ConstantsHandler{
		constantsRepo: constantsRepo,
	}
}

func (h *ConstantsHandler) GetJobs(c echo.Context) error {
	jobs, err := h.constantsRepo.GetJobs(c.Request().Context())
	if err != nil {
		return responses.InternalServerError(c, "Failed to get jobs")
	}

	return responses.Success(c, jobs)
}

func (h *ConstantsHandler) GetTitles(c echo.Context) error {
	titles, err := h.constantsRepo.GetTitles(c.Request().Context())
	if err != nil {
		return responses.InternalServerError(c, "Failed to get titles")
	}

	return responses.Success(c, titles)
}

func (h *ConstantsHandler) GetSectors(c echo.Context) error {
	sectors, err := h.constantsRepo.GetSectors(c.Request().Context())
	if err != nil {
		return responses.InternalServerError(c, "Failed to get sectors")
	}

	return responses.Success(c, sectors)
}

func (h *ConstantsHandler) GetCountries(c echo.Context) error {
	countries, err := h.constantsRepo.GetCountries(c.Request().Context())
	if err != nil {
		return responses.InternalServerError(c, "Failed to get countries")
	}

	return responses.Success(c, countries)
}

func (h *ConstantsHandler) GetCurrencies(c echo.Context) error {
	currencies, err := h.constantsRepo.GetCurrencies(c.Request().Context())
	if err != nil {
		return responses.InternalServerError(c, "Failed to get currencies")
	}

	return responses.Success(c, currencies)
}