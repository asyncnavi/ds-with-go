package main

import (
	"ds-go/linked_list"
	"fmt"
)

func main() {
	list := &linked_list.LList{}
	list.InsertAtBack(1)
	list.InsertAtBack(9)
	list.InsertAtBack(9)
	list.InsertAtBack(9)
	list.InsertAtBack(8)
	list.InsertAtBack(10)

	fmt.Println(linked_list.SumOfLastNNode(list, 5))
}
