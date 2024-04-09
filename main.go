package main

import "ds-go/stack"

func main() {

	st := stack.LStack{}
	st.New()
	st.Push(20)
	st.Push(50)
	st.Push(30)
	st.Print()
}
