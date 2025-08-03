package repo

import (
	"context"
	"fmt"
	"strings"

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

type AnalyticsFilter struct {
	Position string
	Level    string
	Currency string
}

func (r *AnalyticsRepo) GetTotalEntries(ctx context.Context, filter *AnalyticsFilter) (int64, error) {
	collection := r.db.Collection("salary_entries")

	query := bson.M{}
	if filter != nil {
		query = r.buildFilterQuery(filter)
	}

	count, err := collection.CountDocuments(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("failed to count entries: %w", err)
	}
	return count, nil
}

func (r *AnalyticsRepo) GetAverageSalaryByPosition(ctx context.Context, filter *AnalyticsFilter) ([]model.SalaryByCategory, error) {
	return r.getAverageSalaryByField(ctx, "$position", filter)
}

func (r *AnalyticsRepo) GetAverageSalaryByLevel(ctx context.Context, filter *AnalyticsFilter) ([]model.SalaryByCategory, error) {
	return r.getAverageSalaryByField(ctx, "$level", filter)
}

func (r *AnalyticsRepo) GetAverageSalaryByExperience(ctx context.Context, filter *AnalyticsFilter) ([]model.SalaryByCategory, error) {
	return r.getAverageSalaryByField(ctx, "$experience", filter)
}

func (r *AnalyticsRepo) GetAverageSalaryByCompany(ctx context.Context, filter *AnalyticsFilter) ([]model.SalaryByCategory, error) {
	return r.getAverageSalaryByField(ctx, "$company", filter)
}

func (r *AnalyticsRepo) GetAverageSalaryByCity(ctx context.Context, filter *AnalyticsFilter) ([]model.SalaryByCategory, error) {
	return r.getAverageSalaryByField(ctx, "$city", filter)
}

func (r *AnalyticsRepo) GetAverageSalaryByCompanySize(ctx context.Context, filter *AnalyticsFilter) ([]model.SalaryByCategory, error) {
	return r.getAverageSalaryByField(ctx, "$company_size", filter)
}

func (r *AnalyticsRepo) GetAverageSalaryByWorkType(ctx context.Context, filter *AnalyticsFilter) ([]model.SalaryByCategory, error) {
	return r.getAverageSalaryByField(ctx, "$work_type", filter)
}

func (r *AnalyticsRepo) GetAverageSalaryByCurrency(ctx context.Context, filter *AnalyticsFilter) ([]model.SalaryByCategory, error) {
	return r.getAverageSalaryByField(ctx, "$currency", filter)
}

func (r *AnalyticsRepo) GetAverageSalaryByTech(ctx context.Context, filter *AnalyticsFilter) ([]model.SalaryByTech, error) {
	collection := r.db.Collection("salary_entries")

	pipeline := []bson.M{}

	baseMatch := bson.M{
		"tech_stack": bson.M{"$exists": true, "$ne": nil, "$not": bson.M{"$size": 0}},
	}

	if filter != nil {
		additionalFilters := r.buildFilterQuery(filter)
		for key, value := range additionalFilters {
			baseMatch[key] = value
		}
	}

	pipeline = append(pipeline, bson.M{"$match": baseMatch})

	pipeline = append(pipeline, []bson.M{
		{
			"$unwind": "$tech_stack",
		},
		{
			"$match": bson.M{
				"tech_stack": bson.M{"$ne": "", "$exists": true},
			},
		},
		{
			"$group": bson.M{
				"_id":     "$tech_stack",
				"average": bson.M{"$avg": "$salary_min"},
				"min":     bson.M{"$min": "$salary_min"},
				"max":     bson.M{"$max": "$salary_min"},
				"count":   bson.M{"$sum": 1},
			},
		},
		{
			"$sort": bson.M{"_id": 1},
		},
	}...)

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to aggregate salary by tech: %w", err)
	}
	defer cursor.Close(ctx)

	var results []model.SalaryByTech
	if err = cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode results: %w", err)
	}

	return results, nil
}

func (r *AnalyticsRepo) getAverageSalaryByField(ctx context.Context, field string, filter *AnalyticsFilter) ([]model.SalaryByCategory, error) {
	collection := r.db.Collection("salary_entries")

	pipeline := []bson.M{}

	fieldName := strings.TrimPrefix(field, "$")
	baseMatch := bson.M{
		fieldName: bson.M{"$ne": "", "$exists": true},
	}

	if filter != nil {
		additionalFilters := r.buildFilterQuery(filter)
		for key, value := range additionalFilters {
			baseMatch[key] = value
		}
	}

	pipeline = append(pipeline, bson.M{"$match": baseMatch})

	pipeline = append(pipeline, []bson.M{
		{
			"$group": bson.M{
				"_id":     field,
				"average": bson.M{"$avg": "$salary_min"},
				"min":     bson.M{"$min": "$salary_min"},
				"max":     bson.M{"$max": "$salary_min"},
				"count":   bson.M{"$sum": 1},
			},
		},
		{
			"$sort": bson.M{"_id": 1},
		},
	}...)

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to aggregate salary by field %s: %w", field, err)
	}
	defer cursor.Close(ctx)

	var results []model.SalaryByCategory
	if err = cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode results: %w", err)
	}

	return results, nil
}

