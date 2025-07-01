package generics

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	elementIdx := len(s.elements) - 1
	element := s.elements[elementIdx]
	s.elements = s.elements[:elementIdx]
	return element, true
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Len() int {
	return len(s.elements)
}
