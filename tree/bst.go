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

func RSearch(root *Node, key int) bool {
	if root == nil {
		return false
	} else if root.data == key {
		return true
	} else if root.data > key {
		RSearch(root.left, key)
	} else {
		RSearch(root.right, key)
	}
}
