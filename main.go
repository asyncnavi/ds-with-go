package main

import (
	"ds-go/stack"
	"fmt"
)

func main() {

	fmt.Print(stack.InfixToPost("a+b*(c^d-e)"))
}
