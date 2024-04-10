package stack

func changeBracket(exp *string) {
	r := []rune(*exp)

	for i := 0; i < len(r); i++ {
		if r[i] == '(' {
			r[i] = ')'
		} else if r[i] == ')' {
			r[i] = '('
		}
	}

	*exp = string(r)
}

func InfixToPre(code string) string {
	newStr := Reverse(code)

	changeBracket(&newStr)

	postFix := InfixToPost(newStr)

	return Reverse(postFix)
}
