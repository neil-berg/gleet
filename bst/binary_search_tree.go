package bst

import (
	"errors"
	"fmt"
)

// Node is an element on the tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BST is the binary search tree
type BST struct {
	Root *Node
}

// Insert inserts a new element with Value v in the tree
func (t *BST) Insert(v int) {
	node := &Node{Value: v}

	var recurInsert func(currentNode, newNode *Node)

	recurInsert = func(currentNode, newNode *Node) {
		if currentNode.Value == newNode.Value {
			fmt.Println(newNode.Value, "already exists")
			return
		}

		if newNode.Value < currentNode.Value {
			if currentNode.Left == nil {
				currentNode.Left = newNode
			} else {
				recurInsert(currentNode.Left, newNode)
			}
		} else {
			if currentNode.Right == nil {
				currentNode.Right = newNode
			} else {
				recurInsert(currentNode.Right, newNode)
			}
		}

	}

	if t.Root == nil {
		t.Root = node // initial node
	} else {
		recurInsert(t.Root, node)
	}
}

// MaxDepth is the number of nodes along the longest path from the Root node to
// the farthest leaf node.
//             5
//           /   \
//          3     7
//         / \   / \
//        2   4 6   10
// MaxDepth = 3
func MaxDepth(n *Node) int {
	if n == nil {
		return 0
	}

	leftMaxDepth := MaxDepth(n.Left)
	rightMaxDepth := MaxDepth(n.Right)

	if leftMaxDepth > rightMaxDepth {
		return leftMaxDepth + 1
	}
	return rightMaxDepth + 1

}

// MaxValue returns the max value in the tree
func (t *BST) MaxValue() (int, error) {
	if t.Root == nil {
		return 0, errors.New("Empty tree")
	}

	curr := t.Root
	for curr.Right != nil {
		curr = curr.Right
	}
	return curr.Value, nil
}

// MinValue returns the min value in the tree
func (t *BST) MinValue() (int, error) {
	if t.Root == nil {
		return 0, errors.New("empty tree")
	}

	curr := t.Root
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr.Value, nil
}

// ValueExists checks whether value v exists in the tree
func (t *BST) ValueExists(v int) bool {
	if t.Root == nil {
		return false
	}

	var recurExists func(n *Node, v int) bool

	recurExists = func(n *Node, v int) bool {
		if n == nil {
			return false
		}

		if v < n.Value {
			return recurExists(n.Left, v)
		}
		if v > n.Value {
			return recurExists(n.Right, v)
		}
		// Neither greater or less than means this node's value is the target value
		return true

	}

	return recurExists(t.Root, v)
}

// InOrderTraversal traverses the tree in-order (left, node, right)
func (t *BST) InOrderTraversal() []int {
	var values []int
	var recurInOrder func(n *Node)

	recurInOrder = func(n *Node) {
		if n != nil {
			recurInOrder(n.Left)
			values = append(values, n.Value)
			recurInOrder(n.Right)
		}
	}

	recurInOrder(t.Root)
	return values
}

// PreOrderTraversal traverses a tree pre-order (node, left, right)
func (t *BST) PreOrderTraversal() []int {
	var values []int
	var recurPreOrder func(n *Node)

	recurPreOrder = func(n *Node) {
		if n != nil {
			values = append(values, n.Value)
			recurPreOrder(n.Left)
			recurPreOrder(n.Right)
		}
	}

	recurPreOrder(t.Root)
	return values
}

// PostOrderTraversal traverses the tree in post-order (left, right, node)
func (t *BST) PostOrderTraversal() []int {
	var values []int
	var recurPostOrder func(n *Node)

	recurPostOrder = func(n *Node) {
		if n != nil {
			recurPostOrder(n.Left)
			recurPostOrder(n.Right)
			values = append(values, n.Value)
		}
	}

	recurPostOrder(t.Root)
	return values
}

// BreadthFirstTraversal traverses the tree level-wise
func (t *BST) BreadthFirstTraversal() []int {
	// Temp queue as we traverse various nodes
	var queue []*Node
	// Values are the resulting nodes in the BFS sequence
	var values []int

	if t.Root == nil {
		return []int{}
	}

	queue = append(queue, t.Root)
	for len(queue) > 0 {
		// Dequeue and store initial node's value, then enqueue its children
		curr := queue[0]
		queue = queue[1:]
		values = append(values, curr.Value)
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}

	return values
}

// Print displays the tree
func (t *BST) Print() {
	fmt.Println("Root:", t.Root.Value)
}
