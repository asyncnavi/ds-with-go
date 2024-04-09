package stack

func PrefixToPost(expr string) string {
	st := &Stack[string]{}
	st.New(len(expr))

	for i := len(expr) - 1; i >= 0; i-- {
		if IsOperator(string(expr[i])) {
			op1 := st.Pop().(string)
			op2 := st.Pop().(string)

			res := op1 + op2 + string(expr[i])
			st.Push(res)

		} else {
			st.Push(string(expr[i]))
		}

	}

	return st.Peek().(string)
}
