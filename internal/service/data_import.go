package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/eminsonlu/salystic/internal/model"
	"github.com/eminsonlu/salystic/internal/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataImportService struct {
	salaryRepo repo.SalaryEntryRepository
}

func NewDataImportService(salaryRepo repo.SalaryEntryRepository) *DataImportService {
	return &DataImportService{
		salaryRepo: salaryRepo,
	}
}

type ImportDataEntry struct {
	Level       string `json:"level"`
	Position    string `json:"position"`
	TechStack   string `json:"tech_stack"`
	Experience  string `json:"experience"`
	Gender      string `json:"gender"`
	Company     string `json:"company"`
	CompanySize string `json:"company_size"`
	WorkType    string `json:"work_type"`
	City        string `json:"city"`
	Currency    string `json:"currency"`
	Salary      string `json:"salary"`
	RaisePeriod string `json:"raise_period"`
}

func (s *DataImportService) ImportFromJSON(ctx context.Context, filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var entries []ImportDataEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	for i, entry := range entries {
		salaryEntry, err := s.convertToSalaryEntry(entry)
		if err != nil {
			return fmt.Errorf("failed to convert entry %d: %w", i, err)
		}

		if err := s.salaryRepo.Create(ctx, salaryEntry); err != nil {
			return fmt.Errorf("failed to save entry %d: %w", i, err)
		}
	}

	return nil
}

func (s *DataImportService) AnalyzeTechStacks(ctx context.Context, filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var entries []ImportDataEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	techCounts := make(map[string]int)
	invalidTechs := make(map[string]int)

	for _, entry := range entries {
		if entry.TechStack == "" {
			continue
		}

		parsedTechs := s.parseTechStack(entry.TechStack)

		for _, tech := range parsedTechs {
			techCounts[tech]++
		}

		if len(parsedTechs) == 0 && entry.TechStack != "" {
			invalidTechs[entry.TechStack]++
		}
	}

	fmt.Printf("\n=== TECH STACK ANALYSIS ===\n")
	fmt.Printf("Total entries: %d\n", len(entries))
	fmt.Printf("Valid technologies found: %d\n", len(techCounts))
	fmt.Printf("Invalid/unparsed tech strings: %d\n\n", len(invalidTechs))

	fmt.Println("Top 20 Valid Technologies:")
	type techCount struct {
		tech  string
		count int
	}

	var sortedTechs []techCount
	for tech, count := range techCounts {
		sortedTechs = append(sortedTechs, techCount{tech, count})
	}

	for i := 0; i < len(sortedTechs)-1; i++ {
		for j := 0; j < len(sortedTechs)-i-1; j++ {
			if sortedTechs[j].count < sortedTechs[j+1].count {
				sortedTechs[j], sortedTechs[j+1] = sortedTechs[j+1], sortedTechs[j]
			}
		}
	}

	for i, tc := range sortedTechs {
		if i >= 20 {
			break
		}
		fmt.Printf("%2d. %-20s: %d entries\n", i+1, tc.tech, tc.count)
	}

	if len(invalidTechs) > 0 {
		fmt.Println("\nTop 20 Invalid/Unparsed Tech Strings:")
		var sortedInvalid []techCount
		for tech, count := range invalidTechs {
			sortedInvalid = append(sortedInvalid, techCount{tech, count})
		}

		for i := 0; i < len(sortedInvalid)-1; i++ {
			for j := 0; j < len(sortedInvalid)-i-1; j++ {
				if sortedInvalid[j].count < sortedInvalid[j+1].count {
					sortedInvalid[j], sortedInvalid[j+1] = sortedInvalid[j+1], sortedInvalid[j]
				}
			}
		}

		for i, tc := range sortedInvalid {
			if i >= 20 {
				break
			}
			fmt.Printf("%2d. %-40s: %d entries\n", i+1, tc.tech, tc.count)
		}
	}

	fmt.Println("\n=== END ANALYSIS ===")
	fmt.Println()
	return nil
}

