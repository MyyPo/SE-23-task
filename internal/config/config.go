package config

import (
	"fmt"
	"os"

	"github.com/myypo/btcinform/internal/constants"
)

type Config struct {
	DBPath *string
}

func NewConfig() (*Config, error) {
	dbPath, exists := os.LookupEnv(constants.ExchangeRateURL)
	if !exists {
		return nil, fmt.Errorf("Not set env: %s", constants.DBPathEnvKey)
	}
	return &Config{
		DBPath: &dbPath,
	}, nil
}
