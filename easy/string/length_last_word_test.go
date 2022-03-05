package easy

import "testing"

func TestLengthLastWord(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "hello world", want: 5},
		{s: "hello world   ", want: 5},
		{s: "   fly me   to   the moon  ", want: 4},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := LengthLastWord(tt.s)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
