package lib

type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		elements: map[T]struct{}{},
	}
}

func (s *Set[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

func (s *Set[T]) Exists(element T) bool {
	_, ok := s.elements[element]
	return ok
}

func (s *Set[T]) Delete(element T) {
	delete(s.elements, element)
}

func (s *Set[T]) Slice() []T {
	var ret []T
	for key := range s.elements {
		ret = append(ret, key)
	}
	return ret
}
