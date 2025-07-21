package handlers

import (
	"net/http"
	"github.com/eminsonlu/salystic/internal/model"
	"github.com/eminsonlu/salystic/internal/service"
	"github.com/eminsonlu/salystic/pkg/responses"

	"github.com/labstack/echo/v4"
)

type SalaryHandler struct {
	salaryService service.SalaryEntryService
}

func NewSalaryHandler(salaryService service.SalaryEntryService) *SalaryHandler {
	return &SalaryHandler{
		salaryService: salaryService,
	}
}

func (h *SalaryHandler) CreateEntry(c echo.Context) error {
	userID := c.Get("user_id").(string)

	var req model.CreateSalaryEntryRequest
	if err := c.Bind(&req); err != nil {
		return responses.BadRequest(c, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return responses.BadRequest(c, err.Error())
	}

	entry, err := h.salaryService.CreateEntry(c.Request().Context(), userID, &req)
	if err != nil {
		return responses.InternalServerError(c, "Failed to create salary entry")
	}

	return c.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Data:    entry,
	})
}

func (h *SalaryHandler) GetEntry(c echo.Context) error {
	userID := c.Get("user_id").(string)
	entryID := c.Param("id")

	if entryID == "" {
		return responses.BadRequest(c, "Entry ID is required")
	}

	entry, err := h.salaryService.GetEntry(c.Request().Context(), userID, entryID)
	if err != nil {
		if err.Error() == "salary entry not found" {
			return responses.NotFound(c, "Salary entry not found")
		}
		return responses.InternalServerError(c, "Failed to get salary entry")
	}

	return responses.Success(c, entry)
}

func (h *SalaryHandler) GetUserEntries(c echo.Context) error {
	userID := c.Get("user_id").(string)

	entries, err := h.salaryService.GetUserEntries(c.Request().Context(), userID)
	if err != nil {
		return responses.InternalServerError(c, "Failed to get salary entries")
	}

	return responses.Success(c, entries)
}

func (h *SalaryHandler) UpdateEntry(c echo.Context) error {
	userID := c.Get("user_id").(string)
	entryID := c.Param("id")

	if entryID == "" {
		return responses.BadRequest(c, "Entry ID is required")
	}

	var req model.UpdateSalaryEntryRequest
	if err := c.Bind(&req); err != nil {
		return responses.BadRequest(c, "Invalid request body")
	}

	entry, err := h.salaryService.UpdateEntry(c.Request().Context(), userID, entryID, &req)
	if err != nil {
		if err.Error() == "salary entry not found" {
			return responses.NotFound(c, "Salary entry not found")
		}
		return responses.InternalServerError(c, "Failed to update salary entry")
	}

	return responses.Success(c, entry)
}

func (h *SalaryHandler) DeleteEntry(c echo.Context) error {
	userID := c.Get("user_id").(string)
	entryID := c.Param("id")

	if entryID == "" {
		return responses.BadRequest(c, "Entry ID is required")
	}

	err := h.salaryService.DeleteEntry(c.Request().Context(), userID, entryID)
	if err != nil {
		if err.Error() == "salary entry not found" {
			return responses.NotFound(c, "Salary entry not found")
		}
		return responses.InternalServerError(c, "Failed to delete salary entry")
	}

	return responses.SuccessWithMessage(c, "Salary entry deleted successfully", nil)
}

func (h *SalaryHandler) AddRaise(c echo.Context) error {
	userID := c.Get("user_id").(string)
	entryID := c.Param("id")

	if entryID == "" {
		return responses.BadRequest(c, "Entry ID is required")
	}

	var req model.CreateRaiseRequest
	if err := c.Bind(&req); err != nil {
		return responses.BadRequest(c, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return responses.BadRequest(c, err.Error())
	}

	err := h.salaryService.AddRaise(c.Request().Context(), userID, entryID, &req)
	if err != nil {
		if err.Error() == "salary entry not found" {
			return responses.NotFound(c, "Salary entry not found")
		}
		return responses.InternalServerError(c, "Failed to add raise")
	}

	return responses.SuccessWithMessage(c, "Raise added successfully", nil)
}

func (h *SalaryHandler) GetRaises(c echo.Context) error {
	userID := c.Get("user_id").(string)
	entryID := c.Param("id")

	if entryID == "" {
		return responses.BadRequest(c, "Entry ID is required")
	}

	raises, err := h.salaryService.GetRaises(c.Request().Context(), userID, entryID)
	if err != nil {
		if err.Error() == "salary entry not found" {
			return responses.NotFound(c, "Salary entry not found")
		}
		return responses.InternalServerError(c, "Failed to get raises")
	}

	return responses.Success(c, raises)
}