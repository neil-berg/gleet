package tree

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
	root *Node
}

// Insert inserts a new element with Value v in the tree
func (t *BST) Insert(v int) {
	node := &Node{Value: v}

	if t.root == nil {
		t.root = node // initial node
	} else {
		insert(t.root, node)
	}
}

// insert recursively traverses the tree to insert a new node
func insert(currentNode, newNode *Node) {
	if currentNode.Value == newNode.Value {
		fmt.Println("Node already exists")
		return
	}

	if newNode.Value < currentNode.Value {
		if currentNode.Left == nil {
			currentNode.Left = newNode
		} else {
			insert(currentNode.Left, newNode)
		}
	} else {
		if currentNode.Right == nil {
			currentNode.Right = newNode
		} else {
			insert(currentNode.Right, newNode)
		}
	}
}

// MaxDepth is the number of nodes along the longest path from the root node to
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
	if t.root == nil {
		return 0, errors.New("Empty tree")
	}

	curr := t.root
	for curr.Right != nil {
		curr = curr.Right
	}
	return curr.Value, nil
}

// MinValue returns the min value in the tree
func (t *BST) MinValue() (int, error) {
	if t.root == nil {
		return 0, errors.New("empty tree")
	}

	curr := t.root
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr.Value, nil
}

// ValueExists checks whether a value exists in the tree
func (t *BST) ValueExists(v int) bool {
	if t.root == nil {
		return false
	}
	return ExistsTraversal(t.root, v)
}

// ExistsTraversal is an internal recursion to traversal search
func ExistsTraversal(n *Node, v int) bool {
	if n == nil {
		return false
	}

	if v < n.Value {
		return ExistsTraversal(n.Left, v)
	}
	if v > n.Value {
		return ExistsTraversal(n.Right, v)
	}
	return true
}

// Print displays the tree
func (t *BST) Print() {
	fmt.Println("root:", t.root.Value)
}
