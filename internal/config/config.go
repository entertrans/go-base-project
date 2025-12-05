package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    AppEnv    string
    AppPort   string
    DBDriver  string
    DBHost    string
    DBPort    string
    DBUser    string
    DBPass    string
    DBName    string
    JWTSecret string
}

func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

    return &Config{
        AppEnv:    getEnv("APP_ENV", "development"),
        AppPort:   getEnv("APP_PORT", "8080"),
        DBDriver:  getEnv("DB_DRIVER", "mysql"),
        DBHost:    getEnv("DB_HOST", "localhost"),
        DBPort:    getEnv("DB_PORT", "3306"),
        DBUser:    getEnv("DB_USER", "root"),
        DBPass:    getEnv("DB_PASS", ""),
        DBName:    getEnv("DB_NAME", "go_base"),
        JWTSecret: getEnv("JWT_SECRET", "rahasia"),
    }
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
