package main

import (
	"ds-go/linked_list"
	"fmt"
)

func main() {
	list := &linked_list.CList{}
	list.PushEnd(5)
	list.PushEnd(6)
	list.PushEnd(7)
	list.PushEnd(8)
	list.PushEnd(4)
	list.PushEnd(10)

	list.PrintList()
	firstMid, secondMid, isSecondNull := list.GetMid()

	if isSecondNull {
		fmt.Println("\nMiddle node of the list is :", firstMid)
	} else {
		fmt.Println("\nSince the list has even number of node so it has two mid: ", firstMid, secondMid)
	}

}
