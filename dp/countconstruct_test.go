package dp

import (
	"testing"
)

func TestCountConstruct(t *testing.T) {
	cache := make(map[string]int)
	got := CountConstruct("hellohello", []string{"he", "llo", "h", "ello", "hello"}, cache)
	if got != 9 {
		t.Errorf("got %v, want %v", got, 9)
	}
}
