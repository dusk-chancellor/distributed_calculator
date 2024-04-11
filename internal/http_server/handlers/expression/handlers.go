package handlers

import (
	"calculator_yandex/internal/storage"
	"context"
	"encoding/json"
	"log"
	"mime"
	"net/http"
	"time"
)

type Request struct {
	Expression string `json:"expression"`
}

type ResponseData struct {
	ID 		   int64  `json:"id"`
	Expression string `json:"expression"`
	Answer 	   string `json:"answer"`
	Date 	   string `json:"date"`
	Status 	   string `json:"status"`
}

type ExpressionSaver interface {
	InsertExpression(ctx context.Context, expr *storage.Expression) (int64, error)
	SelectExpressions(ctx context.Context) ([]storage.Expression, error)
}

func CreateExpressionHandler(ctx context.Context, expressionSaver ExpressionSaver) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
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

		id, err := expressionSaver.InsertExpression(ctx, &expressionStruct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")

		log.Printf("Successful CreateExpressionHandler operation; id = %d", id)
	}	
}

func GetExpressionsHandler(ctx context.Context, expressionSaver ExpressionSaver) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		allExpressions, err := expressionSaver.SelectExpressions(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var respData []ResponseData

		for _, expr := range allExpressions {
			resp := ResponseData{
				ID: expr.ID,
				Expression: expr.Expression,
				Answer: expr.Answer,
				Date: expr.Date,
				Status: expr.Status,
			}

			respData = append(respData, resp)
		}

		jsonData, err := json.Marshal(respData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)

		log.Print("Successful GetExpressionsHandler operation")
	}
}