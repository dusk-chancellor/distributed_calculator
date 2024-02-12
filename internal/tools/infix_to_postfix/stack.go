package infix_to_postfix

import "strings"

// Stack наш с вами стек для элементов
// top - самый последний элемент в стеке
// size - думаю понадобится в дальнейшем при синхронизации
type Stack struct {
	top  *Element
	size int
}

// Element структура для каждого элемента в стеке
// value понятно - само значение
// next - замена верхнего элемента последующим (нужна только для Pop())
type Element struct {
	value interface{}
	next  *Element
}

// Empty есть там че нить или нет, проверяем
func (s *Stack) Empty() bool {
	return s.size == 0
}

// Top самый верхний элемент стека
func (s *Stack) Top() interface{} {
	return s.top.value
}

// Push заносим элемент в стек
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Pop выносим элемент из стека
// да знаю можно было срезАть слайс, но я без слайса реализовал и из-за этого костыли :)
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

// IsOperator проверка на + - * /
func IsOperator(c uint8) bool {
	return strings.ContainsAny(string(c), "+ & - & * & /")
}

// IsOperand проверка на цифры
func IsOperand(c uint8) bool {
	return c >= '0' && c <= '9'
}

// Precedence для определения порядка выполнения арифметических действий
func Precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return -1
}

// HasHigherPrecedence определяет какая из операций имеет высший приоритет
func HasHigherPrecedence(op1, op2 string) bool {
	op1Prec := Precedence(op1)
	op2Prec := Precedence(op2)
	return op1Prec >= op2Prec
}
