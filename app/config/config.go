package config

import (
	"os"
)

type MongoDbCredential struct {
	AuthSource string
	Username   string
	Password   string
	Url        string
	ServerPort string
}

type Config struct {
	MongoDb MongoDbCredential
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		MongoDb: MongoDbCredential{
			AuthSource: getEnv("AUTH_SERVICE", ""),
			Url:        getEnv("URL_Mongo", ""),
			Username:   getEnv("USERNAME_Mongo", ""),
			Password:   getEnv("PASSWORD_Mongo", ""),
			ServerPort: getEnv("SERVER_PORT", "8081"),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