func (s *DataImportService) convertToSalaryEntry(entry ImportDataEntry) (*model.SalaryEntry, error) {
	techStack := s.parseTechStack(entry.TechStack)

	salaryMin, salaryMax, err := s.parseSalaryRange(entry.Salary)
	if err != nil {
		return nil, fmt.Errorf("failed to parse salary range: %w", err)
	}

	raisePeriod, err := strconv.Atoi(entry.RaisePeriod)
	if err != nil {
		return nil, fmt.Errorf("failed to parse raise period: %w", err)
	}

	currency := s.extractCurrencyCode(entry.Currency)

	now := time.Now()

	return &model.SalaryEntry{
		ID:          primitive.NewObjectID(),
		UserID:      primitive.NewObjectID(),
		Level:       entry.Level,
		Position:    entry.Position,
		TechStack:   techStack,
		Experience:  entry.Experience,
		Gender:      entry.Gender,
		Company:     entry.Company,
		CompanySize: entry.CompanySize,
		WorkType:    s.normalizeWorkType(entry.WorkType),
		City:        entry.City,
		Currency:    currency,
		SalaryRange: entry.Salary,
		SalaryMin:   salaryMin,
		SalaryMax:   salaryMax,
		RaisePeriod: raisePeriod,
		StartTime:   now,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

var validTechStacks = map[string]bool{
	".Net": true, ".Net Core": true, ".Net Framework": true, "Java": true, "Python": true,
	"Go": true, "C#": true, "PHP": true, "Php": true, "NodeJS": true, "Spring Boot": true,
	"Django": true, "Laravel": true, "Rust": true, "Kotlin": true, "Ruby": true,
	"C / C++": true, "C++": true, "C": true, "Scala": true, "Perl": true, "Delphi": true,
	"Pascal": true, "Cobol": true, "Fortran": true, "Groovy": true, "Clojure": true,
	"F#": true, "VB.NET": true, "VB": true, "Visual Basic": true,

	"JavaScript": true, "JavaScript | Html | Css": true, "TypeScript": true, "React": true,
	"Angular": true, "Vue": true, "Vue.js": true, "Next.js": true, "HTML": true, "CSS": true,
	"jQuery": true, "Svelte": true, "Ember JS": true, "Html": true, "Css": true,
	"Bootstrap": true, "SASS": true, "SCSS": true, "Less": true, "Webpack": true,

	"Swift": true, "Objective C": true, "Flutter": true, "React Native": true,
	"Xamarin": true, "Android": true, "iOS": true, "Ionic": true, "Cordova": true,

	"MySQL": true, "PostgreSQL": true, "MongoDB": true, "Oracle": true, "MSSQL": true,
	"SQL": true, "Sql": true, "PL/SQL": true, "Elasticsearch": true, "Redis": true,
	"SQLite": true, "MariaDB": true, "Cassandra": true, "Neo4j": true, "CouchDB": true,
	"DynamoDB": true, "InfluxDB": true, "Firebase": true, "Supabase": true,

	"AWS": true, "Azure": true, "GCP": true, "Docker": true, "Kubernetes": true,
	"Jenkins": true, "Terraform": true, "Ansible": true, "Chef": true, "Puppet": true,
	"GitLab": true, "GitHub": true, "BitBucket": true, "CircleCI": true, "Travis CI": true,
	"Heroku": true, "Vercel": true, "Netlify": true,

	"SAP": true, "ABAP": true, "SAP UI5": true, "Dynamics 365": true, "Salesforce": true,
	"SharePoint": true, "Oracle ERP": true, "Workday": true, "ServiceNow": true,

	"Selenium": true, "Cypress": true, "JMeter": true, "Postman": true, "SOAPUI": true,
	"Jest": true, "Mocha": true, "Chai": true, "Jasmine": true, "Karma": true,
	"TestNG": true, "JUnit": true, "NUnit": true, "Cucumber": true,

	"Unity": true, "Unreal": true, "Unreal Engine": true, "Godot": true, "GameMaker": true,

	"Power BI": true, "Tableau": true, "Apache Spark": true, "Kafka": true,
	"R": true, "MATLAB": true, "SAS": true, "SPSS": true, "Jupyter": true,
	"Pandas": true, "NumPy": true, "TensorFlow": true, "PyTorch": true, "Keras": true,
	"Scikit-learn": true, "OpenCV": true,

	"WordPress": true, "Drupal": true, "Joomla": true, "Magento": true, "Shopify": true,
	"WooCommerce": true, "PrestaShop": true,

	"UiPath": true, "Blue Prism": true, "Automation Anywhere": true, "Power Automate": true,

	"Linux": true, "Git": true, "RabbitMQ": true, "Windows Server": true, "Ubuntu": true,
	"CentOS": true, "RedHat": true, "Debian": true, "MacOS": true, "Unix": true,
	"Jira": true, "Apex": true, "Bash": true, "PowerShell": true, "Shell": true,
	"Active Directory": true, "LDAP": true, "Apache": true, "Nginx": true, "IIS": true,
	"Tomcat": true, "JBoss": true, "WebLogic": true, "WebSphere": true,

	"Hadoop": true, "Hive": true, "Pig": true, "Spark": true, "Storm": true,
	"Flink": true, "NiFi": true, "Talend": true, "Informatica": true, "SSIS": true,
	"Pentaho": true, "Databricks": true, "Snowflake": true, "BigQuery": true,
	"Redshift": true, "Vertica": true, "Teradata": true,

	"ActiveMQ": true, "Apache Pulsar": true, "NATS": true, "ZeroMQ": true,
	"Slack API": true, "Microsoft Teams": true, "WebRTC": true, "Socket.IO": true,

	"Photoshop": true, "Illustrator": true, "Sketch": true, "Figma": true, "Adobe XD": true,
	"InVision": true, "Zeplin": true, "After Effects": true, "Premiere Pro": true,
}

func (s *DataImportService) parseTechStack(techStackStr string) []string {
	if techStackStr == "" {
		return []string{}
	}

	parts := strings.Split(techStackStr, ",")
	var result []string

	for _, part := range parts {
		tech := strings.TrimSpace(part)
		if tech == "" {
			continue
		}

		if validTechStacks[tech] {
			result = append(result, tech)
			continue
		}

		for validTech := range validTechStacks {
			if strings.EqualFold(tech, validTech) {
				result = append(result, validTech)
				continue
			}
		}

		extractedTechs := s.extractValidTechs(tech)
		result = append(result, extractedTechs...)
	}

	return s.removeDuplicates(result)
}

func (s *DataImportService) extractValidTechs(techStr string) []string {
	var result []string
	originalStr := techStr
	techStr = strings.ToLower(techStr)

	words := strings.Fields(techStr)
	for _, word := range words {
		word = strings.TrimSpace(word)
		if word == "" {
			continue
		}

		for validTech := range validTechStacks {
			if strings.EqualFold(word, validTech) {
				result = append(result, validTech)
				break
			}
		}
	}

	if len(result) > 0 {
		return result
	}

	for validTech := range validTechStacks {
		lowerTech := strings.ToLower(validTech)

		if strings.Contains(techStr, lowerTech) {
			if s.isValidTechMatch(originalStr, lowerTech) {
				result = append(result, validTech)
			}
		}
	}

	return result
}

func (s *DataImportService) isValidTechMatch(fullStr, tech string) bool {
	if len(tech) <= 2 {
		words := strings.Fields(fullStr)
		for _, word := range words {
			if strings.ToLower(word) == tech {
				return true
			}
		}
		return false
	}

	return true
}

func (s *DataImportService) removeDuplicates(techs []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, tech := range techs {
		if !seen[tech] {
			seen[tech] = true
			result = append(result, tech)
		}
	}

	return result
}

func (s *DataImportService) parseSalaryRange(salaryStr string) (int64, *int64, error) {
	salaryStr = strings.TrimSpace(salaryStr)

	if strings.HasSuffix(salaryStr, "+") {
		minRaw := strings.TrimSuffix(salaryStr, "+")
		minRaw = strings.ReplaceAll(minRaw, ".", "")

		min, err := strconv.ParseInt(minRaw, 10, 64)
		if err != nil {
			return 0, nil, fmt.Errorf("failed to parse min salary: %w", err)
		}
		return min, nil, nil
	}

	parts := strings.Split(salaryStr, " - ")
	if len(parts) != 2 {
		return 0, nil, fmt.Errorf("invalid salary range format: %s", salaryStr)
	}

	minRaw := strings.ReplaceAll(parts[0], ".", "")
	maxRaw := strings.ReplaceAll(parts[1], ".", "")

	min, err := strconv.ParseInt(minRaw, 10, 64)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to parse min salary: %w", err)
	}

	maxVal, err := strconv.ParseInt(maxRaw, 10, 64)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to parse max salary: %w", err)
	}
	return min, &maxVal, nil
}

func (s *DataImportService) extractCurrencyCode(currencyStr string) string {
	if strings.Contains(currencyStr, "₺") || strings.Contains(currencyStr, "Türk Lirası") {
		return "TRY"
	}
	if strings.Contains(currencyStr, "$") || strings.Contains(currencyStr, "Dolar") {
		return "USD"
	}
	if strings.Contains(currencyStr, "€") || strings.Contains(currencyStr, "Euro") {
		return "EUR"
	}
	if strings.Contains(currencyStr, "£") || strings.Contains(currencyStr, "Sterlin") {
		return "GBP"
	}
	return "TRY"
}

func (s *DataImportService) normalizeWorkType(workType string) string {
	switch {
	case strings.Contains(workType, "Remote") && strings.Contains(workType, "hibrit"):
		return "Hibrit (Ofis + Remote)"
	case strings.Contains(workType, "Hibrit"):
		return "Hibrit (Ofis + Remote)"
	case strings.Contains(workType, "Remote"):
		return "Remote"
	case strings.Contains(workType, "Ofis"):
		return "Ofis"
	default:
		return workType
	}
}