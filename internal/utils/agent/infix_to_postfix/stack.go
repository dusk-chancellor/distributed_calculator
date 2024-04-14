package infix_to_postfix

type Stack struct {
	Top  *Element
	Size int
}

// Element структура для каждого элемента в стеке
// value понятно - само значение
// next - замена верхнего элемента последующим (нужна только для Pop())
type Element struct {
	Value interface{}
	Next  *Element
}

// Empty есть там че нить или нет, проверяем
func (s *Stack) Empty() bool {
	return s.Size == 0
}

// Top самый верхний элемент стека
func (s *Stack) TopFunc() interface{} {
	return s.Top.Value
}

// Push заносим элемент в стек
func (s *Stack) Push(value interface{}) {
	s.Top = &Element{value, s.Top}
	s.Size++
}

// Pop выносим элемент из стека
// да знаю можно было срезАть слайс, но я без слайса реализовал и из-за этого костыли :)
func (s *Stack) Pop() (value interface{}) {
	if s.Size > 0 {
		value, s.Top = s.Top.Value, s.Top.Next
		s.Size--
		return
	}
	return nil
}