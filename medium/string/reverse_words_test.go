package medium

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{s: "a   good   example  ", want: "example good a"},
		{s: "a   good   example  again", want: "again example good a"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := ReverseWords(tt.s)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
