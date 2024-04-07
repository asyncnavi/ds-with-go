package main

import (
	"ds-go/linked_list"
)

func main() {
	list := &linked_list.CList{}
	list.PushEnd(5)
	list.PushEnd(6)
	list.PushEnd(7)
	list.PushBeg(8)
	list.PushBeg(10)
	list.PushBeg(11)

	list.PrintList()
}
