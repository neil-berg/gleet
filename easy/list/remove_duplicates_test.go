package easy

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	head := &ListNode{Value: 1}
	head.Next = &ListNode{Value: 1}
	head.Next.Next = &ListNode{Value: 2}
	head.Next.Next.Next = &ListNode{Value: 3}
	head.Next.Next.Next.Next = &ListNode{Value: 3}
	head.Next.Next.Next.Next.Next = &ListNode{Value: 4}

	newHead := RemoveDuplicates(head)
	var values []int
	curr := newHead
	for curr != nil {
		values = append(values, curr.Value)
		curr = curr.Next
	}

	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(values, want) {
		t.Errorf("got %v, want %v", values, want)
	}
}
