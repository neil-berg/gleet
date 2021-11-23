package datastructures

import (
	"errors"
	"fmt"
)

// Node is an element on the list
type Node struct {
	Next  *Node
	Prev  *Node
	Value interface{}
}

// List is the doubly-link list
type List struct {
	Head *Node
	Tail *Node
	Len  int
}

// New returns an empty list ready for use
func New() *List {
	return &List{}
}

// ValueAt returns the node's value at a given index
func (l *List) ValueAt(index int) interface{} {
	if l.Head == nil {
		return nil
	}
	curr := l.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	return curr.Value
}

// AddToTail adds a new value to the tail of the list
func (l *List) AddToTail(v interface{}) *List {
	node := &Node{Next: nil, Prev: nil, Value: v}

	if l.Head == nil {
		l.Head = node
	} else {
		l.Tail.Next = node
		node.Prev = l.Tail
	}
	l.Tail = node
	l.Len++
	return l
}

// AddToHead adds a new value to the head of the list
func (l *List) AddToHead(v interface{}) *List {
	node := &Node{Value: v}

	if l.Head == nil {
		l.Head = node // Initial node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head.Prev = node
		l.Head = node
	}
	l.Len++
	return l
}

// InsertAt adds a node at the given index
func (l *List) InsertAt(v interface{}, index int) (*List, error) {
	node := &Node{Value: v}

	if index >= l.Len || index < 0 {
		err := errors.New("Index out of range")
		return nil, err
	}

	if index == 0 {
		return l.AddToHead(v), nil
	}

	if index == l.Len-1 {
		return l.AddToTail(v), nil
	}

	curr := l.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	node.Next = curr
	node.Prev = curr.Prev
	curr.Prev.Next = node
	curr.Prev = node

	l.Len++
	return l, nil
}

// RemoveAtHead removes the head node
func (l *List) RemoveAtHead() *List {
	if l.Head == nil {
		return l
	}

	l.Head = l.Head.Next
	l.Head.Prev = nil
	l.Len--
	return l
}

// RemoveAtTail removes the tail node
func (l *List) RemoveAtTail() *List {
	if l.Head == nil {
		return l
	}

	l.Tail = l.Tail.Prev
	l.Tail.Next = nil
	l.Len--
	return l
}

// RemoveAt removes a node at a given index
func (l *List) RemoveAt(index int) (*List, error) {
	if index >= l.Len || index < 0 {
		err := errors.New("invalid index")
		return l, err
	}

	if index == 0 {
		return l.RemoveAtHead(), nil
	}

	if index == l.Len-1 {
		return l.RemoveAtTail(), nil
	}

	curr := l.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	curr.Prev.Next = curr.Next
	curr.Next.Prev = curr.Prev

	l.Len--
	return l, nil
}

// Reverse reverses the list
func (l *List) Reverse() *List {
	if l.Len == 0 {
		return l
	}

	// Move the head to the tail and traverse backwards
	l.Head = l.Tail
	curr := l.Head
	for curr.Prev != nil {
		curr.Next = curr.Prev
		curr.Prev = curr.Next
		curr = curr.Prev
	}
	// Adjust the last node
	curr.Prev = curr.Next
	curr.Next = nil
	l.Tail = curr

	return l
}

// Print prints nodes in the list
func (l *List) Print() {
	curr := l.Head
	fmt.Println("Head\t   ", curr.Value, "->")
	for curr.Next != nil {
		if curr != l.Head {
			fmt.Println("\t", "<-", curr.Value, "->")
		}
		curr = curr.Next
	}
	fmt.Println("Tail\t <-", l.Tail.Value)
	fmt.Println("\nLength:", l.Len)
}
