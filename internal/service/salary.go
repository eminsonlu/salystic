package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/eminsonlu/salystic/internal/model"
	"github.com/eminsonlu/salystic/internal/repo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SalaryEntryService interface {
	CreateEntry(ctx context.Context, userID string, req *model.CreateSalaryEntryRequest) (*model.SalaryEntry, error)
	GetEntry(ctx context.Context, userID, entryID string) (*model.SalaryEntry, error)
	GetUserEntries(ctx context.Context, userID string) ([]*model.SalaryEntry, error)
	UpdateEntry(ctx context.Context, userID, entryID string, req *model.UpdateSalaryEntryRequest) (*model.SalaryEntry, error)
	DeleteEntry(ctx context.Context, userID, entryID string) error
	AddRaise(ctx context.Context, userID, entryID string, req *model.CreateRaiseRequest) error
	GetRaises(ctx context.Context, userID, entryID string) ([]model.Raise, error)
}

type salaryEntryService struct {
	salaryRepo repo.SalaryEntryRepository
}

func NewSalaryEntryService(salaryRepo repo.SalaryEntryRepository) SalaryEntryService {
	return &salaryEntryService{
		salaryRepo: salaryRepo,
	}
}

func (s *salaryEntryService) CreateEntry(ctx context.Context, userID string, req *model.CreateSalaryEntryRequest) (*model.SalaryEntry, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	var salaryRange string
	if req.SalaryMax != nil {
		salaryRange = fmt.Sprintf("%d - %d", req.SalaryMin, *req.SalaryMax)
	} else {
		salaryRange = fmt.Sprintf("%d+", req.SalaryMin)
	}

	entry := &model.SalaryEntry{
		UserID:      userObjID,
		Level:       req.Level,
		Position:    req.Position,
		TechStack:   req.TechStack,
		Experience:  req.Experience,
		Gender:      req.Gender,
		Company:     req.Company,
		CompanySize: req.CompanySize,
		WorkType:    req.WorkType,
		City:        req.City,
		Currency:    req.Currency,
		SalaryRange: salaryRange,
		SalaryMin:   req.SalaryMin,
		SalaryMax:   req.SalaryMax,
		RaisePeriod: req.RaisePeriod,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Raises:      []model.Raise{},
	}

	if err := s.salaryRepo.Create(ctx, entry); err != nil {
		return nil, fmt.Errorf("failed to create salary entry: %w", err)
	}

	return entry, nil
}

func (s *salaryEntryService) GetEntry(ctx context.Context, userID, entryID string) (*model.SalaryEntry, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	entryObjID, err := primitive.ObjectIDFromHex(entryID)
	if err != nil {
		return nil, fmt.Errorf("invalid entry ID: %w", err)
	}

	entry, err := s.salaryRepo.GetByID(ctx, entryObjID, userObjID)
	if err != nil {
		return nil, fmt.Errorf("failed to get salary entry: %w", err)
	}

	if entry == nil {
		return nil, fmt.Errorf("salary entry not found")
	}

	return entry, nil
}

func (s *salaryEntryService) GetUserEntries(ctx context.Context, userID string) ([]*model.SalaryEntry, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	entries, err := s.salaryRepo.GetByUserID(ctx, userObjID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user entries: %w", err)
	}

	return entries, nil
}

func (s *salaryEntryService) UpdateEntry(ctx context.Context, userID, entryID string, req *model.UpdateSalaryEntryRequest) (*model.SalaryEntry, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	entryObjID, err := primitive.ObjectIDFromHex(entryID)
	if err != nil {
		return nil, fmt.Errorf("invalid entry ID: %w", err)
	}

	entry, err := s.salaryRepo.Update(ctx, entryObjID, userObjID, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update salary entry: %w", err)
	}

	if entry == nil {
		return nil, fmt.Errorf("salary entry not found")
	}

	return entry, nil
}

func (s *salaryEntryService) DeleteEntry(ctx context.Context, userID, entryID string) error {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	entryObjID, err := primitive.ObjectIDFromHex(entryID)
	if err != nil {
		return fmt.Errorf("invalid entry ID: %w", err)
	}

	if err := s.salaryRepo.Delete(ctx, entryObjID, userObjID); err != nil {
		return fmt.Errorf("failed to delete salary entry: %w", err)
	}

	return nil
}

func (s *salaryEntryService) AddRaise(ctx context.Context, userID, entryID string, req *model.CreateRaiseRequest) error {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	entryObjID, err := primitive.ObjectIDFromHex(entryID)
	if err != nil {
		return fmt.Errorf("invalid entry ID: %w", err)
	}

	raise := &model.Raise{
		RaiseDate:  req.RaiseDate,
		NewSalary:  req.NewSalary,
		Percentage: req.Percentage,
	}

	if err := s.salaryRepo.AddRaise(ctx, entryObjID, userObjID, raise); err != nil {
		return fmt.Errorf("failed to add raise: %w", err)
	}

	return nil
}

func (s *salaryEntryService) GetRaises(ctx context.Context, userID, entryID string) ([]model.Raise, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	entryObjID, err := primitive.ObjectIDFromHex(entryID)
	if err != nil {
		return nil, fmt.Errorf("invalid entry ID: %w", err)
	}

	raises, err := s.salaryRepo.GetRaises(ctx, entryObjID, userObjID)
	if err != nil {
		return nil, fmt.Errorf("failed to get raises: %w", err)
	}

	return raises, nil
}

func parseSalaryRange(salaryRange string) (int64, *int64) {
	salaryRange = strings.TrimSpace(salaryRange)
	if salaryRange == "" {
		return 0, nil
	}

	cleaned := strings.ReplaceAll(salaryRange, ".", "")

	if strings.HasSuffix(cleaned, "+") {
		numStr := strings.TrimSuffix(cleaned, "+")
		if min, err := strconv.ParseInt(strings.TrimSpace(numStr), 10, 64); err == nil {
			return min, nil
		}
		return 0, nil
	}

	parts := strings.Split(cleaned, "-")
	if len(parts) == 2 {
		minStr := strings.TrimSpace(parts[0])
		maxStr := strings.TrimSpace(parts[1])

		min, err1 := strconv.ParseInt(minStr, 10, 64)
		max, err2 := strconv.ParseInt(maxStr, 10, 64)
		if err1 == nil && err2 == nil {
			return min, &max
		}
		return 0, nil
	}

	if val, err := strconv.ParseInt(cleaned, 10, 64); err == nil {
		return val, &val
	}

	return 0, nil
}
