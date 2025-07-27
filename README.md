# Salystic Backend

A minimal salary benchmarking backend called Salystic for software engineers built with Go, Echo, and MongoDB. Users authenticate via LinkedIn and can submit and view their own salary entries.

## 🚀 Features

* **Anonymous LinkedIn Authentication & User Management**
  * OAuth2 flow with minimal data collection
  * Pseudonymized LinkedIn IDs using HMAC-SHA256
  * JWT-based session management
  * LinkedIn OAuth2 login

* **Salary Entries Management**
  * Submit comprehensive salary entries with validation
  * Required fields: `position`, `level`, `experience`, `company`, `city`, `salary`, `currency`, `techStacks`, `companySize`, `workType`
  * Career progression tracking with raise history
  * Private user access to own entries only

* **Career Progression Tracking**
  * **Job Change Analysis**
    * Track salary differences between job changes
    * Identify if salary increase was the motivation for job switch
    * Compare previous job salary with new job starting salary
  * **Raise Tracking**
    * Record raise history within the same company
    * Track when raises occurred and new salary amounts
    * Calculate annual raise frequency and percentages
    * Monitor salary growth patterns over time

* **Comprehensive Analytics Dashboard**
  * **Public Analytics** - Transparent salary insights for community benefit
  * Multi-dimensional salary breakdowns by:
    * Position, Level, Experience, Company, City
    * Technology Stack, Company Size, Work Type, Currency
  * Statistical insights: Average, Min, Max salaries with entry counts
  * Top-paying positions and technologies charts
  * Salary range distributions
  * Career progression analytics (raises, job changes)
  * Query filtering by position, level, and currency
  * Real-time data updates

* **Advanced Security & Privacy**
  * **Viper Configuration Management** - Centralized config with environment variable binding
  * **Custom CORS Policy** - Public access for analytics, restricted access for private APIs
  * **Data Privacy Protection**
    * Only LinkedIn subject identifier (sub) stored in database (pseudonymized with HMAC-SHA256)
    * Profile data (name, email, picture) received from LinkedIn OAuth but only sent to frontend for display
    * No personal information permanently stored in database
    * Anonymous data aggregation with minimum thresholds

* **Constants & Data Management**
  * Automatic seeding of positions, levels, tech stacks, experiences, companies, cities, etc.
  * 400+ predefined constants for consistent data entry
  * Multi-language support (Turkish cities and companies)
  * Data import capabilities via JSON files

## 🧰 Implementation Principles

Salystic follows these core backend principles:

* **Request Parsing:** Cleanly parse and bind incoming HTTP requests to Go structs.
* **Validation:** Validate all request payloads using field tags and validator middleware.
* **Response Types/Errors:** Standardize API responses and error formats for consistency.
* **Type-Safe Handlers:** Leverage Go's type system to ensure compile-time safety in handlers.
* **Dependency Inversion:** Architect services with interfaces to decouple dependencies and facilitate testing.

### Database Optimizations
- Connection pooling configuration
- Compound indexes for common queries
- Read preference for analytics

## 🎯 IMPLEMENTED SOLID Principles

- **Single Responsibility**: Each service handles one domain concern
- **Open/Closed**: Extensible through interfaces, not modification
- **Liskov Substitution**: Repository interfaces allow swapping implementations
- **Interface Segregation**: Small, focused interfaces (Reader, Writer, Validator)
- **Dependency Inversion**: All dependencies injected via interfaces

## 🔒 Data Privacy & Security

### Current Implementation
- **Public Analytics**: Aggregated salary data is publicly accessible to encourage transparency
- **Private Entries**: Individual salary entries are protected by JWT authentication
- **Anonymous Data**: Only LinkedIn subject identifier (sub) is pseudonymized and stored using HMAC-SHA256
- **Limited Data Storage**: Profile information (name, email, picture) is received from LinkedIn but only sent to frontend for display, not stored in database
- **Minimum Data Threshold**: Analytics only shown when category has 5+ entries

### Security Considerations
⚠️ **Important**: Analytics endpoints are currently public.

