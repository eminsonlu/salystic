package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eminsonlu/salystic/internal/config"
	"github.com/eminsonlu/salystic/internal/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) GetLinkedInAuthURL(state string) string {
	args := m.Called(state)
	return args.String(0)
}

func (m *MockAuthService) AuthenticateWithLinkedIn(ctx context.Context, code string) (*model.AuthResponse, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.AuthResponse), args.Error(1)
}

func (m *MockAuthService) ValidateToken(tokenString string) (*model.JWTClaims, error) {
	args := m.Called(tokenString)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.JWTClaims), args.Error(1)
}

func (m *MockAuthService) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockAuthService) Logout(ctx context.Context, userID string) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

func TestNewAuthHandler(t *testing.T) {
	mockService := &MockAuthService{}
	cfg := &config.Config{}

	handler := NewAuthHandler(mockService, cfg)

	assert.NotNil(t, handler)
	assert.Equal(t, mockService, handler.authService)
	assert.Equal(t, cfg, handler.config)
}

func TestLinkedInLogin(t *testing.T) {
	e := echo.New()
	mockService := &MockAuthService{}
	cfg := &config.Config{}
	handler := NewAuthHandler(mockService, cfg)

	mockService.On("GetLinkedInAuthURL", mock.AnythingOfType("string")).Return("https://linkedin.com/auth")

	req := httptest.NewRequest(http.MethodGet, "/auth/linkedin", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.LinkedInLogin(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, rec.Code)
	assert.Equal(t, "https://linkedin.com/auth", rec.Header().Get("Location"))

	cookies := rec.Result().Cookies()
	assert.Len(t, cookies, 1)
	assert.Equal(t, "oauth_state", cookies[0].Name)
	assert.NotEmpty(t, cookies[0].Value)
	assert.True(t, cookies[0].HttpOnly)

	mockService.AssertExpectations(t)
}

func TestLinkedInCallback_Success(t *testing.T) {
	e := echo.New()
	mockService := &MockAuthService{}
	cfg := &config.Config{
		FrontendCallbackURL: "http://frontend.com/callback",
	}
	handler := NewAuthHandler(mockService, cfg)

	mockAuthResponse := &model.AuthResponse{
		User: &model.User{
			ID:              primitive.NewObjectID(),
			PseudonymizedID: "pseudo123",
		},
		AccessToken: "access_token",
		ExpiresIn:   3600,
	}

	mockService.On("AuthenticateWithLinkedIn", mock.Anything, "auth_code").Return(mockAuthResponse, nil)

	req := httptest.NewRequest(http.MethodGet, "/auth/linkedin/callback?code=auth_code&state=test_state", nil)
	req.AddCookie(&http.Cookie{Name: "oauth_state", Value: "test_state"})
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.LinkedInCallback(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, rec.Code)
	assert.Contains(t, rec.Header().Get("Location"), "success=true")
	assert.Contains(t, rec.Header().Get("Location"), "token=access_token")
	assert.Contains(t, rec.Header().Get("Location"), "expires_in=3600")

	mockService.AssertExpectations(t)
}

func TestLinkedInCallback_MissingCode(t *testing.T) {
	e := echo.New()
	mockService := &MockAuthService{}
	cfg := &config.Config{
		FrontendCallbackURL: "http://frontend.com/callback",
	}
	handler := NewAuthHandler(mockService, cfg)

	req := httptest.NewRequest(http.MethodGet, "/auth/linkedin/callback?state=test_state", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.LinkedInCallback(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, rec.Code)
	assert.Contains(t, rec.Header().Get("Location"), "error=missing_code")
}

func TestLinkedInCallback_InvalidState(t *testing.T) {
	e := echo.New()
	mockService := &MockAuthService{}
	cfg := &config.Config{
		FrontendCallbackURL: "http://frontend.com/callback",
	}
	handler := NewAuthHandler(mockService, cfg)

	req := httptest.NewRequest(http.MethodGet, "/auth/linkedin/callback?code=auth_code&state=invalid_state", nil)
	req.AddCookie(&http.Cookie{Name: "oauth_state", Value: "correct_state"})
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.LinkedInCallback(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, rec.Code)
	assert.Contains(t, rec.Header().Get("Location"), "error=invalid_state")
}

func TestMe_Success(t *testing.T) {
	e := echo.New()
	mockService := &MockAuthService{}
	cfg := &config.Config{}
	handler := NewAuthHandler(mockService, cfg)

	userID := primitive.NewObjectID().Hex()
	mockUser := &model.User{
		ID:              primitive.NewObjectID(),
		PseudonymizedID: "pseudo123",
	}

	mockService.On("GetUserByID", mock.Anything, userID).Return(mockUser, nil)

	req := httptest.NewRequest(http.MethodGet, "/auth/me", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user_id", userID)

	err := handler.Me(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.True(t, response["success"].(bool))
	assert.NotNil(t, response["data"])

	mockService.AssertExpectations(t)
}

func TestLogout_Success(t *testing.T) {
	e := echo.New()
	mockService := &MockAuthService{}
	cfg := &config.Config{}
	handler := NewAuthHandler(mockService, cfg)

	userID := primitive.NewObjectID().Hex()

	mockService.On("Logout", mock.Anything, userID).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/auth/logout", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user_id", userID)

	err := handler.Logout(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.True(t, response["success"].(bool))
	assert.Equal(t, "Successfully logged out", response["message"])

	mockService.AssertExpectations(t)
}

func TestGenerateRandomState(t *testing.T) {
	state1, err1 := generateRandomState()
	state2, err2 := generateRandomState()

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NotEmpty(t, state1)
	assert.NotEmpty(t, state2)
	assert.NotEqual(t, state1, state2)
	assert.Equal(t, 64, len(state1))
}
