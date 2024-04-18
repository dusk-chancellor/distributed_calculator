package infix_to_postfix

import (
	"strings"

)

// ToPostfix - функция превращающая обычные выражения в постфиксную запись
// например: принимает s = "2+ 2 *2" -> возвращает "2 2 2 * +"
// с пробелами для того, чтобы можно было значения больше 9 тоже считать
func ToPostfix(s string) string {
	// объявления стека, постфиксной строки
	var stack Stack
	postfix := ""
	length := len(s)

	for i := 0; i < length; i++ {
		char := string(s[i]) // элемент строки

		if char == " " { // пробелы нам не нужны
			continue
		}

		if char == "(" { // скобки открываем, в стек пушим
			stack.Push(char)
		} else if char == ")" { // скобки закрываем
			for !stack.Empty() { // пока стек не освободится идет цикл,
				// где мы забираем самый верхний элемент и
				// добавляем его в наш постфикс внутри скобки как бы
				str, _ := stack.TopFunc().(string)
				if str == "(" { // когда дойдем до открывавшейся скобочки, то bye bye
					break
				}
				postfix += " " + str
				stack.Pop()
			}
			stack.Pop()
		} else if !IsOperator(s[i]) { // если элемент - это число
			j := i
			number := ""

			for ; j < length && IsOperand(s[j]); j++ { // проходимся по нему
				// чтобы записать его полностью
				number += string(s[j])
			}
			postfix += " " + number // в постфикс
			i = j - 1
		} else { // ни одно условие не прошло? значит это оператор(наверное)!
			for !stack.Empty() {
				top, _ := stack.TopFunc().(string)
				if top == "(" || !HasHigherPrecedence(top, char) { // проверим на всякий
					break
				}
				postfix += " " + top
				stack.Pop()
			} // поп после записи, пуш после цикленка
			stack.Push(char)
		}
	}
	for !stack.Empty() { // дописываем оставшиеся элементы (они у нас уже в красивой очереди стоят)
		str, _ := stack.Pop().(string)
		postfix += " " + str
	}
	return strings.TrimSpace(postfix) // чтобы           вот таких пробелов не было в начале и конце
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