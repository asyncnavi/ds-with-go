package stack

func Reverse(s string) string {
	res := ""
	st := Stack[rune]{}
	maxLen := len(s)
	st.New(maxLen)

	for i := 0; i < maxLen; i++ {
		st.Push(rune(s[i]))
	}

	for !st.IsEmpty() {
		popped := st.Pop().(rune)
		res += string(popped)
	}
	return res
}
