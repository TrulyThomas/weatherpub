package main

import (
	"github.com/TrulyThomas/weatherpub/internal/config"
	"github.com/TrulyThomas/weatherpub/pkg/util"
)

func main() {
	logger := util.NewLogger()
	logger.Info().Msg("Hello, World!")

	cfg, err := config.NewConfig("config.yaml")
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to load config")
	}

	logger.Info().Interface("config", cfg).Msg("Loaded config")
}
