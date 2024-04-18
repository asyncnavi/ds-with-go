package tree

import (
	"bufio"
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
		fmt.Println("%d ", root.data)
		Preorder(root.left)
		Preorder(root.right)
	}
}

func CreateTree() *Node {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the root value:")
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	root := &Node{data: x, left: nil, right: nil}

	q := []*Node{root}

	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		fmt.Println("Enter left child (or -1 to skip):")
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		if x != -1 {
			t := &Node{data: x, left: nil, right: nil}
			p.left = t
			q = append(q, t)
		}

		fmt.Println("Enter right child (or -1 to skip):")
		scanner.Scan()
		x, _ = strconv.Atoi(scanner.Text())
		if x != -1 {
			t := &Node{data: x, left: nil, right: nil}
			p.right = t
			q = append(q, t)
		}
	}

	return root
}
