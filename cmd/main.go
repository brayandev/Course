package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/Course/course"

	"github.com/facebookgo/grace/gracehttp"
)

func main() {
	cfg := newConfig()

	logger, lErr := course.ConfigLog(zap.NewAtomicLevelAt(cfg.LogLevel.Value)).Build()
	if lErr != nil {
		panic(lErr)
	}

	db, dbErr := course.NewMongoDB(cfg.MongoDBEndpoint)
	if dbErr != nil {
		logger.Error("Error to create a new connection for db", zap.Error(dbErr))
	}

	repository := course.NewRepository(db, cfg.MongoDBName, cfg.MongoDBCollectionName)
	service := course.NewService(repository)

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
