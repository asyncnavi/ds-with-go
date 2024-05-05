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
		fmt.Printf("%d ", node.data)

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
		fmt.Printf("%d ", root.data)
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
		fmt.Printf("%d ", curr.data)

		// Move to the right subtree
		curr = curr.right
	}
}

func InOrderR(root *Node) {
	if root != nil {
		PreOrderR(root.left)
		fmt.Printf("%d", root.data)
		PreOrderR(root.right)
	}
}
