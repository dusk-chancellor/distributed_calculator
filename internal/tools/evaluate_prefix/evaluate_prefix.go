// этот файл сделан на время проверки работы вычисления префиксов
// позже, если конечно успею до дедлайна, то
// это все дело пойдет агентам-вычислителям

package evaluate_prefix

import (
	"strconv"
	"strings"
)

// Stack наш стек для работы с постфиксами
type Stack struct {
	items []float64
}

// Push пушит числа в стек
func (s *Stack) Push(item float64) {
	s.items = append(s.items, item)
}

// Pop забирает (буквально) элемент со стека
func (s *Stack) Pop() float64 {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Calculate - функция на время проверки работоспособности постфиксной записи
func Calculate(op1, op2 float64, operator string) float64 {
	switch operator {
	case "+":
		return op2 + op1
	case "-":
		return op2 - op1
	case "*":
		return op2 * op1
	case "/":
		return op2 / op1
	}
	return 0
}

// EvaluatePostfix - пока что функция для вычисления всех значений со стека
// но в будущем координально изменится под работу для агентов (воркер)
func EvaluatePostfix(expression string) float64 {
	var stack Stack
	tokens := strings.Split(expression, " ")

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			// если токен - оператор, то забираем 2 последних элемента со стека
			op1 := stack.Pop()
			op2 := stack.Pop()
			res := Calculate(op1, op2, token)
			stack.Push(res)
		} else {
			// если токен не оператор - то операнд(число), пушим в стек :)
			op, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0
			}
			stack.Push(op)
		}
	}
	// в этом случае, последний элемент стека и есть наш ответ
	return stack.Pop()
}
