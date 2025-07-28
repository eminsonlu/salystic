package repo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/eminsonlu/salystic/internal/model"
	"github.com/eminsonlu/salystic/pkg/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConstantsRepository interface {
	SeedConstants(ctx context.Context) error
	GetPositions(ctx context.Context) ([]string, error)
	GetLevels(ctx context.Context) ([]string, error)
	GetTechStacks(ctx context.Context) ([]string, error)
	GetExperiences(ctx context.Context) ([]string, error)
	GetCompanies(ctx context.Context) ([]string, error)
	GetCompanySizes(ctx context.Context) ([]string, error)
	GetWorkTypes(ctx context.Context) ([]string, error)
	GetCities(ctx context.Context) ([]string, error)
	GetCurrencies(ctx context.Context) ([]string, error)
}

type constantsRepository struct {
	db           *database.MongoDB
	positions    *mongo.Collection
	levels       *mongo.Collection
	techStacks   *mongo.Collection
	experiences  *mongo.Collection
	companies    *mongo.Collection
	companySizes *mongo.Collection
	workTypes    *mongo.Collection
	cities       *mongo.Collection
	currencies   *mongo.Collection
}

func NewConstantsRepository(db *database.MongoDB) ConstantsRepository {
	return &constantsRepository{
		db:           db,
		positions:    db.Database.Collection("positions"),
		levels:       db.Database.Collection("levels"),
		techStacks:   db.Database.Collection("tech_stacks"),
		experiences:  db.Database.Collection("experiences"),
		companies:    db.Database.Collection("companies"),
		companySizes: db.Database.Collection("company_sizes"),
		workTypes:    db.Database.Collection("work_types"),
		cities:       db.Database.Collection("cities"),
		currencies:   db.Database.Collection("currencies"),
	}
}

func (r *constantsRepository) SeedConstants(ctx context.Context) error {
	log.Println("Starting constants seeding...")

	if err := r.seedPositions(ctx); err != nil {
		return fmt.Errorf("failed to seed positions: %w", err)
	}

	if err := r.seedLevels(ctx); err != nil {
		return fmt.Errorf("failed to seed levels: %w", err)
	}

	if err := r.seedTechStacks(ctx); err != nil {
		return fmt.Errorf("failed to seed tech stacks: %w", err)
	}

	if err := r.seedExperiences(ctx); err != nil {
		return fmt.Errorf("failed to seed experiences: %w", err)
	}

	if err := r.seedCompanies(ctx); err != nil {
		return fmt.Errorf("failed to seed companies: %w", err)
	}

	if err := r.seedCompanySizes(ctx); err != nil {
		return fmt.Errorf("failed to seed company sizes: %w", err)
	}

	if err := r.seedWorkTypes(ctx); err != nil {
		return fmt.Errorf("failed to seed work types: %w", err)
	}

	if err := r.seedCities(ctx); err != nil {
		return fmt.Errorf("failed to seed cities: %w", err)
	}

	if err := r.seedCurrencies(ctx); err != nil {
		return fmt.Errorf("failed to seed currencies: %w", err)
	}

	log.Println("Constants seeding completed successfully")
	return nil
}

