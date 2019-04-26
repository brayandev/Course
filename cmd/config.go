package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap/zapcore"
)

// Config access envs.
type Config struct {
	AppPort                int            `envconfig:"APP_PORT"`
	LogLevel               logLevelConfig `envconfig:"LOG_LEVEL" default:"info"`
	ServerHTTPReadTimeout  time.Duration  `envconfig:"SERVER_HTTP_READ_TIMEOUT" default:"1s"`
	ServerHTTPWriteTimeout time.Duration  `envoncifg:"SERVER_HTTP_WRITE_TIMEOUT" default:"1s"`
}

func newConfig() *Config {
	cfg := &Config{}
	err := envconfig.Process("local.env", cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

// LogLevelConfig log level config.
type logLevelConfig struct {
	Value zapcore.Level
}
