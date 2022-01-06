package medium

import (
	"math"
)

// ListNode is a node in the list
type ListNode struct {
	Value int
	Next  *ListNode
}

// List is the linked list
type List struct {
	Head *ListNode
	Len  int
}

// AddTwoRevLists adds two linked lists in reverse order and returns the new
// linked list.
// E.g.
// L1: 2 -> 4 -> 3 (342)
// L2: 5 -> 6 -> 4 (465)
// Out: 7 -> 0 -> 8 (342 + 465 = 807)
func AddTwoRevLists(l1, l2 *List) *List {
	// Integer formed by the reverse of nodes in l1 and l2, respectively
	n1 := 0
	n2 := 0

	// Current node while walking through l1 and l2
	curr1 := l1.Head
	curr2 := l2.Head

	// Step count while walking through l1
	i1 := 0
	for curr1 != nil {
		n1 = n1 + curr1.Value*int(math.Pow(10, float64(i1)))
		curr1 = curr1.Next
		i1++
	}

	// step count while walking through l2
	i2 := 0
	for curr2 != nil {
		n2 = n2 + curr2.Value*int(math.Pow(10, float64(i2)))
		curr2 = curr2.Next
		i2++
	}

	sum := n1 + n2

	out := &List{}
	tailOut := out.Head

	// Edge case of sum equally zero to begin with
	if sum == 0 {
		return &List{Head: &ListNode{Value: 0}, Len: 1}
	}

	// Iterate over the sum and populate the list each pass
	for sum != 0 {
		val := sum % 10
		node := &ListNode{Value: val}
		if out.Head == nil {
			out.Head = node
		} else {
			tailOut.Next = node
		}
		tailOut = node
		sum = sum / 10
		out.Len++
	}

	return out
}
