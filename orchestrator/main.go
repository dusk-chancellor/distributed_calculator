package main

import (
	"calculator_yandex/orchestrator/fetch"
	"fmt"
)

func main() {
	data, err := fetch.FetchExpressionFromClientServer()
	if err != nil {
		return
	}
	fmt.Println(data)
}
