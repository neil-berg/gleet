package easy

// ListNode is a node in a linked list
type ListNode struct {
	Value int
	Next  *ListNode
}

// RemoveDuplicates removes duplicate nodes in a sorted linked list.
func RemoveDuplicates(head *ListNode) *ListNode {
	prev := head
	curr := head.Next
	for curr != nil {
		if curr.Value == prev.Value {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return head
}
