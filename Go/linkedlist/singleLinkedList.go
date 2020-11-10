package linkedlist

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type SinglyLinkedList struct {
	Head *Node
	Len  int
}

func (l *SinglyLinkedList) Print() {
	pointer := l.Head
	fmt.Println("*** PRINTING ****")
	for pointer != nil {
		fmt.Println(pointer.Val)
		pointer = pointer.Next
	}
}

func (l *SinglyLinkedList) InsertAtIndex(index, val int) {
	count := 0
	n := &Node{Val: val, Next: nil}
	pointer := l.Head
	previous := pointer

	for ; pointer != nil; count++ {
		if count == index {
			n.Next = previous.Next
			previous.Next = n
			return
		}
		previous = pointer
		pointer = pointer.Next

	}
}

func (l *SinglyLinkedList) InsertAtEnd(val int) {
	n := &Node{Val: val, Next: nil}
	if l.Head == nil {
		l.Head = n
		return
	}

	pointer := l.Head
	for pointer.Next != nil {
		pointer = pointer.Next
	}
	pointer.Next = n
}

func (l *SinglyLinkedList) InsertAtStart(val int) {
	n := &Node{Val: val, Next: nil}
	if l.Head == nil {
		l.Head = n
		return
	}

	tempHead := l.Head
	l.Head = n
	l.Head.Next = tempHead
}

func (l *SinglyLinkedList) DeleteAtStart() {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
}

func (l *SinglyLinkedList) DeleteAtEnd() {
	if l.Head == nil {
		return
	}
	pointer := l.Head
	previous := pointer
	for pointer.Next != nil {
		previous = pointer
		pointer = pointer.Next
	}
	previous.Next = nil
}

func (l *SinglyLinkedList) DeleteAtIndex(index int) {
	count := 0
	pointer := l.Head
	previous := pointer

	for ; pointer != nil; count++ {
		if count == index {
			previous.Next = pointer.Next
		}
		previous = pointer
		pointer = pointer.Next
	}
}
