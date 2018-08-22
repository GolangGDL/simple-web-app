package main

import (
	"net/http"
	"time"

	"github.com/workshopgo/pkg/routes"
)

func main() {
	mux := routes.Init()

	server := http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 8175,
	}

	server.ListenAndServe()
}
