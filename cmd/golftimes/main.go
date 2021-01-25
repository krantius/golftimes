package main

import (
	"log"
	"net/http"
	"time"

	"github.com/krantius/golftimes/internal/app"
)

func main() {
	s := &app.Server{API: &app.API{
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}}

	http.HandleFunc("/", s.Courses)

	log.Print("http listening on 8080")
	http.ListenAndServe(":8080", nil)
}
