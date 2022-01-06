package easy

import "testing"

func TestLongestSubstrNoRepeat(t *testing.T) {
	tests := map[string]struct {
		s    string
		want int
	}{
		"repeat in window":      {s: "abcdab", want: 4},
		"repeat outside window": {s: "abdba", want: 3},
		"two chars":             {s: "ab", want: 2},
		"all same":              {s: "bbbb", want: 1},
		"empty":                 {s: "", want: 0},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			got := LongestSubstrNoRepeat(test.s)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
