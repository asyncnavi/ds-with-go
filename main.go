package main

import "ds-go/queue"

func main() {
	q := queue.LQueue[string]{}
	q.New()
	q.Enqueue("Hello")
	q.Enqueue("Navraj")
	q.Enqueue("Sandhu")
	println(q.Dequeue().(string))
	q.Print()
}
