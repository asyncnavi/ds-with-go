package linked_list


import "fmt"


type DNode struct {
	prev *DNode
	data int
	next *DNode
}

type DList struct {
	head *DNode
}

func (list *DList) PrintList() {
	cursor := list.head

	for cursor != nil {
		fmt.Printf("%d ,", cursor.data)
		cursor = cursor.next
	}
}

func (list *DList) InsertAtEnd(data int) {
  
	if list.head == nil {
		list.head = &DNode{prev: nil, data: data, next: nil}
		return
	}

	cursor := list.head

	for cursor.next != nil {
		cursor = cursor.next
	}
	newNode := &DNode{prev: cursor, data: data, next: nil}
	cursor.next = newNode
}
