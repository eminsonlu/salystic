package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/eminsonlu/salystic/internal/model"
	"github.com/eminsonlu/salystic/internal/repo"
	"github.com/jellydator/ttlcache/v3"
	"golang.org/x/sync/errgroup"
)

type AnalyticsService struct {
	analyticsRepo *repo.AnalyticsRepo
	cache         *ttlcache.Cache[string, *model.Analytics]
	cacheCareer   *ttlcache.Cache[string, *model.CareerAnalytics]
}

func NewAnalyticsService(analyticsRepo *repo.AnalyticsRepo) *AnalyticsService {
	return NewAnalyticsServiceWithTTL(analyticsRepo, 10*time.Minute)
}

func NewAnalyticsServiceWithTTL(analyticsRepo *repo.AnalyticsRepo, cacheTTL time.Duration) *AnalyticsService {
	cache := ttlcache.New(
		ttlcache.WithTTL[string, *model.Analytics](cacheTTL),
		ttlcache.WithCapacity[string, *model.Analytics](100),
	)
	go cache.Start()

	cacheCareer := ttlcache.New(
		ttlcache.WithTTL[string, *model.CareerAnalytics](cacheTTL),
		ttlcache.WithCapacity[string, *model.CareerAnalytics](100),
	)
	go cacheCareer.Start()

	return &AnalyticsService{
		analyticsRepo: analyticsRepo,
		cache:         cache,
		cacheCareer:   cacheCareer,
	}
}

func (s *AnalyticsService) generateCacheKey(level, position, currency string) string {
	key := fmt.Sprintf("analytics:%s:%s:%s", level, position, currency)
	hash := md5.Sum([]byte(key))
	return fmt.Sprintf("%x", hash)
}

