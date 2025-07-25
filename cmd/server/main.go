package main

import (
	"context"
	"log"

	"github.com/eminsonlu/salystic/internal/api/routes"
	"github.com/eminsonlu/salystic/internal/auth"
	"github.com/eminsonlu/salystic/internal/config"
	"github.com/eminsonlu/salystic/internal/repo"
	"github.com/eminsonlu/salystic/internal/service"
	"github.com/eminsonlu/salystic/pkg/database"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.NewMongoDB(cfg.MongoURI, cfg.MongoDB, cfg.MongoUser, cfg.MongoPass)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer db.Close()

	log.Println("Connected to MongoDB successfully")

	constantsRepo := repo.NewConstantsRepository(db)
	if err := constantsRepo.SeedConstants(context.Background()); err != nil {
		log.Fatalf("Failed to seed constants: %v", err)
	}

	userRepo := repo.NewUserRepository(db)

	jwtManager, err := auth.NewJWTManager(cfg.JWTSecret, cfg.JWTExpiry, cfg.HMACSecret)
	if err != nil {
		log.Fatalf("Failed to create JWT manager: %v", err)
	}

	linkedinOAuth := auth.NewLinkedInOAuth(cfg.LinkedInClientID, cfg.LinkedInClientSecret, cfg.LinkedInRedirectURL, cfg.HMACSecret)
	authService := service.NewAuthService(userRepo, linkedinOAuth, jwtManager)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	routes.SetupRoutes(e, db, authService, cfg)

	log.Printf("Server starting on port %s", cfg.Port)
	if err := e.Start(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}