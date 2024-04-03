package linked_list

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LList struct {
	head *Node
}

// Printing the list

func (list *LList) PrintList() {
	current := list.head // First Element

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

// Check if is empty or not

func (list *LList) IsEmpty() bool {
	return list.head == nil
}

// Count the number of element

func (list *LList) Count() {

	count := 0
	current := list.head

	for current != nil {
		count++
		current = current.next
	}

	fmt.Println("Count : ")
}

// Inserting at back

func (list *LList) InsertAtBack(data int) {
	newNode := &Node{data: data, next: nil}
	if list.IsEmpty() {
		list.head = newNode
		return
	}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

// Insert at front

func (list *LList) InsertAtFront(data int) {
	// if list is empty
	if list.head == nil {
		list.head = &Node{data: data, next: nil}
		return
	}

	newNode := &Node{data: data, next: list.head}

	list.head = newNode
}

// Insert at a specific index

func (list *LList) InsertAt(data, index int) {
	if list.head == nil && index != 0 {
		fmt.Println("Linked List is empty cannot add to specific index rather than 0")
		return
	}

	if index == 0 {
		list.InsertAtFront(data)
		return
	}

	current := list.head

	for i := 0; i < index-1; i++ {
		if current.next != nil {
			current = current.next
		} else {
			fmt.Println("Invalid index")
			return
		}
	}

	if current.next == nil {
		current.next = &Node{data: data, next: nil}
		return
	}

	newNode := &Node{data: data, next: current.next}

	current.next = newNode
}

// Find the middle

func (list *LList) GetMid() int {
	slowPtr := list.head
	fastPtr := list.head

	for fastPtr != nil && fastPtr.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
	}

	return slowPtr.data
}

// Delete : At the end

func (list *LList) RemoveAtEnd() {
	if list.IsEmpty() {
		return
	}

	if list.head.next == nil {
		list.head = nil
		return
	}
	current := list.head

	for current.next.next != nil {
		current = current.next
	}

	current.next = nil
}

// Delete:  from Front

func (list *LList) DeleteFromFront() {
	if list.IsEmpty() {
		return
	}

	if list.head.next == nil {
		list.head = nil
		return
	}

	list.head = list.head.next

}

// Delete : The middle Element

func (list *LList) DeleteAtMiddle() {
	if list.IsEmpty() {
		return
	}

	if list.head.next == nil {
		list.head = nil
		return
	}

	var prev *Node
	slowPtr := list.head
	fastPtr := list.head

	for fastPtr != nil && fastPtr.next != nil {
		prev = slowPtr
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
	}

	prev.next = slowPtr.next

}

// Delete : at specific index

func (list *LList) RemoveAt(index int) {
	if list.IsEmpty() {
		fmt.Println("No node to delete : List is empty")
		return
	}

	if list.head.next == nil {
		list.head = nil
		return
	}

	current := list.head

	for i := 0; i < index-1; i++ {
		if current.next != nil {
			current = current.next
		} else {
			fmt.Println("Invalid index")
			return
		}
	}

	if current.next.next == nil {
		current.next = nil
		return
	}

	current.next = current.next.next
}
