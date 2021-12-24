package medium

import (
	"reflect"
	"testing"
)

func newList(nums []int) *List {
	out := &List{}
	tail := out.Head
	for _, v := range nums {
		node := &ListNode{Value: v}
		if out.Head == nil {
			out.Head = node
		} else {
			tail.Next = node
		}
		tail = node
		out.Len++
	}
	return out
}

func getValues(l *List) []int {
	curr := l.Head
	var values []int
	for curr != nil {
		values = append(values, curr.Value)
		curr = curr.Next
	}
	return values
}

func TestAddTwoRevLists(t *testing.T) {
	tests := map[string]struct {
		l1nums []int
		l2nums []int
		want   []int
	}{
		"first": {
			l1nums: []int{2, 4, 3},
			l2nums: []int{5, 6, 4},
			want:   []int{7, 0, 8},
		},
		"second": {
			l1nums: []int{9, 9, 9, 9, 9, 9, 9},
			l2nums: []int{9, 9, 9, 9},
			want:   []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		"zero": {
			l1nums: []int{0},
			l2nums: []int{0},
			want:   []int{0},
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			l1 := newList(test.l1nums)
			l2 := newList(test.l2nums)
			got := getValues(AddTwoRevLists(l1, l2))
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}

}
