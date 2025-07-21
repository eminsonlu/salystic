package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SalaryEntry struct {
	ID                primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	UserID            primitive.ObjectID  `bson:"user_id" json:"user_id"`
	Country           string              `bson:"country" json:"country"`
	Currency          string              `bson:"currency" json:"currency"`
	Sector            string              `bson:"sector" json:"sector"`
	Job               string              `bson:"job" json:"job"`
	Title             string              `bson:"title" json:"title"`
	Salary            int64               `bson:"salary" json:"salary"`
	StartTime         time.Time           `bson:"start_time" json:"startTime"`
	EndTime           *time.Time          `bson:"end_time,omitempty" json:"endTime,omitempty"`
	PreviousJobSalary *int64              `bson:"previous_job_salary,omitempty" json:"previousJobSalary,omitempty"`
	Raises            []Raise             `bson:"raises,omitempty" json:"raises,omitempty"`
	CreatedAt         time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt         time.Time           `bson:"updated_at" json:"updated_at"`
}

type Raise struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	RaiseDate  time.Time          `bson:"raise_date" json:"raiseDate"`
	NewSalary  int64              `bson:"new_salary" json:"newSalary"`
	Percentage float64            `bson:"percentage" json:"percentage"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
}

type CreateSalaryEntryRequest struct {
	Country           string     `json:"country" validate:"required"`
	Currency          string     `json:"currency" validate:"required"`
	Sector            string     `json:"sector" validate:"required"`
	Job               string     `json:"job" validate:"required"`
	Title             string     `json:"title" validate:"required"`
	Salary            int64      `json:"salary" validate:"required,min=1"`
	StartTime         time.Time  `json:"startTime" validate:"required"`
	EndTime           *time.Time `json:"endTime,omitempty"`
	PreviousJobSalary *int64     `json:"previousJobSalary,omitempty"`
}

type UpdateSalaryEntryRequest struct {
	Country           *string    `json:"country,omitempty"`
	Currency          *string    `json:"currency,omitempty"`
	Sector            *string    `json:"sector,omitempty"`
	Job               *string    `json:"job,omitempty"`
	Title             *string    `json:"title,omitempty"`
	Salary            *int64     `json:"salary,omitempty"`
	StartTime         *time.Time `json:"startTime,omitempty"`
	EndTime           *time.Time `json:"endTime,omitempty"`
	PreviousJobSalary *int64     `json:"previousJobSalary,omitempty"`
}

type CreateRaiseRequest struct {
	RaiseDate  time.Time `json:"raiseDate" validate:"required"`
	NewSalary  int64     `json:"newSalary" validate:"required,min=1"`
	Percentage float64   `json:"percentage" validate:"required"`
}