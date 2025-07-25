package repo

import (
	"context"
	"fmt"
	"github.com/eminsonlu/salystic/internal/model"
	"github.com/eminsonlu/salystic/pkg/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SalaryEntryRepository interface {
	Create(ctx context.Context, entry *model.SalaryEntry) error
	GetByID(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (*model.SalaryEntry, error)
	GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]*model.SalaryEntry, error)
	Update(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID, update *model.UpdateSalaryEntryRequest) (*model.SalaryEntry, error)
	Delete(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) error
	AddRaise(ctx context.Context, entryID primitive.ObjectID, userID primitive.ObjectID, raise *model.Raise) error
	GetRaises(ctx context.Context, entryID primitive.ObjectID, userID primitive.ObjectID) ([]model.Raise, error)
}

type salaryEntryRepository struct {
	collection *mongo.Collection
}

func NewSalaryEntryRepository(db *database.MongoDB) SalaryEntryRepository {
	return &salaryEntryRepository{
		collection: db.Database.Collection("salary_entries"),
	}
}

func (r *salaryEntryRepository) Create(ctx context.Context, entry *model.SalaryEntry) error {
	entry.CreatedAt = time.Now()
	entry.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, entry)
	if err != nil {
		return fmt.Errorf("failed to create salary entry: %w", err)
	}

	entry.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *salaryEntryRepository) GetByID(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (*model.SalaryEntry, error) {
	var entry model.SalaryEntry
	filter := bson.M{"_id": id, "user_id": userID}
	
	err := r.collection.FindOne(ctx, filter).Decode(&entry)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get salary entry: %w", err)
	}
	return &entry, nil
}

func (r *salaryEntryRepository) GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]*model.SalaryEntry, error) {
	filter := bson.M{"user_id": userID}
	opts := options.Find().SetSort(bson.M{"created_at": -1})
	
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find salary entries: %w", err)
	}
	defer cursor.Close(ctx)

	var entries []*model.SalaryEntry
	if err = cursor.All(ctx, &entries); err != nil {
		return nil, fmt.Errorf("failed to decode salary entries: %w", err)
	}

	return entries, nil
}

func (r *salaryEntryRepository) Update(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID, update *model.UpdateSalaryEntryRequest) (*model.SalaryEntry, error) {
	filter := bson.M{"_id": id, "user_id": userID}
	
	updateDoc := bson.M{"$set": bson.M{"updated_at": time.Now()}}
	setDoc := updateDoc["$set"].(bson.M)
	
	if update.Level != nil {
		setDoc["level"] = *update.Level
	}
	if update.Position != nil {
		setDoc["position"] = *update.Position
	}
	if update.TechStack != nil {
		setDoc["tech_stack"] = update.TechStack
	}
	if update.Experience != nil {
		setDoc["experience"] = *update.Experience
	}
	if update.Gender != nil {
		setDoc["gender"] = *update.Gender
	}
	if update.Company != nil {
		setDoc["company"] = *update.Company
	}
	if update.CompanySize != nil {
		setDoc["company_size"] = *update.CompanySize
	}
	if update.WorkType != nil {
		setDoc["work_type"] = *update.WorkType
	}
	if update.City != nil {
		setDoc["city"] = *update.City
	}
	if update.Currency != nil {
		setDoc["currency"] = *update.Currency
	}
	if update.SalaryMin != nil {
		setDoc["salary_min"] = *update.SalaryMin
	}
	if update.SalaryMax != nil {
		setDoc["salary_max"] = *update.SalaryMax
	}
	if update.RaisePeriod != nil {
		setDoc["raise_period"] = *update.RaisePeriod
	}
	if update.StartTime != nil {
		setDoc["start_time"] = *update.StartTime
	}
	if update.EndTime != nil {
		setDoc["end_time"] = *update.EndTime
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var entry model.SalaryEntry
	err := r.collection.FindOneAndUpdate(ctx, filter, updateDoc, opts).Decode(&entry)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to update salary entry: %w", err)
	}

	return &entry, nil
}

func (r *salaryEntryRepository) Delete(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) error {
	filter := bson.M{"_id": id, "user_id": userID}
	
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete salary entry: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("salary entry not found")
	}

	return nil
}

func (r *salaryEntryRepository) AddRaise(ctx context.Context, entryID primitive.ObjectID, userID primitive.ObjectID, raise *model.Raise) error {
	filter := bson.M{"_id": entryID, "user_id": userID}
	
	raise.ID = primitive.NewObjectID()
	raise.CreatedAt = time.Now()

	update := bson.M{
		"$push": bson.M{"raises": raise},
		"$set":  bson.M{"updated_at": time.Now()},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to add raise: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("salary entry not found")
	}

	return nil
}

func (r *salaryEntryRepository) GetRaises(ctx context.Context, entryID primitive.ObjectID, userID primitive.ObjectID) ([]model.Raise, error) {
	filter := bson.M{"_id": entryID, "user_id": userID}
	projection := bson.M{"raises": 1}
	
	var result struct {
		Raises []model.Raise `bson:"raises"`
	}

	err := r.collection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("salary entry not found")
		}
		return nil, fmt.Errorf("failed to get raises: %w", err)
	}

	if result.Raises == nil {
		result.Raises = []model.Raise{}
	}

	return result.Raises, nil
}
