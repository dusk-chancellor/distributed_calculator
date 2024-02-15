package handlers

import (
	"calculator_yandex/internal/storage"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func SetTimeoutsHandler(w http.ResponseWriter, r *http.Request) {
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

	timeout := values.Get("timeout") // получаем значение выражения
	if timeout == "" {
		http.Error(w, "Missing 'timeout' parameter in request body", http.StatusBadRequest)
		return
	}

	if err := storage.SetNewTimeout(timeout); err != nil { // добавляем
		http.Error(w, "Could not set new expression", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Timeout added successfully"))
}

func GetTimeoutsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { // только GET!
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	timeouts, err := storage.GetTimeouts() // получаем выражения
	if err != nil {
		http.Error(w, "Failed to retrieve expressions"+err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(timeouts) // маршалируем
	if err != nil {
		http.Error(w, "Failed to marshal expressions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData) // ставим в хедер тип данных и записываем в врайтер
}
