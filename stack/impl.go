package stack

import "fmt"

type Stack[C any] struct {
	top   int
	max   int
	nodes []C
}

func (stack *Stack[T]) Init(max int) {
	stack.top = -1
	stack.max = max
	stack.nodes = make([]T, max)
}

func (stack *Stack[T]) Print() {
	for i := 0; i <= stack.top; i++ {
		fmt.Print(stack.nodes[i])
	}
	fmt.Println()
}

func (stack *Stack[T]) Push(data T) {
	if stack.top >= stack.max-1 {
		fmt.Println("Stack overflow")
		return
	}
	stack.top++
	stack.nodes[stack.top] = data
}

func (stack *Stack[T]) Pop() {
	if stack.top == -1 {
		fmt.Println("Stack is underflow")
		return
	}
	stack.top--
}

func (stack *Stack[T]) Peek() {
	if stack.top == -1 {
		fmt.Println("Stack is underflow")
		return
	}
	println(stack.nodes[stack.top])
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.top == -1
}
