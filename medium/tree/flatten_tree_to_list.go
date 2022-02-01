package medium

// TreeNode describes a node on a binary tree
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// FlattenTreeToList converts a binary tree via pre-order traversal into a
// linked list using the same TreeNode shape.
func FlattenTreeToList(root *TreeNode) *TreeNode {
	var head *TreeNode
	var tail *TreeNode

	var recurPreOrder func(node *TreeNode)
	recurPreOrder = func(node *TreeNode) {
		if node != nil {
			listNode := &TreeNode{Value: node.Value}
			if head == nil {
				head = listNode
			} else {
				tail.Right = listNode
			}
			tail = listNode
			recurPreOrder(node.Left)
			recurPreOrder(node.Right)
		}
	}

	recurPreOrder(root)

	return head
}

// Uses reverse pre-order...fancy

// func FlattenTreeToList(root *TreeNode) *TreeNode {
// 	var head *TreeNode

// 	var recurRevPreOrder func(node *TreeNode)
// 	recurRevPreOrder = func(node *TreeNode) {
// 		if node.Right != nil {
// 			recurRevPreOrder(node.Right)
// 		}
// 		if node.Left != nil {
// 			recurRevPreOrder(node.Left)
// 		}
// 		node.Right = head
// 		head = node
// 	}

// 	recurRevPreOrder(root)

// 	return head
// }
