package fetch

import (
	"encoding/json"
	"io"
	"net/http"
)

type Data struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Status     string `json:"status"`
}

func FetchExpressionFromClientServer() (Data, error) {
	url := "http://localhost:8081/getexpressionslist"

	resp, err := http.Get(url)
	if err != nil {
		return Data{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Data{}, err
	}

	var expressionsList Data
	if err := json.Unmarshal(body, &expressionsList); err != nil {
		return Data{}, err
	}

	return expressionsList, nil
}
