package model

import "time"

type Analytics struct {
	TotalEntries         int64                          `json:"totalEntries"`
	AverageSalaryByJob   map[string]map[string]float64  `json:"averageSalaryByJob"`
	AverageSalaryBySector map[string]float64             `json:"averageSalaryBySector"`
	LastUpdated          time.Time                      `json:"lastUpdated"`
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
	AveragePerYear            float64 `json:"averagePerYear"`
	AveragePercentage         float64 `json:"averagePercentage"`
	MedianTimeBetweenRaises   int     `json:"medianTimeBetweenRaises"`
}

type SalaryByJobTitle struct {
	Job     string  `bson:"_id.job" json:"job"`
	Title   string  `bson:"_id.title" json:"title"`
	Average float64 `bson:"average" json:"average"`
	Count   int64   `bson:"count" json:"count"`
}

type SalaryBySector struct {
	Sector  string  `bson:"_id" json:"sector"`
	Average float64 `bson:"average" json:"average"`
	Count   int64   `bson:"count" json:"count"`
}

type JobChangeData struct {
	PreviousSalary int64 `bson:"previousJobSalary"`
	CurrentSalary  int64 `bson:"salary"`
}

type RaiseData struct {
	EntryID    string    `bson:"_id"`
	Raises     []Raise   `bson:"raises"`
	StartTime  time.Time `bson:"startTime"`
	EndTime    *time.Time `bson:"endTime"`
}