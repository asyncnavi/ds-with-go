package queue

import (
	"fmt"
)

type Queue[T any] struct {
	entries []T
	front   int
	rear    int
	max     int
}

func (q *Queue[T]) New(max int) {
	q.rear = -1
	q.front = -1
	q.max = max
	q.entries = make([]T, max)
}

func (q *Queue[T]) Print() {

	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("%v, ", q.entries[i])
	}
}

func (q *Queue[T]) Enqueue(data T) {
	if q.max == q.rear {
		fmt.Println("Queue is full")
		return
	}
	q.rear++
	q.entries[q.rear] = data
}

func (q *Queue[T]) Dequeue() interface{} {
	if q.rear == q.front {
		fmt.Println("Queue is empty")
		return nil
	}
	q.front = (q.front + 1) % q.max
	item := q.entries[q.front]
	return item
}

func (q *Queue[T]) IsEmpty() bool {
	return q.rear == q.front
}
