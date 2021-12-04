package heap

import (
	"reflect"
	"testing"
)

func TestMaxUpHeapify(t *testing.T) {
	heap := &Heap{arr: []int{5, 10, 3, 8, 9}}
	heap.MaxUpHeapify()
	got := heap.arr
	want := []int{10, 9, 3, 8, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMaxDownHeapify(t *testing.T) {
	heap := &Heap{arr: []int{5, 10, 3, 8, 9}}
	heap.MaxDownHeapify()
	got := heap.arr
	want := []int{10, 9, 3, 8, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMinUpHeapify(t *testing.T) {
	heap := &Heap{arr: []int{5, 10, 3, 8, 9}}
	heap.MinUpHeapify()
	got := heap.arr
	want := []int{3, 8, 5, 10, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMinDownHeapify(t *testing.T) {
	heap := &Heap{arr: []int{5, 10, 3, 8, 9}}
	heap.MinDownHeapify()
	got := heap.arr
	want := []int{3, 8, 5, 10, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMaxInsert(t *testing.T) {
	heap := &Heap{arr: []int{5, 10, 3, 8, 9}}
	heap.MaxUpHeapify()
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
	heap.MinUpHeapify()
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
	heap.MaxUpHeapify()
	heap.MaxDelete()
	got := heap.arr
	want := []int{7, 6, 4, 2, 1, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMinDelete(t *testing.T) {
	heap := &Heap{arr: []int{3, 5, 7, 10, 9, 8, 12}}
	heap.MinUpHeapify()
	heap.MinDelete()
	got := heap.arr
	want := []int{5, 9, 7, 10, 12, 8}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMaxAscSort(t *testing.T) {
	heap := &Heap{arr: []int{20, 15, 12, 14, 3, 8, 10}}
	got := heap.MaxSortAsc()
	want := []int{3, 8, 10, 12, 14, 15, 20}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMaxDescSort(t *testing.T) {
	heap := &Heap{arr: []int{20, 15, 12, 14, 3, 8, 10}}
	got := heap.MaxSortDesc()
	want := []int{20, 15, 14, 12, 10, 8, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMinAscSort(t *testing.T) {
	heap := &Heap{arr: []int{12, 6, 8, 3, 9, 11}}
	got := heap.MinAscSort()
	want := []int{3, 6, 8, 9, 11, 12}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMinDescSort(t *testing.T) {
	heap := &Heap{arr: []int{12, 6, 8, 3, 9, 11}}
	got := heap.MinDescSort()
	want := []int{12, 11, 9, 8, 6, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
