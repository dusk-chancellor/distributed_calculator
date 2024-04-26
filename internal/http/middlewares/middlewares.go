package middlewares

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/dusk-chancellor/distributed_calculator/internal/utils/orchestrator/jwts"
)

func AuthorizeJWTToken(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			log.Printf("no cookie found")
			return
		}
						
		tokenString := cookie.Value
						
		tokenValue, err := jwts.VerifyJWTToken(tokenString)
		if err != nil {
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			log.Printf("error: %v", err)
			return
		}
						
		userID, err := strconv.ParseInt(tokenValue, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("error: %v", err)
			return
		}
			
		ctx := context.WithValue(r.Context(), "userid", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}