func (r *constantsRepository) seedPositions(ctx context.Context) error {
	positions := []model.Position{
		{Name: "AI Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Agile Coach", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Back-end Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Blockchain Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Business Analyst", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Business Intelligence", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "CTO", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Chief Data Officer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Cloud Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Computer Vision Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Consultant", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Cyber Security", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Data Analyst", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Data Architect", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Data Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Data Scientist", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Database Administrator (DBA)", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "DevOps Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Digital Transformation Executive", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Director of Software Development", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "ERP Developer & Consultant", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Embedded Software Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Front-end Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Full Stack Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Game Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "HR", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "IT", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "IT Manager", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Lead Product", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Machine Learning Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Mobile Application Developer (Android)", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Mobile Application Developer (Full Stack & Cross)", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Mobile Application Developer (iOS)", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Network Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Playable Ads Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "No-Code Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Platform Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Product Manager", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Product Owner", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Program Manager", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Project Manager", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "QA / Automation", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "QA / Manuel Test", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "QA Manager", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "R&D Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "R&D Manager", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "RPA Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Robotic Software Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "SAP / ABAP Developer & Consultant", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Site Reliability Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Software Architect", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Software Development Manager / Engineering Manager", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Software Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Solution Architect", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Solution Developer & Architect", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Support Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "System Admin", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "System Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Team / Tech Lead", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "UI/UX Designer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	insertedCount := 0
	for _, position := range positions {
		filter := bson.M{"name": position.Name}
		update := bson.M{
			"$setOnInsert": bson.M{
				"name":      position.Name,
				"createdAt": time.Now(),
				"updatedAt": time.Now(),
			},
		}

		upsert := true
		result, err := r.positions.UpdateOne(ctx, filter, update, &options.UpdateOptions{
			Upsert: &upsert,
		})
		if err != nil {
			return err
		}

		if result.UpsertedID != nil {
			insertedCount++
		}
	}

	if insertedCount > 0 {
		log.Printf("Seeded %d new positions", insertedCount)
	} else {
		log.Println("All positions already exist, no new positions added")
	}
	return nil
}

func (r *constantsRepository) seedLevels(ctx context.Context) error {
	levels := []model.Level{
		{Name: "Junior", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Middle", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Senior", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	insertedCount := 0
	for _, level := range levels {
		filter := bson.M{"name": level.Name}
		update := bson.M{
			"$setOnInsert": bson.M{
				"name":      level.Name,
				"createdAt": time.Now(),
				"updatedAt": time.Now(),
			},
		}

		upsert := true
		result, err := r.levels.UpdateOne(ctx, filter, update, &options.UpdateOptions{
			Upsert: &upsert,
		})
		if err != nil {
			return err
		}

		if result.UpsertedID != nil {
			insertedCount++
		}
	}

	if insertedCount > 0 {
		log.Printf("Seeded %d new levels", insertedCount)
	} else {
		log.Println("All levels already exist, no new levels added")
	}
	return nil
}

func (r *constantsRepository) seedTechStacks(ctx context.Context) error {
	techStacks := []model.TechStack{
		{Name: ".Net", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: ".Net Core", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: ".Net Framework", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Java", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Python", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Go", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "C#", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "PHP", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "NodeJS", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Spring Boot", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Django", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Laravel", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Rust", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Kotlin", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		{Name: "JavaScript", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "TypeScript", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "React", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Angular", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Vue", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Next.js", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "HTML", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "CSS", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "jQuery", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Svelte", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		{Name: "Swift", Category: "Mobile", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Kotlin", Category: "Mobile", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Flutter", Category: "Mobile", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "React Native", Category: "Mobile", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Xamarin", Category: "Mobile", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		{Name: "MySQL", Category: "Database", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "PostgreSQL", Category: "Database", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "MongoDB", Category: "Database", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Oracle", Category: "Database", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "MSSQL", Category: "Database", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "SQL", Category: "Database", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "PL/SQL", Category: "Database", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Elasticsearch", Category: "Database", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		{Name: "AWS", Category: "Cloud", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Azure", Category: "Cloud", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "GCP", Category: "Cloud", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Docker", Category: "DevOps", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Kubernetes", Category: "DevOps", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Jenkins", Category: "DevOps", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Terraform", Category: "DevOps", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Ansible", Category: "DevOps", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		{Name: "SAP", Category: "Enterprise", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "ABAP", Category: "Enterprise", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "SAP UI5", Category: "Enterprise", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Dynamics 365", Category: "Enterprise", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		{Name: "Selenium", Category: "Testing", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Cypress", Category: "Testing", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "JMeter", Category: "Testing", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Postman", Category: "Testing", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		{Name: "Power BI", Category: "Analytics", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Tableau", Category: "Analytics", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Apache Spark", Category: "BigData", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Kafka", Category: "BigData", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "R", Category: "Analytics", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "MATLAB", Category: "Analytics", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		{Name: "Linux", Category: "System", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Git", Category: "VersionControl", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Redis", Category: "Cache", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "RabbitMQ", Category: "MessageQueue", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	insertedCount := 0
	for _, techStack := range techStacks {
		filter := bson.M{"name": techStack.Name}
		update := bson.M{
			"$setOnInsert": bson.M{
				"name":      techStack.Name,
				"category":  techStack.Category,
				"createdAt": time.Now(),
				"updatedAt": time.Now(),
			},
		}

		upsert := true
		result, err := r.techStacks.UpdateOne(ctx, filter, update, &options.UpdateOptions{
			Upsert: &upsert,
		})
		if err != nil {
			return err
		}

		if result.UpsertedID != nil {
			insertedCount++
		}
	}

	if insertedCount > 0 {
		log.Printf("Seeded %d new tech stacks", insertedCount)
	} else {
		log.Println("All tech stacks already exist, no new tech stacks added")
	}
	return nil
}

func (r *constantsRepository) seedExperiences(ctx context.Context) error {
	experiences := []model.Experience{
		{Range: "0 - 1 Yıl", MinYears: 0, MaxYears: &[]int{1}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Range: "1 - 3 Yıl", MinYears: 1, MaxYears: &[]int{3}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Range: "3 - 5 Yıl", MinYears: 3, MaxYears: &[]int{5}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Range: "5 - 7 Yıl", MinYears: 5, MaxYears: &[]int{7}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Range: "7 - 10 Yıl", MinYears: 7, MaxYears: &[]int{10}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Range: "10 - 12 Yıl", MinYears: 10, MaxYears: &[]int{12}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Range: "12 - 14 Yıl", MinYears: 12, MaxYears: &[]int{14}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Range: "15 Yıl ve üzeri", MinYears: 15, MaxYears: nil, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	insertedCount := 0
	for _, experience := range experiences {
		filter := bson.M{"range": experience.Range}
		update := bson.M{
			"$setOnInsert": bson.M{
				"range":     experience.Range,
				"minYears":  experience.MinYears,
				"maxYears":  experience.MaxYears,
				"createdAt": time.Now(),
				"updatedAt": time.Now(),
			},
		}

		upsert := true
		result, err := r.experiences.UpdateOne(ctx, filter, update, &options.UpdateOptions{
			Upsert: &upsert,
		})
		if err != nil {
			return err
		}

		if result.UpsertedID != nil {
			insertedCount++
		}
	}

	if insertedCount > 0 {
		log.Printf("Seeded %d new experiences", insertedCount)
	} else {
		log.Println("All experiences already exist, no new experiences added")
	}
	return nil
}

func (r *constantsRepository) seedCompanies(ctx context.Context) error {
	companies := []model.Company{
		{Name: "AI (Yapay Zeka)", Sector: "AI (Yapay Zeka)", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Akıllı Şehirler", Sector: "Akıllı Şehirler", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Arge", Sector: "Arge", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Banka", Sector: "Banka", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Beyaz Eşya", Sector: "Beyaz Eşya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Borsa", Sector: "Borsa", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Cloud", Sector: "Cloud", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Coğrafi Bilgi sistemleri", Sector: "Coğrafi Bilgi sistemleri", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Danışmanlık", Sector: "Danışmanlık", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Data", Sector: "Data", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Dernek & Vakıf", Sector: "Dernek & Vakıf", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Dijital / Reklam Ajansı", Sector: "Dijital / Reklam Ajansı", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Diğer", Sector: "Diğer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "E-Ticaret", Sector: "E-Ticaret", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "ERP", Sector: "ERP", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "ERP & CRM", Sector: "ERP & CRM", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Elektronik", Sector: "Elektronik", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Emlak", Sector: "Emlak", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Endüstri", Sector: "Endüstri", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Enerji", Sector: "Enerji", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Entegratör", Sector: "Entegratör", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Eğitim", Sector: "Eğitim", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Eğlence", Sector: "Eğlence", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Fabrika", Sector: "Fabrika", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Fintech", Sector: "Fintech", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Freelancer", Sector: "Freelancer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Gıda & Yemek", Sector: "Gıda & Yemek", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Havacılık", Sector: "Havacılık", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Hizmet", Sector: "Hizmet", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Holding", Sector: "Holding", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Hukuk", Sector: "Hukuk", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "IoT", Sector: "IoT", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Kamu", Sector: "Kamu", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Kripto", Sector: "Kripto", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Kurumsal", Sector: "Kurumsal", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Lojistik", Sector: "Lojistik", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Makina", Sector: "Makina", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Medya", Sector: "Medya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Mobil Uygulama", Sector: "Mobil Uygulama", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Mobilya", Sector: "Mobilya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Muhasebe", Sector: "Muhasebe", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Otomasyon", Sector: "Otomasyon", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Otomotiv", Sector: "Otomotiv", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Outsource", Sector: "Outsource", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Oyun", Sector: "Oyun", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Pazarlama", Sector: "Pazarlama", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Perakende", Sector: "Perakende", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Reklam", Sector: "Reklam", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "SaaS", Sector: "SaaS", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Saas", Sector: "Saas", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Sanayi", Sector: "Sanayi", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Savunma Sanayi", Sector: "Savunma Sanayi", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Sağlık", Sector: "Sağlık", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Siber Güvenlik", Sector: "Siber Güvenlik", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Sigorta", Sector: "Sigorta", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Startup", Sector: "Startup", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Tarım", Sector: "Tarım", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Tekstil & Giyim & Moda", Sector: "Tekstil & Giyim & Moda", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Telekomünikasyon", Sector: "Telekomünikasyon", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Turizm", Sector: "Turizm", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Yazılım Evi & Danışmanlık", Sector: "Yazılım Evi & Danışmanlık", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Üniversite", Sector: "Üniversite", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Üretim", Sector: "Üretim", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "İK & Danışmanlık", Sector: "İK & Danışmanlık", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "İnşaat", Sector: "İnşaat", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Şans Oyunları & Bahis", Sector: "Şans Oyunları & Bahis", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	insertedCount := 0
	for _, company := range companies {
		filter := bson.M{"name": company.Name}
		update := bson.M{
			"$setOnInsert": bson.M{
				"name":      company.Name,
				"sector":    company.Sector,
				"createdAt": time.Now(),
				"updatedAt": time.Now(),
			},
		}

		upsert := true
		result, err := r.companies.UpdateOne(ctx, filter, update, &options.UpdateOptions{
			Upsert: &upsert,
		})
		if err != nil {
			return err
		}

		if result.UpsertedID != nil {
			insertedCount++
		}
	}
	log.Printf("Seeded %d companies", insertedCount)
	return nil
}

func (r *constantsRepository) seedCompanySizes(ctx context.Context) error {
	companySizes := []model.CompanySize{
		{Range: "1 - 5 Kişi", MinSize: 1, MaxSize: &[]int{5}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Range: "6 - 10 Kişi", MinSize: 6, MaxSize: &[]int{10}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Range: "11 - 20 Kişi", MinSize: 11, MaxSize: &[]int{20}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Range: "21 - 50 Kişi", MinSize: 21, MaxSize: &[]int{50}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Range: "51 - 100 Kişi", MinSize: 51, MaxSize: &[]int{100}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Range: "101 - 249 Kişi", MinSize: 101, MaxSize: &[]int{249}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Range: "250+", MinSize: 250, MaxSize: nil, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	insertedCount := 0
	for _, companySize := range companySizes {
		filter := bson.M{"range": companySize.Range}
		update := bson.M{
			"$setOnInsert": bson.M{
				"range":     companySize.Range,
				"minSize":   companySize.MinSize,
				"maxSize":   companySize.MaxSize,
				"createdAt": time.Now(),
				"updatedAt": time.Now(),
			},
		}

		upsert := true
		result, err := r.companySizes.UpdateOne(ctx, filter, update, &options.UpdateOptions{
			Upsert: &upsert,
		})
		if err != nil {
			return err
		}

		if result.UpsertedID != nil {
			insertedCount++
		}
	}
	log.Printf("Seeded %d company sizes", insertedCount)
	return nil
}

func (r *constantsRepository) seedWorkTypes(ctx context.Context) error {
	workTypes := []model.WorkType{
		{Name: "Ofis", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Remote", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Hibrit (Ofis + Remote)", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	insertedCount := 0
	for _, workType := range workTypes {
		filter := bson.M{"name": workType.Name}
		update := bson.M{
			"$setOnInsert": bson.M{
				"name":      workType.Name,
				"createdAt": time.Now(),
				"updatedAt": time.Now(),
			},
		}

		upsert := true
		result, err := r.workTypes.UpdateOne(ctx, filter, update, &options.UpdateOptions{
			Upsert: &upsert,
		})
		if err != nil {
			return err
		}

		if result.UpsertedID != nil {
			insertedCount++
		}
	}
	log.Printf("Seeded %d work types", insertedCount)
	return nil
}

func (r *constantsRepository) seedCities(ctx context.Context) error {
	cities := []model.City{
		{Name: "* ABD", Country: "ABD", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Almanya", Country: "Almanya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Arnavutluk", Country: "Arnavutluk", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Avrupa", Country: "Avrupa", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Avusturya", Country: "Avusturya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Azerbaycan", Country: "Azerbaycan", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Belçika", Country: "Belçika", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Birleşik Arap Emirlikleri", Country: "Birleşik Arap Emirlikleri", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Bosna Hersek", Country: "Bosna Hersek", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Dubai", Country: "Dubai", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Estonya", Country: "Estonya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Finlandiya", Country: "Finlandiya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Fransa", Country: "Fransa", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Güney Kore", Country: "Güney Kore", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Hollanda", Country: "Hollanda", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Hong Kong", Country: "Hong Kong", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Japonya", Country: "Japonya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* KKTC", Country: "KKTC", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Kanada", Country: "Kanada", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Karadağ", Country: "Karadağ", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Letonya", Country: "Letonya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Lüksemburg", Country: "Lüksemburg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Malta", Country: "Malta", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Meksika", Country: "Meksika", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Polonya", Country: "Polonya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Portekiz", Country: "Portekiz", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Slovakya", Country: "Slovakya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Tayland", Country: "Tayland", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Türkiye", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* Çekya", Country: "Çekya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* İngiltere", Country: "İngiltere", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* İrlanda", Country: "İrlanda", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* İspanya", Country: "İspanya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* İsveç", Country: "İsveç", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* İsviçre", Country: "İsviçre", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "* İtalya", Country: "İtalya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Adana", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Adıyaman", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Afyonkarahisar", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Aksaray", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Amasya", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Ankara", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Antalya", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Ardahan", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Artvin", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Aydın", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Ağrı", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Balıkesir", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Batman", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Bilecik", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Bitlis", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Bolu", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Bursa", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Denizli", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Diyarbakır", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Düzce", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Edirne", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Elazığ", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Erzincan", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Erzurum", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Eskişehir", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Gaziantep", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Giresun", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Hakkari", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Hatay", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Isparta", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Iğdır", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Kahramanmaraş", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Karabük", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Karaman", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Kars", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Kastamonu", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Kayseri", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Kilis", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Kocaeli", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Konya", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Kütahya", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Kırklareli", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Kırşehir", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Malatya", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Manisa", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Mardin", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Mersin", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Muğla", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Muş", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Nevşehir", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Niğde", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Ordu", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Osmaniye", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Rize", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Sakarya", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Samsun", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Sinop", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Sivas", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Tekirdağ", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Tokat", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Trabzon", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Uşak", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Van", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Yalova", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Yozgat", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Zonguldak", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Çanakkale", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Çorum", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "İstanbul", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "İzmir", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Şanlıurfa", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	insertedCount := 0
	for _, city := range cities {
		filter := bson.M{"name": city.Name, "country": city.Country}
		update := bson.M{
			"$setOnInsert": bson.M{
				"name":      city.Name,
				"country":   city.Country,
				"createdAt": time.Now(),
				"updatedAt": time.Now(),
			},
		}

		upsert := true
		result, err := r.cities.UpdateOne(ctx, filter, update, &options.UpdateOptions{
			Upsert: &upsert,
		})
		if err != nil {
			return err
		}

		if result.UpsertedID != nil {
			insertedCount++
		}
	}
	log.Printf("Seeded %d cities", insertedCount)
	return nil
}

func (r *constantsRepository) seedCurrencies(ctx context.Context) error {
	currencies := []model.Currency{
		model.Currency{Name: "Türk Lirası", Code: "TRY", Symbol: "₺", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Currency{Name: "Dolar", Code: "USD", Symbol: "$", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Currency{Name: "Euro", Code: "EUR", Symbol: "€", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Currency{Name: "Sterlin", Code: "GBP", Symbol: "£", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	insertedCount := 0
	for _, currency := range currencies {
		filter := bson.M{"code": currency.Code}
		update := bson.M{
			"$setOnInsert": bson.M{
				"name":      currency.Name,
				"code":      currency.Code,
				"symbol":    currency.Symbol,
				"createdAt": time.Now(),
				"updatedAt": time.Now(),
			},
		}

		upsert := true
		result, err := r.currencies.UpdateOne(ctx, filter, update, &options.UpdateOptions{
			Upsert: &upsert,
		})
		if err != nil {
			return err
		}

		if result.UpsertedID != nil {
			insertedCount++
		}
	}
	log.Printf("Seeded %d currencies", insertedCount)
	return nil
}

func (r *constantsRepository) GetPositions(ctx context.Context) ([]string, error) {
	cursor, err := r.positions.Find(ctx, bson.M{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find positions: %w", err)
	}
	defer cursor.Close(ctx)

	var names []string
	for cursor.Next(ctx) {
		var position struct {
			Name string `bson:"name"`
		}
		if err := cursor.Decode(&position); err != nil {
			return nil, fmt.Errorf("failed to decode position: %w", err)
		}
		names = append(names, position.Name)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return names, nil
}

func (r *constantsRepository) GetLevels(ctx context.Context) ([]string, error) {
	cursor, err := r.levels.Find(ctx, bson.M{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find levels: %w", err)
	}
	defer cursor.Close(ctx)

	var names []string
	for cursor.Next(ctx) {
		var level struct {
			Name string `bson:"name"`
		}
		if err := cursor.Decode(&level); err != nil {
			return nil, fmt.Errorf("failed to decode level: %w", err)
		}
		names = append(names, level.Name)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return names, nil
}

func (r *constantsRepository) GetTechStacks(ctx context.Context) ([]string, error) {
	cursor, err := r.techStacks.Find(ctx, bson.M{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find tech stacks: %w", err)
	}
	defer cursor.Close(ctx)

	var names []string
	for cursor.Next(ctx) {
		var techStack struct {
			Name string `bson:"name"`
		}
		if err := cursor.Decode(&techStack); err != nil {
			return nil, fmt.Errorf("failed to decode tech stack: %w", err)
		}
		names = append(names, techStack.Name)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return names, nil
}

func (r *constantsRepository) GetExperiences(ctx context.Context) ([]string, error) {
	cursor, err := r.experiences.Find(ctx, bson.M{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find experiences: %w", err)
	}
	defer cursor.Close(ctx)

	var names []string
	for cursor.Next(ctx) {
		var experience struct {
			Range string `bson:"range"`
		}
		if err := cursor.Decode(&experience); err != nil {
			return nil, fmt.Errorf("failed to decode experience: %w", err)
		}
		names = append(names, experience.Range)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return names, nil
}

func (r *constantsRepository) GetCompanies(ctx context.Context) ([]string, error) {
	cursor, err := r.companies.Find(ctx, bson.M{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find companies: %w", err)
	}
	defer cursor.Close(ctx)

	var names []string
	for cursor.Next(ctx) {
		var company struct {
			Name string `bson:"name"`
		}
		if err := cursor.Decode(&company); err != nil {
			return nil, fmt.Errorf("failed to decode company: %w", err)
		}
		names = append(names, company.Name)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return names, nil
}

func (r *constantsRepository) GetCompanySizes(ctx context.Context) ([]string, error) {
	cursor, err := r.companySizes.Find(ctx, bson.M{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find company sizes: %w", err)
	}
	defer cursor.Close(ctx)

	var names []string
	for cursor.Next(ctx) {
		var companySize struct {
			Range string `bson:"range"`
		}
		if err := cursor.Decode(&companySize); err != nil {
			return nil, fmt.Errorf("failed to decode company size: %w", err)
		}
		names = append(names, companySize.Range)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return names, nil
}

func (r *constantsRepository) GetWorkTypes(ctx context.Context) ([]string, error) {
	cursor, err := r.workTypes.Find(ctx, bson.M{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find work types: %w", err)
	}
	defer cursor.Close(ctx)

	var names []string
	for cursor.Next(ctx) {
		var workType struct {
			Name string `bson:"name"`
		}
		if err := cursor.Decode(&workType); err != nil {
			return nil, fmt.Errorf("failed to decode work type: %w", err)
		}
		names = append(names, workType.Name)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return names, nil
}

func (r *constantsRepository) GetCities(ctx context.Context) ([]string, error) {
	cursor, err := r.cities.Find(ctx, bson.M{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find cities: %w", err)
	}
	defer cursor.Close(ctx)

	var names []string
	for cursor.Next(ctx) {
		var city struct {
			Name string `bson:"name"`
		}
		if err := cursor.Decode(&city); err != nil {
			return nil, fmt.Errorf("failed to decode city: %w", err)
		}
		names = append(names, city.Name)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return names, nil
}

func (r *constantsRepository) GetCurrencies(ctx context.Context) ([]string, error) {
	cursor, err := r.currencies.Find(ctx, bson.M{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find currencies: %w", err)
	}
	defer cursor.Close(ctx)

	var names []string
	for cursor.Next(ctx) {
		var currency struct {
			Code string `bson:"code"`
		}
		if err := cursor.Decode(&currency); err != nil {
			return nil, fmt.Errorf("failed to decode currency: %w", err)
		}
		names = append(names, currency.Code)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return names, nil
}
