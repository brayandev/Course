package main

import (
	"github.com/Course/course"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func createRouter(service course.Service, logger *zap.Logger) chi.Router {
	router := chi.NewRouter()
	router.Use(contextMiddleware)

	router.With(accessLogMiddleware(logger)).Route("/course", func(router chi.Router) {
		router.Post("/", createCourse(service, logger))
	})

	return router
}
