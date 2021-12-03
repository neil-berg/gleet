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

func TestCreateMinHeap(t *testing.T) {
	heap := &Heap{arr: []int{5, 10, 3, 8, 9}}
	heap.CreateMinHeap()
	got := heap.arr
	want := []int{3, 8, 5, 10, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMaxInsert(t *testing.T) {
	heap := &Heap{arr: []int{5, 10, 3, 8, 9}}
	heap.CreateMaxHeap()
	heap.MaxInsert(11)
	heap.MaxInsert(12)
	got := heap.arr
	want := []int{12, 9, 11, 8, 5, 3, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMinInsert(t *testing.T) {
	heap := &Heap{arr: []int{5, 10, 3, 8, 9}}
	heap.CreateMinHeap()
	heap.MinInsert(4)
	heap.MinInsert(1)
	got := heap.arr
	want := []int{1, 8, 3, 10, 9, 5, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMaxDelete(t *testing.T) {
	heap := &Heap{arr: []int{10, 7, 4, 6, 1, 3, 2}}
	heap.CreateMaxHeap()
	heap.MaxDelete()
	got := heap.arr
	want := []int{7, 6, 4, 2, 1, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMinDelete(t *testing.T) {
	heap := &Heap{arr: []int{3, 5, 7, 10, 9, 8, 12}}
	heap.CreateMinHeap()
	heap.MinDelete()
	got := heap.arr
	want := []int{5, 9, 7, 10, 12, 8}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
