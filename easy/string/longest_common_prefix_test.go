package easy

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		words []string
		want  string
	}{
		{words: []string{"flower", "flow", "flaw"}, want: "fl"},
		{words: []string{"flower", "flow", "f"}, want: "f"},
		{words: []string{"flower", "plow", "claw"}, want: ""},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := LongestCommonPrefix(tt.words)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
