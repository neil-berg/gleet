package easy

import (
	"reflect"
	"testing"
)

// newLinkedList takes an array and turns it into a linked list
// E.g. [1,2,3] => 1 -> 2 -> 3
func newLinkedList(a []int) *LinkedList {
	list := &LinkedList{}

	curr := list.Head
	for _, v := range a {
		node := &Node{Value: v, Next: nil}
		if list.Length == 0 {
			list.Head = node
		} else {
			curr.Next = node
		}
		curr = node

		list.Length++
	}

	return list
}

func TestMergedLinkedLists(t *testing.T) {

	tests := map[string]struct {
		l1   *LinkedList
		l2   *LinkedList
		want []int
	}{
		"non-empty": {
			l1:   newLinkedList([]int{1, 2, 5}),
			l2:   newLinkedList([]int{0, 2, 6, 7, 8, 9}),
			want: []int{0, 1, 2, 2, 5, 6, 7, 8, 9},
		},
		"l1 empty": {
			l1:   newLinkedList([]int{}),
			l2:   newLinkedList([]int{0, 2, 6, 7, 8, 9}),
			want: []int{0, 2, 6, 7, 8, 9},
		},
		"l2 empty": {
			l1:   newLinkedList([]int{1, 2, 5}),
			l2:   newLinkedList([]int{}),
			want: []int{1, 2, 5},
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			got := getValues(MergeLinkedLists(test.l1, test.l2))
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
