package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data string
	left *Node
	right *Node
}

func create_tree(arr []string) *Node {
	if len(arr)==0 {
		return nil
	}
	if !check_operator(arr[0]) {
		node := Node{arr[0], nil, nil}
		head := create_tree(arr[1:])
		if head == nil {
			return &node
		}
		head.left = &node
		return head
	}
		node := Node{arr[0], nil, nil}
		right := create_tree(arr[1:])
		node.right = right
		return &node

}
func check_operator(s string) bool {
	if s=="+" || s=="-" {
		return true
	}
	return false
}
func preorder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.data+ " ")
	preorder(node.left)
	preorder(node.right)
}
func postorder(node *Node) {
	if node == nil {
		return
	}
	postorder(node.left)
	postorder(node.right)
	fmt.Print(node.data+ " ")
}
func main() {
	s := "a + b - c"
	arr :=  strings.Split(s, " ")
	root := create_tree(arr)
	preorder(root)
	fmt.Println()
	postorder(root)
	fmt.Println()
}
