package easy

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	head := &ListNode{Value: 1}
	head.Next = &ListNode{Value: 2}
	head.Next.Next = &ListNode{Value: 3}

	revHead := ReverseList(head)

	var values []int
	curr := revHead
	for curr != nil {
		values = append(values, curr.Value)
		curr = curr.Next
	}

	want := []int{3, 2, 1}
	if !reflect.DeepEqual(values, want) {
		t.Errorf("got %v, want %v", values, want)
	}
}