func (s *AnalyticsService) GetGeneralAnalytics(ctx context.Context, level, position, currency string) (*model.Analytics, error) {
	cacheKey := s.generateCacheKey(level, position, currency)

	if cached := s.cache.Get(cacheKey); cached != nil {
		return cached.Value(), nil
	}

	filter := &repo.AnalyticsFilter{
		Level:    level,
		Position: position,
		Currency: currency,
	}

	g, ctx := errgroup.WithContext(ctx)

	var (
		combinedResult *repo.CombinedAnalyticsResult
		salaryByTech   []model.SalaryByTech
	)

	g.Go(func() error {
		var err error
		combinedResult, err = s.analyticsRepo.GetCombinedAnalytics(ctx, filter)
		return err
	})

	g.Go(func() error {
		var err error
		salaryByTech, err = s.analyticsRepo.GetAverageSalaryByTech(ctx, filter)
		return err
	})

	if err := g.Wait(); err != nil {
		return nil, fmt.Errorf("failed to fetch analytics data: %w", err)
	}

	var totalEntries int64
	if len(combinedResult.TotalCount) > 0 {
		totalEntries = combinedResult.TotalCount[0].Total
	}

	var averageSalary float64
	if len(combinedResult.OverallAverage) > 0 {
		averageSalary = combinedResult.OverallAverage[0].Average
	}

	salaryByPosition := combinedResult.ByPosition
	salaryByLevel := combinedResult.ByLevel
	salaryByExperience := combinedResult.ByExperience
	salaryByCompany := combinedResult.ByCompany
	salaryByCity := combinedResult.ByCity
	salaryByCompanySize := combinedResult.ByCompanySize
	salaryByWorkType := combinedResult.ByWorkType
	salaryByCurrency := combinedResult.ByCurrency

	averageByPositionMap := make(map[string]float64)
	minByPositionMap := make(map[string]float64)
	maxByPositionMap := make(map[string]float64)
	for _, item := range salaryByPosition {
		averageByPositionMap[item.Category] = item.Average
		minByPositionMap[item.Category] = item.Min
		maxByPositionMap[item.Category] = item.Max
	}

	averageByLevelMap := make(map[string]float64)
	minByLevelMap := make(map[string]float64)
	maxByLevelMap := make(map[string]float64)
	for _, item := range salaryByLevel {
		averageByLevelMap[item.Category] = item.Average
		minByLevelMap[item.Category] = item.Min
		maxByLevelMap[item.Category] = item.Max
	}

	averageByTechMap := make(map[string]float64)
	minByTechMap := make(map[string]float64)
	maxByTechMap := make(map[string]float64)
	for _, item := range salaryByTech {
		averageByTechMap[item.Tech] = item.Average
		minByTechMap[item.Tech] = item.Min
		maxByTechMap[item.Tech] = item.Max
	}

	averageByExperienceMap := make(map[string]float64)
	minByExperienceMap := make(map[string]float64)
	maxByExperienceMap := make(map[string]float64)
	for _, item := range salaryByExperience {
		averageByExperienceMap[item.Category] = item.Average
		minByExperienceMap[item.Category] = item.Min
		maxByExperienceMap[item.Category] = item.Max
	}

	averageByCompanyMap := make(map[string]float64)
	minByCompanyMap := make(map[string]float64)
	maxByCompanyMap := make(map[string]float64)
	for _, item := range salaryByCompany {
		averageByCompanyMap[item.Category] = item.Average
		minByCompanyMap[item.Category] = item.Min
		maxByCompanyMap[item.Category] = item.Max
	}

	averageByCityMap := make(map[string]float64)
	minByCityMap := make(map[string]float64)
	maxByCityMap := make(map[string]float64)
	for _, item := range salaryByCity {
		averageByCityMap[item.Category] = item.Average
		minByCityMap[item.Category] = item.Min
		maxByCityMap[item.Category] = item.Max
	}

	averageByCompanySizeMap := make(map[string]float64)
	minByCompanySizeMap := make(map[string]float64)
	maxByCompanySizeMap := make(map[string]float64)
	for _, item := range salaryByCompanySize {
		averageByCompanySizeMap[item.Category] = item.Average
		minByCompanySizeMap[item.Category] = item.Min
		maxByCompanySizeMap[item.Category] = item.Max
	}

	averageByWorkTypeMap := make(map[string]float64)
	minByWorkTypeMap := make(map[string]float64)
	maxByWorkTypeMap := make(map[string]float64)
	for _, item := range salaryByWorkType {
		averageByWorkTypeMap[item.Category] = item.Average
		minByWorkTypeMap[item.Category] = item.Min
		maxByWorkTypeMap[item.Category] = item.Max
	}

	averageByCurrencyMap := make(map[string]float64)
	minByCurrencyMap := make(map[string]float64)
	maxByCurrencyMap := make(map[string]float64)
	for _, item := range salaryByCurrency {
		averageByCurrencyMap[item.Category] = item.Average
		minByCurrencyMap[item.Category] = item.Min
		maxByCurrencyMap[item.Category] = item.Max
	}

	topPayingPositions := s.buildTopPayingChart(salaryByPosition, 10)
	topPayingTechs := s.buildTopPayingTechChart(salaryByTech, 10)
	salaryRanges := s.buildSalaryRanges(salaryByPosition, currency)

	analytics := &model.Analytics{
		TotalEntries:               totalEntries,
		AverageSalary:              averageSalary,
		AverageSalaryByPosition:    averageByPositionMap,
		MinSalaryByPosition:        minByPositionMap,
		MaxSalaryByPosition:        maxByPositionMap,
		AverageSalaryByLevel:       averageByLevelMap,
		MinSalaryByLevel:           minByLevelMap,
		MaxSalaryByLevel:           maxByLevelMap,
		AverageSalaryByTech:        averageByTechMap,
		MinSalaryByTech:            minByTechMap,
		MaxSalaryByTech:            maxByTechMap,
		AverageSalaryByExperience:  averageByExperienceMap,
		MinSalaryByExperience:      minByExperienceMap,
		MaxSalaryByExperience:      maxByExperienceMap,
		AverageSalaryByCompany:     averageByCompanyMap,
		MinSalaryByCompany:         minByCompanyMap,
		MaxSalaryByCompany:         maxByCompanyMap,
		AverageSalaryByCity:        averageByCityMap,
		MinSalaryByCity:            minByCityMap,
		MaxSalaryByCity:            maxByCityMap,
		AverageSalaryByCompanySize: averageByCompanySizeMap,
		MinSalaryByCompanySize:     minByCompanySizeMap,
		MaxSalaryByCompanySize:     maxByCompanySizeMap,
		AverageSalaryByWorkType:    averageByWorkTypeMap,
		MinSalaryByWorkType:        minByWorkTypeMap,
		MaxSalaryByWorkType:        maxByWorkTypeMap,
		AverageSalaryByCurrency:    averageByCurrencyMap,
		MinSalaryByCurrency:        minByCurrencyMap,
		MaxSalaryByCurrency:        maxByCurrencyMap,
		TopPayingPositions:         topPayingPositions,
		TopPayingTechs:             topPayingTechs,
		SalaryRanges:               salaryRanges,
		LastUpdated:                time.Now(),
	}

	s.cache.Set(cacheKey, analytics, ttlcache.DefaultTTL)
	return analytics, nil
}

