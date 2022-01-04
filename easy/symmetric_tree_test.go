package easy

import "testing"

func TestSymmetricTree(t *testing.T) {
	tree := &Tree{Root: &TreeNode{Value: 1}}
	tree.Root.Left = &TreeNode{Value: 2}
	tree.Root.Left.Left = &TreeNode{Value: 3}
	// tree.Root.Left.Right = &TreeNode{Value: 4}
	tree.Root.Right = &TreeNode{Value: 2}
	tree.Root.Right.Left = &TreeNode{Value: 3}
	// tree.Root.Right.Right = &TreeNode{Value: 3}

	got := SymmetricTree(tree.Root)
	if got != false {
		t.Errorf("incorrect")
	}
}
func TestSymmetricTreeBFS(t *testing.T) {
	tree := &Tree{Root: &TreeNode{Value: 1}}
	tree.Root.Left = &TreeNode{Value: 2}
	tree.Root.Left.Left = &TreeNode{Value: 3}
	tree.Root.Left.Right = &TreeNode{Value: 4}
	tree.Root.Right = &TreeNode{Value: 2}
	tree.Root.Right.Left = &TreeNode{Value: 4}
	tree.Root.Right.Right = &TreeNode{Value: 3}

	got := SymmetricTreeBFS(tree.Root)
	if got != true {
		t.Errorf("incorrect")
	}
}