### Implemented Security Features
```go
// Custom CORS Policy - Different rules for public vs private APIs
func CORSWithConfig(frontendURL string) echo.MiddlewareFunc {
    // Public analytics APIs - allow all origins
    if strings.HasPrefix(path, "/api/v1/analytics") ||
       strings.HasPrefix(path, "/api/v1/constants") {
        corsConfig := middleware.CORSConfig{
            AllowOrigins: []string{"*"},
            AllowMethods: []string{"GET", "OPTIONS"},
        }
    }
    // Private APIs - restrict to frontend domain
    corsConfig := middleware.CORSConfig{
        AllowOrigins: []string{frontendURL},
        AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowCredentials: true,
    }
}

// Viper Configuration Management
viper.AutomaticEnv()
viper.SetDefault("SERVER_PORT", "8080")
viper.BindEnv("server.port", "PORT")
```

## 📦 Prerequisites

* Go 1.23+
* MongoDB 4.4+
* LinkedIn OAuth2 application credentials

## ⚙️ Configuration

1. Copy `.env.example` to `.env`:

   ```bash
   cp .env.example .env
   ```
2. Update environment variables in `.env`:

   ```dotenv
    # Server Configuration
    PORT=8080

    # Database Configuration
    MONGO_URI=mongodb://localhost:27017
    MONGO_DB=salarydb
    MONGO_USER=admin
    MONGO_PASS=admin

    # JWT Configuration
    JWT_SECRET=your_jwt_secret_here_change_this
    JWT_EXPIRY=24h

    # LinkedIn OAuth Configuration
    LINKEDIN_CLIENT_ID=your_linkedin_client_id
    LINKEDIN_CLIENT_SECRET=your_linkedin_client_secret
    LINKEDIN_REDIRECT_URL=http://localhost:8080/auth/linkedin/callback
    FRONTEND_CALLBACK_URL=http://localhost:3000/auth/callback

    # Security & CORS
    HMAC_SECRET=your_hmac_secret_here_change_this
    FRONTEND_URL=http://localhost:3000
   ```

## 🏃 Running Locally

```bash
go mod download
go run ./cmd/server
```

Server runs at `http://localhost:8080`.

### Quick Security Check
After starting the server, verify security settings:
- Analytics endpoints (`/api/v1/analytics`) should be accessible without auth
- Individual entries (`/api/v1/entries`) should return 401 without JWT token
- Implementing rate limiting for production deployment

## 📂 Project Structure

```
salystic-backend/
├── cmd/
│   ├── server/         # Main application entrypoint
│   └── import/         # Data import utility
├── internal/
│   ├── api/
│   │   ├── handlers/   # HTTP request handlers
│   │   ├── middleware/ # Custom middleware (auth, CORS)
│   │   └── routes/     # Route definitions
│   ├── auth/          # JWT and LinkedIn OAuth logic
│   ├── config/        # Viper configuration management
│   ├── model/         # Data models and structures
│   ├── repo/          # Repository implementations
│   └── service/       # Business logic layer
├── pkg/
│   ├── database/      # MongoDB connection utilities
│   └── responses/     # Standardized API responses
├── .env.example       # Environment configuration template
├── Dockerfile         # Container configuration
├── go.mod
└── README.md
```

## 🛠️ API Endpoints

> **Note**: Analytics endpoints are publicly accessible to promote transparency. Individual salary entries remain private and require authentication.

### Authentication

| Method | Path                      | Auth Required |
| ------ | ------------------------- | ------------- |
| GET    | `/auth/linkedin`          | No            |
| GET    | `/auth/linkedin/callback` | No            |
| GET    | `/auth/me`                | JWT           |
| POST   | `/auth/logout`            | JWT           |

### Salary Entries

| Method | Path                  | Auth Required | Description           |
| ------ | --------------------- | ------------- | --------------------- |
| POST   | `/api/v1/entries`     | JWT           | Submit a salary entry |
| GET    | `/api/v1/entries`     | JWT           | List user's entries   |
| GET    | `/api/v1/entries/:id` | JWT           | Get entry details     |
| PUT    | `/api/v1/entries/:id` | JWT           | Update a salary entry |
| DELETE | `/api/v1/entries/:id` | JWT           | Delete a salary entry |

### Career Progression 🆕

| Method | Path                              | Auth Required | Description                        |
| ------ | --------------------------------- | ------------- | ---------------------------------- |
| POST   | `/api/v1/entries/:id/raises`     | JWT           | Add raise record to current job    |
| GET    | `/api/v1/entries/:id/raises`     | JWT           | Get raise history for an entry     |

### Analytics (Public)

