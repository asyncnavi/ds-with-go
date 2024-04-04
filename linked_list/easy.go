package linked_list

import "fmt"

// Q1 Check if the linked list are identical or no

func IsIdentical(l1 *LList, l2 *LList) bool {
	current1 := l1.Head
	current2 := l2.Head

	for current1 != nil && current2 != nil {
		if current1.data != current2.data {
			return false
		}

		current1 = current1.next
		current2 = current2.next
	}
	return true
}
func IsIdenticalR(l1 *Node, l2 *Node) bool {

	if l1 == nil && l2 == nil {
		return true
	}

	if l1 != nil && l2 != nil {
		return l1.data == l2.data && IsIdenticalR(l1.next, l2.next)
	}

	return true
}

// Q2 Print the middle of the given linked list

func PrintMiddle(list *LList) {
	sPtr := list.Head
	fPtr := list.Head

	for fPtr.next != nil && fPtr.next.next != nil {
		sPtr = sPtr.next
		fPtr = fPtr.next.next
	}

	fmt.Printf("Middle of the linked list is : %d", sPtr.data)
}

// Q3 Get the nth element

func GetNth(list *LList, n int) {
	if list.Head == nil {
		fmt.Println("List is empty")
	}

	current := list.Head
	for i := 0; i < n-1; i++ {
		if current.next == nil || current.next.next == nil {
			fmt.Println("Invalid Index")
			return
		}

		current = current.next
	}
	fmt.Printf("Element at position %d is %d", n, current.data)
}

// Q4

func NthFromEnd(list *LList, n int) {
	if list.IsEmpty() {
		fmt.Println("List is empty")
	}

	size := list.GetOrPrintSize(false)
	current := list.Head
	for i := 0; i < size-n; i++ {
		if current.next == nil || current.next.next == nil {
			fmt.Println("Invalid Index")
			return
		}
		current = current.next
	}

	fmt.Printf("%d", current.data)
}

// Q 5

func MoveLastToFront(list *LList) {
	current := list.Head
	prev := list.Head
	for current.next != nil {
		prev = current
		current = current.next
	}

	prev.next = nil
	current.next = list.Head
	list.Head = current

}

// Q 6 Flip tail and head

func FlipTailAndHead(list *LList) {
	current := list.Head
	prev := list.Head

	for current != nil && current.next != nil {
		prev = current
		current = current.next
	}

	current.next = list.Head.next
	list.Head.next = nil
	prev.next = list.Head
	list.Head = current

}

// Q7
func MakeMiddleNodeHead(list *LList) {

	if list.Head == nil || list.Head.next == nil {
		return
	}

	var prev *Node
	slowPtr := list.Head
	fastPtr := list.Head

	for fastPtr != nil && fastPtr.next.next != nil {
		prev = slowPtr
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
	}
	prev.next = prev.next.next
	slowPtr.next = list.Head
	list.Head = slowPtr

}
