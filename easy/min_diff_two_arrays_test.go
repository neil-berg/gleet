package easy

import "testing"

func TestMinDiffTwoArrays(t *testing.T) {
	a := []int{5, 1, 9, 12}
	b := []int{30, 8, 3, 25}
	got := MinDiffTwoArrays(a, b)
	want := 1
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
