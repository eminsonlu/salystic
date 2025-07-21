package service

import (
	"context"
	"fmt"
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

	entry := &model.SalaryEntry{
		UserID:            userObjID,
		Country:           req.Country,
		Currency:          req.Currency,
		Sector:            req.Sector,
		Job:               req.Job,
		Title:             req.Title,
		Salary:            req.Salary,
		StartTime:         req.StartTime,
		EndTime:           req.EndTime,
		PreviousJobSalary: req.PreviousJobSalary,
		Raises:            []model.Raise{},
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