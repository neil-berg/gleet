package graph

import (
	"fmt"
	"testing"
)

// func TestDepthFirstTraversal(t *testing.T) {
// 	graph := &Graph{
// 		AdjList: AdjacencyList{
// 			1: []int{2, 3},
// 			2: []int{},
// 			3: []int{4},
// 			4: []int{5},
// 			5: []int{6, 7},
// 			6: []int{},
// 			7: []int{},
// 		},
// 	}

// 	// got := graph.DepthFirstTraversalRecur(1)
// 	fmt.Println(graph.BFS(1))
// 	// fmt.Println(got2)
// }
func TestHasPathBFS(t *testing.T) {
	graph := &Graph{
		AdjList: AdjacencyList{
			1: []int{2, 4},
			2: []int{3},
			3: []int{},
			4: []int{2, 5, 6},
			5: []int{},
			6: []int{},
		},
	}

	// got := graph.DepthFirstTraversalRecur(1)
	fmt.Println(graph.HasPathBFS(2, 5))
	// fmt.Println(got2)
}
