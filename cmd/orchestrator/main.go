package main

import (
	"context"
	"log"
	"net/http"

	exprHandler "github.com/dusk-chancellor/distributed_calculator/internal/http/handlers/expression"
	authHandler "github.com/dusk-chancellor/distributed_calculator/internal/http/handlers/auth"
	"github.com/dusk-chancellor/distributed_calculator/internal/http/middlewares"
	"github.com/dusk-chancellor/distributed_calculator/internal/storage"
	"github.com/dusk-chancellor/distributed_calculator/internal/utils/orchestrator/manager"
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
	mux.Handle("/auth/", http.StripPrefix("/auth", http.FileServer(http.Dir("frontend/auth"))))

	mux.Handle("/signup/", authHandler.RegisterUserHandler(ctx, db))
	mux.Handle("/login/", authHandler.LoginUserHandler(ctx, db))
	mux.Handle("POST /expression/", middlewares.JWTMiddleware(exprHandler.CreateExpressionHandler(ctx, db)))
	mux.Handle("GET /expression/", middlewares.JWTMiddleware(exprHandler.GetExpressionsHandler(ctx, db)))
	mux.Handle("DELETE /expression/{id}/", middlewares.JWTMiddleware(exprHandler.DeleteExpressionHandler(ctx, db)))

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go manager.RunManager(ctx, db)

	log.Printf("Running Orchestrator server at %s", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("error: %v", err)
	}

	log.Print("Something went wrong...")
}
