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

func PrintTree(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node == nil {
				fmt.Print(" ʘ ")
			} else {
				fmt.Printf(" %d ", node.data)

				// Enqueue the left and right children of the current node
				queue = append(queue, node.left)
				queue = append(queue, node.right)
			}
		}

		// Move to the next line after printing a level
		fmt.Println()
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

	for !q.IsEmpty() {
		p := q.Dequeue().(*Node)

		fmt.Printf("Enter left child for %d (or -1 to skip):", p.data)
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		if x != -1 {
			t := &Node{data: x, left: nil, right: nil}
			p.left = t
			q.Enqueue(t)
		}

		fmt.Printf("Enter right child for %d (or -1 to skip):", p.data)
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
