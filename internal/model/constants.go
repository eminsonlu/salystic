package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Level struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type Position struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type TechStack struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Category  string             `bson:"category" json:"category"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type Experience struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Range     string             `bson:"range" json:"range"`
	MinYears  int                `bson:"min_years" json:"min_years"`
	MaxYears  *int               `bson:"max_years,omitempty" json:"max_years,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type Company struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Sector    string             `bson:"sector" json:"sector"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type CompanySize struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Range     string             `bson:"range" json:"range"`
	MinSize   int                `bson:"min_size" json:"min_size"`
	MaxSize   *int               `bson:"max_size,omitempty" json:"max_size,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type WorkType struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type City struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Country   string             `bson:"country" json:"country"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type Currency struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Code      string             `bson:"code" json:"code"`
	Symbol    string             `bson:"symbol" json:"symbol"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

var (
	Levels = []string{"Junior", "Middle", "Senior"}
	
	ExperienceRanges = []string{
		"0 - 1 Yıl",
		"1 - 3 Yıl",
		"3 - 5 Yıl",
		"5 - 7 Yıl",
		"7 - 10 Yıl",
		"10 - 12 Yıl",
		"12 - 14 Yıl",
		"15 Yıl ve üzeri",
	}
	
	CompanySizes = []string{
		"1 - 5 Kişi",
		"6 - 10 Kişi",
		"11 - 20 Kişi",
		"21 - 50 Kişi",
		"51 - 100 Kişi",
		"101 - 249 Kişi",
		"250+",
	}
	
	WorkTypes = []string{
		"Ofis",
		"Remote",
		"Hibrit (Ofis + Remote)",
	}
	
	Genders = []string{"Erkek", "Kadın"}
	
	Currencies = []struct {
		Name   string
		Code   string
		Symbol string
	}{
		{"Türk Lirası", "TRY", "₺"},
		{"Dolar", "USD", "$"},
		{"Euro", "EUR", "€"},
		{"Sterlin", "GBP", "£"},
	}
	
	RaisePeriods = []int{1, 2, 3, 4}
)