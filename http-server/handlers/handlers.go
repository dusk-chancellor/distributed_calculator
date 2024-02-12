package handlers

import (
	"calculator_yandex/internal/storage"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// GetExpressionsHandler - метод получения всех данных сразу и их запись в врайтер
func GetExpressionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { // только GET!
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	exprs, err := storage.GetStoredExpressions() // получаем выражения
	if err != nil {
		http.Error(w, "Failed to retrieve expressions"+err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(exprs) // маршалируем
	if err != nil {
		http.Error(w, "Failed to marshal expressions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData) // ставим в хедер тип данных и записываем в врайтер
}

// SetExpressionHandler - чтобы юзер мог добавлять свои выражения curl запросом
// пока реализовано так, в дальнейшем будет полноценный фронт-енд (надеюсь)
func SetExpressionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { //только POST!
		http.Error(w, "Invalid request method. Expected POST", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body) // читаем че там юзер прислал
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	values, err := url.ParseQuery(string(bodyBytes)) // парсим че там
	if err != nil {
		http.Error(w, "Failed to parse query", http.StatusInternalServerError)
		return
	}

	expr := values.Get("expr") // получаем значение выражения
	if expr == "" {
		http.Error(w, "Missing 'expr' parameter in request body", http.StatusBadRequest)
		return
	}

	if err := storage.SetNewExpression(expr); err != nil { // добавляем
		http.Error(w, "Could not set new expression", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Expression added successfully")) // в хедер - все ок по httpски, в врайтер - все ок по человечески
}

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
