package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ExpressionReq struct {
	Expression string `json:"expression"`
}

func SendExpressionHandler(w http.ResponseWriter, r *http.Request) {
	var req ExpressionReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return
	}

	fmt.Println("Received:", req.Expression)

	w.Write([]byte("Expression received"))
}

func main() {
	http.HandleFunc("/sendexpression", SendExpressionHandler)
	http.Handle("/", http.FileServer(http.Dir("frontend")))

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}
}