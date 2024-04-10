package stack

func prec(s string) int {
	switch s {
	case "^":
		return 3
	case "/", "*":
		return 2
	case "+", "-":
		return 1
	default:
		return -1
	}
}

func InfixToPost(code string) string {
	st := Stack[string]{}
	st.New(len(code))
	postfix := ""

	for _, char := range code {
		opchar := string(char)

		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			postfix = postfix + opchar
		} else if char == '(' {
			st.Push(opchar)
		} else if char == ')' {
			for st.Peek().(string) != "(" {
				postfix = postfix + st.Peek().(string)
				st.Pop()
			}
			st.Pop()
		} else {
			for !st.IsEmpty() && prec(opchar) <= prec(st.Peek().(string)) {
				postfix = postfix + st.Peek().(string)
				st.Pop()
			}
			st.Push(opchar)
		}
	}

	for !st.IsEmpty() {
		postfix = postfix + st.Peek().(string)
		st.Pop()
	}
	return postfix

}
