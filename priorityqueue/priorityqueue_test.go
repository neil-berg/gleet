package priorityqueue

import (
	"reflect"
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	pq := &PriorityQueue{
		Items: []PQNode{
			PQNode{Priority: 5, Name: "Alice"},
			PQNode{Priority: 6, Name: "Bill"},
			PQNode{Priority: 3, Name: "Bob"},
			PQNode{Priority: 7, Name: "Sam"},
			PQNode{Priority: 9, Name: "Joe"},
		},
	}

	pq.NewPriorityQueue()

	var priorities []int
	for _, node := range pq.Items {
		priorities = append(priorities, node.Priority)
	}

	want := []int{9, 7, 3, 5, 6}
	if !reflect.DeepEqual(priorities, want) {
		t.Errorf("got %v, want %v", priorities, want)
	}
}

func TestInsert(t *testing.T) {
	pq := &PriorityQueue{
		Items: []PQNode{
			PQNode{Priority: 5, Name: "Alice"},
			PQNode{Priority: 6, Name: "Bill"},
			PQNode{Priority: 3, Name: "Bob"},
			PQNode{Priority: 7, Name: "Sam"},
			PQNode{Priority: 9, Name: "Joe"},
		},
	}

	pq.NewPriorityQueue()
	pq.Insert(PQNode{Priority: 13, Name: "Jane"})

	var priorities []int
	for _, node := range pq.Items {
		priorities = append(priorities, node.Priority)
	}

	want := []int{13, 7, 9, 5, 6, 3}
	if !reflect.DeepEqual(priorities, want) {
		t.Errorf("got %v, want %v", priorities, want)
	}
}

func TestDelete(t *testing.T) {
	pq := &PriorityQueue{
		Items: []PQNode{
			PQNode{Priority: 5, Name: "Alice"},
			PQNode{Priority: 6, Name: "Bill"},
			PQNode{Priority: 3, Name: "Bob"},
			PQNode{Priority: 7, Name: "Sam"},
			PQNode{Priority: 9, Name: "Joe"},
		},
	}

	pq.NewPriorityQueue()
	pq.Delete()

	var priorities []int
	for _, node := range pq.Items {
		priorities = append(priorities, node.Priority)
	}

	want := []int{7, 6, 5, 3}
	if !reflect.DeepEqual(priorities, want) {
		t.Errorf("got %v, want %v", priorities, want)
	}
}
