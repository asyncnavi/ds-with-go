package main

import "ds-go/queue"

func main() {
	q := queue.Queue[string]{}
	q.New(10)
	q.Enqueue("Hello")
	q.Enqueue("Navraj")
	q.Enqueue("Sandhu")

	q.Print()
}
