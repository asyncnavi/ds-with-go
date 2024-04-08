package main

import (
	"ds-go/stack"
)

func main() {
	st := &stack.Stack{}
	st.Init(10)
	st.Push(2)
	st.Push(5)
	st.Push(8)
	st.Peek()

}
