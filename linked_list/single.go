package linked_list

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LList struct {
	Head *Node
}

// Printing the list

func (list *LList) PrintList() {
	current := list.Head // First Element
	fmt.Print("Linked List : [ ")
	for current != nil {
		if current.next == nil {
			fmt.Printf("%d", current.data)
		} else {
			fmt.Printf("%d, ", current.data)
		}
		current = current.next
	}
	fmt.Print(" ]\n")
}

// Check if is empty or not

func (list *LList) IsEmpty() bool {
	return list.Head == nil
}

/*
		GetOrPrintSize  print or return the size of Linked List
		if `p = true` print and return
	    else return
*/
func (list *LList) GetOrPrintSize(p bool) int {

	count := 0
	current := list.Head

	for current != nil {
		count++
		current = current.next
	}
	if p {
		fmt.Printf("List Size : %d \n", count)
		return count
	}

	return count
}

// Inserting at back

func (list *LList) InsertAtBack(data int) {
	newNode := &Node{data: data, next: nil}
	if list.IsEmpty() {
		list.Head = newNode
		return
	}

	current := list.Head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

// Insert at front

func (list *LList) InsertAtFront(data int) {
	// if list is empty
	if list.Head == nil {
		list.Head = &Node{data: data, next: nil}
		return
	}

	newNode := &Node{data: data, next: list.Head}

	list.Head = newNode
}

// Insert at a specific index

func (list *LList) InsertAt(data, index int) {
	if list.Head == nil && index != 0 {
		fmt.Println("Linked List is empty cannot add to specific index rather than 0")
		return
	}

	if index == 0 {
		list.InsertAtFront(data)
		return
	}

	current := list.Head

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
	slowPtr := list.Head
	fastPtr := list.Head

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

	if list.Head.next == nil {
		list.Head = nil
		return
	}
	current := list.Head

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

	if list.Head.next == nil {
		list.Head = nil
		return
	}

	list.Head = list.Head.next

}

// Delete : The middle Element

func (list *LList) DeleteAtMiddle() {
	if list.IsEmpty() {
		return
	}

	if list.Head.next == nil {
		list.Head = nil
		return
	}

	var prev *Node
	slowPtr := list.Head
	fastPtr := list.Head

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

	if list.Head.next == nil {
		list.Head = nil
		return
	}

	current := list.Head

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

// Reverse Linked List

func (list *LList) Reverse() {
	current := list.Head
	var prev *Node = nil
	var next *Node = nil
	for current != nil {
		next = current.next
		current.next = prev

		prev = current
		current = next
	}
	list.Head = prev
}
