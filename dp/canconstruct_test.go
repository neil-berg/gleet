package dp

import (
	"testing"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		target string
		words  []string
		want   bool
	}{
		{target: "abcdef", words: []string{"ab", "abc", "cd", "def", "abcd"}, want: true},
		{target: "homebrew", words: []string{"home", "bre", "w"}, want: true},
		{target: "test", words: []string{"tes", "s", "st"}, want: false},
		{target: "enterapotentpot", words: []string{"a", "p", "ent", "enter", "ot", "o", "t"}, want: true},
		{target: "hellohello", words: []string{"he", "llo", "h", "ello", "hello"}, want: true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			cache := make(map[string]bool)
			got := CanConstruct(tt.target, tt.words, cache)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
