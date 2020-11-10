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
	l.Print()            // 100 3 5 200
	l.DeleteAtIndex(2)   // 100 3 200
	l.DeleteAtStart()    // 3 200
	l.InsertAtStart(100) // 100 3 200
	l.DeleteAtEnd()      // 100 3
	l.Print()

	fmt.Println("\n***** Doubly LinkedList *******")
	l2 := linkedlist.DoublyLinkedList{}
	l2.InsertAtEnd(5)     //5
	l2.InsertAtStart(100) //100 5
	l2.Print("Insertion at Start & End")
	l2.InsertAtEnd(200) // 100 5 200
	l2.Print("Insertion at End")
	l2.DeleteAtStart()      // 5 200
	l2.InsertAtIndex(1, 10) // 5 10 200
	l2.Print("Deletion at Start and Insert at Index")
	l2.DeleteAtEnd()
	l2.Print("Delete at End") // 5 10
	l2.DeleteAtIndex(1)
	l2.Print("Delete at Index 1") // 5

}
