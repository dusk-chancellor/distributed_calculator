package handlers

import (
	"calculator_yandex/internal/storage"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func GetExpressionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	exprs, err := storage.GetStoredExpressions()
	if err != nil {
		http.Error(w, "Failed to retrieve expressions"+err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(exprs)
	if err != nil {
		http.Error(w, "Failed to marshal expressions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}

func SetExpressionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Expected POST", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	values, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		http.Error(w, "Failed to parse query", http.StatusInternalServerError)
		return
	}

	expr := values.Get("expr")
	if expr == "" {
		http.Error(w, "Missing 'expr' parameter in request body", http.StatusBadRequest)
		return
	}

	if err := storage.SetNewExpression(expr); err != nil {
		http.Error(w, "Could not set new expression", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Expression added successfully"))
}

func ClearExpressionsHandler(w http.ResponseWriter, r *http.Request) {
	err := storage.ClearExpressionsList()
	if err != nil {
		http.Error(w, "Could not clear expressions list", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Expressions list cleared successfully"))
}
