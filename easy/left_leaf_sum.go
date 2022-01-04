package easy

// LeftLeafSum calculates the sum of all left leaves in a binary tree.
func LeftLeafSum(root *TreeNode) int {
	sum := 0

	var recurInOrder func(n *TreeNode)

	recurInOrder = func(n *TreeNode) {
		if n != nil {
			recurInOrder(n.Left)
			if n.Left != nil && n.Left.Left == nil {
				sum += n.Left.Value
			}
			recurInOrder(n.Right)
		}
	}

	recurInOrder(root)

	return sum
}
