package datas

type Stack[T any] interface {
	Push(value T)
	Pop() *T
	Peek() *T
	IsEmpty() bool
	Flush() []T
	FlushReverse() []T
}

type element[T any] struct {
	value T
	next  *element[T]
}

type stack[T any] struct {
	head *element[T]
}

func NewStack[T any](value *T) Stack[T] {
	if value == nil {
		return &stack[T]{
			head: nil,
		}
	}
	return &stack[T]{
		head: &element[T]{
			value: *value,
			next:  nil,
		},
	}
}

func (s *stack[T]) Push(value T) {
	var nh = element[T]{
		value: value,
		next:  s.head,
	}

	s.head = &nh
}

func (s *stack[T]) Pop() *T {
	if s.head == nil {
		return nil
	}
	res := s.head.value
	s.head = s.head.next
	return &res
}

func (s *stack[T]) Peek() *T {
	if s.head == nil {
		return nil
	}
	res := s.head.value
	return &res
}

func (s *stack[T]) IsEmpty() bool {
	return s.Peek() == nil
}

func (s *stack[T]) Flush() (result []T) {
	for s.Peek() != nil {
		result = append(result, *s.Pop())
	}

	return result
}

func (s *stack[T]) FlushReverse() (result []T) {
	if s.IsEmpty() {
		return []T{}
	}
	v := s.Pop()
	if s.IsEmpty() {
		result = append(result, *v)
		return result
	}
	result = s.FlushReverse()
	result = append(result, *v)
	return result
}
