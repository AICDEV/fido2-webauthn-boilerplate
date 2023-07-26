package utils

import (
	"git.eon-cds.de/repos/dlab/wad-fido2/backend/models"
	"github.com/kelseyhightower/envconfig"
)

const (
	ENV_PREFIX = "WAD"
)

// ParseEnv required environment variables found in *./models/env.model.go. If some values are missing the program will end up in panic
func ParseEnv() *models.EnvConfig {
	envConfig := &models.EnvConfig{}

	err := envconfig.Process(ENV_PREFIX, envConfig)

	if err != nil {
		panic(err)
	}

	return envConfig
}
