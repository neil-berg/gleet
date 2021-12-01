package heap

import (
	"reflect"
	"testing"
)

func TestCreateMaxHeap(t *testing.T) {
	heap := &Heap{arr: []int{5, 10, 3, 8, 9}}
	heap.CreateMaxHeap()
	got := heap.arr
	want := []int{10, 9, 3, 8, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
