package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/eminsonlu/salystic/internal/model"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type MockSalaryEntryService struct {
	mock.Mock
}

func (m *MockSalaryEntryService) CreateEntry(ctx context.Context, userID string, req *model.CreateSalaryEntryRequest) (*model.SalaryEntry, error) {
	args := m.Called(ctx, userID, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.SalaryEntry), args.Error(1)
}

func (m *MockSalaryEntryService) GetEntry(ctx context.Context, userID, entryID string) (*model.SalaryEntry, error) {
	args := m.Called(ctx, userID, entryID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.SalaryEntry), args.Error(1)
}

func (m *MockSalaryEntryService) GetUserEntries(ctx context.Context, userID string) ([]*model.SalaryEntry, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.SalaryEntry), args.Error(1)
}

func (m *MockSalaryEntryService) UpdateEntry(ctx context.Context, userID, entryID string, req *model.UpdateSalaryEntryRequest) (*model.SalaryEntry, error) {
	args := m.Called(ctx, userID, entryID, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.SalaryEntry), args.Error(1)
}

func (m *MockSalaryEntryService) DeleteEntry(ctx context.Context, userID, entryID string) error {
	args := m.Called(ctx, userID, entryID)
	return args.Error(0)
}

func (m *MockSalaryEntryService) AddRaise(ctx context.Context, userID, entryID string, req *model.CreateRaiseRequest) error {
	args := m.Called(ctx, userID, entryID, req)
	return args.Error(0)
}

func (m *MockSalaryEntryService) GetRaises(ctx context.Context, userID, entryID string) ([]model.Raise, error) {
	args := m.Called(ctx, userID, entryID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Raise), args.Error(1)
}

func TestNewSalaryHandler(t *testing.T) {
	mockService := &MockSalaryEntryService{}
	handler := NewSalaryHandler(mockService)

	assert.NotNil(t, handler)
	assert.Equal(t, mockService, handler.salaryService)
}

func TestCreateEntry_Success(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	mockService := &MockSalaryEntryService{}
	handler := NewSalaryHandler(mockService)

	userID := primitive.NewObjectID().Hex()
	now := time.Now()

	reqBody := model.CreateSalaryEntryRequest{
		Level:       "Senior",
		Position:    "Backend Developer",
		TechStack:   []string{"Go", "MongoDB"},
		Experience:  "5 - 7 Yıl",
		Gender:      "Erkek",
		Company:     "Technology",
		CompanySize: "101 - 249 Kişi",
		WorkType:    "Remote",
		City:        "İstanbul",
		Currency:    "USD",
		SalaryMin:   150000,
		RaisePeriod: 1,
		StartTime:   now,
	}

	expectedEntry := &model.SalaryEntry{
		ID:          primitive.NewObjectID(),
		Level:       "Senior",
		Position:    "Backend Developer",
		TechStack:   []string{"Go", "MongoDB"},
		Experience:  "5 - 7 Yıl",
		Gender:      "Erkek",
		Company:     "Technology",
		CompanySize: "101 - 249 Kişi",
		WorkType:    "Remote",
		City:        "İstanbul",
		Currency:    "USD",
		SalaryRange: "150.000 - 150.999",
		SalaryMin:   150000,
		SalaryMax:   nil,
		RaisePeriod: 1,
		StartTime:   now,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	mockService.On("CreateEntry", mock.Anything, userID, mock.AnythingOfType("*model.CreateSalaryEntryRequest")).Return(expectedEntry, nil)

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/entries", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user_id", userID)

	err := handler.CreateEntry(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.True(t, response["success"].(bool))
	assert.NotNil(t, response["data"])

	mockService.AssertExpectations(t)
}

func TestGetUserEntries_Success(t *testing.T) {
	e := echo.New()
	mockService := &MockSalaryEntryService{}
	handler := NewSalaryHandler(mockService)

	userID := primitive.NewObjectID().Hex()
	salaryMax := int64(150999)
	entries := []*model.SalaryEntry{
		{
			ID:          primitive.NewObjectID(),
			Level:       "Senior",
			Position:    "Backend Developer",
			TechStack:   []string{"Go", "MongoDB"},
			Experience:  "5 - 7 Yıl",
			Gender:      "Erkek",
			Company:     "Technology",
			CompanySize: "101 - 249 Kişi",
			WorkType:    "Remote",
			City:        "İstanbul",
			Currency:    "USD",
			SalaryRange: "150.000 - 150.999",
			SalaryMin:   150000,
			SalaryMax:   &salaryMax,
			RaisePeriod: 1,
		},
	}

	mockService.On("GetUserEntries", mock.Anything, userID).Return(entries, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/entries", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user_id", userID)

	err := handler.GetUserEntries(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.True(t, response["success"].(bool))
	assert.NotNil(t, response["data"])

	mockService.AssertExpectations(t)
}

func TestDeleteEntry_Success(t *testing.T) {
	e := echo.New()
	mockService := &MockSalaryEntryService{}
	handler := NewSalaryHandler(mockService)

	userID := primitive.NewObjectID().Hex()
	entryID := primitive.NewObjectID().Hex()

	mockService.On("DeleteEntry", mock.Anything, userID, entryID).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/entries/"+entryID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user_id", userID)
	c.SetParamNames("id")
	c.SetParamValues(entryID)

	err := handler.DeleteEntry(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.True(t, response["success"].(bool))
	assert.Equal(t, "Salary entry deleted successfully", response["message"])

	mockService.AssertExpectations(t)
}