package trees

import "fmt"

// TestBSTExecution : testing all the functions of BST
func TestBSTExecution() {
	fmt.Println("\n ************ Binary Search Tree *************")
	tree := NewBSTTree()
	tree.insert(5)
	tree.insert(10)
	tree.insert(20)
	tree.insert(2)
	printInorder(tree.root) //2 5 10 20
	fmt.Println("Found 100:", search(tree.root, 100))
	fmt.Println("Found 20: ", search(tree.root, 20))
	found, _ := delete(tree.root, 100)
	fmt.Println("\nDeleted 100:", found)
	printInorder(tree.root) //2 5 10 20
	found, _ = delete(tree.root, 5)
	fmt.Println("\nDeleted 5:", found)
	printInorder(tree.root) //2 10 20
	tree.insert(9)
	found, _ = delete(tree.root, 9)
	fmt.Println("\nDeleted 9:", found)
	printInorder(tree.root) //2 5 10 20

}
