package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string // Changed to "host:port" format for SQL Server
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	// Load environment variables from the .env file
	godotenv.Load()

	// Return the config with updated defaults for SQL Server
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "SA"),                                                            // Default SQL Server user is 'SA'
		DBPassword: getEnv("DB_PASSWORD", "YourStrong!Pass123"),                                        // Default SQL Server password (can be changed)
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_ADDRESS", "localhost"), getEnv("DB_PORT", "1433")), // SQL Server default port
		DBName:     getEnv("DB_NAME", "ecommerce"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
