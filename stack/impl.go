package stack

import "fmt"

type Stack struct {
	top   int
	max   int
	nodes []int
}

func (stack *Stack) Init(max int) {
	stack.top = -1
	stack.max = max
	stack.nodes = make([]int, max)
}

func (stack *Stack) Print() {
	for i := 0; i <= stack.top; i++ {
		fmt.Printf("%d, ", stack.nodes[i])
	}
	fmt.Println()
}

func (stack *Stack) Push(data int) {
	if stack.top >= stack.max-1 {
		fmt.Println("Stack overflow")
		return
	}
	stack.top++
	stack.nodes[stack.top] = data
}

func (stack *Stack) Pop() int {
	if stack.top == -1 {
		fmt.Println("Stack is empty")
		return -1
	}
	popped := stack.nodes[stack.top]
	stack.top--
	return popped
}

func (stack *Stack) Peek() int {
	if stack.top == -1 {
		fmt.Println("Stack is empty")
		return -1
	}
	return stack.nodes[stack.top]
}

func (stack *Stack) IsEmpty() bool {
	return stack.top == -1
}
