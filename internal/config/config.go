package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                 string
	MongoURI             string
	MongoDB              string
	MongoUser            string
	MongoPass            string
	JWTSecret            string
	JWTExpiry            string
	LinkedInClientID     string
	LinkedInClientSecret string
	HMACSecret           string
	LinkedInRedirectURL  string
	FrontendCallbackURL  string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		Port:                 getEnv("PORT", "8080"),
		MongoURI:             getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:              getEnv("MONGO_DB", "salarydb"),
		MongoUser:            getEnv("MONGO_USER", "admin"),
		MongoPass:            getEnv("MONGO_PASS", "admin"),
		JWTSecret:            getEnv("JWT_SECRET", "your_jwt_secret"),
		JWTExpiry:            getEnv("JWT_EXPIRY", "24h"),
		LinkedInClientID:     getEnv("LINKEDIN_CLIENT_ID", ""),
		LinkedInClientSecret: getEnv("LINKEDIN_CLIENT_SECRET", ""),
		HMACSecret:           getEnv("HMAC_SECRET", "your_hmac_secret"),
		LinkedInRedirectURL:  getEnv("LINKEDIN_REDIRECT_URL", "http://localhost:8080/auth/linkedin/callback"),
		FrontendCallbackURL:  getEnv("FRONTEND_CALLBACK_URL", "http://localhost:3000/auth/callback"),
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
