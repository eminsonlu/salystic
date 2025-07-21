package routes

import (
	"github.com/eminsonlu/salystic/internal/api/handlers"
	authMiddleware "github.com/eminsonlu/salystic/internal/api/middleware"
	"github.com/eminsonlu/salystic/internal/config"
	"github.com/eminsonlu/salystic/internal/repo"
	"github.com/eminsonlu/salystic/internal/service"
	"github.com/eminsonlu/salystic/pkg/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(e *echo.Echo, db *database.MongoDB, authService service.AuthService, cfg *config.Config) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	salaryRepo := repo.NewSalaryEntryRepository(db)
	constantsRepo := repo.NewConstantsRepository(db)
	analyticsRepo := repo.NewAnalyticsRepo(db.Database)

	salaryService := service.NewSalaryEntryService(salaryRepo)
	analyticsService := service.NewAnalyticsService(analyticsRepo)

	healthHandler := handlers.NewHealthHandler(db)
	authHandler := handlers.NewAuthHandler(authService, cfg)
	salaryHandler := handlers.NewSalaryHandler(salaryService)
	constantsHandler := handlers.NewConstantsHandler(constantsRepo)
	analyticsHandler := handlers.NewAnalyticsHandler(analyticsService)
	authMW := authMiddleware.NewAuthMiddleware(authService)

	e.GET("/health", healthHandler.Health)

	authGroup := e.Group("/auth")
	authGroup.GET("/linkedin", authHandler.LinkedInLogin)
	authGroup.GET("/linkedin/callback", authHandler.LinkedInCallback)
	authGroup.GET("/me", authHandler.Me, authMW.RequireAuth)
	authGroup.POST("/logout", authHandler.Logout, authMW.RequireAuth)

	api := e.Group("/api/v1")
	api.GET("/health", healthHandler.Health)

	entriesGroup := api.Group("/entries", authMW.RequireAuth)
	entriesGroup.POST("", salaryHandler.CreateEntry)
	entriesGroup.GET("", salaryHandler.GetUserEntries)
	entriesGroup.GET("/:id", salaryHandler.GetEntry)
	entriesGroup.PUT("/:id", salaryHandler.UpdateEntry)
	entriesGroup.DELETE("/:id", salaryHandler.DeleteEntry)
	entriesGroup.POST("/:id/raises", salaryHandler.AddRaise)
	entriesGroup.GET("/:id/raises", salaryHandler.GetRaises)

	constantsGroup := api.Group("/constants")
	constantsGroup.GET("/jobs", constantsHandler.GetJobs)
	constantsGroup.GET("/titles", constantsHandler.GetTitles)
	constantsGroup.GET("/sectors", constantsHandler.GetSectors)
	constantsGroup.GET("/countries", constantsHandler.GetCountries)
	constantsGroup.GET("/currencies", constantsHandler.GetCurrencies)

	analyticsGroup := api.Group("/analytics")
	analyticsGroup.GET("", analyticsHandler.GetGeneralAnalytics)
	analyticsGroup.GET("/career", analyticsHandler.GetCareerAnalytics)
}
