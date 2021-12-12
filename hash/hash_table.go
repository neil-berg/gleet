package hash

import (
	"errors"
	"fmt"
)

// NodeData is the data of a linked list node
type NodeData struct {
	Code  int
	Value string
}

// Node is element in the linked list of values for a key in the table
type Node struct {
	Data *NodeData
	Next *Node
}

// LinkedList is the linked list of values for a key in the table
type LinkedList struct {
	Head   *Node
	Length int
}

// Table is the container of a hash table
type Table struct {
	Map        map[int]LinkedList
	NumEntries int
	NumBuckets int
}

// Add adds an items to the linked list
func (ll LinkedList) Add(node *Node) LinkedList {
	if ll.Head == nil {
		ll.Head = node
	} else {
		curr := ll.Head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = node
	}
	ll.Length++
	return ll
}

// Remove removes a node from the linked list if it exists
func (ll LinkedList) Remove(node *Node) LinkedList {
	if ll.Head == nil {
		return ll
	}

	if ll.Head.Data.Value == node.Data.Value {
		ll.Head = nil
		return ll
	}

	curr := ll.Head
	for curr.Next != nil {
		if curr.Next.Data.Value == node.Data.Value {
			curr.Next = curr.Next.Next
			return ll
		}
		curr = curr.Next
	}

	return ll
}

// Find walks the list and returns the node if found, else an error
func (ll LinkedList) Find(node *Node) (*Node, error) {
	if ll.Head == nil {
		return nil, errors.New("Not found")
	}

	curr := ll.Head
	for curr != nil {
		if curr.Data.Value == node.Data.Value && curr.Data.Code == node.Data.Code {
			return node, nil
		}
		curr = curr.Next
	}
	return nil, errors.New("Not found")
}

// GetHashCodeAndIndex takes a string and returns its hash code and index in the table
func (ht *Table) GetHashCodeAndIndex(s string) (int, int) {
	code := 0
	for _, c := range s {
		code += int(c)
	}

	index := code % ht.NumBuckets
	return code, index
}

// NewTable creates a new table with specified number of buckets
func NewTable(numBuckets int) *Table {
	return &Table{
		Map:        make(map[int]LinkedList),
		NumEntries: 0,
		NumBuckets: numBuckets,
	}
}

// Insert insert a new item to the hash table
func (ht *Table) Insert(s string) {
	code, index := ht.GetHashCodeAndIndex(s)
	node := &Node{
		Data: &NodeData{Code: code, Value: s},
	}

	ht.Map[index] = ht.Map[index].Add(node)
	ht.NumEntries++

	loadFactor := float32(ht.NumEntries / ht.NumBuckets)
	if loadFactor > 0.7 {
		fmt.Println("resizing....")
		resized := Resize(ht)
		ht.Map = resized.Map
		ht.NumEntries = resized.NumEntries
		ht.NumBuckets = resized.NumBuckets
	}
}

// Delete deletes the item from the hash table
func (ht *Table) Delete(s string) {
	code, index := ht.GetHashCodeAndIndex(s)
	node := &Node{
		Data: &NodeData{Code: code, Value: s},
	}

	ht.Map[index] = ht.Map[index].Remove(node)
	ht.NumEntries--
}

// Get searches for a node in the table
func (ht *Table) Get(s string) (*Node, error) {
	code, index := ht.GetHashCodeAndIndex(s)
	node := &Node{
		Data: &NodeData{Code: code, Value: s},
	}
	return ht.Map[index].Find(node)
}

// Resize creates a new table with twice the number of buckets and shifts
// existing entries into the new table
func Resize(ht *Table) *Table {
	resized := NewTable(ht.NumBuckets * 2)

	for i := 0; i < ht.NumBuckets; i++ {
		ll := ht.Map[i]
		curr := ll.Head
		for curr != nil {
			ht.Delete(curr.Data.Value)
			resized.Insert(curr.Data.Value)
			curr = curr.Next
		}
	}

	return resized
}

// Display displays the table
func (ht *Table) Display() {
	fmt.Println("Buckets:", ht.NumBuckets, "Entries:", ht.NumEntries)
	for i := 0; i < ht.NumBuckets; i++ {
		data := []NodeData{}

		list := ht.Map[i]
		curr := list.Head
		for curr != nil {
			data = append(data, *curr.Data)
			curr = curr.Next
		}

		fmt.Printf("%d: %v \n", i, data)
	}
}
