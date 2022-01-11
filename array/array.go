package array

import (
	"math"
)

// Array is a list with push, pop, insertAt, and reverse methods
type Array struct {
	items []interface{}
	len   int
}

// Push adds an item to the right of the array
func (a *Array) Push(item interface{}) *Array {
	a.items = append(a.items, item)
	a.len++
	return a
}

// Pop removes the last item in the array
func (a *Array) Pop() *Array {
	a.items = a.items[:len(a.items)-1]
	a.len--
	return a
}

// InsertAt adds an item at a certain index
func (a *Array) InsertAt(item interface{}, i int) *Array {
	items1 := append(a.items[0:i], item)
	items2 := a.items[i+1:]
	a.items = append(items1, items2...)
	a.len++
	return a
}

// Reverse reverses the items
func (a *Array) Reverse() *Array {
	f := int(math.Floor(float64(a.len / 2)))

	for i := 0; i < f; i++ {
		first := a.items[i]
		second := a.items[a.len-1-i]
		a.items[i], a.items[a.len-1-i] = second, first
	}
	return a
}

// BinarySearch tries to find value v (an int) in a sorted array
func (a *Array) BinarySearch(v int) bool {
	min := 0
	max := a.len - 1

	for min <= max {
		mid := (min + max) / 2
		if a.items[mid].(int) < v {
			min = min + 1
		} else if a.items[mid].(int) > v {
			max = max - 1
		} else {
			return true
		}
	}
	return false
}
