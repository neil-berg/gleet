package easy

// ListPalindrome checks if the values in a linked list for a palindrome
func ListPalindrome(head *ListNode) bool {
	var values []int

	curr := head
	for curr != nil {
		values = append(values, curr.Value)
		curr = curr.Next
	}

	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		if values[i] != values[j] {
			return false
		}
	}
	return true
}
