package heap

// Heap is the shape of a heap
type Heap struct {
	arr []int
}

// LeftChildIdx returns the index of the left child for a node at index i
func (h *Heap) LeftChildIdx(i int) int {
	return 2*i + 1
}

// RightChildIdx returns the index of the right child for a node at index i
func (h *Heap) RightChildIdx(i int) int {
	return 2*i + 2
}

// Swap swaps the first and second elements in the heap array
func (h *Heap) Swap(first, second int) {
	h.arr[first], h.arr[second] = h.arr[second], h.arr[first]
}

// MaxHeapify converts a 0-based array into a max-heap
func (h *Heap) MaxHeapify(i int) {
	// For node at index i:
	leftChildIdx := h.LeftChildIdx(i)
	rightChildIdx := h.RightChildIdx(i)

	// Track the index of the largest value between current parent node (i-th
	// index) and its children
	largestIdx := i

	if leftChildIdx < len(h.arr) && h.arr[leftChildIdx] > h.arr[largestIdx] {
		largestIdx = leftChildIdx
	}
	if rightChildIdx < len(h.arr) && h.arr[rightChildIdx] > h.arr[largestIdx] {
		largestIdx = rightChildIdx
	}

	if largestIdx != i {
		// Swap the parent with its larger child node
		h.Swap(i, largestIdx)
		// Recursively heapify (downward) possibly affected subtrees after the swap
		h.MaxHeapify(largestIdx)
	}
}

// CreateMaxHeap reverse walks an array, starting at the last non-leaf node index
// (i.e. the first parent node index) and turns it into a max-heap.
func (h *Heap) CreateMaxHeap() {
	// Last non-leaf node index in a heap is N/2-1
	for i := len(h.arr)/2 - 1; i >= 0; i-- {
		h.MaxHeapify(i)
	}
}

// MaxInsert inserts a node with value v into the max-heap
func (h *Heap) MaxInsert(v int) {
	h.arr = append(h.arr, v)
	h.CreateMaxHeap()
}

// MaxDelete deletes the first element in the heap and down max heapifies
func (h *Heap) MaxDelete() {
	// Pop initial element and shift last element to the front
	h.arr = append([]int{h.arr[len(h.arr)-1]}, h.arr[1:len(h.arr)-1]...)
	// Down heapify, root to last parent node
	for i := 0; i < len(h.arr)/2-1; i++ {
		h.MaxHeapify(i)
	}
}

// MinHeapify converts a 0-based array into a min-heap
func (h *Heap) MinHeapify(i int) {
	rightChildIdx := h.RightChildIdx(i)
	leftChildIdx := h.LeftChildIdx(i)

	smallestIdx := i
	if leftChildIdx < len(h.arr) && h.arr[leftChildIdx] < h.arr[smallestIdx] {
		smallestIdx = leftChildIdx
	}
	if rightChildIdx < len(h.arr) && h.arr[rightChildIdx] < h.arr[smallestIdx] {
		smallestIdx = rightChildIdx
	}

	if smallestIdx != i {
		h.Swap(i, smallestIdx)
		h.MinHeapify(smallestIdx)
	}
}

// CreateMinHeap reverse walks an array, starting at the last non-leaf node
// index (i.e. the first parent node index) and turns it into a min-heap.
func (h *Heap) CreateMinHeap() {
	for i := len(h.arr)/2 - 1; i >= 0; i-- {
		h.MinHeapify(i)
	}
}

// MinInsert inserts a node with value v into the min-heap
func (h *Heap) MinInsert(v int) {
	h.arr = append(h.arr, v)
	h.CreateMinHeap()
}

// MinDelete removes the root node and down min-heapifies
func (h *Heap) MinDelete() {
	h.arr = append([]int{h.arr[len(h.arr)-1]}, h.arr[1:len(h.arr)-1]...)
	for i := 0; i <= len(h.arr)/2-1; i++ {
		h.MinHeapify(i)
	}
}
