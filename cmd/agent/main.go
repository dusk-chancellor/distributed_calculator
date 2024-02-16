package main

import (
	"calculator_yandex/http-server/agent_handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/setexpr", agent_handlers.SetExprHandler)

	if err := http.ListenAndServe(":8081", mux); err != nil {
		return
	}
}
