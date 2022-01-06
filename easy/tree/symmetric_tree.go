package easy

// SymmetricTree checks whether a tree is a mirror if itself (i,e, symmetric around its center).
func SymmetricTree(root *TreeNode) bool {
	// Root only edge case
	if root.Left == nil && root.Right == nil {
		return true
	}

	var recurMirror func(n1, n2 *TreeNode) bool

	recurMirror = func(n1, n2 *TreeNode) bool {
		if n1 == nil && n2 == nil {
			return true
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Value != n2.Value {
			return false
		}

		// Check subtrees
		return recurMirror(n1.Left, n2.Right) && recurMirror(n1.Right, n2.Left)
	}

	return recurMirror(root.Left, root.Right)
}

// SymmetricTreeBFS checks for symmetry using BFS level checks
func SymmetricTreeBFS(root *TreeNode) bool {
	queue := []*TreeNode{root.Left, root.Right}

	for len(queue) > 0 {
		// Pick first 2 items and then remove them from the queue
		n1 := queue[0]
		n2 := queue[1]
		queue = queue[2:]

		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Value != n2.Value {
			return false
		}

		// Enqueue children in an order to match the mirror
		queue = append(queue, n1.Left, n2.Right, n1.Right, n2.Left)
	}

	return true
}
