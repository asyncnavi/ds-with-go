package main

import (
	"ds-go/stack"
	"fmt"
)

func main() {
	name := "({[})"

	fmt.Print(stack.MatchBraces(name))
}
