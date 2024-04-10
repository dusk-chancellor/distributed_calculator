package handlers

import (
	"calculator_yandex/internal/storage"
	"encoding/json"
	"mime"
	"net/http"
	"time"
)

type Request struct {
	Expression string `json:"expression"`
}


func CreateExpressionHandler(w http.ResponseWriter, r *http.Request) {
	
	date := time.Now()

	contentType := r.Header.Get("Content-Type")

	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediaType != "application/json" {
		http.Error(w, "expected Content-Type to be application/json", http.StatusUnsupportedMediaType)
		return
	}

	jsonDec := json.NewDecoder(r.Body)
	jsonDec.DisallowUnknownFields()

	var req Request
	if err := jsonDec.Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var expressionStruct = storage.Expression{
		Expression: req.Expression,
		Date: date.Format("2006/01/02 15:04:05"),
		Status: "stored",
	}
	
}