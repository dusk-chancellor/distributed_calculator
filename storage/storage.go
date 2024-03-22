package storage

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

type Expression struct {
	ID     string `json:"id"`
	Expr   string `json:"expression"`
	Date   string `json:"date"`
	Status string `json:"status"`
}

type ExpressionStore struct {
	sync.Mutex

	exprs  map[int]Expression
	id int
}

func NewStorage() *ExpressionStore {
	return &ExpressionStore{
		exprs:  make(map[int]Expression),
		id: 0,
	}
}

func (e *ExpressionStore) SetExpression(expr string, date time.Time, status string) {
	e.Lock()
	defer e.Unlock()

	e.id++
	formattedDate := date.Format("2006/01/02 15:04:05")
	newExpr := Expression{
		ID: e.generateID(),
		Expr: expr,
		Date: formattedDate,
		Status: status,
	}

	e.exprs[e.id] = newExpr
	log.Print("Successfully stored an expression")
}

func (e *ExpressionStore) GetExpression(id int) (Expression, error) {
	e.Lock()
	defer e.Unlock()

	expr, ok := e.exprs[id]
	if !ok {
		return Expression{}, fmt.Errorf("expression with id #%v not found", id)
	}
	return expr, nil
}

func (e *ExpressionStore) GetAllExpressions() []Expression {
	e.Lock()
	defer e.Unlock()

	allExprs := make([]Expression, 0, len(e.exprs))
	for _, expr := range e.exprs {
		allExprs = append(allExprs, expr)
	}
	return allExprs
}

func (e *ExpressionStore) DeleteExpression(id int) error {
	e.Lock()
	defer e.Unlock()

	if _, ok := e.exprs[id]; !ok {
		return fmt.Errorf("expression with id #%v not found", id)
	}

	delete(e.exprs, id)
	log.Print("Expression deleted")
	return nil
}

func (e *ExpressionStore) DeleteAllExpressions() {
	e.Lock()
	defer e.Unlock()

	e.exprs = make(map[int]Expression)
	e.id = 0
}

func (e *ExpressionStore) generateID() string {
	switch {
	case e.id < 10:
		return "00" + strconv.Itoa(e.id)
	case e.id < 100:
		return "0" + strconv.Itoa(e.id)
	default:
		return strconv.Itoa(e.id)
	}
}

