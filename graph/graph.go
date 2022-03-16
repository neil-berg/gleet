package graph

// Graph is a the structure of graph
type Graph struct {
	// Adjacency is the graph's adjacency map
	Adjacency map[int][]int
	// Visited is the graph's visited map
	Visited map[int]bool
	// Directed is whether the graph is directed or not
	Directed bool
}

// TraverseDepthFirst walks a graph depth-first starting at the src node
func (g *Graph) TraverseDepthFirst(src int) []int {
	stack := []int{src}

	values := []int{}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		g.Visited[current] = true
		values = append(values, current)
		stack = stack[:len(stack)-1]

		for _, neighbor := range g.Adjacency[current] {
			if !g.Visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}

	return values
}

// TraverseDepthFirstRecur walks a graph depth-first recursively starting at the src node
func (g *Graph) TraverseDepthFirstRecur(src int) []int {
	var values []int

	var recur func(node int)
	recur = func(node int) {
		g.Visited[node] = true
		values = append(values, node)
		neighbors := g.Adjacency[node]
		for _, neighbor := range neighbors {
			if !g.Visited[neighbor] {
				recur(neighbor)
			}
		}
	}

	recur(src)
	return values
}

// TraverseBreadth traverses a graph breadth-first
func (g *Graph) TraverseBreadth(src int) []int {
	var queue []int
	var values []int

	g.Visited[src] = true
	queue = append(queue, src)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		values = append(values, current)
		for _, node := range g.Adjacency[current] {
			if !g.Visited[node] {
				queue = append(queue, node)
				g.Visited[node] = true
			}
		}
	}

	return values
}

// HasPathBFS determines if a path exists between the src and dest nodes
func (g *Graph) HasPathBFS(src, dest int) bool {
	if src == dest {
		return true
	}

	var queue []int

	g.Visited[src] = true
	queue = append(queue, src)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, neighbor := range g.Adjacency[curr] {
			if neighbor == dest {
				return true
			}
			if !g.Visited[neighbor] {
				queue = append(queue, neighbor)
				g.Visited[neighbor] = true
			}
		}
	}

	return false
}

// HasPathDFS determines if a path exists via depth-first search
func (g *Graph) HasPathDFS(src, dst int) bool {
	if src == dst {
		return true
	}

	if g.Visited[src] {
		return false
	}

	g.Visited[src] = true

	for _, neighbor := range g.Adjacency[src] {
		if g.HasPathDFS(neighbor, dst) {
			return true
		}
	}

	return false
}
