package sorting

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := map[string]struct {
		a    []int
		want []int
	}{
		"first":  {a: []int{5, 7, 2, 1, 3}, want: []int{1, 2, 3, 5, 7}},
		"second": {a: []int{-5, 7, -2, 1, -3, 4}, want: []int{-5, -3, -2, 1, 4, 7}},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			got := MergeSort(test.a)
			want := test.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
