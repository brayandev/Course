package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Course/course"

	"github.com/facebookgo/grace/gracehttp"
)

func main() {
	cfg := newConfig()

	repository := course.NewRepository()
	service := course.NewService(repository)

	router := createRouter(service)
	sErr := gracehttp.Serve(&http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.AppPort),
		Handler:      router,
		ReadTimeout:  cfg.ServerHTTPReadTimeout,
		WriteTimeout: cfg.ServerHTTPWriteTimeout,
	})
	if sErr != nil {
		log.Println("failed on start server.", sErr)
	}
}
