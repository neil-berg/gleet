package easy

// InvertTree inverst a binary tree.
func InvertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		InvertTree(root.Left)
		InvertTree(root.Right)
	}

	return root
}
