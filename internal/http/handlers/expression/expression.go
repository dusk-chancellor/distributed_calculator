package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dusk-chancellor/distributed_calculator/internal/storage"
)

// Handlers for operations with expressions

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

// must be somewhere else than here
const (
	null   = "null"
	stored = "stored"
)

type ExpressionInteractor interface { // Methods for interactions with database
	InsertExpression(ctx context.Context, expr *storage.Expression) (int64, error)
	SelectExpressionsByID(ctx context.Context, userID int64) ([]storage.Expression, error)
	DeleteExpression(ctx context.Context, id int64) error
}

// CreateExpressionHandler - post method handler which stores an expression
func CreateExpressionHandler(ctx context.Context, expressionSaver ExpressionInteractor) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		date := time.Now()

		jsonDec := json.NewDecoder(r.Body)

		var req Request
		if err := jsonDec.Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userID, ok := r.Context().Value("userid").(int64)
		if !ok {
			http.Error(w, "userID not received", http.StatusBadRequest)
			log.Printf("userID not received: %d", userID)
			return
		}

		var expressionStruct = storage.Expression{
			UserID: userID,
			Expression: req.Expression,
			Answer: null,
			Date: date.Format("2006/01/02 15:04:05"),
			Status: stored,
		}

		id, err := expressionSaver.InsertExpression(ctx, &expressionStruct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		log.Printf("Successful CreateExpressionHandler operation; id = %d", id)
	}
}

// GetExpressionHandler - get method handler which writes all expressions from database
func GetExpressionsHandler(ctx context.Context, expressionSaver ExpressionInteractor) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		userID, ok := r.Context().Value("userid").(int64)
		if !ok {
			http.Error(w, "userID not received", http.StatusBadRequest)
			log.Printf("userID not received: %d", userID)
			return
		}

		allExpressions, err := expressionSaver.SelectExpressionsByID(ctx, userID)
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


		json.NewEncoder(w).Encode(respData)
		log.Print("Successful GetExpressionsHandler operation")
	}
}

// DeleteExpressionHandler
func DeleteExpressionHandler(ctx context.Context, expressionSaver ExpressionInteractor) http.HandlerFunc {
	
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = expressionSaver.DeleteExpression(ctx, int64(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		log.Print("Successful DeleteExpressionHandler operation")
	}
}