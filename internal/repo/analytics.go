package repo

import (
	"context"
	"fmt"

	"github.com/eminsonlu/salystic/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AnalyticsRepo struct {
	db *mongo.Database
}

func NewAnalyticsRepo(db *mongo.Database) *AnalyticsRepo {
	return &AnalyticsRepo{
		db: db,
	}
}

func (r *AnalyticsRepo) GetTotalEntries(ctx context.Context) (int64, error) {
	collection := r.db.Collection("salary_entries")
	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, fmt.Errorf("failed to count entries: %w", err)
	}
	return count, nil
}

func (r *AnalyticsRepo) GetAverageSalaryByJob(ctx context.Context) ([]model.SalaryByJobTitle, error) {
	collection := r.db.Collection("salary_entries")
	
	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id": bson.M{
					"job":   "$job",
					"title": "$title",
				},
				"average": bson.M{"$avg": "$salary"},
				"count":   bson.M{"$sum": 1},
			},
		},
		{
			"$match": bson.M{
				"count": bson.M{"$gte": 5}, // Only show categories with 5+ entries
			},
		},
		{
			"$sort": bson.M{"_id.job": 1, "_id.title": 1},
		},
	}
	
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to aggregate salary by job: %w", err)
	}
	defer cursor.Close(ctx)
	
	var results []model.SalaryByJobTitle
	if err = cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode results: %w", err)
	}
	
	return results, nil
}

func (r *AnalyticsRepo) GetAverageSalaryBySector(ctx context.Context) ([]model.SalaryBySector, error) {
	collection := r.db.Collection("salary_entries")
	
	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id":     "$sector",
				"average": bson.M{"$avg": "$salary"},
				"count":   bson.M{"$sum": 1},
			},
		},
		{
			"$match": bson.M{
				"count": bson.M{"$gte": 5}, // Only show categories with 5+ entries
			},
		},
		{
			"$sort": bson.M{"_id": 1},
		},
	}
	
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to aggregate salary by sector: %w", err)
	}
	defer cursor.Close(ctx)
	
	var results []model.SalaryBySector
	if err = cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode results: %w", err)
	}
	
	return results, nil
}

func (r *AnalyticsRepo) GetJobChangeData(ctx context.Context) ([]model.JobChangeData, error) {
	collection := r.db.Collection("salary_entries")
	
	filter := bson.M{
		"previousJobSalary": bson.M{"$exists": true, "$ne": nil, "$gt": 0},
	}
	
	projection := bson.M{
		"previousJobSalary": 1,
		"salary":            1,
	}
	
	cursor, err := collection.Find(ctx, filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, fmt.Errorf("failed to find job change data: %w", err)
	}
	defer cursor.Close(ctx)
	
	var results []model.JobChangeData
	if err = cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode job change data: %w", err)
	}
	
	return results, nil
}

func (r *AnalyticsRepo) GetRaiseData(ctx context.Context) ([]model.RaiseData, error) {
	collection := r.db.Collection("salary_entries")
	
	filter := bson.M{
		"raises": bson.M{"$exists": true, "$ne": nil, "$not": bson.M{"$size": 0}},
	}
	
	projection := bson.M{
		"raises":    1,
		"startTime": 1,
		"endTime":   1,
	}
	
	cursor, err := collection.Find(ctx, filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, fmt.Errorf("failed to find raise data: %w", err)
	}
	defer cursor.Close(ctx)
	
	var results []model.RaiseData
	if err = cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode raise data: %w", err)
	}
	
	return results, nil
}