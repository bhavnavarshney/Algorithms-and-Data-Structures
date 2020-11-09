package main

import (
	"fmt"

	"github.com/bhavnavarshney/Algorithms-and-Data-Structures/Go/linkedlist"
)

func main() {
	fmt.Println("***** Singly LinkedList *****")
	l := linkedlist.SinglyLinkedList{}
	l.InsertAtEnd(5)
	l.InsertAtStart(100)
	l.InsertAtEnd(200)
	l.InsertAtIndex(1, 3)
	l.Print()          // 100 3 5 200
	l.DeleteAtIndex(2) // 100 3 200
	l.DeleteAtStart()  // 3 200
	l.DeleteAtEnd()    // 3
	l.Print()

}
