package structures

type Stack[T any] struct {
	elements []T
}

func NewStack[T any](elements ...T) *Stack[T] {
	return &Stack[T]{
		elements: elements,
	}
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() T {
	lastIndex := len(s.elements) - 1
	ret := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return ret
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Peek() T {
	lastIndex := len(s.elements) - 1
	return s.elements[lastIndex]
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}
