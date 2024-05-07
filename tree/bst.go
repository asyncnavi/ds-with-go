package tree

func search(root *Node, key int) bool {

	for root != nil {
		if root.data > key {
			root = root.left
		} else if root.data < key {
			root = root.right
		} else {
			return true
		}
	}

	return false
}
