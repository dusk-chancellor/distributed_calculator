package manage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Expression struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Date       string `json:"date"`
	Status     string `json:"status"`
	Answer     string `json:"answer"`
}

type Expressions struct {
	Expressions []*Expression `json:"expressions"`
}

const expressionsFile = "database/expressions.json"

func ReceiveAnswer(id, answer, receivedError string) error {
	exprs, err := getFromExpressionsFile(expressionsFile)
	if err != nil {
		return err
	}

	for _, expr := range exprs.Expressions {
		if expr.ID == id {
			if receivedError == "error" {
				expr.Status = "undefined (division by zero)"
				return saveToFile(expressionsFile, exprs)
			} else {
				expr.Status = "done"
				expr.Answer = answer
				return saveToFile(expressionsFile, exprs)
			}
		}
	}
	return err
}

func saveToFile(fileName string, val interface{}) error {
	data, err := json.Marshal(val)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fileName, data, 0644); err != nil {
		return err
	}

	return nil
}

func getFromExpressionsFile(fileName string) (*Expressions, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	var expressions Expressions
	if err := json.Unmarshal(data, &expressions); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	return &expressions, nil
}
