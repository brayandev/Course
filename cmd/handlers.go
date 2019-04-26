package main

import (
	"net/http"

	"github.com/Course/course"
	"go.uber.org/zap"
)

func createCourse(service course.Service, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
