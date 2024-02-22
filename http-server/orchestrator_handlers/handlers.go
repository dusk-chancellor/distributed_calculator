package orchestrator_handlers

import (
	"calculator_yandex/internal/manage"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetAnswerHandler(w http.ResponseWriter, r *http.Request) {
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

	idAndExprAndError := strings.Split(string(body), ":")
	log.Println("GetAnswerHandler -", idAndExprAndError)

	if err := manage.ReceiveAnswer(idAndExprAndError[0], idAndExprAndError[1], idAndExprAndError[2]); err != nil {
		http.Error(w, "Cannot receive answer from agent", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[GetAnswerHandler]Expression received and processed"))
}
