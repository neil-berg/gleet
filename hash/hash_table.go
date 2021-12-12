package hash

import "fmt"

// Node is element in the linked list of values for a key in the table
type Node struct {
	Data []interface{}
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
	node := &Node{Data: []interface{}{code, s}}

	ht.Map[index] = ht.Map[index].Add(node)
	ht.NumEntries++

	loadFactor := float32(ht.NumEntries / ht.NumBuckets)
	if loadFactor > 0.7 {
		ht.Resize()
	}
}

// Delete deletes the item from the hash table
// func (ht *Table) Delete(s string) {
// 	code := GetHashCode(s)
// 	index := code % ht.NumBuckets
// 	node := &Node{Data: []interface{}{code, s}}

// 	ht.Map[index] = ht.Map[index].Remove(node)
// 	ht.NumEntries--
// }

// Resize resizes the table
func (ht *Table) Resize() {
	// TOOD - resize logic
	// 1. Create new table with 2* numBuckets
	// 2. For each item in existing table, remove, and then re-insert into new table
	// fmt.Println("resizing table!")
}

// Display displays the table
func (ht *Table) Display() {
	for i := 0; i < ht.NumBuckets; i++ {
		values := []interface{}{}

		list := ht.Map[i]
		curr := list.Head
		for curr != nil {
			values = append(values, curr.Data)
			curr = curr.Next
		}

		fmt.Printf("%d: %v \n", i, values)
	}
}
