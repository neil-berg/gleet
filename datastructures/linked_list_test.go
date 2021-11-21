package datastructures

import (
	"testing"
)

func TestListValueAt(t *testing.T) {
	l := New()
	nAdd := 10
	for i := 0; i < nAdd; i++ {
		l.AddToTail(i)
	}
	l.Print()

	targetIdx := 3
	got := l.ValueAt(targetIdx)
	want := 3
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestListAddToTail(t *testing.T) {
	l := New()

	nAdd := 9
	for i := 0; i < nAdd; i++ {
		l.AddToTail(i)
	}
	l.Print()

	if l.Len != nAdd {
		t.Errorf("bad length")
	}
}

func TestListAddToHead(t *testing.T) {
	l := New()

	nAdd := 5
	for i := 0; i < nAdd; i++ {
		l.AddToHead(i)
	}
	l.Print()

	if l.Len != nAdd {
		t.Errorf("bad length")
	}
}

func TestListInsertAt(t *testing.T) {
	l := New()

	nAdd := 5
	for i := 0; i < nAdd; i++ {
		l.AddToTail(i)
	}
	insertIdx := 2
	insertVal := 25
	l.InsertAt(insertVal, insertIdx)
	l.Print()

	tests := map[string]struct {
		got  int
		want int
	}{
		"length":   {got: l.Len, want: nAdd + 1},
		"inserted": {got: l.ValueAt(insertIdx).(int), want: insertVal},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.got != test.want {
				t.Errorf("got %v, want %v", test.got, test.want)
			}
		})
	}
}
