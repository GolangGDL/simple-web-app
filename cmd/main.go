package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GolangGdl/simple-web-app/pkg/routes"
)

func main() {
	mux := routes.Init()

	srv := http.Server{
		Addr:           ":8081",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 8175,
	}
	fmt.Println("Listen on port " + srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
