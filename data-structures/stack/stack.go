package stack

import "fmt"

type StackGeneric[T any] struct {
	Data []T
	Cur  int
}

func New[T any]() *StackGeneric[T] {
	return &StackGeneric[T]{
		Data: make([]T, 0),
		Cur:  0,
	}
}

func (s *StackGeneric[T]) Push(item T) {
	s.Data = append(s.Data, item)
	s.Cur++
}

// 1-2-3, cur = 3
// pop -> (3 | [1, 2])
func (s *StackGeneric[T]) Pop() (*T, error) {
	if len(s.Data) == 0 {
		return nil, fmt.Errorf("cannot pop empty stack")
	}
	item := s.Data[s.Cur-1]
	s.Data = s.Data[:s.Cur-1]
	s.Cur--
	return &item, nil
}
