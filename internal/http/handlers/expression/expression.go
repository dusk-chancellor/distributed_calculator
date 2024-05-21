package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dusk-chancellor/distributed_calculator/internal/storage"
	"github.com/dusk-chancellor/distributed_calculator/internal/utils/orchestrator/manager"
	_ "github.com/joho/godotenv/autoload"
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
var (
	null   = "null"
	stored = "stored"
)

// CreateExpressionHandler - post method handler which stores an expression
func CreateExpressionHandler(ctx context.Context, expressionSaver storage.ExpressionInteractor) http.HandlerFunc {

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

		go manager.Manage(ctx, expressionSaver, agentAddress())

		w.WriteHeader(http.StatusCreated)
		log.Printf("Successful CreateExpressionHandler operation; id = %d", id)
	}
}

// GetExpressionHandler - get method handler which writes all expressions from database
func GetExpressionsHandler(ctx context.Context, expressionSaver storage.ExpressionInteractor) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		userID, ok := r.Context().Value("userid").(int64)
		if !ok {
			http.Error(w, "userID not received", http.StatusBadRequest)
			log.Printf("userID not received: %d", userID)
			return
		}

		go manager.Manage(ctx, expressionSaver, agentAddress())

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
func DeleteExpressionHandler(ctx context.Context, expressionSaver storage.ExpressionInteractor) http.HandlerFunc {
	
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		go manager.Manage(ctx, expressionSaver, agentAddress())

		pathValues := strings.Split(r.URL.Path, "/")
		if len(pathValues) < 3 || pathValues[2] == "" {
			http.Error(w, "Invalid URL format", http.StatusBadRequest)
			return
		}

		expressionID, err := strconv.ParseInt(pathValues[2], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = expressionSaver.DeleteExpression(ctx, int64(expressionID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		log.Print("Successful DeleteExpressionHandler operation")
	}
}

func agentAddress() string {
	agentHost, ok := os.LookupEnv("AGENT_HOST")
	if !ok {
		log.Print("AGENT_HOST not set, using 0.0.0.0")
		agentHost = "0.0.0.0"
	}

	agentPort, ok := os.LookupEnv("AGENT_PORT")
	if !ok {
		log.Print("AGENT_PORT not set, using 5000")
		agentPort = "5000"
	}
	return fmt.Sprintf("%s:%s", agentHost, agentPort)
}