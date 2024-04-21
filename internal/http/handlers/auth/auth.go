package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInteractor interface {
	RegisterUser(ctx context.Context, uname, pswrd string) error
	LoginUser(ctx context.Context, uname, pswrd string) (string, error)
}

func RegisterUserHandler(ctx context.Context, userInteractor UserInteractor) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf("error: %v", err)
			return
		}

		if err := userInteractor.RegisterUser(ctx, req.Username, req.Password); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("error: %v", err)
			return
		}

		log.Print("success RegisterUserHandler")
		w.WriteHeader(http.StatusCreated)
	}
}

func LoginUserHandler(ctx context.Context, userInteractor UserInteractor) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf("error: %v", err)
			return
		}

		token, err := userInteractor.LoginUser(ctx, req.Username, req.Password)
		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			log.Printf("error: %v", err)
			return
		}

		response := map[string]string{
			"token": token,
		}

		json.NewEncoder(w).Encode(response)
		log.Print("success LoginUserHandler")
	}
}