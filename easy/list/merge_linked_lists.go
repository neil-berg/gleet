package easy

// Node is an element in the linked list
type Node struct {
	Value int
	Next  *Node
}

// LinkedList is the container of elements
type LinkedList struct {
	Head   *Node
	Length int
}

// getValues returns the values of a linked list
func getValues(head *Node) []int {
	var values []int

	curr := head
	for curr != nil {
		values = append(values, curr.Value)
		curr = curr.Next
	}

	return values
}

// MergeLinkedLists merges two linked lists together (l2 into l1)
// You are given the heads of two sorted linked lists list1 and list2.
// Merge the two lists in a one sorted list. The list should be made by splicing
// together the nodes of the first two lists.
// Ex: Input: list1 = [2,3,5], list2 = [1,4]
// Output [1,2,3,4,5]
func MergeLinkedLists(l1, l2 *LinkedList) *Node {
	// Handle edge cases
	if l1.Head == nil {
		return l2.Head
	}

	if l2.Head == nil {
		return l1.Head
	}

	if l1.Head == nil && l2.Head == nil {
		return &Node{}
	}

	// New head of the merged list
	var head *Node

	// Initially the tail is the head, it'll advance each insertion
	tail := head

	// Current positions as we walk through l1 and l2
	c1 := l1.Head
	c2 := l2.Head

	for c1 != nil && c2 != nil {
		var min *Node

		// Pick min node and advance position in the list containing it
		if c1.Value < c2.Value {
			min = c1
			c1 = c1.Next
		} else {
			min = c2
			c2 = c2.Next
		}

		// Set head if empty, otherwise add to the tail
		if head == nil {
			head = min
		} else {
			tail.Next = min
		}

		// Update tail position
		tail = min
	}

	// Add any remaining list nodes if one list has been walked through
	if c1 != nil {
		tail.Next = c1
	} else {
		tail.Next = c2
	}

	return head
}
