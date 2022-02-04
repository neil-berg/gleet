package medium

import "testing"

func TestReverseBetweenParens(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{s: "(u(love)i)", want: "iloveu"},
		{s: "(ed(et(oc))el)", want: "leetcode"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := ReverseBetweenParens(tt.s)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
