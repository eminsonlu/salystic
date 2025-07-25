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

func (h *ConstantsHandler) GetPositions(c echo.Context) error {
	positions, err := h.constantsRepo.GetPositions(c.Request().Context())
	if err != nil {
		return responses.InternalServerError(c, "Failed to get positions")
	}

	return responses.Success(c, positions)
}

func (h *ConstantsHandler) GetLevels(c echo.Context) error {
	levels, err := h.constantsRepo.GetLevels(c.Request().Context())
	if err != nil {
		return responses.InternalServerError(c, "Failed to get levels")
	}

	return responses.Success(c, levels)
}

func (h *ConstantsHandler) GetTechStacks(c echo.Context) error {
	techStacks, err := h.constantsRepo.GetTechStacks(c.Request().Context())
	if err != nil {
		return responses.InternalServerError(c, "Failed to get tech stacks")
	}

	return responses.Success(c, techStacks)
}

func (h *ConstantsHandler) GetExperiences(c echo.Context) error {
	experiences, err := h.constantsRepo.GetExperiences(c.Request().Context())
	if err != nil {
		return responses.InternalServerError(c, "Failed to get experiences")
	}

	return responses.Success(c, experiences)
}

func (h *ConstantsHandler) GetCompanies(c echo.Context) error {
	companies, err := h.constantsRepo.GetCompanies(c.Request().Context())
	if err != nil {
		return responses.InternalServerError(c, "Failed to get companies")
	}

	return responses.Success(c, companies)
}

func (h *ConstantsHandler) GetCompanySizes(c echo.Context) error {
	companySizes, err := h.constantsRepo.GetCompanySizes(c.Request().Context())
	if err != nil {
		return responses.InternalServerError(c, "Failed to get company sizes")
	}

	return responses.Success(c, companySizes)
}

func (h *ConstantsHandler) GetWorkTypes(c echo.Context) error {
	workTypes, err := h.constantsRepo.GetWorkTypes(c.Request().Context())
	if err != nil {
		return responses.InternalServerError(c, "Failed to get work types")
	}

	return responses.Success(c, workTypes)
}

func (h *ConstantsHandler) GetCities(c echo.Context) error {
	cities, err := h.constantsRepo.GetCities(c.Request().Context())
	if err != nil {
		return responses.InternalServerError(c, "Failed to get cities")
	}

	return responses.Success(c, cities)
}

func (h *ConstantsHandler) GetCurrencies(c echo.Context) error {
	currencies, err := h.constantsRepo.GetCurrencies(c.Request().Context())
	if err != nil {
		return responses.InternalServerError(c, "Failed to get currencies")
	}

	return responses.Success(c, currencies)
}
