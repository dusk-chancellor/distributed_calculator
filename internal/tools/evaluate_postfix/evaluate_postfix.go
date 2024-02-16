package evaluate_postfix

import (
	"strconv"
	"strings"
)

// этот файл сделан на время проверки работы вычисления префиксов
// позже, если конечно успею до дедлайна, то
// это все дело пойдет агенту-вычислителю

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

// Calculate - вычисляет
func Calculate(ch chan float64, op1, op2 float64, operator string) {
	switch operator {
	case "+":
		ch <- op2 + op1
	case "-":
		ch <- op2 - op1
	case "*":
		ch <- op2 * op1
	case "/":
		if op1 == 0 {
			ch <- -1
			return
		}
		ch <- op2 / op1
	}
	ch <- 0
}

// EvaluatePostfix - для вычисления всех значений со стека
func EvaluatePostfix(expression string) float64 {
	var stack Stack
	tokens := strings.Split(expression, " ")
	ch := make(chan float64)

	go func() {
		for _, token := range tokens {
			if token == "+" || token == "-" || token == "*" || token == "/" {
				// если токен - оператор, то забираем 2 последних элемента со стека
				op1 := stack.Pop()
				op2 := stack.Pop()
				go Calculate(ch, op1, op2, token)
				stack.Push(<-ch)
			} else {
				// если токен не оператор - то операнд(число), пушим в стек :)
				op, _ := strconv.ParseFloat(token, 64)

				stack.Push(op)
			}
		}
	}()
	// в этом случае, последний элемент стека и есть наш ответ
	return stack.Pop()
}
