package main

import (
	"ds-go/linked_list"
)

func main() {
	list := &linked_list.DList{}
	list.InsertAtEnd(5)
	list.InsertAtEnd(6)
	list.InsertAtEnd(7)
	list.InsertAtEnd(8)
	list.InsertAtEnd(4)
	list.InsertAtEnd(10)

	list.PrintList()

}
