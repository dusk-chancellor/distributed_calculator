package main

import (
	"context"
	"log"
	"net/http"

	authHandler "github.com/dusk-chancellor/distributed_calculator/internal/http/handlers/auth"
	exprHandler "github.com/dusk-chancellor/distributed_calculator/internal/http/handlers/expression"
	"github.com/dusk-chancellor/distributed_calculator/internal/storage"
	"github.com/dusk-chancellor/distributed_calculator/internal/http/middlewares"
	"github.com/dusk-chancellor/distributed_calculator/internal/utils/orchestrator/manager"
)

// Orchestrator - the main app server, which directly operates with database
// It can be described as "manager" of services

func main() {
	ctx := context.Background()

	db, err := storage.New("./database/storage.db")
	if err != nil {
		panic(err)
	}

	addr := "0.0.0.0:8080"

	mux := http.NewServeMux()

	mainPageHandler := middlewares.AuthorizeJWTToken(http.FileServer(http.Dir("frontend/main")))
	authPageHandler := http.StripPrefix("/auth", http.FileServer(http.Dir("frontend/auth")))

	mux.Handle("/", mainPageHandler)
	mux.Handle("/auth/", authPageHandler)

	mux.Handle("POST /auth/signup/", authHandler.RegisterUserHandler(ctx, db))
	mux.Handle("POST /auth/login/", authHandler.LoginUserHandler(ctx, db))

	mux.Handle("POST /expression/", middlewares.AuthorizeJWTToken(exprHandler.CreateExpressionHandler(ctx, db)))
	mux.Handle("GET /expression/", middlewares.AuthorizeJWTToken(exprHandler.GetExpressionsHandler(ctx, db)))
	mux.Handle("DELETE /expression/{id}/", middlewares.AuthorizeJWTToken(exprHandler.DeleteExpressionHandler(ctx, db)))

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go manager.RunManager(ctx, db)

	log.Printf("running Orchestrator server at %s", addr)
	go log.Fatal(server.ListenAndServe())

	log.Print("Something went wrong...")
}
