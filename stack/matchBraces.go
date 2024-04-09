package stack

import "fmt"

func MatchBraces(expr string) bool {
	st := Stack[rune]{}
	maxChar := len(expr)
	st.New(maxChar)

	for i := 0; i < maxChar; i++ {
		fmt.Println(st.Peek())
		if st.IsEmpty() {
			st.Push(rune(expr[i]))
		} else if (st.Peek() == '(' && rune(expr[i]) == ')') || (st.Peek() == '{' && rune(expr[i]) == '}') || (st.Peek() == '[' && rune(expr[i]) == ']') {
			popped := st.Pop()
			fmt.Print(popped)
		} else {
			st.Push(rune(expr[i]))
		}
		st.Print()

	}
	if st.IsEmpty() {
		return true
	}

	return false
}
