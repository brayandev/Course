package main

import (
	"net/http"

	"github.com/Course/course"
)

func createRouter(svc course.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