func (s *AnalyticsService) GetCareerAnalytics(ctx context.Context) (*model.CareerAnalytics, error) {
	cacheKey := "career_analytics"

	if cached := s.cacheCareer.Get(cacheKey); cached != nil {
		return cached.Value(), nil
	}

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

	s.cacheCareer.Set(cacheKey, &model.CareerAnalytics{
		JobChanges: jobChangeAnalytics,
		Raises:     raiseAnalytics,
	}, ttlcache.DefaultTTL)

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
		if len(item.Raises) > 0 {
			salaryMax := item.SalaryMax
			if salaryMax == nil {
				salaryMax = &item.SalaryMin
			}
			initialSalary := (item.SalaryMin + *salaryMax) / 2
			latestRaise := item.Raises[len(item.Raises)-1]
			if latestRaise.NewSalary > initialSalary {
				withIncrease++
				increase := float64(latestRaise.NewSalary-initialSalary) / float64(initialSalary) * 100
				totalIncrease += increase
			}
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
	var timeBetweenRaises []int
	var totalYears float64

	for _, entry := range data {
		if len(entry.Raises) == 0 {
			continue
		}

		endTime := time.Now()
		if entry.EndTime != nil {
			endTime = *entry.EndTime
		}
		years := endTime.Sub(entry.StartTime).Hours() / (24 * 365)
		totalYears += years

		totalRaises += len(entry.Raises)

		for i, raise := range entry.Raises {
			totalPercentage += raise.Percentage

			if i > 0 {
				prevRaise := entry.Raises[i-1]
				months := int(raise.CreatedAt.Sub(prevRaise.CreatedAt).Hours() / (24 * 30))
				timeBetweenRaises = append(timeBetweenRaises, months)
			} else {
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
		AveragePerYear:          roundToTwoDecimals(averagePerYear),
		AveragePercentage:       roundToTwoDecimals(averagePercentage),
		MedianTimeBetweenRaises: medianTime,
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

func (s *AnalyticsService) GetAvailablePositions(ctx context.Context) ([]string, error) {
	return s.analyticsRepo.GetAvailablePositions(ctx)
}

func (s *AnalyticsService) GetAvailableLevels(ctx context.Context) ([]string, error) {
	return s.analyticsRepo.GetAvailableLevels(ctx)
}

func roundToTwoDecimals(value float64) float64 {
	return float64(int(value*100+0.5)) / 100
}

func (s *AnalyticsService) buildTopPayingChart(data []model.SalaryByCategory, limit int) []model.ChartDataPoint {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Average > data[j].Average
	})

	if len(data) > limit {
		data = data[:limit]
	}

	result := make([]model.ChartDataPoint, len(data))
	for i, item := range data {
		result[i] = model.ChartDataPoint{
			Name:  item.Category,
			Value: int(math.Round(item.Average)),
			Count: item.Count,
		}
	}
	return result
}

func (s *AnalyticsService) buildTopPayingTechChart(data []model.SalaryByTech, limit int) []model.ChartDataPoint {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Average > data[j].Average
	})

	if len(data) > limit {
		data = data[:limit]
	}

	result := make([]model.ChartDataPoint, len(data))
	for i, item := range data {
		result[i] = model.ChartDataPoint{
			Name:  item.Tech,
			Value: int(math.Round(item.Average)),
			Count: item.Count,
		}
	}
	return result
}

func (s *AnalyticsService) buildSalaryRanges(positionData []model.SalaryByCategory, currency string) []model.ChartDataPoint {
	var ranges []struct {
		min   float64
		max   float64
		label string
	}

	if currency == "TRY" {
		ranges = []struct {
			min   float64
			max   float64
			label string
		}{
			{0, 50000, "Under ₺50K"},
			{50000, 100000, "₺50K - ₺100K"},
			{100000, 200000, "₺100K - ₺200K"},
			{200000, 300000, "₺200K - ₺300K"},
			{300000, 500000, "₺300K - ₺500K"},
			{500000, math.Inf(1), "₺500K+"},
		}
	} else {
		ranges = []struct {
			min   float64
			max   float64
			label string
		}{
			{0, 50000, "Under $50K"},
			{50000, 75000, "$50K - $75K"},
			{75000, 100000, "$75K - $100K"},
			{100000, 150000, "$100K - $150K"},
			{150000, 200000, "$150K - $200K"},
			{200000, math.Inf(1), "$200K+"},
		}
	}

	result := make([]model.ChartDataPoint, len(ranges))
	for i, rangeItem := range ranges {
		count := int64(0)
		for _, pos := range positionData {
			if pos.Average >= rangeItem.min && pos.Average < rangeItem.max {
				count += pos.Count
			}
		}
		result[i] = model.ChartDataPoint{
			Name:  rangeItem.label,
			Value: int(count),
		}
	}
	return result
}
