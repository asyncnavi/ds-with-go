package main

import "ds-go/tree"

func main() {
	t := tree.CreateTree()
	tree.Preorder(t)
}
