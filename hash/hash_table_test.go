package hash

import (
	"testing"
)

func TestGetHashCodeAndIndex(t *testing.T) {
	ht := NewTable(3)
	code, index := ht.GetHashCodeAndIndex("Ann")
	wantCode := 285
	wantIndex := 0

	if code != wantCode && index != wantIndex {
		t.Errorf("got: (%v, %v), want: (%v, %v)", code, index, wantCode, wantIndex)
	}
}

func TestInsert(t *testing.T) {
	ht := NewTable(5)
	s := "Sam"
	code, index := ht.GetHashCodeAndIndex(s)
	ht.Insert(s)

	gotNode, err := ht.Get(s)
	gotValue := gotNode.Data.Value
	gotCode := gotNode.Data.Code

	if gotValue != s && gotCode != code && err != nil {
		t.Errorf("Stored wrong data, got (%v, %v), want (%v, %v)", gotCode, gotValue, s, index)
	}
}

func TestDelete(t *testing.T) {
	ht := NewTable(5)
	s := "Sam"
	ht.Insert(s)
	ht.Delete(s)

	gotNode, err := ht.Get(s)

	if gotNode != nil && err == nil {
		t.Errorf("Failed to delete node")
	}
}

func TestResize(t *testing.T) {
	numBuckets := 5
	ht := NewTable(5)
	ht.Insert("Sam")
	ht.Insert("Joe")
	ht.Insert("Pam")
	ht.Insert("Ben")
	ht.Insert("Ann") // Triggers resize (285)

	if ht.NumBuckets != numBuckets*2 {
		t.Errorf("Wrong num buckets, got %v, want %v", ht.NumBuckets, numBuckets*2)
	}

	if ht.NumEntries != 5 {
		t.Errorf("Wrong num buckets, got %v, want %v", ht.NumEntries, 5)
	}

	// ht.Display()
}
