package model

import "time"

type Analytics struct {
	TotalEntries              int64              `json:"totalEntries"`
	AverageSalary             float64            `json:"averageSalary"`
	AverageSalaryByPosition   map[string]float64 `json:"averageSalaryByPosition"`
	MinSalaryByPosition       map[string]float64 `json:"minSalaryByPosition"`
	MaxSalaryByPosition       map[string]float64 `json:"maxSalaryByPosition"`
	AverageSalaryByLevel      map[string]float64 `json:"averageSalaryByLevel"`
	MinSalaryByLevel          map[string]float64 `json:"minSalaryByLevel"`
	MaxSalaryByLevel          map[string]float64 `json:"maxSalaryByLevel"`
	AverageSalaryByTech       map[string]float64 `json:"averageSalaryByTech"`
	MinSalaryByTech           map[string]float64 `json:"minSalaryByTech"`
	MaxSalaryByTech           map[string]float64 `json:"maxSalaryByTech"`
	AverageSalaryByExperience map[string]float64 `json:"averageSalaryByExperience"`
	MinSalaryByExperience     map[string]float64 `json:"minSalaryByExperience"`
	MaxSalaryByExperience     map[string]float64 `json:"maxSalaryByExperience"`
	AverageSalaryByCompany    map[string]float64 `json:"averageSalaryByCompany"`
	MinSalaryByCompany        map[string]float64 `json:"minSalaryByCompany"`
	MaxSalaryByCompany        map[string]float64 `json:"maxSalaryByCompany"`
	AverageSalaryByCity       map[string]float64 `json:"averageSalaryByCity"`
	MinSalaryByCity           map[string]float64 `json:"minSalaryByCity"`
	MaxSalaryByCity           map[string]float64 `json:"maxSalaryByCity"`
	AverageSalaryByCompanySize map[string]float64 `json:"averageSalaryByCompanySize"`
	MinSalaryByCompanySize     map[string]float64 `json:"minSalaryByCompanySize"`
	MaxSalaryByCompanySize     map[string]float64 `json:"maxSalaryByCompanySize"`
	AverageSalaryByWorkType   map[string]float64 `json:"averageSalaryByWorkType"`
	MinSalaryByWorkType       map[string]float64 `json:"minSalaryByWorkType"`
	MaxSalaryByWorkType       map[string]float64 `json:"maxSalaryByWorkType"`
	AverageSalaryByCurrency   map[string]float64 `json:"averageSalaryByCurrency"`
	MinSalaryByCurrency       map[string]float64 `json:"minSalaryByCurrency"`
	MaxSalaryByCurrency       map[string]float64 `json:"maxSalaryByCurrency"`
	TopPayingPositions        []ChartDataPoint   `json:"topPayingPositions"`
	TopPayingTechs            []ChartDataPoint   `json:"topPayingTechs"`
	SalaryRanges              []ChartDataPoint   `json:"salaryRanges"`
	LastUpdated               time.Time          `json:"lastUpdated"`
}

type ChartDataPoint struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
	Count int64  `json:"count,omitempty"`
}

type CareerAnalytics struct {
	JobChanges JobChangeAnalytics `json:"jobChanges"`
	Raises     RaiseAnalytics     `json:"raises"`
}

type JobChangeAnalytics struct {
	AverageSalaryIncrease  float64 `json:"averageSalaryIncrease"`
	PercentageWithIncrease float64 `json:"percentageWithIncrease"`
}

type RaiseAnalytics struct {
	AveragePerYear          float64 `json:"averagePerYear"`
	AveragePercentage       float64 `json:"averagePercentage"`
	MedianTimeBetweenRaises int     `json:"medianTimeBetweenRaises"`
}

type SalaryByCategory struct {
	Category string  `bson:"_id" json:"category"`
	Average  float64 `bson:"average" json:"average"`
	Min      float64 `bson:"min" json:"min"`
	Max      float64 `bson:"max" json:"max"`
	Count    int64   `bson:"count" json:"count"`
}

type SalaryByTech struct {
	Tech    string  `bson:"_id" json:"tech"`
	Average float64 `bson:"average" json:"average"`
	Min     float64 `bson:"min" json:"min"`
	Max     float64 `bson:"max" json:"max"`
	Count   int64   `bson:"count" json:"count"`
}

type JobChangeData struct {
	SalaryMin int64   `bson:"salary_min"`
	SalaryMax *int64  `bson:"salary_max"`
	Raises    []Raise `bson:"raises"`
}

type RaiseData struct {
	EntryID   string     `bson:"_id"`
	Raises    []Raise    `bson:"raises"`
	StartTime time.Time  `bson:"startTime"`
	EndTime   *time.Time `bson:"endTime"`
}