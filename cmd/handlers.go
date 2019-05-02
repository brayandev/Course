package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	api "github.com/Course/course-api"
	"go.uber.org/zap"
)

func createCourse(service api.Service, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var course api.Course
		route := "create-course"
		err := json.NewDecoder(r.Body).Decode(&course)
		if err != nil {
			writeError(w, err)
			api.LogError(r.Context(), logger, route, "invalid body request", err)
			return
		}
		courseID, cErr := service.CreateCourse(r.Context(), course)
		if cErr != nil {
			writeError(w, cErr)
			api.LogError(r.Context(), logger, route, "error on create course", cErr)
		}
		writeResponse(w, http.StatusOK, course)
		fmt.Println(courseID)
	}
}

func writeResponse(w http.ResponseWriter, code int, content versionable) error {
	if content == nil {
		w.WriteHeader(code)
		return nil
	}
	contentType := "application/json"
	if content.Version() != "" {
		contentType = fmt.Sprintf("application/%s+json", content.Version())
	}
	w.Header().Set("Content-Type", fmt.Sprintf("%s; charset=utf-8", contentType))
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(content)
	if err != nil {
		return err
	}
	return nil
}

func writeError(w http.ResponseWriter, err error) {
	switch tErr := err.(type) {
	case *api.Error:
		writeResponse(w, getErrorHTTPCode(tErr), tErr)
	default:
		writeResponse(w, http.StatusInternalServerError, api.NewUnknownError(err.Error()))
	}
}

func getErrorHTTPCode(err *api.Error) int {
	switch err.ErrType {
	case api.ErrorInvalidRequest:
		return http.StatusBadRequest

	default:
		return http.StatusInternalServerError
	}
}

type versionable interface {
	Version() string
}
