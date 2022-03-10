package easy

// MiddleNode returns the middle node's value in a linked list
func MiddleNode(head *ListNode) int {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow.Value
}
