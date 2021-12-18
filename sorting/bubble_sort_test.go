package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := map[string]struct {
		a    []int
		want []int
	}{
		"pos": {a: []int{5, 3, 7, 2, 6}, want: []int{2, 3, 5, 6, 7}},
		"neg": {a: []int{-5, -3, -7, -2, -6}, want: []int{-7, -6, -5, -3, -2}},
		"mix": {a: []int{-5, 3, -7, 2, -6}, want: []int{-7, -6, -5, 2, 3}},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			got := BubbleSort(test.a, len(test.a))
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
