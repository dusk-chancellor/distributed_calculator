package storage

import (
	"calculator_yandex/http-server/orchestrator_handlers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
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

var logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

func NewExpressionsList() *Expressions {
	return &Expressions{Expressions: []*Expression{}}
}

func SetNewExpression(expr string) error {
	exprs, err := getFromExpressionsFile(expressionsFile)
	if err != nil {
		logger.Printf("Cannot getFromExpressionsFile: %s", err)
		return err
	}
	var newExpr Expression
	for _, ch := range expr {
		if ch >= '9' && ch <= '0' && (ch != '+' && ch != '-' && ch != '*' && ch != '/') {
			logger.Printf("Wrong Expression Format")
			newExpr = Expression{
				ID:         "",
				Expression: expr,
				Status:     "invalid",
				Date:       time.Now().String(),
			}
			exprs.Expressions = append(exprs.Expressions, &newExpr)

			return saveToFile(expressionsFile, exprs)
		}
	}
	if err := orchestrator_handlers.PostExpressionToAgent(expr); err != nil {
		return err
	}

	newID := fmt.Sprintf("%d", len(exprs.Expressions)+1)
	newExpr = Expression{
		ID:         newID,
		Expression: expr,
		Status:     "in process",
		Date:       time.Now().String(),
	}

	exprs.Expressions = append(exprs.Expressions, &newExpr)

	return saveToFile(expressionsFile, exprs)
}

func getFromExpressionsFile(fileName string) (*Expressions, error) {
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
		logger.Printf("Cannot unmarshal data: %s", err)
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	logger.Printf("List retrieved from %s", expressionsFile)
	return &expressions, nil
}

func GetStoredExpressions() (*Expressions, error) {
	return getFromExpressionsFile(expressionsFile)
}
