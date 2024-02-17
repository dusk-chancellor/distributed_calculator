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

func ReceiveAnswer(id, answer string) error {
	exprs, err := getFromExpressionsFile(expressionsFile)
	if err != nil {
		return err
	}

	var updateExpr Expression
	for _, expr := range exprs.Expressions {
		if expr.ID == id {
			updateExpr = Expression{
				ID:         id,
				Expression: expr.Expression,
				Date:       expr.Date,
				Status:     "done",
				Answer:     answer,
			}
			if err := deleteFromFile(id); err != nil {
				return err
			}
			exprs.Expressions = append(exprs.Expressions, &updateExpr)
			return saveToFile(expressionsFile, exprs)
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

func deleteFromFile(id string) error {
	exprs, err := getFromExpressionsFile(expressionsFile)
	if err != nil {
		return err
	}

	index := -1
	for i, expr := range exprs.Expressions {
		if expr.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("ID not found: %s", id)
	}

	exprs.Expressions = append(exprs.Expressions[:index], exprs.Expressions[index+1:]...)

	return saveToFile(expressionsFile, exprs)
}
