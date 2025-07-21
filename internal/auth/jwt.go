package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/eminsonlu/salystic/internal/model"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secret     string
	expiry     time.Duration
	hmacSecret string
}

func NewJWTManager(secret, expiryStr, hmacSecret string) (*JWTManager, error) {
	expiry, err := time.ParseDuration(expiryStr)
	if err != nil {
		return nil, fmt.Errorf("invalid expiry duration: %w", err)
	}

	return &JWTManager{
		secret:     secret,
		expiry:     expiry,
		hmacSecret: hmacSecret,
	}, nil
}

func (j *JWTManager) GenerateToken(user *model.User) (string, int64, error) {
	now := time.Now()
	expiresAt := now.Add(j.expiry)

	claims := model.JWTClaims{
		UserID:          user.ID.Hex(),
		PseudonymizedID: user.PseudonymizedID,
		StandardClaims: model.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: expiresAt.Unix(),
			Issuer:    "salystic-backend",
			Subject:   user.ID.Hex(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":          claims.UserID,
		"pseudonymized_id": claims.PseudonymizedID,
		"iat":              claims.IssuedAt,
		"exp":              claims.ExpiresAt,
		"iss":              claims.Issuer,
		"sub":              claims.Subject,
	})

	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", 0, fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, expiresAt.Unix(), nil
}

func (j *JWTManager) ValidateToken(tokenString string) (*model.JWTClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid user_id claim")
	}

	pseudonymizedID, ok := claims["pseudonymized_id"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid pseudonymized_id claim")
	}

	iat, ok := claims["iat"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid iat claim")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid exp claim")
	}

	iss, ok := claims["iss"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid iss claim")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid sub claim")
	}

	return &model.JWTClaims{
		UserID:          userID,
		PseudonymizedID: pseudonymizedID,
		StandardClaims: model.StandardClaims{
			IssuedAt:  int64(iat),
			ExpiresAt: int64(exp),
			Issuer:    iss,
			Subject:   sub,
		},
	}, nil
}

func (j *JWTManager) PseudonymizeLinkedInID(linkedInID string) string {
	h := hmac.New(sha256.New, []byte(j.hmacSecret))
	h.Write([]byte(linkedInID))
	h.Write([]byte(strconv.FormatInt(time.Now().Unix()/86400, 10)))
	return hex.EncodeToString(h.Sum(nil))
}
