package linkedlist

import "fmt"

type DoublyNode struct {
	Val  int
	Next *DoublyNode
	Prev *DoublyNode
}

type DoublyLinkedList struct {
	Head *DoublyNode
	Len  int
}

func (l *DoublyLinkedList) Print(msg string) {
	pointer := l.Head
	fmt.Println("\n*** PRINTING:", msg, " ****\n")
	for pointer != nil {
		fmt.Print(pointer.Prev, "[", pointer.Val, "]", pointer.Next, "| -> |")
		pointer = pointer.Next
	}
}

func (l *DoublyLinkedList) InsertAtEnd(val int) {
	n := &DoublyNode{Val: val, Next: nil, Prev: nil}
	if l.Head == nil {
		l.Head = n
		return
	}

	pointer := l.Head
	previous := pointer
	for pointer != nil {
		previous = pointer
		pointer = pointer.Next
	}
	previous.Next = n
	n.Prev = previous
}

func (l *DoublyLinkedList) InsertAtStart(val int) {
	n := &DoublyNode{Val: val, Prev: nil, Next: nil}
	if l.Head == nil {
		l.Head = n
		return
	}

	l.Head.Prev = n
	tempHead := l.Head
	l.Head = n
	l.Head.Next = tempHead
}

func (l *DoublyLinkedList) DeleteAtStart() {
	if l.Head == nil {
		return
	}
	previous := l.Head.Prev
	l.Head = l.Head.Next
	l.Head.Prev = previous
}

func (l *DoublyLinkedList) InsertAtIndex(index, val int) {
	count := 0
	n := &DoublyNode{Val: val, Next: nil, Prev: nil}
	pointer := l.Head
	previous := pointer

	for ; pointer != nil; count++ {
		if count == index {
			n.Next = previous.Next
			n.Prev = previous
			previous.Next = n
			pointer.Prev = n
			return
		}
		previous = pointer
		pointer = pointer.Next

	}
}

func (l *DoublyLinkedList) DeleteAtEnd() {
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

func (l *DoublyLinkedList) DeleteAtIndex(index int) {
	count := 0
	pointer := l.Head
	previous := pointer

	for ; pointer != nil; count++ {
		if count == index {
			previous.Next = pointer.Next
			pointer.Prev = previous
		}
		previous = pointer
		pointer = pointer.Next
	}
}