func (r *AnalyticsRepo) GetJobChangeData(ctx context.Context) ([]model.JobChangeData, error) {
	collection := r.db.Collection("salary_entries")

	filter := bson.M{
		"raises": bson.M{"$exists": true, "$ne": nil, "$not": bson.M{"$size": 0}},
	}

	projection := bson.M{
		"salary_min": 1,
		"salary_max": 1,
		"raises":     1,
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

func (r *AnalyticsRepo) GetOverallAverageSalary(ctx context.Context, filter *AnalyticsFilter) (float64, error) {
	collection := r.db.Collection("salary_entries")

	pipeline := []bson.M{}

	if filter != nil {
		matchStage := bson.M{"$match": r.buildFilterQuery(filter)}
		pipeline = append(pipeline, matchStage)
	}

	pipeline = append(pipeline, bson.M{
		"$group": bson.M{
			"_id":     nil,
			"average": bson.M{"$avg": "$salary_min"},
		},
	})

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, fmt.Errorf("failed to get overall average salary: %w", err)
	}
	defer cursor.Close(ctx)

	var result struct {
		Average float64 `bson:"average"`
	}
	if cursor.Next(ctx) {
		if err = cursor.Decode(&result); err != nil {
			return 0, fmt.Errorf("failed to decode average salary: %w", err)
		}
	}

	return result.Average, nil
}

func (r *AnalyticsRepo) GetAvailablePositions(ctx context.Context) ([]string, error) {
	collection := r.db.Collection("salary_entries")

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"position": bson.M{"$ne": "", "$exists": true},
			},
		},
		{
			"$group": bson.M{
				"_id": "$position",
			},
		},
		{
			"$sort": bson.M{"_id": 1},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to get available positions: %w", err)
	}
	defer cursor.Close(ctx)

	var results []struct {
		Position string `bson:"_id"`
	}
	if err = cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode positions: %w", err)
	}

	positions := make([]string, len(results))
	for i, result := range results {
		positions[i] = result.Position
	}

	return positions, nil
}

func (r *AnalyticsRepo) GetAvailableLevels(ctx context.Context) ([]string, error) {
	collection := r.db.Collection("salary_entries")

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"level": bson.M{"$ne": "", "$exists": true},
			},
		},
		{
			"$group": bson.M{
				"_id": "$level",
			},
		},
		{
			"$sort": bson.M{"_id": 1},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to get available levels: %w", err)
	}
	defer cursor.Close(ctx)

	var results []struct {
		Level string `bson:"_id"`
	}
	if err = cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode levels: %w", err)
	}

	levels := make([]string, len(results))
	for i, result := range results {
		levels[i] = result.Level
	}

	return levels, nil
}

type CombinedAnalyticsResult struct {
	TotalCount []struct {
		Total int64 `bson:"total"`
	} `bson:"totalCount"`
	OverallAverage []struct {
		Average float64 `bson:"average"`
	} `bson:"overallAverage"`
	ByPosition    []model.SalaryByCategory `bson:"byPosition"`
	ByLevel       []model.SalaryByCategory `bson:"byLevel"`
	ByExperience  []model.SalaryByCategory `bson:"byExperience"`
	ByCompany     []model.SalaryByCategory `bson:"byCompany"`
	ByCity        []model.SalaryByCategory `bson:"byCity"`
	ByCompanySize []model.SalaryByCategory `bson:"byCompanySize"`
	ByWorkType    []model.SalaryByCategory `bson:"byWorkType"`
	ByCurrency    []model.SalaryByCategory `bson:"byCurrency"`
}

