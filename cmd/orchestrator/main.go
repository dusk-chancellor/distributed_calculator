// ЮХУУ запускаем оркестратор
package main

import (
	"calculator_yandex/http-server/curl_handlers"
	"calculator_yandex/http-server/orchestrator_handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // с mux удобнее
	mux.HandleFunc("/getexpressionslist", curl_handlers.GetExpressionsHandler)
	mux.HandleFunc("/setnewexpression", curl_handlers.SetExpressionHandler)
	mux.HandleFunc("/clearexpressionslist", curl_handlers.ClearExpressionsHandler)
	mux.HandleFunc("/settimeout", curl_handlers.SetTimeoutsHandler)
	mux.HandleFunc("/gettimeouts", curl_handlers.GetTimeoutsHandler)
	mux.HandleFunc("/getanswer", orchestrator_handlers.GetAnswerHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
