package main

import (
	exprHandler "calculator_yandex/internal/http_server/handlers/expression"
	"calculator_yandex/internal/storage"
	"context"
	"log"
	"net/http"
)

// Orchestrator - the main app server, which directly operates with database
// It can be described as "manager" of services

func main() {
	ctx := context.TODO()

	db, err := storage.New("./database/storage.db")
	if err != nil {
		panic(err)
	}
	
	addr := "localhost:8080" // must be in config file

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("frontend/main")))
	mux.Handle("POST /expression/", exprHandler.CreateExpressionHandler(ctx, db))
	mux.Handle("GET /expression/", exprHandler.GetExpressionsHandler(ctx, db))
	mux.Handle("DELETE /expression/{id}/", exprHandler.DeleteExpressionHandler(ctx, db))

	server := &http.Server{
		Addr: addr,
		Handler: mux,
	}

	log.Printf("Orchestrator started at %s", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("error: %v", err)
	}

	log.Print("Something went wrong...")
}