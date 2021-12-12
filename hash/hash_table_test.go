package hash

import "testing"

func TestInsert(t *testing.T) {
	ht := NewTable(5)
	ht.Insert("Sam")
	ht.Insert("Joe")
	ht.Insert("Pam")
	ht.Insert("Ben")
	ht.Insert("Ann")
	ht.Insert("Pete")
	ht.Insert("Jan")
	ht.Insert("Bill")
	ht.Display()
}
