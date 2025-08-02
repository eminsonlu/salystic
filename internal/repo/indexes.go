package repo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IndexRepo struct {
	db *mongo.Database
}

func NewIndexRepo(db *mongo.Database) *IndexRepo {
	return &IndexRepo{
		db: db,
	}
}

func (r *IndexRepo) CreateAnalyticsIndexes(ctx context.Context) error {
	salaryCollection := r.db.Collection("salary_entries")

	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "position", Value: 1},
				{Key: "level", Value: 1},
				{Key: "currency", Value: 1},
			},
			Options: options.Index().SetName("analytics_filter_idx"),
		},
		{
			Keys: bson.D{
				{Key: "tech_stack", Value: 1},
				{Key: "position", Value: 1},
				{Key: "level", Value: 1},
			},
			Options: options.Index().SetName("tech_analytics_idx"),
		},
		{
			Keys: bson.D{
				{Key: "raises", Value: 1},
				{Key: "startTime", Value: 1},
				{Key: "endTime", Value: 1},
			},
			Options: options.Index().SetName("career_analytics_idx"),
		},
		{
			Keys: bson.D{
				{Key: "position", Value: 1},
				{Key: "salary_min", Value: 1},
			},
			Options: options.Index().SetName("position_salary_idx"),
		},
		{
			Keys: bson.D{
				{Key: "level", Value: 1},
				{Key: "salary_min", Value: 1},
			},
			Options: options.Index().SetName("level_salary_idx"),
		},
		{
			Keys: bson.D{
				{Key: "experience", Value: 1},
				{Key: "salary_min", Value: 1},
			},
			Options: options.Index().SetName("experience_salary_idx"),
		},
		{
			Keys: bson.D{
				{Key: "company", Value: 1},
				{Key: "salary_min", Value: 1},
			},
			Options: options.Index().SetName("company_salary_idx"),
		},
		{
			Keys: bson.D{
				{Key: "city", Value: 1},
				{Key: "salary_min", Value: 1},
			},
			Options: options.Index().SetName("city_salary_idx"),
		},
		{
			Keys: bson.D{
				{Key: "company_size", Value: 1},
				{Key: "salary_min", Value: 1},
			},
			Options: options.Index().SetName("company_size_salary_idx"),
		},
		{
			Keys: bson.D{
				{Key: "work_type", Value: 1},
				{Key: "salary_min", Value: 1},
			},
			Options: options.Index().SetName("work_type_salary_idx"),
		},
		{
			Keys: bson.D{
				{Key: "currency", Value: 1},
				{Key: "salary_min", Value: 1},
			},
			Options: options.Index().SetName("currency_salary_idx"),
		},
		{
			Keys: bson.D{
				{Key: "position", Value: 1},
			},
			Options: options.Index().SetName("position_idx"),
		},
		{
			Keys: bson.D{
				{Key: "level", Value: 1},
			},
			Options: options.Index().SetName("level_idx"),
		},
	}

	log.Println("Creating analytics indexes for salary_entries collection...")

	indexNames, err := salaryCollection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		return fmt.Errorf("failed to create analytics indexes: %w", err)
	}

	log.Printf("Successfully created %d indexes: %v", len(indexNames), indexNames)
	return nil
}

func (r *IndexRepo) CreateUserIndexes(ctx context.Context) error {
	userCollection := r.db.Collection("users")

	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "linkedin_id", Value: 1},
			},
			Options: options.Index().SetName("linkedin_id_idx").SetUnique(true).SetSparse(true),
		},
	}

	log.Println("Creating user indexes...")

	indexNames, err := userCollection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		return fmt.Errorf("failed to create user indexes: %w", err)
	}

	log.Printf("Successfully created user indexes: %v", indexNames)
	return nil
}

func (r *IndexRepo) CreateAllIndexes(ctx context.Context) error {
	if err := r.CreateAnalyticsIndexes(ctx); err != nil {
		return fmt.Errorf("failed to create analytics indexes: %w", err)
	}

	if err := r.CreateUserIndexes(ctx); err != nil {
		return fmt.Errorf("failed to create user indexes: %w", err)
	}

	log.Println("All database indexes created successfully")
	return nil
}

func (r *IndexRepo) ListIndexes(ctx context.Context, collectionName string) error {
	collection := r.db.Collection(collectionName)

	cursor, err := collection.Indexes().List(ctx)
	if err != nil {
		return fmt.Errorf("failed to list indexes for %s: %w", collectionName, err)
	}
	defer cursor.Close(ctx)

	log.Printf("Indexes for collection '%s':", collectionName)
	for cursor.Next(ctx) {
		var index bson.M
		if err := cursor.Decode(&index); err != nil {
			log.Printf("Error decoding index: %v", err)
			continue
		}
		log.Printf("  - %s: %v", index["name"], index["key"])
	}

	return nil
}
