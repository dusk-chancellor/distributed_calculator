// ЮХУУ запускаем сервак storage
package main

import (
	"calculator_yandex/http-server/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // с mux удобнее
	mux.HandleFunc("/getexpressionslist", handlers.GetExpressionsHandler)
	mux.HandleFunc("/setnewexpression", handlers.SetExpressionHandler)
	mux.HandleFunc("/clearexpressionslist", handlers.ClearExpressionsHandler)
	mux.HandleFunc("/settimeout", handlers.SetTimeoutsHandler)
	mux.HandleFunc("/gettimeouts", handlers.GetTimeoutsHandler)

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		return
	}
}
