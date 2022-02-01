package medium

import (
	"reflect"
	"testing"
)

// getListValues in a helper to walk the "list" of TreeNodes and return its values
func getListValues(head *TreeNode) []int {
	var values []int
	curr := head
	for curr != nil {
		values = append(values, curr.Value)
		curr = curr.Right
	}

	return values
}

func TestFlattenTreeToList(t *testing.T) {
	root := &TreeNode{Value: 1}
	root.Left = &TreeNode{Value: 2}
	root.Left.Left = &TreeNode{Value: 3}
	root.Left.Right = &TreeNode{Value: 4}
	root.Right = &TreeNode{Value: 5}
	root.Right.Right = &TreeNode{Value: 6}

	head := FlattenTreeToList(root)
	got := getListValues(head)
	want := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
