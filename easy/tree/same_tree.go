package easy

// TreeNode describes a node of a tree
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Tree is a data structure of tree nodes
type Tree struct {
	Root *TreeNode
}

// SameTree checks if binary trees a, b are the same (given their roots).
// Two binary trees are considered the same if they are structurally identical,
// and the nodes have the same value.
func SameTree(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Value == b.Value {
		return SameTree(a.Left, b.Left) && SameTree(a.Right, b.Right)
	}
	return false
}
