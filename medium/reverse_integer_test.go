package medium

import "testing"

func TestReverseInteger(t *testing.T) {
	tests := map[string]struct {
		n    int
		want int
	}{
		"pos":   {n: 123, want: 321},
		"neg":   {n: -123, want: -321},
		"zeros": {n: -123, want: -321},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			got := ReverseInteger(test.n)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
