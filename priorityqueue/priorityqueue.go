package priorityqueue

// PQNode is a node in the priority queue
type PQNode struct {
	Priority int // larger value = higher priority
	Name     string
}

// PriorityQueue implements the priority queue data structure
type PriorityQueue struct {
	Items []PQNode
}

// NewPriorityQueue takes a slice of pointers to priority queue nodes and returns
// a priority queue (aka max-heap) from them.
func (pq *PriorityQueue) NewPriorityQueue() {
	for i := len(pq.Items) / 2; i >= 0; i-- {
		pq.MaxHeapify(i)
	}
}

// MaxHeapify creates a max heap of items in the priority queue
func (pq *PriorityQueue) MaxHeapify(i int) {
	maxIdx := i

	leftIdx := 2*i + 1
	rightIdx := 2*i + 2

	if leftIdx < len(pq.Items) && pq.Items[leftIdx].Priority > pq.Items[maxIdx].Priority {
		maxIdx = leftIdx
	}

	if rightIdx < len(pq.Items) && pq.Items[rightIdx].Priority > pq.Items[maxIdx].Priority {
		maxIdx = rightIdx
	}

	if maxIdx != i {
		// Swap and then recursively max heapify again for potentially affected children
		pq.Items[maxIdx], pq.Items[i] = pq.Items[i], pq.Items[maxIdx]
		pq.MaxHeapify(maxIdx)
	}
}

// Insert adds a node to the priority queue
func (pq *PriorityQueue) Insert(node PQNode) {
	// Heapify from the bottom-up
	pq.Items = append(pq.Items, node)
	for i := len(pq.Items) / 2; i >= 0; i-- {
		pq.MaxHeapify(i)
	}
}

// Delete removes the top node in the priority queue
func (pq *PriorityQueue) Delete() {
	if len(pq.Items) == 1 {
		pq.Items = []PQNode{}
		return
	}
	// Pop and then down-max-heapify
	pq.Items = pq.Items[1:]
	for i := 0; i <= len(pq.Items)/2; i++ {
		pq.MaxHeapify(i)
	}
}
