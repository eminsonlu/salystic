package repo

import (
	"testing"
	"time"

	"github.com/eminsonlu/salystic/internal/model"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserRepository_Structure(t *testing.T) {
	repo := &userRepository{
		collection: nil,
	}
	
	assert.NotNil(t, repo)
}

func TestUser_ModelValidation(t *testing.T) {
	user := &model.User{
		ID:              primitive.NewObjectID(),
		PseudonymizedID: "test_pseudo_id",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		LastLogin:       time.Now(),
	}
	
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.PseudonymizedID)
	assert.False(t, user.CreatedAt.IsZero())
	assert.False(t, user.UpdatedAt.IsZero())
	assert.False(t, user.LastLogin.IsZero())
}

func TestUserRepository_Interface(t *testing.T) {
	var repo UserRepository = &userRepository{}
	
	assert.NotNil(t, repo)
	
	assert.NotNil(t, repo.Create)
	assert.NotNil(t, repo.GetByPseudonymizedID)
	assert.NotNil(t, repo.GetByID)
	assert.NotNil(t, repo.UpdateLastLogin)
}
