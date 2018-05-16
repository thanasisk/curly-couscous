package main

import "fmt"
import "math/rand"

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (root *Node) insert(newNode *Node) {
	if root.key == newNode.key {
		return
	}
	if root.key < newNode.key {
		if root.right == nil {
			root.right = newNode
		} else {
			root.right.insert(newNode)
		}
	} else {
		if root.left == nil {
			root.left = newNode
		} else {
			root.left.insert(newNode)
		}
	}
}

func (root *Node) inOrder(fn func(n *Node)) {

	if root.left != nil {
		root.left.inOrder(fn)
	}

	fn(root)

	if root.right != nil {
		root.right.inOrder(fn)
	}

}

type Bst struct {
	root   *Node
	length int
}

func (bst *Bst) insert(value int) {
	node := &Node{key: value}
	bst.length++

	if bst.root == nil {
		bst.root = node
	} else {
		bst.root.insert(node)
	}
}

func (bst *Bst) inorderPrint() {
	bst.root.inOrder(func(node *Node) {
		fmt.Println(node.key)
	})
}
func main() {
	tree := new(Bst)
	for i := 0; i < 0x20; i++ {
		tree.insert(rand.Int())
	}
	tree.inorderPrint()
}
