package main

import "fmt"

type node struct {
	value string
	left  *node
	right *node
}

func newNode(val string) *node {
	return &node{value: val, left: nil, right: nil}
}
func preOrder(root *node) {
	if root == nil {
		return
	}
	fmt.Printf(root.value + " ")
	preOrder(root.left)
	preOrder(root.right)
}
func postOrder(root *node) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Printf(root.value + " ")

}
func inOrder(root *node) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Printf(root.value + " ")
	inOrder(root.right)
}
func main() {
	nodeA := newNode("a")
	nodeB := newNode("b")
	nodeC := newNode("c")
	nodePlus := newNode("+")
	nodeMinus := newNode("-")
	root := nodePlus
	nodePlus.left = nodeA
	nodePlus.right = nodeMinus
	nodeMinus.left = nodeB
	nodeMinus.right = nodeC
	fmt.Println("Inorder")
	inOrder(root)
	fmt.Println("")
	fmt.Println("PreOrder")
	preOrder(root)
	fmt.Println("")
	fmt.Println("PostOrder")
	postOrder(root)
	fmt.Println("")

}
