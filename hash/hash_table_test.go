package hash

import (
	"fmt"
	"testing"
)

// setupTable is a local helper to setup a hash table that looks like
// 0: Sam ->
//
//
//
// func seupTable() *Table {

// }

func TestInsert(t *testing.T) {
	ht := NewTable(5)
	ht.Insert("Sam")
	// ht.Insert("Sam")
	// ht.Insert("Joe")
	// ht.Insert("Pam")
	// ht.Insert("Ben")
	// ht.Insert("Ann")
	// ht.Insert("Pete")
	// ht.Insert("Jan")
	// ht.Insert("Bill")
	// ht.Delete("Pam")
	// ht.Delete("Jane")
	// ht.Delete("Ann")
	ht.Display()
}

func TestGet(t *testing.T) {
	ht := NewTable(5)
	ht.Insert("Sam")
	ht.Insert("Joe")
	node, _ := ht.Get("Sam")
	fmt.Println(node.Data.Value)
	// ht.Display()
}

func TestResize(t *testing.T) {
	ht := NewTable(5)
	ht.Insert("Sam")
	ht.Insert("Joe")
	ht.Insert("Pam")
	ht.Insert("Ben")
	ht.Insert("Ann")
	ht.Display()
}
