package calculation

import (
	"fmt"
	"strconv"
	"strings"
)

func Evaluate(expr string) (float64, error) {
	var stack Stack
	tokens := strings.Split(expr, " ")

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			// если токен - оператор, то забираем 2 последних элемента со стека
			op1 := stack.Pop()
			op2 := stack.Pop()
			ans, err := Calculate(op1, op2, token)
			if err != nil {
				return 0, err
			}
			stack.Push(ans)
		} else {
			// если токен не оператор - то операнд(число), пушим в стек :)
			op, _ := strconv.ParseFloat(token, 64)

			stack.Push(op)
		}
	}
	// в этом случае, последний элемент стека и есть наш ответ
	return stack.Pop(), nil
}

// Calculate - вычисляет
func Calculate(op1, op2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return op2 + op1, nil
	case "-":
		return op2 - op1, nil
	case "*":
		return op2 * op1, nil
	case "/":
		if op1 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return op2 / op1, nil
	default:
		return 0, fmt.Errorf("unknown operation")
	}
}