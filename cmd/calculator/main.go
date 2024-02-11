package main

import (
	"calculator_yandex/http-server/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/getexpressionslist", handlers.GetExpressionsHandler)
	mux.HandleFunc("/setnewexpression", handlers.SetExpressionHandler)
	mux.HandleFunc("/clearexpressionslist", handlers.ClearExpressionsHandler)

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		return
	}
}
