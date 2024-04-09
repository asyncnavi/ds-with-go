package main

import "ds-go/stack"

func main() {

	st := stack.Stack[rune]{}
	st.New(10)
	st.Push('c')
	st.Push('x')
	st.Push('v')
	st.Push('t')

	st.Print()

}
