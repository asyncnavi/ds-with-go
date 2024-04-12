package main

import "ds-go/queue"

func reverse(q *queue.LQueue[int]) {
	var value int
	if !q.IsEmpty() {
		value = q.Dequeue().(int)
		reverse(q)
		q.Enqueue(value)
	}
}

func main() {
	q := queue.LQueue[int]{}
	q.New()
	q.Enqueue(20)
	q.Enqueue(10)
	q.Enqueue(30)
	q.Enqueue(40)
	reverse(&q)
	q.Print()
}