func (r *AnalyticsRepo) GetCombinedAnalytics(ctx context.Context, filter *AnalyticsFilter) (*CombinedAnalyticsResult, error) {
	collection := r.db.Collection("salary_entries")

	baseMatch := bson.M{}
	if filter != nil {
		baseMatch = r.buildFilterQuery(filter)
	}

	pipeline := []bson.M{
		{"$match": baseMatch},
		{
			"$facet": bson.M{
				"totalCount": []bson.M{
					{"$count": "total"},
				},
				"overallAverage": []bson.M{
					{
						"$group": bson.M{
							"_id":     nil,
							"average": bson.M{"$avg": "$salary_min"},
						},
					},
				},
				"byPosition": []bson.M{
					{
						"$match": bson.M{
							"position": bson.M{"$ne": "", "$exists": true},
						},
					},
					{
						"$group": bson.M{
							"_id":     "$position",
							"average": bson.M{"$avg": "$salary_min"},
							"min":     bson.M{"$min": "$salary_min"},
							"max":     bson.M{"$max": "$salary_min"},
							"count":   bson.M{"$sum": 1},
						},
					},
					{"$sort": bson.M{"_id": 1}},
				},
				"byLevel": []bson.M{
					{
						"$match": bson.M{
							"level": bson.M{"$ne": "", "$exists": true},
						},
					},
					{
						"$group": bson.M{
							"_id":     "$level",
							"average": bson.M{"$avg": "$salary_min"},
							"min":     bson.M{"$min": "$salary_min"},
							"max":     bson.M{"$max": "$salary_min"},
							"count":   bson.M{"$sum": 1},
						},
					},
					{"$sort": bson.M{"_id": 1}},
				},
				"byExperience": []bson.M{
					{
						"$match": bson.M{
							"experience": bson.M{"$ne": "", "$exists": true},
						},
					},
					{
						"$group": bson.M{
							"_id":     "$experience",
							"average": bson.M{"$avg": "$salary_min"},
							"min":     bson.M{"$min": "$salary_min"},
							"max":     bson.M{"$max": "$salary_min"},
							"count":   bson.M{"$sum": 1},
						},
					},
					{"$sort": bson.M{"_id": 1}},
				},
				"byCompany": []bson.M{
					{
						"$match": bson.M{
							"company": bson.M{"$ne": "", "$exists": true},
						},
					},
					{
						"$group": bson.M{
							"_id":     "$company",
							"average": bson.M{"$avg": "$salary_min"},
							"min":     bson.M{"$min": "$salary_min"},
							"max":     bson.M{"$max": "$salary_min"},
							"count":   bson.M{"$sum": 1},
						},
					},
					{"$sort": bson.M{"_id": 1}},
				},
				"byCity": []bson.M{
					{
						"$match": bson.M{
							"city": bson.M{"$ne": "", "$exists": true},
						},
					},
					{
						"$group": bson.M{
							"_id":     "$city",
							"average": bson.M{"$avg": "$salary_min"},
							"min":     bson.M{"$min": "$salary_min"},
							"max":     bson.M{"$max": "$salary_min"},
							"count":   bson.M{"$sum": 1},
						},
					},
					{"$sort": bson.M{"_id": 1}},
				},
				"byCompanySize": []bson.M{
					{
						"$match": bson.M{
							"company_size": bson.M{"$ne": "", "$exists": true},
						},
					},
					{
						"$group": bson.M{
							"_id":     "$company_size",
							"average": bson.M{"$avg": "$salary_min"},
							"min":     bson.M{"$min": "$salary_min"},
							"max":     bson.M{"$max": "$salary_min"},
							"count":   bson.M{"$sum": 1},
						},
					},
					{"$sort": bson.M{"_id": 1}},
				},
				"byWorkType": []bson.M{
					{
						"$match": bson.M{
							"work_type": bson.M{"$ne": "", "$exists": true},
						},
					},
					{
						"$group": bson.M{
							"_id":     "$work_type",
							"average": bson.M{"$avg": "$salary_min"},
							"min":     bson.M{"$min": "$salary_min"},
							"max":     bson.M{"$max": "$salary_min"},
							"count":   bson.M{"$sum": 1},
						},
					},
					{"$sort": bson.M{"_id": 1}},
				},
				"byCurrency": []bson.M{
					{
						"$match": bson.M{
							"currency": bson.M{"$ne": "", "$exists": true},
						},
					},
					{
						"$group": bson.M{
							"_id":     "$currency",
							"average": bson.M{"$avg": "$salary_min"},
							"min":     bson.M{"$min": "$salary_min"},
							"max":     bson.M{"$max": "$salary_min"},
							"count":   bson.M{"$sum": 1},
						},
					},
					{"$sort": bson.M{"_id": 1}},
				},
			},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to execute combined analytics query: %w", err)
	}
	defer cursor.Close(ctx)

	var result CombinedAnalyticsResult
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("failed to decode combined analytics result: %w", err)
		}
	}

	return &result, nil
}

func (r *AnalyticsRepo) buildFilterQuery(filter *AnalyticsFilter) bson.M {
	query := bson.M{}

	if filter == nil {
		return query
	}

	if filter.Position != "" {
		query["position"] = filter.Position
	}

	if filter.Level != "" {
		query["level"] = filter.Level
	}

	if filter.Currency != "" {
		query["currency"] = filter.Currency
	}

	return query
}
