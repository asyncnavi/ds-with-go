package queue

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LQueue[T any] struct {
	front *Node[T]
	rear  *Node[T]
}

func (q *LQueue[T]) New() {
	q.front = nil
	q.rear = nil
}

func (q *LQueue[T]) Print() {
	if q.front == nil {
		fmt.Println("Queue is empty")
		return
	}

	current := q.front

	for current != nil {
		fmt.Printf("%v, ", current.data)
		current = current.next
	}
}

func (q *LQueue[T]) Enqueue(data T) {
	newNode := &Node[T]{data: data, next: nil}

	if newNode == nil {
		fmt.Println("Cannot add more element since queue is full")
		return
	}
	if q.front == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
}

func (q *LQueue[T]) Dequeue() interface{} {
	if q.front == nil {
		fmt.Println("Queue is empty")
		return nil
	}
	next := q.front.next
	data := q.front.data
	q.front = next
	return data
}
