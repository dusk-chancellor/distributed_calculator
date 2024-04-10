package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("Running frontend server on localhost:8000")

	if err := http.ListenAndServe(":8000", http.FileServer(http.Dir("cmd/frontend"))); err != nil {
		log.Fatalf("error: %v", err)
		return
	}
}
