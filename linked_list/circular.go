package linked_list

import "fmt"

// singly circular list

type CList struct {
	data int
	head *Node
}

func (list *CList) PrintList() {
	current := list.head
	for current.next != list.head {
		fmt.Printf("%d, ", current.data)
		current = current.next
	}

	fmt.Printf("%d, ", current.data)
}

func (list *CList) PushEnd(data int) {

	if list.head == nil {
		newNode := &Node{data: data, next: nil}
		list.head = newNode
		list.head.next = newNode
		return
	}

	current := list.head

	newNode := &Node{data: data, next: list.head}
	for current.next != list.head {
		current = current.next
	}

	current.next = newNode

}

func (list *CList) PushBeg(data int) {

	if list.head == nil {
		newNode := &Node{data: data, next: nil}
		list.head = newNode
		list.head.next = newNode
		return
	}

	newNode := &Node{data: data, next: list.head}
	current := list.head
	for current.next != list.head {
		current = current.next
	}
	list.head = newNode
	current.next = newNode

}

func (list *CList) GetMid() (int, int, bool) {

	if list.head == nil {
		fmt.Printf("List is Empty")
	}

	slowCursor := list.head
	fastCursor := list.head

	for fastCursor.next != list.head && fastCursor.next.next != list.head {
		slowCursor = slowCursor.next
		fastCursor = fastCursor.next.next
	}

	if fastCursor.next != list.head {
		return slowCursor.data, slowCursor.next.data, false
	}
	return slowCursor.data, 0, true
}
