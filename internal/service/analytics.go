package service

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/eminsonlu/salystic/internal/model"
	"github.com/eminsonlu/salystic/internal/repo"
)

type AnalyticsService struct {
	analyticsRepo *repo.AnalyticsRepo
}

func NewAnalyticsService(analyticsRepo *repo.AnalyticsRepo) *AnalyticsService {
	return &AnalyticsService{
		analyticsRepo: analyticsRepo,
	}
}

func (s *AnalyticsService) GetGeneralAnalytics(ctx context.Context) (*model.Analytics, error) {
	totalEntries, err := s.analyticsRepo.GetTotalEntries(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get total entries: %w", err)
	}
	
	salaryByJob, err := s.analyticsRepo.GetAverageSalaryByJob(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get salary by job: %w", err)
	}
	
	salaryBySector, err := s.analyticsRepo.GetAverageSalaryBySector(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get salary by sector: %w", err)
	}
	
	// Transform job data into nested map
	averageByJob := make(map[string]map[string]float64)
	for _, item := range salaryByJob {
		if _, exists := averageByJob[item.Job]; !exists {
			averageByJob[item.Job] = make(map[string]float64)
		}
		averageByJob[item.Job][item.Title] = item.Average
	}
	
	// Transform sector data into map
	averageBySector := make(map[string]float64)
	for _, item := range salaryBySector {
		averageBySector[item.Sector] = item.Average
	}
	
	return &model.Analytics{
		TotalEntries:          totalEntries,
		AverageSalaryByJob:    averageByJob,
		AverageSalaryBySector: averageBySector,
		LastUpdated:           time.Now(),
	}, nil
}

func (s *AnalyticsService) GetCareerAnalytics(ctx context.Context) (*model.CareerAnalytics, error) {
	jobChangeData, err := s.analyticsRepo.GetJobChangeData(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get job change data: %w", err)
	}
	
	raiseData, err := s.analyticsRepo.GetRaiseData(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get raise data: %w", err)
	}
	
	jobChangeAnalytics := s.calculateJobChangeAnalytics(jobChangeData)
	raiseAnalytics := s.calculateRaiseAnalytics(raiseData)
	
	return &model.CareerAnalytics{
		JobChanges: jobChangeAnalytics,
		Raises:     raiseAnalytics,
	}, nil
}

func (s *AnalyticsService) calculateJobChangeAnalytics(data []model.JobChangeData) model.JobChangeAnalytics {
	if len(data) == 0 {
		return model.JobChangeAnalytics{}
	}
	
	var totalIncrease float64
	var withIncrease int
	
	for _, item := range data {
		if item.CurrentSalary > item.PreviousSalary {
			withIncrease++
			increase := float64(item.CurrentSalary-item.PreviousSalary) / float64(item.PreviousSalary) * 100
			totalIncrease += increase
		}
	}
	
	averageIncrease := float64(0)
	if withIncrease > 0 {
		averageIncrease = totalIncrease / float64(withIncrease)
	}
	
	percentageWithIncrease := float64(withIncrease) / float64(len(data)) * 100
	
	return model.JobChangeAnalytics{
		AverageSalaryIncrease:  roundToTwoDecimals(averageIncrease),
		PercentageWithIncrease: roundToTwoDecimals(percentageWithIncrease),
	}
}

func (s *AnalyticsService) calculateRaiseAnalytics(data []model.RaiseData) model.RaiseAnalytics {
	if len(data) == 0 {
		return model.RaiseAnalytics{}
	}
	
	var totalRaises int
	var totalPercentage float64
	var timeBetweenRaises []int // in months
	var totalYears float64
	
	for _, entry := range data {
		if len(entry.Raises) == 0 {
			continue
		}
		
		// Calculate entry duration
		endTime := time.Now()
		if entry.EndTime != nil {
			endTime = *entry.EndTime
		}
		years := endTime.Sub(entry.StartTime).Hours() / (24 * 365)
		totalYears += years
		
		totalRaises += len(entry.Raises)
		
		// Calculate percentage and time between raises
		for i, raise := range entry.Raises {
			totalPercentage += raise.Percentage
			
			if i > 0 {
				prevRaise := entry.Raises[i-1]
				months := int(raise.CreatedAt.Sub(prevRaise.CreatedAt).Hours() / (24 * 30))
				timeBetweenRaises = append(timeBetweenRaises, months)
			} else {
				// Time from job start to first raise
				months := int(raise.CreatedAt.Sub(entry.StartTime).Hours() / (24 * 30))
				timeBetweenRaises = append(timeBetweenRaises, months)
			}
		}
	}
	
	averagePerYear := float64(0)
	if totalYears > 0 {
		averagePerYear = float64(totalRaises) / totalYears
	}
	
	averagePercentage := float64(0)
	if totalRaises > 0 {
		averagePercentage = totalPercentage / float64(totalRaises)
	}
	
	medianTime := calculateMedian(timeBetweenRaises)
	
	return model.RaiseAnalytics{
		AveragePerYear:            roundToTwoDecimals(averagePerYear),
		AveragePercentage:         roundToTwoDecimals(averagePercentage),
		MedianTimeBetweenRaises:   medianTime,
	}
}

func calculateMedian(values []int) int {
	if len(values) == 0 {
		return 0
	}
	
	sort.Ints(values)
	n := len(values)
	
	if n%2 == 1 {
		return values[n/2]
	}
	
	return (values[n/2-1] + values[n/2]) / 2
}

func roundToTwoDecimals(value float64) float64 {
	return float64(int(value*100+0.5)) / 100
}