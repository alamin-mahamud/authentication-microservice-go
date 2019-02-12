package config

import "os"

type database struct {
	Host     string
	DBName   string
	UserName string
	Password string
}

var (
	DB database
)

// InitDatabase - initializes the database configuration struct
func InitDatabase() {
	DB = database{
		Host:     getFromEnvOrDefault("DB_HOST", "localhost"),
		DBName:   getFromEnvOrDefault("DB_NAME", "temporaryDB"),
		UserName: getFromEnvOrDefault("DB_USER", "root"),
		Password: getFromEnvOrDefault("DB_PASSWORD", "root"),
	}
}

func getFromEnvOrDefault(env string, defaultValue string) string {
	result := os.Getenv("DB_HOST")
	if result == "" && defaultValue != "" {
		result = defaultValue
	}
	return result
}
