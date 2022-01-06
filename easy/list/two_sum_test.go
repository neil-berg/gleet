package easy

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := map[string]struct {
		nums   []int
		target int
		want   []int
	}{
		"all pos":  {nums: []int{2, 7, 11, 3}, target: 9, want: []int{0, 1}},
		"some neg": {nums: []int{5, 0, -1}, target: 4, want: []int{0, 2}},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			got := TwoSum(c.nums, c.target)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got: %v, want: %v", got, c.want)
			}

		})
	}
}
