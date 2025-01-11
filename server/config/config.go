package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPServer HTTPServerConfig
	Mongo      MongoConfig
}

type MongoConfig struct {
	URI      string `env:"MONGO_URI" envDefault:"mongodb+srv://movseek:movseekdb123@movseek.kjfyn.mongodb.net"`
	URI_2    string `env:"MONGO_URI_2" envDefault:"mongodb+srv://movseek:movseekdb123@cluster14271.ogtgfuh.mongodb.net"`
	DBName   string `env:"MONGO_DB_NAME" envDefault:"MovSeek"`
	DBName_2 string `env:"MONGO_DB_NAME_2" envDefault:"Cluster14271"`
}

type HTTPServerConfig struct {
	Port int `env:"PORT" envDefault:"8080"`
}

// Load loads the configuration from the environment variables.

func Load() (*Config, error) {
	var config Config

	// Load .env file if it exists
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	// Parse environment variables into config struct
	if err := env.Parse(&config); err != nil {
		return nil, fmt.Errorf("error parsing env vars: %w", err)
	}

	return &config, nil
}