| Method | Path                        | Auth Required | Description                          |
| ------ | --------------------------- | ------------- | ------------------------------------ |
| GET    | `/api/v1/analytics`         | No            | Comprehensive salary analytics       |
| GET    | `/api/v1/analytics/career`  | No            | Career progression insights          |
| GET    | `/api/v1/analytics/positions` | No          | Available positions list             |
| GET    | `/api/v1/analytics/levels`  | No            | Available levels list                |

### Constants (Public)

| Method | Path                        | Auth Required | Description                          |
| ------ | --------------------------- | ------------- | ------------------------------------ |
| GET    | `/api/v1/constants/positions` | No          | All position options                 |
| GET    | `/api/v1/constants/levels`  | No            | All level options                    |
| GET    | `/api/v1/constants/tech-stacks` | No        | All technology stack options         |
| GET    | `/api/v1/constants/experiences` | No        | All experience range options         |
| GET    | `/api/v1/constants/companies` | No          | All company/sector options           |
| GET    | `/api/v1/constants/company-sizes` | No      | All company size options             |
| GET    | `/api/v1/constants/work-types` | No        | All work type options                |
| GET    | `/api/v1/constants/cities`  | No            | All city options                     |
| GET    | `/api/v1/constants/currencies` | No        | All currency options                 |

## 📝 Request/Response Examples

### Submit Salary Entry
```json
POST /api/v1/entries
{
  "position": "Back-end Developer",
  "level": "Senior",
  "experience": "5 - 7 Yıl",
  "company": "Yazılım Evi & Danışmanlık",
  "companySize": "51 - 100 Kişi",
  "city": "İstanbul",
  "workType": "Remote",
  "techStacks": ["Go", "NodeJS", "React"],
  "salary": 150000,
  "currency": "TRY",
  "startTime": "2024-01-01T00:00:00Z",
  "endTime": null,
  "salaryMin": 120000,  // Optional: previous job salary for tracking
  "salaryMax": 180000   // Optional: current job salary range
}
```

### Add Raise Record
```json
POST /api/v1/entries/:id/raises
{
  "raiseDate": "2024-06-01T00:00:00Z",
  "newSalary": 165000,
  "percentage": 10
}
```

### Analytics Response (Public)
```json
GET /api/v1/analytics?position=Back-end Developer&level=Senior&currency=TRY

{
  "totalEntries": 1543,
  "averageSalary": 142500,
  "averageSalaryByPosition": {
    "Back-end Developer": 135000,
    "Front-end Developer": 125000,
    "Full Stack Developer": 140000
  },
  "minSalaryByPosition": {
    "Back-end Developer": 85000,
    "Front-end Developer": 75000
  },
  "maxSalaryByPosition": {
    "Back-end Developer": 200000,
    "Front-end Developer": 180000
  },
  "averageSalaryByLevel": {
    "Junior": 75000,
    "Middle": 115000,
    "Senior": 155000
  },
  "averageSalaryByTech": {
    "Go": 165000,
    "React": 145000,
    "NodeJS": 140000
  },
  "topPayingPositions": [
    {"name": "AI Engineer", "value": 180000, "count": 45},
    {"name": "DevOps Engineer", "value": 170000, "count": 78}
  ],
  "topPayingTechs": [
    {"name": "Go", "value": 165000, "count": 234},
    {"name": "Rust", "value": 160000, "count": 67}
  ],
  "salaryRanges": [
    {"name": "50k-75k", "value": 50000, "count": 156},
    {"name": "75k-100k", "value": 75000, "count": 289}
  ],
  "lastUpdated": "2024-01-15T10:30:00Z"
}
```

### Career Analytics Response (Public)
```json
GET /api/v1/analytics/career
{
  "jobChanges": {
    "averageSalaryIncrease": 25.5,
    "percentageWithIncrease": 87.3
  },
  "raises": {
    "averagePerYear": 1.2,
    "averagePercentage": 8.5,
    "medianTimeBetweenRaises": 11  // months
  }
}
```

Note: All data is aggregated and anonymous. No individual entries or personal information is exposed.

## ✅ Testing

We have tests for all main parts of the backend.

Run all tests:
```bash
go test ./internal/...
```

Run tests with details:
```bash
go test -v ./internal/...
```

## 🤝 Contributing

1. Fork the repository and clone
2. `git checkout -b feature/...`
3. `go fmt && go test`
4. `git commit -m "feat: ..."`
5. `git push origin feature/...` and open a PR