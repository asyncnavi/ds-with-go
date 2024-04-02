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

// Print Linked List

func (list *LList) PrintList() {
	current := list.head // First Element

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

// check if is empty or not

func (list *LList) IsEmpty() bool {
	return list.head == nil
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

// Insert at

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

	// check if the index already exists or not

	if current.next == nil {
		current.next = &Node{data: data, next: nil}
		return
	}
	// Creating new Node
	newNode := &Node{data: data, next: current.next}

	current.next = newNode
}

// find middle

func (list *LList) GetMid() int {
	slowPtr := list.head
	fastPtr := list.head

	for fastPtr != nil && fastPtr.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
	}

	return slowPtr.data
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

// Reversing the linkedlist
