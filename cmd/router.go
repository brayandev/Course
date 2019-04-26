package main

import (
	"github.com/Course/course"
	"github.com/go-chi/chi"
)

func createRouter(svc course.Service) chi.Router {
	router := chi.NewRouter()
	return router
}
