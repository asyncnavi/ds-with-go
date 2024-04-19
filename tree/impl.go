package tree

import (
	"bufio"
	"ds-go/queue"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func Preorder(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.data)
		Preorder(root.left)
		Preorder(root.right)
	}
}

func CreateTree() *Node {
	scanner := bufio.NewScanner(os.Stdin)
	q := queue.Queue[*Node]{}
	q.New(100)
	fmt.Println("Enter the root value:")
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	root := &Node{data: x, left: nil, right: nil}
	q.Enqueue(root)

	for q.IsEmpty() {
		p := q.Dequeue().(*Node)

		fmt.Println("Enter left child (or -1 to skip):")
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		if x != -1 {
			t := &Node{data: x, left: nil, right: nil}
			p.left = t
			q.Enqueue(t)
		}

		fmt.Println("Enter right child (or -1 to skip):")
		scanner.Scan()
		x, _ = strconv.Atoi(scanner.Text())
		if x != -1 {
			t := &Node{data: x, left: nil, right: nil}
			p.right = t
			q.Enqueue(t)
		}
	}

	return root
}
