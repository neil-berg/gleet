package graph

// AdjacencyList is an adjacency list
type AdjacencyList map[int][]int

// Graph is a graph
type Graph struct {
	AdjList AdjacencyList
}

// DFS is a depth-first traversal
func (g *Graph) DFS(initialNode int) []int {
	stack := []int{initialNode}

	result := []int{}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		result = append(result, current)
		stack = stack[:len(stack)-1]

		for _, neighbor := range g.AdjList[current] {
			stack = append(stack, neighbor)
		}
	}

	return result
}

// DFSRecur uses recursion to traverse depth-wise
func (g *Graph) DFSRecur(initialNode int) []int {
	var values []int

	var recur func(node int)
	recur = func(node int) {
		values = append(values, node)
		neighbors := g.AdjList[node]
		for _, neighbor := range neighbors {
			recur(neighbor)
		}
	}

	recur(initialNode)
	return values
}

// BFS traverses a graph breadth-first
func (g *Graph) BFS(initialNode int) []int {
	visited := make(map[int]bool)
	var queue []int
	var values []int

	visited[initialNode] = true
	queue = append(queue, initialNode)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		values = append(values, current)
		for _, node := range g.AdjList[current] {
			if !visited[node] {
				queue = append(queue, node)
				visited[node] = true
			}
		}
	}

	return values
}

// HasPathBFS determines if a path exists between the src and dest nodes
func (g *Graph) HasPathBFS(src, dest int) bool {
	visited := make(map[int]bool)
	var queue []int

	visited[src] = true
	queue = append(queue, src)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, neighbor := range g.AdjList[curr] {
			if neighbor == dest {
				return true
			}
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}

	return false
}
