package calculation

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