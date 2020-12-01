package trees

import "fmt"

// Node : Holds individual node
type Node struct {
	left  *Node
	right *Node
	data  int
}

// BST : Holds the Tree
type BST struct {
	root *Node
}

// NewBSTTree : Gives new BST Tree
func NewBSTTree() *BST {
	return &BST{
		root: nil,
	}
}

func (b *BST) insert(key int) {
	newNode := &Node{data: key}
	if b.root == nil {
		b.root = newNode
		return
	}

	_ = insertNode(b.root, key)
}

func insertNode(root *Node, key int) *Node {
	if root == nil {
		return &Node{data: key}
	}
	if key < root.data {
		root.left = insertNode(root.left, key)
	} else {
		root.right = insertNode(root.right, key)
	}
	return root
}

func printInorder(root *Node) {
	if root != nil {
		printInorder(root.left)
		fmt.Print(root.data, "->")
		printInorder(root.right)
	}
}

func search(root *Node, key int) bool {
	if root == nil {
		return false
	}
	if root.data == key {
		return true
	}
	if key < root.data {
		search(root.left, key)
	}
	return search(root.right, key)
}

func delete(root *Node, key int) (bool, *Node) {
	var found bool
	if root == nil {
		return false, nil
	}
	if root.data == key {
		if root.left == nil && root.right == nil {
			return true, nil
		}
		if root.left != nil && root.right == nil {
			return true, root.left
		}
		if root.right != nil && root.left == nil {
			return true, root.right
		}
		if root.right != nil && root.left != nil {
			successor := getInorderSuccessor(root.right)
			root.data = successor.data
			_, root.right = delete(root.right, root.data)
			return true, root
		}
	}

	if key < root.data {
		found, root.left = delete(root.left, key)
		if found {
			return found, root
		}
	} else {
		found, root.right = delete(root.right, key)
		if found {
			return found, root
		}
	}
	return false, root
}

func getInorderSuccessor(root *Node) *Node {
	for root != nil && root.left != nil {
		root = root.left
	}
	return root
}
