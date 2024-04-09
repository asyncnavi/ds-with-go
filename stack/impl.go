package stack

import (
	"fmt"
)

type Stack[T any] struct {
	top   int
	max   int
	items []T
}

func (s *Stack[T]) New(max int) {
	s.top = -1
	s.max = max
	s.items = make([]T, max)
}

func (s *Stack[T]) Print() {
	top := s.top
	fmt.Print("Stack: [")

	for top != -1 {
		fmt.Printf("%v, ", s.items[top])
		top--
	}
	fmt.Print("]")

}

func (s *Stack[T]) Push(item T) {
	if s.top >= s.max-1 {
		fmt.Println("Stack is overflowed")
		return
	}

	s.top++
	s.items[s.top] = item
}

func (s *Stack[T]) Pop() interface{} {
	if s.top == -1 {
		fmt.Println("Stack is underflow")
		return nil
	}
	popped := s.items[s.top]
	s.top--
	return popped
}

func (s *Stack[T]) Peek() interface{} {
	if s.top == -1 {
		fmt.Println("Stack is underflow")
		return nil
	}

	return s.items[s.top]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == -1
}
