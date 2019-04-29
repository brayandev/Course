package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Course/course"
	"go.uber.org/zap"
)

func createCourse(service course.Service, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var course course.Course
		err := json.NewDecoder(r.Body).Decode(course)
		if err != nil {
			logger.Error("error on decoder body of course", zap.NamedError("error", err))
		}
		courseID, err := service.CreateCourse(r.Context(), course)
		if err != nil {
			logger.Error("fail on create course", zap.NamedError("error", err))
		}
		fmt.Println(courseID)
	}
}
