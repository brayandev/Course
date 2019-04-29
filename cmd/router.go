package main

import (
	api "github.com/Course/course-api"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func createRouter(service api.Service, logger *zap.Logger) chi.Router {
	router := chi.NewRouter()

	router.Use(contextMiddleware)
	router.With(accessLogMiddleware(logger)).Route("/course", func(router chi.Router) {
		router.Post("/", createCourse(service, logger))
	})

	return router
}
