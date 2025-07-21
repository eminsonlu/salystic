package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PseudonymizedID string             `bson:"pseudonymized_id" json:"pseudonymized_id"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
	LastLogin       time.Time          `bson:"last_login" json:"last_login"`
}

type JWTClaims struct {
	UserID          string `json:"user_id"`
	PseudonymizedID string `json:"pseudonymized_id"`
	StandardClaims
}

type StandardClaims struct {
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
	Issuer    string `json:"iss"`
	Subject   string `json:"sub"`
}

type LinkedInProfile struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

type AuthResponse struct {
	User        *User            `json:"user"`
	Profile     *LinkedInProfile `json:"profile"`
	AccessToken string           `json:"access_token"`
	TokenType   string           `json:"token_type"`
	ExpiresIn   int64            `json:"expires_in"`
}
