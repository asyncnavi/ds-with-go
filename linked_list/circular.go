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
