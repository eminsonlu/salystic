# Salystic Backend

A minimal salary benchmarking backend called Salystic for software engineers built with Go, Echo, and MongoDB. Users authenticate via LinkedIn and can submit and view their own salary entries.

## 🚀 Features

* **Anonymous LinkedIn Authentication & User Management**
  * OAuth2 flow with minimal data collection
  * Pseudonymized LinkedIn IDs using HMAC-SHA256
  * JWT-based session management
  * LinkedIn OAuth2 login

* **Salary Entries**
  * Submit salary entries with required fields:
    * `country`, `currency`, `sector`, `job`, `title`, `salary`, `startTime`, `endTime`

* **Career Progression Tracking** 🆕
  * **Job Change Analysis**
    * Track salary differences between job changes
    * Identify if salary increase was the motivation for job switch
    * Compare previous job salary with new job starting salary
  * **Raise Tracking**
    * Record raise history within the same company
    * Track when raises occurred and new salary amounts
    * Calculate annual raise frequency and percentages
    * Monitor salary growth patterns over time

* **Data Privacy**
  * Only LinkedIn subject identifier (sub) stored in database (pseudonymized with HMAC-SHA256)
  * Profile data (name, email, picture) received from LinkedIn OAuth but only sent to frontend for display
  * No personal information permanently stored in database
  * Anonymous data aggregation

* **Constants Seeding**
  * Migration on startup seeds jobs, titles, sectors, countries, and currencies if not present

* **Analytics**
  * View total entries count (public access)
  * View average salary by job and sector (public access)
  * Career progression insights (avg raises per year, job change patterns) 🆕
  * All analytics show aggregated, anonymous data only

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

### Recommended Security Enhancements
```go
// Add rate limiting middleware
rateLimiter := middleware.NewRateLimiter(
    100,    // requests
    "1h",   // per hour
)
e.GET("/api/v1/analytics", handler.GetAnalytics, rateLimiter)

// Add cache headers for public endpoints
e.Use(middleware.CacheControl(300)) // 5 minutes

// CORS configuration
e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"https://yourdomain.com"},
    AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
}))
```

## 📦 Prerequisites

* Go 1.20+
* MongoDB 4.4+
* LinkedIn OAuth2 application credentials

## ⚙️ Configuration

1. Copy `.env.example` to `.env`:

   ```bash
   cp .env.example .env
   ```
2. Update environment variables in `.env`:

   ```dotenv
    PORT=8080
    MONGO_URI=mongodb://localhost:27017
    MONGO_DB=salarydb
    MONGO_USER=admin
    MONGO_PASS=admin
    JWT_SECRET=your_jwt_secret_here_change_this
    JWT_EXPIRY=24h
    LINKEDIN_CLIENT_ID=your_linkedin_client_id
    LINKEDIN_CLIENT_SECRET=your_linkedin_client_secret
    HMAC_SECRET=your_hmac_secret_here_change_this
    LINKEDIN_REDIRECT_URL=http://localhost:8080/auth/linkedin/callback
    FRONTEND_CALLBACK_URL=http://localhost:3000/auth/callback
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
├── cmd/server          # Application entrypoint
├── internal
│   ├── api             # HTTP handlers
│   ├── auth            # Auth logic & middleware
│   ├── model           # Data models
│   ├── repo            # Repository implementations
│   └── service         # Business logic
│       ├── salary      # Salary entry management
│       └── career      # Career progression tracking 🆕
├── config              # Configuration loader
├── pkg                 # Utilities
├── .env.example
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
| GET    | `/api/v1/analytics`         | No            | Total entries & averages             |
| GET    | `/api/v1/analytics/career`  | No            | Career progression insights 🆕       |

## 📝 Request/Response Examples

### Submit Salary Entry with Previous Job Info
```json
POST /api/v1/entries
{
  "country": "US",
  "currency": "USD",
  "sector": "Technology",
  "job": "Backend Developer",
  "title": "Senior",
  "salary": 150000,
  "startTime": "2024-01-01T00:00:00Z",
  "endTime": null,
  "previousJobSalary": 120000  // Optional: for job change tracking
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
GET /api/v1/analytics

{
  "totalEntries": 1543,
  "averageSalaryByJob": {
    "Backend Developer": {
      "Junior": 75000,
      "Mid": 95000,
      "Senior": 135000
    }
  },
  "averageSalaryBySector": {
    "Technology": 115000,
    "Finance": 125000,
    "Healthcare": 95000
  },
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
