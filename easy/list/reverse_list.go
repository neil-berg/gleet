package easy

// ReverseList reverses a singly-linked list
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		curr := head
		head = curr.Next
		curr.Next = prev
		prev = curr
	}

	return prev
}
