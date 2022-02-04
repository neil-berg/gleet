package easy

import "testing"

func TestListPalindrome(t *testing.T) {
	head := &ListNode{Value: 1}
	head.Next = &ListNode{Value: 2}
	head.Next.Next = &ListNode{Value: 3}
	head.Next.Next.Next = &ListNode{Value: 2}
	head.Next.Next.Next.Next = &ListNode{Value: 1}

	head2 := &ListNode{Value: 1}
	head2.Next = &ListNode{Value: 2}
	head2.Next.Next = &ListNode{Value: 3}
	head2.Next.Next.Next = &ListNode{Value: 3}
	head2.Next.Next.Next.Next = &ListNode{Value: 3}

	tests := []struct {
		node *ListNode
		want bool
	}{
		{node: head, want: true},
		{node: head2, want: false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := ListPalindrome(tt.node)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
