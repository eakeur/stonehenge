package config

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadConfigurations() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Config{}, err
	}

	return Config{
		Database: DatabaseConfigurations{
			User:           os.Getenv("DB_USER"),
			Password:       os.Getenv("DB_PASS"),
			Host:           os.Getenv("DB_HOST"),
			Port:           os.Getenv("DB_PORT"),
			Name:           os.Getenv("DB_NAME"),
			SSLMode:        os.Getenv("DB_SSL"),
			MigrationsPath: os.Getenv("DB_MIG"),
		},
		Access: AccessConfigurations{
			ExpirationTime: os.Getenv("ACCESS_EXPIRES"),
			SigningKey:     os.Getenv("ACCESS_SIGNING_KEY"),
		},
		Server: ServerConfigurations{
			Hostname: os.Getenv("API_HOSTNAME"),
			ListenPort: os.Getenv("API_PORT"),
		},
		Logger: LoggerConfigurations{
			Environment: os.Getenv("ENVIRONMENT"),
		},
	}, nil
}