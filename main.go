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
	list.InsertAtFront(100)
	list.InsertAt(15, 3)
	list.PrintList()
	list.GetMid()
}
