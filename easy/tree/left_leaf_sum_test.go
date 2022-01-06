package easy

import (
	"testing"
)

func TestLeftLeafSum(t *testing.T) {
	tree := &Tree{Root: &TreeNode{Value: 2}}
	tree.Root.Left = &TreeNode{Value: 4}
	tree.Root.Right = &TreeNode{Value: 3}
	tree.Root.Right.Left = &TreeNode{Value: 7}
	sum := LeftLeafSum(tree.Root)
	if sum != 11 {
		t.Errorf("wrong sum")
	}
}
