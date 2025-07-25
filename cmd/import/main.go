package main

import (
	"context"
	"log"
	"os"

	"github.com/eminsonlu/salystic/internal/config"
	"github.com/eminsonlu/salystic/internal/repo"
	"github.com/eminsonlu/salystic/internal/service"
	"github.com/eminsonlu/salystic/pkg/database"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run cmd/import/main.go <json_file_path>")
	}

	jsonFilePath := os.Args[1]

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.NewMongoDB(cfg.MongoURI, cfg.MongoDB, cfg.MongoUser, cfg.MongoPass)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	salaryRepo := repo.NewSalaryEntryRepository(db)
	importService := service.NewDataImportService(salaryRepo)

	ctx := context.Background()

	log.Printf("Starting import from file: %s", jsonFilePath)

	log.Println("Analyzing tech stacks in the data...")
	if err := importService.AnalyzeTechStacks(ctx, jsonFilePath); err != nil {
		log.Printf("Warning: Failed to analyze tech stacks: %v", err)
	}

	if err := importService.ImportFromJSON(ctx, jsonFilePath); err != nil {
		log.Fatalf("Import failed: %v", err)
	}

	log.Println("Import completed successfully!")
}