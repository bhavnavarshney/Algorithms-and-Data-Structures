package main

import (
	"fmt"

	"github.com/bhavnavarshney/Algorithms-and-Data-Structures/Go/hashing"
	"github.com/bhavnavarshney/Algorithms-and-Data-Structures/Go/linkedlist"
	"github.com/bhavnavarshney/Algorithms-and-Data-Structures/Go/searching"
	"github.com/bhavnavarshney/Algorithms-and-Data-Structures/Go/strings"
	"github.com/bhavnavarshney/Algorithms-and-Data-Structures/Go/trees"
)

func main() {
	var ch int
	for {
		fmt.Println("\n0: Exit 1. SinglyLinkedList 2: DoublyLinkedList 3. HashMap 4. BinarySearchTree 5. KMP 6. Searching")
		fmt.Scanf("%d", &ch)
		if ch == 0 {
			return
		}
		switch ch {
		case 1:
			linkedlist.TestSinglyLinkedList()
			break
		case 2:
			linkedlist.TestDoublyLinkedList()
			break
		case 3:
			hashing.TestHashing()
			break
		case 4:
			trees.TestBSTExecution()
			break
		case 5:
			strings.TestKMP()
			break
		case 6:
			searching.TestSearching()
			break
		}
	}

}
