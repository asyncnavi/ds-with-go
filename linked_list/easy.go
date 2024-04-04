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

// Q7 Make middle node the head node

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

// Q 8

func Reverse(list *LList) {
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

// Q 9

func AddOneUtil(list *LList) {
	// Reverse the list
	Reverse(list)

	cursor := list.Head
	var prev *Node
	carry := 1
	var sum int
	for cursor != nil {
		sum = carry + cursor.data

		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		sum = sum % 10
		cursor.data = sum
		prev = cursor
		cursor = cursor.next
	}

	if carry > 0 {
		prev.next = &Node{data: carry, next: nil}
	}

	Reverse(list)
}

// Q 10

func getNodeValue(node *Node) int {
	if node != nil {
		return node.data
	} else {
		return 0
	}
}

func AddTwoList(l1 *LList, l2 *LList) *LList {
	Reverse(l1)
	Reverse(l2)

	curr1 := l1.Head
	curr2 := l2.Head
	carry := 0
	var sum int
	var temp *Node
	var prev *Node
	res := &LList{}
	for curr1 != nil || curr2 != nil {
		sum = carry + getNodeValue(curr1) + getNodeValue(curr2)
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		sum = sum % 10
		temp = &Node{data: sum, next: nil}

		if res.Head == nil {
			res.Head = temp
		} else {
			prev.next = temp
		}
		prev = temp

		if curr1 != nil {
			curr1 = curr1.next
		}

		if curr2 != nil {
			curr2 = curr2.next
		}
	}

	if carry > 0 {
		temp.next = &Node{data: carry, next: nil}
	}
	Reverse(res)
	Reverse(l1)
	Reverse(l2)
	return res
}
