package service

import (
	"context"
	"fmt"
	"github.com/eminsonlu/salystic/internal/auth"
	"github.com/eminsonlu/salystic/internal/model"
	"github.com/eminsonlu/salystic/internal/repo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService interface {
	GetLinkedInAuthURL(state string) string
	AuthenticateWithLinkedIn(ctx context.Context, code string) (*model.AuthResponse, error)
	ValidateToken(tokenString string) (*model.JWTClaims, error)
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
	Logout(ctx context.Context, userID string) error
}

type authService struct {
	userRepo      repo.UserRepository
	linkedinOAuth *auth.LinkedInOAuth
	jwtManager    *auth.JWTManager
}

func NewAuthService(userRepo repo.UserRepository, linkedinOAuth *auth.LinkedInOAuth, jwtManager *auth.JWTManager) AuthService {
	return &authService{
		userRepo:      userRepo,
		linkedinOAuth: linkedinOAuth,
		jwtManager:    jwtManager,
	}
}

func (s *authService) GetLinkedInAuthURL(state string) string {
	return s.linkedinOAuth.GetAuthURL(state)
}

func (s *authService) AuthenticateWithLinkedIn(ctx context.Context, code string) (*model.AuthResponse, error) {
	token, err := s.linkedinOAuth.ExchangeCodeForToken(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code for token: %w", err)
	}

	profile, err := s.linkedinOAuth.GetLinkedInProfile(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("failed to get LinkedIn profile: %w", err)
	}

	pseudonymizedID := s.linkedinOAuth.PseudonymizeLinkedInID(profile.Sub)

	user, err := s.userRepo.GetByPseudonymizedID(ctx, pseudonymizedID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by pseudonymized ID: %w", err)
	}

	if user == nil {
		user = &model.User{
			PseudonymizedID: pseudonymizedID,
		}

		if err := s.userRepo.Create(ctx, user); err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}
	} else {
		if err := s.userRepo.UpdateLastLogin(ctx, user.ID); err != nil {
			return nil, fmt.Errorf("failed to update last login: %w", err)
		}
	}

	accessToken, expiresAt, err := s.jwtManager.GenerateToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &model.AuthResponse{
		User:        user,
		Profile:     profile,
		AccessToken: accessToken,
		TokenType:   "Bearer",
		ExpiresIn:   expiresAt,
	}, nil
}

func (s *authService) ValidateToken(tokenString string) (*model.JWTClaims, error) {
	return s.jwtManager.ValidateToken(tokenString)
}

func (s *authService) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *authService) Logout(ctx context.Context, userID string) error {
	return nil
}
