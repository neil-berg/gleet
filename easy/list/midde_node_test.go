package easy

import "testing"

func TestMiddleNode(t *testing.T) {
	head1 := &ListNode{Value: 1}
	head1.Next = &ListNode{Value: 2}
	head1.Next.Next = &ListNode{Value: 3}
	head1.Next.Next.Next = &ListNode{Value: 4}
	head1.Next.Next.Next.Next = &ListNode{Value: 5}

	head2 := &ListNode{Value: 1}
	head2.Next = &ListNode{Value: 2}
	head2.Next.Next = &ListNode{Value: 3}
	head2.Next.Next.Next = &ListNode{Value: 4}
	head2.Next.Next.Next.Next = &ListNode{Value: 5}
	head2.Next.Next.Next.Next.Next = &ListNode{Value: 6}

	tests := []struct {
		head *ListNode
		want int
	}{
		{head: head1, want: 3},
		{head: head2, want: 4},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := MiddleNode(tt.head)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}
