package medium

import "testing"

func TestLongestConsecutiveSequence(t *testing.T) {
	tests := map[string]struct {
		a    []int
		want int
	}{
		"positives": {a: []int{100, 4, 200, 7, 6, 5}, want: 4},
		"negatives": {a: []int{-100, -101, -3, -6}, want: 2},
		"mix signs": {a: []int{0, -1, 2, -2}, want: 3},
		"length 1":  {a: []int{1}, want: 1},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			got := LongestConsecutiveSequence(test.a)
			want := test.want
			if got != want {
				t.Errorf("got %v, wnat %v", got, want)
			}
		})
	}
}
