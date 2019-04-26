package main

import (
	"github.com/Course/course"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func createRouter(svc course.Service, logger *zap.Logger) chi.Router {
	router := chi.NewRouter()
	router.Use(contextMiddleware)
	router.With(accessLogMiddleware(logger))

	return router
}
