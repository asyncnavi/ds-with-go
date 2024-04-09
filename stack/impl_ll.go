package stack

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LStack struct {
	top *Node
}

func (s *LStack) New() {
	s.top = nil
}

func (s *LStack) Print() {
	cursor := s.top
	fmt.Print("Stack:")
	for cursor != nil {
		fmt.Print(cursor.data, " ")
		cursor = cursor.next
	}
	fmt.Println()
}

func (s *LStack) Push(data interface{}) {
	newNode := &Node{data: data, next: s.top}
	s.top = newNode
}

func (s *LStack) Pop() interface{} {
	if s.IsEmpty() {
		fmt.Println("Stack is underflow")
		return nil
	}
	popped := s.top.data
	s.top = s.top.next
	return popped
}

func (s *LStack) Peak() interface{} {
	if s.IsEmpty() {
		fmt.Println("Stack is underflow")
		return nil
	}

	data := s.top.data
	return data
}

func (s *LStack) IsEmpty() bool {
	return s.top == nil
}
