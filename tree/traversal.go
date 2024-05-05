package tree

import "fmt"

func PreOrderR(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.data)
		PreOrderR(root.left)
		PreOrderR(root.right)
	}
}

func InOrderR(root *Node) {
	if root != nil {
		PreOrderR(root.left)
		fmt.Printf("%d", root.data)
		PreOrderR(root.right)
	}
}
