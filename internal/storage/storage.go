package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Expression struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Status     string `json:"status"`
}

type Expressions struct {
	Expressions []*Expression `json:"expressions"`
}

const fileName = "expressions.json"

var logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

func NewExpressionsList() *Expressions {
	return &Expressions{Expressions: []*Expression{}}
}

func getFromFile() (*Expressions, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		logger.Printf("File does not exist: %s", err)
		return NewExpressionsList(), nil
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		logger.Printf("Cannot read the file: %s", err)
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	if len(data) == 0 {
		logger.Println("File is empty, returning new Expressions object")
		return NewExpressionsList(), nil
	}

	var expressions Expressions
	if err := json.Unmarshal(data, &expressions); err != nil {
		logger.Printf("Cannot unmarshal data to expression: %s", err)
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	logger.Printf("Expressions list retrieved from %s", fileName)
	return &expressions, nil
}

func GetStoredExpressions() (*Expressions, error) {
	return getFromFile()
}

func SetNewExpression(expr string) error {
	exprs, err := getFromFile()
	if err != nil {
		logger.Printf("Cannot getFromFile: %s", err)
		return err
	}

	newID := fmt.Sprintf("%d", len(exprs.Expressions)+1)
	newExpr := &Expression{
		ID:         newID,
		Expression: expr,
		Status:     "-",
	}
	exprs.Expressions = append(exprs.Expressions, newExpr)

	return saveToFile(exprs)
}

func saveToFile(exprs *Expressions) error {
	data, err := json.Marshal(exprs)
	if err != nil {
		logger.Printf("Failed to marshal expressions: %v", err)
		return err
	}

	if err := ioutil.WriteFile(fileName, data, 0644); err != nil {
		logger.Printf("Failed to write to file: %v", err)
		return err
	}
	logger.Printf("Successfully wrote to file: %s", fileName)

	return nil
}

func ClearExpressionsList() error {
	defaultState := NewExpressionsList()

	jsonDefault, err := json.Marshal(defaultState)
	if err != nil {
		logger.Printf("Failed to marshal default state: %v", err)
		return err
	}

	err = ioutil.WriteFile(fileName, jsonDefault, 0644)
	if err != nil {
		logger.Printf("Failed to write default state to file: %v", err)
		return err
	}
	logger.Printf("Expressions list cleared")
	return nil
}
