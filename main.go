package main

import (
	"ds-go/stack"
	"fmt"
)

func main() {

	fmt.Print(stack.PrefixToPost("+a*bc"))
}
