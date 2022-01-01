package sorting

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	tests := map[string]struct {
		a    []int
		want []int
	}{
		"first":  {a: []int{5, 3, 7, 1}, want: []int{1, 3, 5, 7}},
		"second": {a: []int{-5, 3, -7, 1}, want: []int{-7, -5, 1, 3}},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			Quicksort(test.a, 0, len(test.a)-1)
			if !reflect.DeepEqual(test.a, test.want) {
				t.Errorf("got %v, want %v", test.a, test.want)
			}
		})
	}
}
