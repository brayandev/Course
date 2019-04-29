package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	api "github.com/Course/course-api"
	"github.com/facebookgo/grace/gracehttp"
)

func main() {
	cfg := newConfig()

	logger, lErr := api.ConfigLog(zap.NewAtomicLevelAt(cfg.LogLevel.Value)).Build()
	if lErr != nil {
		panic(lErr)
	}

	db, dbErr := api.NewMongoDB(cfg.MongoDBEndpoint)
	if dbErr != nil {
		logger.Error("Error to create a new connection for db", zap.Error(dbErr))
	}

	repository := api.NewRepository(db, cfg.MongoDBName, cfg.MongoDBCollectionName)
	service := api.NewService(repository)

	router := createRouter(service, logger)
	sErr := gracehttp.Serve(&http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.AppPort),
		Handler:      router,
		ReadTimeout:  cfg.ServerHTTPReadTimeout,
		WriteTimeout: cfg.ServerHTTPWriteTimeout,
	})
	if sErr != nil {
		logger.Error("failed on server start", zap.NamedError("error", sErr))
	}
}
