package tree

import (
	"ds-go/stack"
	"fmt"
)

func PreOrder(root *Node) {
	if root == nil {
		return
	}

	st := &stack.LStack{}
	st.New()

	st.Push(root)

	for !st.IsEmpty() {
		node := st.Pop().(*Node)
		fmt.Printf("%d ", node.Data)

		// Push the right child first so that it gets processed after the left child
		if node.right != nil {
			st.Push(node.right)
		}
		if node.left != nil {
			st.Push(node.left)
		}
	}
}

func PreOrderR(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.Data)
		PreOrderR(root.left)
		PreOrderR(root.right)
	}
}
func InOrder(root *Node) {
	if root == nil {
		return
	}

	st := &stack.LStack{}
	st.New()

	curr := root

	for curr != nil || !st.IsEmpty() {
		// Reach the leftmost node of the current subtree
		for curr != nil {
			st.Push(curr)
			curr = curr.left
		}

		// Current node is nil, so we pop from stack
		curr = st.Pop().(*Node)
		fmt.Printf("%d ", curr.Data)

		// Move to the right subtree
		curr = curr.right
	}
}

func InOrderR(root *Node) {
	if root != nil {
		PreOrderR(root.left)
		fmt.Printf("%d", root.Data)
		PreOrderR(root.right)
	}
}

func LevelTraversal(root *Node) {
	if root == nil {
		return
	}

	q := []*Node{root}

	for len(q) > 0 {
		levelSize := len(q)

		for i := 0; i < levelSize; i++ {
			node := q[0]

			// Enqueue current( first ) element
			q := q[1:]

			if node == nil {
				fmt.Println("*")
			} else {
				fmt.Printf("%d", node.Data)

				q = append(q, node.left)
				q = append(q, node.right)
			}
		}
	}
	fmt.Println()
}
