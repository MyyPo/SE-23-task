package config

import (
	"fmt"
	"os"

	"github.com/myypo/btcinform/internal/constants"
)

type Config interface {
	GetDBPath() *string
	GetSMTPHost() *string
	GetSMTPPort() *string
	GetEmailUsername() *string
	GetEmailPassword() *string
}

type ConfigImpl struct {
	DBPath        *string
	SMTPHost      *string
	SMTPPort      *string
	EmailUsername *string
	EmailPassword *string
}

func NewConfigImpl() (*ConfigImpl, error) {
	dbPath, err := getEnvValue(constants.DBPathEnvKey)
	if err != nil {
		return nil, err
	}
	smtpHost, err := getEnvValue(constants.SMTPHostEnvKey)
	if err != nil {
		return nil, err
	}
	smtpPort, err := getEnvValue(constants.SMTPPortEnvKey)
	if err != nil {
		return nil, err
	}
	emailUsername, err := getEnvValue(constants.EmailUsernameEnvKey)
	if err != nil {
		return nil, err
	}
	emailPassword, err := getEnvValue(constants.EmailPasswordEnvKey)
	if err != nil {
		return nil, err
	}
	return &ConfigImpl{
		DBPath:        dbPath,
		SMTPHost:      smtpHost,
		SMTPPort:      smtpPort,
		EmailUsername: emailUsername,
		EmailPassword: emailPassword,
	}, nil
}

func getEnvValue(envKey string) (*string, error) {
	envVal, exists := os.LookupEnv(envKey)
	if !exists {
		return nil, fmt.Errorf("Not set env: %s", envKey)
	}
	return &envVal, nil
}

func (c *ConfigImpl) GetDBPath() *string {
	return c.DBPath
}

func (c *ConfigImpl) GetSMTPHost() *string {
	return c.SMTPHost
}

func (c *ConfigImpl) GetSMTPPort() *string {
	return c.SMTPPort
}

func (c *ConfigImpl) GetEmailUsername() *string {
	return c.EmailUsername
}

func (c *ConfigImpl) GetEmailPassword() *string {
	return c.EmailPassword
}
