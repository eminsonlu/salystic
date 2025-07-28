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
	count, err := r.positions.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Positions already seeded, skipping...")
		return nil
	}

	positions := []interface{}{
		model.Position{Name: "AI Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Agile Coach", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Back-end Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Blockchain Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Business Analyst", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Business Intelligence", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "CTO", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Chief Data Officer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Cloud Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Computer Vision Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Consultant", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Cyber Security", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Data Analyst", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Data Architect", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Data Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Data Scientist", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Database Administrator (DBA)", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "DevOps Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Digital Transformation Executive", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Director of Software Development", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "ERP Developer & Consultant", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Embedded Software Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Front-end Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Full Stack Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Game Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "HR", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "IT", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "IT Manager", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Lead Product", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Machine Learning Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Mobile Application Developer (Android)", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Mobile Application Developer (Full Stack & Cross)", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Mobile Application Developer (iOS)", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Network Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Playable Ads Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "No-Code Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Platform Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Product Manager", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Product Owner", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Program Manager", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Project Manager", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "QA / Automation", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "QA / Manuel Test", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "QA Manager", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "R&D Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "R&D Manager", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "RPA Developer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Robotic Software Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "SAP / ABAP Developer & Consultant", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Site Reliability Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Software Architect", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Software Development Manager / Engineering Manager", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Software Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Solution Architect", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Solution Developer & Architect", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Support Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "System Admin", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "System Engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "Team / Tech Lead", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Position{Name: "UI/UX Designer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	_, err = r.positions.InsertMany(ctx, positions)
	if err != nil {
		return err
	}

	log.Printf("Seeded %d positions", len(positions))
	return nil
}

func (r *constantsRepository) seedLevels(ctx context.Context) error {
	count, err := r.levels.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Levels already seeded, skipping...")
		return nil
	}

	levels := []interface{}{
		model.Level{Name: "Junior", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Level{Name: "Middle", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Level{Name: "Senior", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	_, err = r.levels.InsertMany(ctx, levels)
	if err != nil {
		return err
	}

	log.Printf("Seeded %d levels", len(levels))
	return nil
}

func (r *constantsRepository) seedTechStacks(ctx context.Context) error {
	count, err := r.techStacks.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Tech stacks already seeded, skipping...")
		return nil
	}

	techStacks := []interface{}{
		model.TechStack{Name: ".Net", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: ".Net Core", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: ".Net Framework", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Java", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Python", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Go", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "C#", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "PHP", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "NodeJS", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Spring Boot", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Django", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Laravel", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Rust", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Kotlin", Category: "Backend", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		model.TechStack{Name: "JavaScript", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "TypeScript", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "React", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Angular", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Vue", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Next.js", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "HTML", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "CSS", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "jQuery", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Svelte", Category: "Frontend", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		model.TechStack{Name: "Swift", Category: "Mobile", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Kotlin", Category: "Mobile", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Flutter", Category: "Mobile", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "React Native", Category: "Mobile", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Xamarin", Category: "Mobile", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		model.TechStack{Name: "MySQL", Category: "Database", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "PostgreSQL", Category: "Database", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "MongoDB", Category: "Database", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Oracle", Category: "Database", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "MSSQL", Category: "Database", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "SQL", Category: "Database", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "PL/SQL", Category: "Database", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Elasticsearch", Category: "Database", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		model.TechStack{Name: "AWS", Category: "Cloud", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Azure", Category: "Cloud", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "GCP", Category: "Cloud", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Docker", Category: "DevOps", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Kubernetes", Category: "DevOps", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Jenkins", Category: "DevOps", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Terraform", Category: "DevOps", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Ansible", Category: "DevOps", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		model.TechStack{Name: "SAP", Category: "Enterprise", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "ABAP", Category: "Enterprise", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "SAP UI5", Category: "Enterprise", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Dynamics 365", Category: "Enterprise", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		model.TechStack{Name: "Selenium", Category: "Testing", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Cypress", Category: "Testing", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "JMeter", Category: "Testing", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Postman", Category: "Testing", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		model.TechStack{Name: "Power BI", Category: "Analytics", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Tableau", Category: "Analytics", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Apache Spark", Category: "BigData", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Kafka", Category: "BigData", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "R", Category: "Analytics", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "MATLAB", Category: "Analytics", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		model.TechStack{Name: "Linux", Category: "System", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Git", Category: "VersionControl", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "Redis", Category: "Cache", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.TechStack{Name: "RabbitMQ", Category: "MessageQueue", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	_, err = r.techStacks.InsertMany(ctx, techStacks)
	if err != nil {
		return err
	}

	log.Printf("Seeded %d tech stacks", len(techStacks))
	return nil
}

func (r *constantsRepository) seedExperiences(ctx context.Context) error {
	count, err := r.experiences.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Experiences already seeded, skipping...")
		return nil
	}

	experiences := []interface{}{
		model.Experience{Range: "0 - 1 Yıl", MinYears: 0, MaxYears: &[]int{1}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Experience{Range: "1 - 3 Yıl", MinYears: 1, MaxYears: &[]int{3}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Experience{Range: "3 - 5 Yıl", MinYears: 3, MaxYears: &[]int{5}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Experience{Range: "5 - 7 Yıl", MinYears: 5, MaxYears: &[]int{7}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Experience{Range: "7 - 10 Yıl", MinYears: 7, MaxYears: &[]int{10}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Experience{Range: "10 - 12 Yıl", MinYears: 10, MaxYears: &[]int{12}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Experience{Range: "12 - 14 Yıl", MinYears: 12, MaxYears: &[]int{14}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Experience{Range: "15 Yıl ve üzeri", MinYears: 15, MaxYears: nil, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	_, err = r.experiences.InsertMany(ctx, experiences)
	if err != nil {
		return err
	}

	log.Printf("Seeded %d experiences", len(experiences))
	return nil
}

func (r *constantsRepository) seedCompanies(ctx context.Context) error {
	count, err := r.companies.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Companies already seeded, skipping...")
		return nil
	}

	companies := []interface{}{
		model.Company{Name: "AI (Yapay Zeka)", Sector: "AI (Yapay Zeka)", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Akıllı Şehirler", Sector: "Akıllı Şehirler", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Arge", Sector: "Arge", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Banka", Sector: "Banka", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Beyaz Eşya", Sector: "Beyaz Eşya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Borsa", Sector: "Borsa", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Cloud", Sector: "Cloud", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Coğrafi Bilgi sistemleri", Sector: "Coğrafi Bilgi sistemleri", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Danışmanlık", Sector: "Danışmanlık", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Data", Sector: "Data", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Dernek & Vakıf", Sector: "Dernek & Vakıf", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Dijital / Reklam Ajansı", Sector: "Dijital / Reklam Ajansı", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Diğer", Sector: "Diğer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "E-Ticaret", Sector: "E-Ticaret", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "ERP", Sector: "ERP", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "ERP & CRM", Sector: "ERP & CRM", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Elektronik", Sector: "Elektronik", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Emlak", Sector: "Emlak", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Endüstri", Sector: "Endüstri", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Enerji", Sector: "Enerji", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Entegratör", Sector: "Entegratör", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Eğitim", Sector: "Eğitim", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Eğlence", Sector: "Eğlence", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Fabrika", Sector: "Fabrika", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Fintech", Sector: "Fintech", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Freelancer", Sector: "Freelancer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Gıda & Yemek", Sector: "Gıda & Yemek", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Havacılık", Sector: "Havacılık", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Hizmet", Sector: "Hizmet", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Holding", Sector: "Holding", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Hukuk", Sector: "Hukuk", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "IoT", Sector: "IoT", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Kamu", Sector: "Kamu", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Kripto", Sector: "Kripto", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Kurumsal", Sector: "Kurumsal", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Lojistik", Sector: "Lojistik", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Makina", Sector: "Makina", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Medya", Sector: "Medya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Mobil Uygulama", Sector: "Mobil Uygulama", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Mobilya", Sector: "Mobilya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Muhasebe", Sector: "Muhasebe", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Otomasyon", Sector: "Otomasyon", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Otomotiv", Sector: "Otomotiv", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Outsource", Sector: "Outsource", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Oyun", Sector: "Oyun", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Pazarlama", Sector: "Pazarlama", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Perakende", Sector: "Perakende", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Reklam", Sector: "Reklam", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "SaaS", Sector: "SaaS", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Saas", Sector: "Saas", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Sanayi", Sector: "Sanayi", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Savunma Sanayi", Sector: "Savunma Sanayi", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Sağlık", Sector: "Sağlık", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Siber Güvenlik", Sector: "Siber Güvenlik", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Sigorta", Sector: "Sigorta", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Startup", Sector: "Startup", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Tarım", Sector: "Tarım", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Tekstil & Giyim & Moda", Sector: "Tekstil & Giyim & Moda", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Telekomünikasyon", Sector: "Telekomünikasyon", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Turizm", Sector: "Turizm", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Yazılım Evi & Danışmanlık", Sector: "Yazılım Evi & Danışmanlık", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Üniversite", Sector: "Üniversite", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Üretim", Sector: "Üretim", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "İK & Danışmanlık", Sector: "İK & Danışmanlık", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "İnşaat", Sector: "İnşaat", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Company{Name: "Şans Oyunları & Bahis", Sector: "Şans Oyunları & Bahis", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	_, err = r.companies.InsertMany(ctx, companies)
	if err != nil {
		return err
	}

	log.Printf("Seeded %d companies", len(companies))
	return nil
}

func (r *constantsRepository) seedCompanySizes(ctx context.Context) error {
	count, err := r.companySizes.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Company sizes already seeded, skipping...")
		return nil
	}

	companySizes := []interface{}{
		model.CompanySize{Range: "1 - 5 Kişi", MinSize: 1, MaxSize: &[]int{5}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.CompanySize{Range: "6 - 10 Kişi", MinSize: 6, MaxSize: &[]int{10}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.CompanySize{Range: "11 - 20 Kişi", MinSize: 11, MaxSize: &[]int{20}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.CompanySize{Range: "21 - 50 Kişi", MinSize: 21, MaxSize: &[]int{50}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.CompanySize{Range: "51 - 100 Kişi", MinSize: 51, MaxSize: &[]int{100}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.CompanySize{Range: "101 - 249 Kişi", MinSize: 101, MaxSize: &[]int{249}[0], CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.CompanySize{Range: "250+", MinSize: 250, MaxSize: nil, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	_, err = r.companySizes.InsertMany(ctx, companySizes)
	if err != nil {
		return err
	}

	log.Printf("Seeded %d company sizes", len(companySizes))
	return nil
}

func (r *constantsRepository) seedWorkTypes(ctx context.Context) error {
	count, err := r.workTypes.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Work types already seeded, skipping...")
		return nil
	}

	workTypes := []interface{}{
		model.WorkType{Name: "Ofis", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.WorkType{Name: "Remote", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.WorkType{Name: "Hibrit (Ofis + Remote)", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	_, err = r.workTypes.InsertMany(ctx, workTypes)
	if err != nil {
		return err
	}

	log.Printf("Seeded %d work types", len(workTypes))
	return nil
}

func (r *constantsRepository) seedCities(ctx context.Context) error {
	count, err := r.cities.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Cities already seeded, skipping...")
		return nil
	}

	cities := []interface{}{
		model.City{Name: "* ABD", Country: "ABD", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Almanya", Country: "Almanya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Arnavutluk", Country: "Arnavutluk", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Avrupa", Country: "Avrupa", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Avusturya", Country: "Avusturya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Azerbaycan", Country: "Azerbaycan", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Belçika", Country: "Belçika", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Birleşik Arap Emirlikleri", Country: "Birleşik Arap Emirlikleri", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Bosna Hersek", Country: "Bosna Hersek", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Dubai", Country: "Dubai", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Estonya", Country: "Estonya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Finlandiya", Country: "Finlandiya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Fransa", Country: "Fransa", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Güney Kore", Country: "Güney Kore", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Hollanda", Country: "Hollanda", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Hong Kong", Country: "Hong Kong", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Japonya", Country: "Japonya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* KKTC", Country: "KKTC", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Kanada", Country: "Kanada", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Karadağ", Country: "Karadağ", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Letonya", Country: "Letonya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Lüksemburg", Country: "Lüksemburg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Malta", Country: "Malta", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Meksika", Country: "Meksika", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Polonya", Country: "Polonya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Portekiz", Country: "Portekiz", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Slovakya", Country: "Slovakya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Tayland", Country: "Tayland", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Türkiye", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* Çekya", Country: "Çekya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* İngiltere", Country: "İngiltere", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* İrlanda", Country: "İrlanda", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* İspanya", Country: "İspanya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* İsveç", Country: "İsveç", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* İsviçre", Country: "İsviçre", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "* İtalya", Country: "İtalya", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		model.City{Name: "Adana", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Adıyaman", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Afyonkarahisar", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Aksaray", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Amasya", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Ankara", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Antalya", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Ardahan", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Artvin", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Aydın", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Ağrı", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Balıkesir", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Batman", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Bilecik", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Bitlis", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Bolu", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Bursa", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Denizli", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Diyarbakır", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Düzce", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Edirne", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Elazığ", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Erzincan", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Erzurum", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Eskişehir", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Gaziantep", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Giresun", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Hakkari", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Hatay", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Isparta", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Iğdır", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Kahramanmaraş", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Karabük", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Karaman", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Kars", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Kastamonu", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Kayseri", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Kilis", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Kocaeli", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Konya", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Kütahya", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Kırklareli", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Kırşehir", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Malatya", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Manisa", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Mardin", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Mersin", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Muğla", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Muş", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Nevşehir", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Niğde", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Ordu", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Osmaniye", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Rize", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Sakarya", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Samsun", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Sinop", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Sivas", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Tekirdağ", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Tokat", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Trabzon", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Uşak", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Van", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Yalova", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Yozgat", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Zonguldak", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Çanakkale", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Çorum", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "İstanbul", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "İzmir", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.City{Name: "Şanlıurfa", Country: "Türkiye", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	_, err = r.cities.InsertMany(ctx, cities)
	if err != nil {
		return err
	}

	log.Printf("Seeded %d cities", len(cities))
	return nil
}

func (r *constantsRepository) seedCurrencies(ctx context.Context) error {
	count, err := r.currencies.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Currencies already seeded, skipping...")
		return nil
	}

	currencies := []interface{}{
		model.Currency{Name: "Türk Lirası", Code: "TRY", Symbol: "₺", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Currency{Name: "Dolar", Code: "USD", Symbol: "$", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Currency{Name: "Euro", Code: "EUR", Symbol: "€", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		model.Currency{Name: "Sterlin", Code: "GBP", Symbol: "£", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	_, err = r.currencies.InsertMany(ctx, currencies)
	if err != nil {
		return err
	}

	log.Printf("Seeded %d currencies", len(currencies))
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
