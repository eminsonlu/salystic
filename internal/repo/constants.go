package repo

import (
	"context"
	"fmt"
	"log"
	"github.com/eminsonlu/salystic/internal/model"
	"github.com/eminsonlu/salystic/pkg/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConstantsRepository interface {
	SeedConstants(ctx context.Context) error
	GetJobs(ctx context.Context) ([]*model.Job, error)
	GetTitles(ctx context.Context) ([]*model.Title, error)
	GetSectors(ctx context.Context) ([]*model.Sector, error)
	GetCountries(ctx context.Context) ([]*model.Country, error)
	GetCurrencies(ctx context.Context) ([]*model.Currency, error)
}

type constantsRepository struct {
	db         *database.MongoDB
	jobs       *mongo.Collection
	titles     *mongo.Collection
	sectors    *mongo.Collection
	countries  *mongo.Collection
	currencies *mongo.Collection
}

func NewConstantsRepository(db *database.MongoDB) ConstantsRepository {
	return &constantsRepository{
		db:         db,
		jobs:       db.Database.Collection("jobs"),
		titles:     db.Database.Collection("titles"),
		sectors:    db.Database.Collection("sectors"),
		countries:  db.Database.Collection("countries"),
		currencies: db.Database.Collection("currencies"),
	}
}

func (r *constantsRepository) SeedConstants(ctx context.Context) error {
	log.Println("Starting constants seeding...")

	if err := r.seedJobs(ctx); err != nil {
		return fmt.Errorf("failed to seed jobs: %w", err)
	}

	if err := r.seedTitles(ctx); err != nil {
		return fmt.Errorf("failed to seed titles: %w", err)
	}

	if err := r.seedSectors(ctx); err != nil {
		return fmt.Errorf("failed to seed sectors: %w", err)
	}

	if err := r.seedCountries(ctx); err != nil {
		return fmt.Errorf("failed to seed countries: %w", err)
	}

	if err := r.seedCurrencies(ctx); err != nil {
		return fmt.Errorf("failed to seed currencies: %w", err)
	}

	log.Println("Constants seeding completed successfully")
	return nil
}

