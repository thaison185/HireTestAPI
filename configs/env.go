package configs

import (
	"os"
)

type Env struct {
	AppName             string
	AppEnv              string
	AppHost             string
	AppPort             string
	JWTSecret           string
	JWTRefreshExpiresIn string
	JWTAccessExpiresIn  string
	DBHost              string
	DBPort              string
	DBUser              string
	DBPassword          string
	DBName              string
	DBSSLMode           string
}

func LoadEnv() Env {
	return Env{
		AppName:             getEnv("APP_NAME", "HireTest API"),
		AppEnv:              getEnv("APP_ENV", "development"),
		AppHost:             getEnv("APP_HOST", "0.0.0.0"),
		AppPort:             getEnv("APP_PORT", "8080"),
		JWTSecret:           getEnv("JWT_SECRET", "super-secret-key"),
		JWTRefreshExpiresIn: getEnv("JWT_REFRESH_EXPIRES_IN", "168h"),
		JWTAccessExpiresIn:  getEnv("JWT_ACCESS_EXPIRES_IN", "24h"),
		DBHost:              getEnv("DB_HOST", "localhost"),
		DBPort:              getEnv("DB_PORT", "5432"),
		DBUser:              getEnv("DB_USER", "postgres"),
		DBPassword:          getEnv("DB_PASSWORD", "postgres"),
		DBName:              getEnv("DB_NAME", "hiretest"),
		DBSSLMode:           getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
