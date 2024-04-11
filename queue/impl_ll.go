package queue

type Node[T any] struct {
	data T
	next *Node[T]
}

type LQueue struct {
	front int
	rear  int
}
