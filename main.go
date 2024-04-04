package main

import (
	"ds-go/linked_list"
)

func main() {
	list := &linked_list.LList{}
	list1 := &linked_list.LList{}
	list.InsertAtBack(1)
	list.InsertAtBack(9)
	list.InsertAtBack(8)
	list.InsertAtBack(9)

	list1.InsertAtBack(1)
	list1.InsertAtBack(9)

	linked_list.AddTwoList(list, list1).PrintList()
}
