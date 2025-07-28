package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SalaryEntry struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	Level       string             `bson:"level" json:"level"`
	Position    string             `bson:"position" json:"position"`
	TechStack   []string           `bson:"tech_stack" json:"tech_stack"`
	Experience  string             `bson:"experience" json:"experience"`
	Gender      string             `bson:"gender" json:"gender"`
	Company     string             `bson:"company" json:"company"`
	CompanySize string             `bson:"company_size" json:"company_size"`
	WorkType    string             `bson:"work_type" json:"work_type"`
	City        string             `bson:"city" json:"city"`
	Currency    string             `bson:"currency" json:"currency"`
	SalaryRange string             `bson:"salary_range" json:"salary_range"`
	SalaryMin   int64              `bson:"salary_min" json:"salary_min"`
	SalaryMax   *int64             `bson:"salary_max,omitempty" json:"salary_max,omitempty"`
	RaisePeriod int                `bson:"raise_period" json:"raise_period"`
	StartTime   time.Time          `bson:"start_time" json:"start_time"`
	EndTime     *time.Time         `bson:"end_time,omitempty" json:"end_time,omitempty"`
	Raises      []Raise            `bson:"raises,omitempty" json:"raises,omitempty"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type Raise struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	RaiseDate  time.Time          `bson:"raise_date" json:"raiseDate"`
	NewSalary  int64              `bson:"new_salary" json:"newSalary"`
	Percentage float64            `bson:"percentage" json:"percentage"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
}

type CreateSalaryEntryRequest struct {
	Level       string     `json:"level" validate:"required"`
	Position    string     `json:"position" validate:"required"`
	TechStack   []string   `json:"tech_stack" validate:"required"`
	Experience  string     `json:"experience" validate:"required"`
	Gender      string     `json:"gender" validate:"required"`
	Company     string     `json:"company" validate:"required"`
	CompanySize string     `json:"company_size" validate:"required"`
	WorkType    string     `json:"work_type" validate:"required"`
	City        string     `json:"city" validate:"required"`
	Currency    string     `json:"currency" validate:"required"`
	SalaryMin   int64      `json:"salary_min" validate:"required,min=1"`
	SalaryMax   *int64     `json:"salary_max,omitempty"`
	RaisePeriod int        `json:"raise_period" validate:"required,min=1,max=4"`
	StartTime   time.Time  `json:"start_time" validate:"required"`
	EndTime     *time.Time `json:"end_time,omitempty"`
}

type UpdateSalaryEntryRequest struct {
	Level       *string    `json:"level,omitempty"`
	Position    *string    `json:"position,omitempty"`
	TechStack   []string   `json:"tech_stack,omitempty"`
	Experience  *string    `json:"experience,omitempty"`
	Gender      *string    `json:"gender,omitempty"`
	Company     *string    `json:"company,omitempty"`
	CompanySize *string    `json:"company_size,omitempty"`
	WorkType    *string    `json:"work_type,omitempty"`
	City        *string    `json:"city,omitempty"`
	Currency    *string    `json:"currency,omitempty"`
	SalaryMin   *int64     `json:"salary_min,omitempty"`
	SalaryMax   *int64     `json:"salary_max,omitempty"`
	SalaryRange *string    `json:"salary_range,omitempty"`
	RaisePeriod *int       `json:"raise_period,omitempty"`
	StartTime   *time.Time `json:"start_time,omitempty"`
	EndTime     *time.Time `json:"end_time,omitempty"`
}

type CreateRaiseRequest struct {
	RaiseDate  time.Time `json:"raiseDate" validate:"required"`
	NewSalary  int64     `json:"newSalary" validate:"required,min=1"`
	Percentage float64   `json:"percentage" validate:"required"`
}
