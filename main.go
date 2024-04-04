package main

import (
	"ds-go/linked_list"
)

func main() {
	list := &linked_list.LList{}
	list.InsertAtBack(40)
	list.InsertAtBack(20)
	list.InsertAtBack(10)
	list.InsertAtBack(5)
	list.InsertAtBack(10)

	list.PrintList()
	linked_list.NthFromEnd(list, 5)
	linked_list.MoveLastToFront(list)
	list.PrintList()
}
