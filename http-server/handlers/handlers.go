package handlers

import (
	"calculator_yandex/internal/storage"
	"net/http"
)

// ClearExpressionsHandler - разом стирает все данные с базы данных (моего json файлика ахахха)
func ClearExpressionsHandler(w http.ResponseWriter, r *http.Request) {
	err := storage.ClearExpressionsList() // взял и стер
	if err != nil {
		http.Error(w, "Could not clear expressions list", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Expressions list cleared successfully")) // все четко
}
