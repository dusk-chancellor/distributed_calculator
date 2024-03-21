package main

// TODO: REST API
// [ ] POST   /expressions/    - creates an expression, an id for it and saves it
// [ ] GET    /expressions/    - returns a list of all expressions
// [ ] GET    /expressions/:id - returns specific expression with specified id
// [ ] DELETE /expressions/    - deletes all expressions
// [ ] DELETE /expressions/:id - deletes expression by its id

import (
	"calculator_yandex/storage"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"time"
)

type expressionServer struct {
	store *storage.ExpressionStore
}

func NewExprServer() *expressionServer {
	return &expressionServer{
		store: storage.NewStorage(),
	}
}

func (es *expressionServer) setExpressionHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling expression set on %s\n", r.URL.Path)
	date := time.Now()

	contentType := r.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		log.Printf("ParseMediaType: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediaType != "application/json" {
		log.Print("mediaType is not json")
		http.Error(w, "expected Content-Type to be application/json", http.StatusUnsupportedMediaType)
		return
	}

	jsonDec := json.NewDecoder(r.Body)
	jsonDec.DisallowUnknownFields()

	type RequestExpression struct {
		Expression string `json:"expression"`
	}

	var re RequestExpression
	if err := jsonDec.Decode(&re); err != nil {
		log.Printf("Could not decode: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	es.store.SetExpression(re.Expression, date, "stored")
	w.WriteHeader(http.StatusOK)
}

func (es *expressionServer) getAllExpressionsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling get all expressions on %s\n", r.URL.Path)

	allExprs := es.store.GetAllExpressions()
	jsoned, err := json.Marshal(allExprs)
	if err != nil {
		log.Printf("Marshalling json: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsoned)
}

func main() {
	mux := http.NewServeMux()
	server := NewExprServer()

	mux.Handle("/", http.FileServer(http.Dir(".")))

	mux.HandleFunc("POST /expression/", server.setExpressionHandler)
	mux.HandleFunc("GET /expression/", server.getAllExpressionsHandler)

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		return
	}
}
