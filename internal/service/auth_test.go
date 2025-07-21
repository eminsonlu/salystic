package service

import (
	"context"
	"testing"

	"github.com/eminsonlu/salystic/internal/model"
	"github.com/eminsonlu/salystic/internal/repo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
)

type LinkedInOAuthInterface interface {
	GetAuthURL(state string) string
	ExchangeCodeForToken(ctx context.Context, code string) (*oauth2.Token, error)
	GetLinkedInProfile(ctx context.Context, token *oauth2.Token) (*model.LinkedInProfile, error)
	PseudonymizeLinkedInID(linkedInID string) string
}

type JWTManagerInterface interface {
	GenerateToken(user *model.User) (string, int64, error)
	ValidateToken(tokenString string) (*model.JWTClaims, error)
}

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	if args.Get(0) != nil {
		user.ID = args.Get(0).(primitive.ObjectID)
	}
	return args.Error(1)
}

func (m *MockUserRepository) GetByPseudonymizedID(ctx context.Context, pseudonymizedID string) (*model.User, error) {
	args := m.Called(ctx, pseudonymizedID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*model.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) UpdateLastLogin(ctx context.Context, id primitive.ObjectID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type MockLinkedInOAuth struct {
	mock.Mock
}

func (m *MockLinkedInOAuth) GetAuthURL(state string) string {
	args := m.Called(state)
	return args.String(0)
}

func (m *MockLinkedInOAuth) ExchangeCodeForToken(ctx context.Context, code string) (*oauth2.Token, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*oauth2.Token), args.Error(1)
}

func (m *MockLinkedInOAuth) GetLinkedInProfile(ctx context.Context, token *oauth2.Token) (*model.LinkedInProfile, error) {
	args := m.Called(ctx, token)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.LinkedInProfile), args.Error(1)
}

func (m *MockLinkedInOAuth) PseudonymizeLinkedInID(linkedInID string) string {
	args := m.Called(linkedInID)
	return args.String(0)
}

type MockJWTManager struct {
	mock.Mock
}

func (m *MockJWTManager) GenerateToken(user *model.User) (string, int64, error) {
	args := m.Called(user)
	return args.String(0), args.Get(1).(int64), args.Error(2)
}

func (m *MockJWTManager) ValidateToken(tokenString string) (*model.JWTClaims, error) {
	args := m.Called(tokenString)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.JWTClaims), args.Error(1)
}

type testAuthService struct {
	userRepo      repo.UserRepository
	linkedinOAuth LinkedInOAuthInterface
	jwtManager    JWTManagerInterface
}

func newTestAuthService(userRepo repo.UserRepository, linkedinOAuth LinkedInOAuthInterface, jwtManager JWTManagerInterface) *testAuthService {
	return &testAuthService{
		userRepo:      userRepo,
		linkedinOAuth: linkedinOAuth,
		jwtManager:    jwtManager,
	}
}

func (s *testAuthService) GetLinkedInAuthURL(state string) string {
	return s.linkedinOAuth.GetAuthURL(state)
}

func (s *testAuthService) ValidateToken(tokenString string) (*model.JWTClaims, error) {
	return s.jwtManager.ValidateToken(tokenString)
}

func TestAuthService_GetLinkedInAuthURL(t *testing.T) {
	mockUserRepo := &MockUserRepository{}
	mockLinkedInOAuth := &MockLinkedInOAuth{}
	mockJWTManager := &MockJWTManager{}

	service := newTestAuthService(mockUserRepo, mockLinkedInOAuth, mockJWTManager)

	state := "test_state"
	expectedURL := "https://linkedin.com/auth?state=test_state"

	mockLinkedInOAuth.On("GetAuthURL", state).Return(expectedURL)

	result := service.GetLinkedInAuthURL(state)

	assert.Equal(t, expectedURL, result)
	mockLinkedInOAuth.AssertExpectations(t)
}

func TestAuthService_ValidateToken(t *testing.T) {
	mockUserRepo := &MockUserRepository{}
	mockLinkedInOAuth := &MockLinkedInOAuth{}
	mockJWTManager := &MockJWTManager{}

	service := newTestAuthService(mockUserRepo, mockLinkedInOAuth, mockJWTManager)

	tokenString := "jwt_token"
	expectedClaims := &model.JWTClaims{
		UserID:          "user_id",
		PseudonymizedID: "pseudo_id",
	}

	mockJWTManager.On("ValidateToken", tokenString).Return(expectedClaims, nil)

	result, err := service.ValidateToken(tokenString)

	assert.NoError(t, err)
	assert.Equal(t, expectedClaims, result)
	mockJWTManager.AssertExpectations(t)
}

func TestAuthService_Structure(t *testing.T) {
	mockUserRepo := &MockUserRepository{}
	mockLinkedInOAuth := &MockLinkedInOAuth{}
	mockJWTManager := &MockJWTManager{}

	service := newTestAuthService(mockUserRepo, mockLinkedInOAuth, mockJWTManager)
	
	assert.NotNil(t, service)
	assert.NotNil(t, service.userRepo)
	assert.NotNil(t, service.linkedinOAuth)
	assert.NotNil(t, service.jwtManager)
}

func TestAuthService_Interfaces(t *testing.T) {
	var userRepo repo.UserRepository = &MockUserRepository{}
	var linkedinOAuth LinkedInOAuthInterface = &MockLinkedInOAuth{}
	var jwtManager JWTManagerInterface = &MockJWTManager{}
	
	assert.NotNil(t, userRepo)
	assert.NotNil(t, linkedinOAuth)
	assert.NotNil(t, jwtManager)
}