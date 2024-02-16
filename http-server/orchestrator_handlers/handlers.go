package orchestrator_handlers

import (
	"bytes"
	"net/http"
)

func PostExpressionToAgent(expr string) error {
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8081/setexpr", bytes.NewBufferString(expr))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "text/plain")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
