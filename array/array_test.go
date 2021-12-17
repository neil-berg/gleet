package array

import (
	"testing"
)

func TestPush(t *testing.T) {
	arr := &Array{}

	startIdx := 0
	endIdx := 5
	for i := startIdx; i < endIdx; i++ {
		arr.Push(i)
	}

	tests := map[string]struct {
		want int
		got  int
	}{
		"length":        {want: endIdx, got: arr.len},
		"initial value": {want: startIdx, got: arr.items[0].(int)},
		"last value":    {want: endIdx - 1, got: arr.items[arr.len-1].(int)},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.got != test.want {
				t.Errorf("want %v, got %v", test.want, test.got)
			}
		})
	}
}

func TestPop(t *testing.T) {
	arr := &Array{}

	startIdx := 0
	endIdx := 10
	for i := startIdx; i < endIdx; i++ {
		arr.Push(i)
	}

	nPop := 4
	for i := 0; i < nPop; i++ {
		arr.Pop()
	}

	tests := map[string]struct {
		want int
		got  int
	}{
		"length":        {want: endIdx - nPop, got: endIdx - nPop},
		"initial value": {want: startIdx, got: arr.items[0].(int)},
		"last value":    {want: endIdx - 1 - nPop, got: arr.items[arr.len-1].(int)},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.got != test.want {
				t.Errorf("want %v, got %v", test.want, test.got)
			}
		})
	}
}

func TestInsertAt(t *testing.T) {
	arr := &Array{}

	startIdx := 0
	endIdx := 10
	for i := startIdx; i < endIdx; i++ {
		arr.Push(i)
	}

	insertVal := 25
	insertIdx := 4
	arr.InsertAt(insertVal, insertIdx)

	tests := map[string]struct {
		want int
		got  int
	}{
		"length":   {want: endIdx - 1, got: endIdx - 1},
		"inserted": {want: insertVal, got: arr.items[insertIdx].(int)},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.got != test.want {
				t.Errorf("want %v, got %v", test.want, test.got)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	arr := &Array{}

	startIdx := 0
	endIdx := 9
	for i := startIdx; i < endIdx; i++ {
		arr.Push(i)
	}

	arr.Reverse()

	tests := map[string]struct {
		want int
		got  int
	}{
		"length":        {want: endIdx, got: arr.len},
		"initial value": {want: endIdx - 1, got: arr.items[0].(int)},
		"last value":    {want: startIdx, got: arr.items[arr.len-1].(int)},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.got != test.want {
				t.Errorf("want %v, got %v", test.want, test.got)
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	arr := &Array{}
	startIdx := 0
	endIdx := 9
	for i := startIdx; i < endIdx; i++ {
		arr.Push(i)
	}

	tests := map[string]struct {
		v    int
		want bool
	}{"exists": {v: 5, want: true},
		"does not exist": {v: -2, want: false}}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			got := arr.BinarySearch(test.v)
			if got != test.want {
				t.Errorf("got %v, wnat %v", got, test.want)
			}
		})
	}
}
