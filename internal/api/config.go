package api

import (
	"fmt"
	"net/url"
	"os"
)

type ServerConfig struct {
	DatabaseURL string
	Environment string
	ServiceName string
	Port        string
	BaseURL     url.URL
}

func LoadConfig() (*ServerConfig, error) {
	s := ServerConfig{}

	if v, ok := os.LookupEnv("DATABASE_URI"); ok {
		s.DatabaseURL = v
	} else {
		return nil, fmt.Errorf("DATABASE_URI not set")
	}

	if v, ok := os.LookupEnv("SERVICE_NAME"); ok {
		s.ServiceName = v
	} else {
		s.ServiceName = "GENERIC_WEB_API_SERVICE"
	}

	if v, ok := os.LookupEnv("API_PORT"); ok {
		s.Port = v
	} else {
		s.Port = "8080"
	}

	if v, ok := os.LookupEnv("DEPLOYMENT_ENVIRONMENT"); ok {
		s.Environment = v
	} else {
		s.Environment = "production"
	}

	if v, ok := os.LookupEnv("BASE_URL"); ok {
		parsed, err := url.Parse(v)
		if err != nil {
			return nil, fmt.Errorf("invalid BASE_URL: %w", err)
		}
		s.BaseURL = *parsed
	} else {
		parsed, _ := url.Parse("https://api.example.com")
		s.BaseURL = *parsed
	}

	return &s, nil
}
