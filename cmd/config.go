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
	ServerHTTPWriteTimeout time.Duration  `envconfig:"SERVER_HTTP_WRITE_TIMEOUT" default:"1s"`
	MongoDBName            string         `envconfig:"MONGO_DB_NAME"`
	MongoDBCollectionName  string         `envconfig:"MONGO_DB_COLLECTION_NAME"`
	MongoDBEndpoint        string         `envconfig:"MONGO_DB_ENDPOINT"`
}

func newConfig() *Config {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}

// LogLevelConfig log level config.
type logLevelConfig struct {
	Value zapcore.Level
}
