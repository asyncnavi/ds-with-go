package main

import (
	"ds-go/linked_list"
)

func main() {
	list := &linked_list.LList{}
	list.InsertAtBack(1)
	list.InsertAtBack(9)
	list.InsertAtBack(8)
	list.InsertAtBack(9)
	list.InsertAtBack(6)
	list.InsertAtBack(7)

	list.PrintList()
	linked_list.PairWiseSwap(list)
	list.PrintList()
}
