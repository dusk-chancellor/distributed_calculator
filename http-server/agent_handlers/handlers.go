package agent_handlers

import (
	"calculator_yandex/internal/calculation"
	"io"
	"net/http"
)

func SetExprHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post method allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	expression := string(body)

	if err := calculation.ReceiveExpression(expression); err != nil {
		http.Error(w, "Cannot ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Expression received and processed"))
}
