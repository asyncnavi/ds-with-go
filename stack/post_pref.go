package stack

func IsOperator(ch string) bool {
	switch ch {
	case "+", "-", "*", "/":
		return true
	}
	return false
}

func PostToPrefix(expr string) string {
	st := &Stack[string]{}
	st.New(len(expr))

	for _, ch := range expr {
		if IsOperator(string(ch)) {
			op1 := st.Pop().(string)
			op2 := st.Pop().(string)

			res := string(ch) + op2 + op1
			st.Push(res)

		} else {
			st.Push(string(ch))
		}

	}

	res := ""

	for !st.IsEmpty() {
		res = res + st.Peek().(string)
		st.Pop()
	}

	return res
}
