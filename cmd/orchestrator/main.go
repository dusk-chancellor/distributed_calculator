package main

import (
	"context"
	"log"
	"net/http"

	authHandler "github.com/dusk-chancellor/distributed_calculator/internal/http/handlers/auth"
	exprHandler "github.com/dusk-chancellor/distributed_calculator/internal/http/handlers/expression"
	"github.com/dusk-chancellor/distributed_calculator/internal/storage"
	"github.com/dusk-chancellor/distributed_calculator/internal/utils/orchestrator/jwts"
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

	addr := "0.0.0.0:8080"

	mux := http.NewServeMux()

	authPageHandler := http.StripPrefix("/auth", http.FileServer(http.Dir("frontend/auth")))

	mux.Handle("/", http.HandlerFunc(mainPageHandler))
	mux.Handle("/auth/", authPageHandler)

	mux.Handle("POST /auth/signup/", authHandler.RegisterUserHandler(ctx, db))
	mux.Handle("POST /auth/login/", authHandler.LoginUserHandler(ctx, db))
	mux.Handle("POST /expression/", exprHandler.CreateExpressionHandler(ctx, db))
	mux.Handle("GET /expression/", exprHandler.GetExpressionsHandler(ctx, db))
	mux.Handle("DELETE /expression/{id}/", exprHandler.DeleteExpressionHandler(ctx, db))

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go manager.RunManager(ctx, db)

	log.Printf("running Orchestrator server at %s", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("error: %v", err)
	}

	log.Print("Something went wrong...")
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("auth_token")
	if err != nil {
		http.Redirect(w, r, "/auth", http.StatusSeeOther)
		log.Printf("no cookie found")
		return
	}

	tokenString := cookie.Value

	_, err = jwts.VerifyJWTToken(tokenString)
	if err != nil {
		http.Redirect(w, r, "/auth", http.StatusSeeOther)
		log.Printf("error: %v", err)
		return
	}

	fileServer := http.FileServer(http.Dir("frontend/main"))

	fileServer.ServeHTTP(w, r)
}
