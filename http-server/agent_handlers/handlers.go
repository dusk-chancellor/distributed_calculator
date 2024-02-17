package agent_handlers

import (
	"calculator_yandex/internal/calculation"
	"io"
	"net/http"
	"strings"
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

	idAndExpr := strings.Split(string(body), ":")

	if err := calculation.ReceiveAndPost(idAndExpr[0], idAndExpr[1]); err != nil {
		http.Error(w, "Cannot receive and post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Expression received and processed"))
}