func (r *constantsRepository) seedJobs(ctx context.Context) error {
	count, err := r.jobs.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Jobs already seeded, skipping...")
		return nil
	}

	jobs := []interface{}{
		model.Job{Name: "Frontend Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Job{Name: "Backend Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Job{Name: "Full Stack Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Job{Name: "DevOps Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Job{Name: "Data Scientist", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Job{Name: "Machine Learning Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Job{Name: "Product Manager", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Job{Name: "Software Architect", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Job{Name: "QA Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Job{Name: "Mobile Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	_, err = r.jobs.InsertMany(ctx, jobs)
	if err != nil {
		return err
	}

	log.Printf("Seeded %d jobs", len(jobs))
	return nil
}

func (r *constantsRepository) seedTitles(ctx context.Context) error {
	count, err := r.titles.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Titles already seeded, skipping...")
		return nil
	}

	titles := []interface{}{
		model.Title{Name: "Intern", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Title{Name: "Junior", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Title{Name: "Mid", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Title{Name: "Senior", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Title{Name: "Lead", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Title{Name: "Principal", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Title{Name: "Staff", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Title{Name: "Manager", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Title{Name: "Director", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	_, err = r.titles.InsertMany(ctx, titles)
	if err != nil {
		return err
	}

	log.Printf("Seeded %d titles", len(titles))
	return nil
}

func (r *constantsRepository) seedSectors(ctx context.Context) error {
	count, err := r.sectors.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Sectors already seeded, skipping...")
		return nil
	}

	sectors := []interface{}{
		model.Sector{Name: "Technology", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Sector{Name: "Finance", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Sector{Name: "Healthcare", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Sector{Name: "Education", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Sector{Name: "E-commerce", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Sector{Name: "Gaming", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Sector{Name: "Media", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Sector{Name: "Government", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Sector{Name: "Consulting", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Sector{Name: "Startup", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	_, err = r.sectors.InsertMany(ctx, sectors)
	if err != nil {
		return err
	}

	log.Printf("Seeded %d sectors", len(sectors))
	return nil
}

func (r *constantsRepository) seedCountries(ctx context.Context) error {
	count, err := r.countries.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Countries already seeded, skipping...")
		return nil
	}

	countries := []interface{}{
		model.Country{Name: "United States", Code: "US", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Country{Name: "United Kingdom", Code: "GB", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Country{Name: "Canada", Code: "CA", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Country{Name: "Germany", Code: "DE", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Country{Name: "France", Code: "FR", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Country{Name: "Netherlands", Code: "NL", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Country{Name: "Australia", Code: "AU", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Country{Name: "Singapore", Code: "SG", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Country{Name: "India", Code: "IN", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Country{Name: "Japan", Code: "JP", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	_, err = r.countries.InsertMany(ctx, countries)
	if err != nil {
		return err
	}

	log.Printf("Seeded %d countries", len(countries))
	return nil
}

func (r *constantsRepository) seedCurrencies(ctx context.Context) error {
	count, err := r.currencies.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Currencies already seeded, skipping...")
		return nil
	}

	currencies := []interface{}{
		model.Currency{Name: "US Dollar", Code: "USD", Symbol: "$", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Currency{Name: "Euro", Code: "EUR", Symbol: "€", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Currency{Name: "British Pound", Code: "GBP", Symbol: "£", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Currency{Name: "Canadian Dollar", Code: "CAD", Symbol: "C$", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Currency{Name: "Australian Dollar", Code: "AUD", Symbol: "A$", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Currency{Name: "Singapore Dollar", Code: "SGD", Symbol: "S$", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Currency{Name: "Indian Rupee", Code: "INR", Symbol: "₹", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Currency{Name: "Japanese Yen", Code: "JPY", Symbol: "¥", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	_, err = r.currencies.InsertMany(ctx, currencies)
	if err != nil {
		return err
	}

	log.Printf("Seeded %d currencies", len(currencies))
	return nil
}

func (r *constantsRepository) GetJobs(ctx context.Context) ([]*model.Job, error) {
	cursor, err := r.jobs.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find jobs: %w", err)
	}
	defer cursor.Close(ctx)

	var jobs []*model.Job
	if err = cursor.All(ctx, &jobs); err != nil {
		return nil, fmt.Errorf("failed to decode jobs: %w", err)
	}

	return jobs, nil
}

func (r *constantsRepository) GetTitles(ctx context.Context) ([]*model.Title, error) {
	cursor, err := r.titles.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find titles: %w", err)
	}
	defer cursor.Close(ctx)

	var titles []*model.Title
	if err = cursor.All(ctx, &titles); err != nil {
		return nil, fmt.Errorf("failed to decode titles: %w", err)
	}

	return titles, nil
}

func (r *constantsRepository) GetSectors(ctx context.Context) ([]*model.Sector, error) {
	cursor, err := r.sectors.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find sectors: %w", err)
	}
	defer cursor.Close(ctx)

	var sectors []*model.Sector
	if err = cursor.All(ctx, &sectors); err != nil {
		return nil, fmt.Errorf("failed to decode sectors: %w", err)
	}

	return sectors, nil
}

func (r *constantsRepository) GetCountries(ctx context.Context) ([]*model.Country, error) {
	cursor, err := r.countries.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find countries: %w", err)
	}
	defer cursor.Close(ctx)

	var countries []*model.Country
	if err = cursor.All(ctx, &countries); err != nil {
		return nil, fmt.Errorf("failed to decode countries: %w", err)
	}

	return countries, nil
}

func (r *constantsRepository) GetCurrencies(ctx context.Context) ([]*model.Currency, error) {
	cursor, err := r.currencies.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find currencies: %w", err)
	}
	defer cursor.Close(ctx)

	var currencies []*model.Currency
	if err = cursor.All(ctx, &currencies); err != nil {
		return nil, fmt.Errorf("failed to decode currencies: %w", err)
	}

	return currencies, nil
}
