package auth

import (
	"testing"
	"time"

	"github.com/eminsonlu/salystic/internal/model"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewJWTManager_Success(t *testing.T) {
	secret := "test_secret"
	expiryStr := "24h"
	hmacSecret := "hmac_secret"

	jwtManager, err := NewJWTManager(secret, expiryStr, hmacSecret)

	assert.NoError(t, err)
	assert.NotNil(t, jwtManager)
	assert.Equal(t, secret, jwtManager.secret)
	assert.Equal(t, hmacSecret, jwtManager.hmacSecret)
	assert.Equal(t, 24*time.Hour, jwtManager.expiry)
}

func TestNewJWTManager_InvalidExpiry(t *testing.T) {
	secret := "test_secret"
	expiryStr := "invalid_expiry"
	hmacSecret := "hmac_secret"

	jwtManager, err := NewJWTManager(secret, expiryStr, hmacSecret)

	assert.Error(t, err)
	assert.Nil(t, jwtManager)
	assert.Contains(t, err.Error(), "invalid expiry duration")
}

func TestGenerateToken_Success(t *testing.T) {
	secret := "test_secret"
	expiryStr := "1h"
	hmacSecret := "hmac_secret"

	jwtManager, err := NewJWTManager(secret, expiryStr, hmacSecret)
	assert.NoError(t, err)

	user := &model.User{
		ID:              primitive.NewObjectID(),
		PseudonymizedID: "pseudo_id_123",
	}

	tokenString, expiresAt, err := jwtManager.GenerateToken(user)

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)
	assert.Greater(t, expiresAt, time.Now().Unix())
	assert.LessOrEqual(t, expiresAt, time.Now().Add(time.Hour).Unix())
}

func TestValidateToken_Success(t *testing.T) {
	secret := "test_secret"
	expiryStr := "1h"
	hmacSecret := "hmac_secret"

	jwtManager, err := NewJWTManager(secret, expiryStr, hmacSecret)
	assert.NoError(t, err)

	user := &model.User{
		ID:              primitive.NewObjectID(),
		PseudonymizedID: "pseudo_id_123",
	}

	tokenString, _, err := jwtManager.GenerateToken(user)
	assert.NoError(t, err)

	claims, err := jwtManager.ValidateToken(tokenString)

	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, user.ID.Hex(), claims.UserID)
	assert.Equal(t, user.PseudonymizedID, claims.PseudonymizedID)
	assert.Equal(t, "salystic-backend", claims.Issuer)
	assert.Equal(t, user.ID.Hex(), claims.Subject)
	assert.Greater(t, claims.ExpiresAt, time.Now().Unix())
	assert.LessOrEqual(t, claims.IssuedAt, time.Now().Unix())
}

func TestValidateToken_InvalidToken(t *testing.T) {
	secret := "test_secret"
	expiryStr := "1h"
	hmacSecret := "hmac_secret"

	jwtManager, err := NewJWTManager(secret, expiryStr, hmacSecret)
	assert.NoError(t, err)

	invalidToken := "invalid.jwt.token"

	claims, err := jwtManager.ValidateToken(invalidToken)

	assert.Error(t, err)
	assert.Nil(t, claims)
	assert.Contains(t, err.Error(), "failed to parse token")
}

func TestValidateToken_WrongSecret(t *testing.T) {
	secret1 := "test_secret_1"
	secret2 := "test_secret_2"
	expiryStr := "1h"
	hmacSecret := "hmac_secret"

	jwtManager1, err := NewJWTManager(secret1, expiryStr, hmacSecret)
	assert.NoError(t, err)

	jwtManager2, err := NewJWTManager(secret2, expiryStr, hmacSecret)
	assert.NoError(t, err)

	user := &model.User{
		ID:              primitive.NewObjectID(),
		PseudonymizedID: "pseudo_id_123",
	}

	tokenString, _, err := jwtManager1.GenerateToken(user)
	assert.NoError(t, err)

	claims, err := jwtManager2.ValidateToken(tokenString)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestValidateToken_ExpiredToken(t *testing.T) {
	secret := "test_secret"
	expiryStr := "1ns"
	hmacSecret := "hmac_secret"

	jwtManager, err := NewJWTManager(secret, expiryStr, hmacSecret)
	assert.NoError(t, err)

	user := &model.User{
		ID:              primitive.NewObjectID(),
		PseudonymizedID: "pseudo_id_123",
	}

	tokenString, _, err := jwtManager.GenerateToken(user)
	assert.NoError(t, err)

	time.Sleep(10 * time.Millisecond)

	claims, err := jwtManager.ValidateToken(tokenString)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestPseudonymizeLinkedInID(t *testing.T) {
	secret := "test_secret"
	expiryStr := "1h"
	hmacSecret := "hmac_secret"

	jwtManager, err := NewJWTManager(secret, expiryStr, hmacSecret)
	assert.NoError(t, err)

	linkedInID1 := "linkedin_id_1"
	linkedInID2 := "linkedin_id_2"

	pseudoID1 := jwtManager.PseudonymizeLinkedInID(linkedInID1)
	pseudoID2 := jwtManager.PseudonymizeLinkedInID(linkedInID2)

	assert.NotEmpty(t, pseudoID1)
	assert.NotEmpty(t, pseudoID2)
	assert.NotEqual(t, pseudoID1, pseudoID2)
	assert.Equal(t, 64, len(pseudoID1))
	assert.Equal(t, 64, len(pseudoID2))
}

func TestPseudonymizeLinkedInID_Consistency(t *testing.T) {
	secret := "test_secret"
	expiryStr := "1h"
	hmacSecret := "hmac_secret"

	jwtManager, err := NewJWTManager(secret, expiryStr, hmacSecret)
	assert.NoError(t, err)

	linkedInID := "linkedin_id_123"

	pseudoID1 := jwtManager.PseudonymizeLinkedInID(linkedInID)
	pseudoID2 := jwtManager.PseudonymizeLinkedInID(linkedInID)

	assert.Equal(t, pseudoID1, pseudoID2)
}

func TestGenerateToken_ValidateToken_RoundTrip(t *testing.T) {
	secret := "test_secret"
	expiryStr := "24h"
	hmacSecret := "hmac_secret"

	jwtManager, err := NewJWTManager(secret, expiryStr, hmacSecret)
	assert.NoError(t, err)

	originalUser := &model.User{
		ID:              primitive.NewObjectID(),
		PseudonymizedID: "pseudo_id_456",
	}

	tokenString, expiresAt, err := jwtManager.GenerateToken(originalUser)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)
	assert.Greater(t, expiresAt, time.Now().Unix())

	claims, err := jwtManager.ValidateToken(tokenString)
	assert.NoError(t, err)
	assert.NotNil(t, claims)

	assert.Equal(t, originalUser.ID.Hex(), claims.UserID)
	assert.Equal(t, originalUser.PseudonymizedID, claims.PseudonymizedID)
	assert.Equal(t, "salystic-backend", claims.Issuer)
	assert.Equal(t, originalUser.ID.Hex(), claims.Subject)
}

func TestValidateToken_MissingClaims(t *testing.T) {
	secret := "test_secret"
	expiryStr := "1h"
	hmacSecret := "hmac_secret"

	jwtManager, err := NewJWTManager(secret, expiryStr, hmacSecret)
	assert.NoError(t, err)

	malformedToken := "not.a.valid.jwt.token.format"

	claims, err := jwtManager.ValidateToken(malformedToken)

	assert.Error(t, err)
	assert.Nil(t, claims)
}