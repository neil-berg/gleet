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
	insertIdx := 3
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

func TestListRemoveAtHead(t *testing.T) {
	l := New()

	nAdd := 5
	for i := 0; i < nAdd; i++ {
		l.AddToTail(i)
	}
	l.RemoveAtHead()
	l.Print()

	if l.Len != nAdd-1 {
		t.Errorf("bad length")
	}
}

func TestListRemoveAtTail(t *testing.T) {
	l := New()

	nAdd := 5
	for i := 0; i < nAdd; i++ {
		l.AddToTail(i)
	}
	l.RemoveAtTail()
	l.Print()

	if l.Len != nAdd-1 {
		t.Errorf("bad length")
	}
}

func TestListRemovedAt(t *testing.T) {
	l := New()

	nAdd := 5
	for i := 0; i < nAdd; i++ {
		l.AddToTail(i)
	}
	l.RemoveAt(2)
	l.Print()

	if l.Len != nAdd-1 {
		t.Errorf("bad length")
	}
}

func TestListReverse(t *testing.T) {
	l := New()

	nAdd := 5
	for i := 0; i < nAdd; i++ {
		l.AddToTail(i)
	}
	l.Reverse()
	l.Print()

	tests := map[string]struct {
		got, want int
	}{
		"length":     {got: nAdd, want: nAdd},
		"fist value": {got: l.ValueAt(0).(int), want: 4},
		"last value": {got: l.ValueAt(nAdd - 1).(int), want: 0},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.got != test.want {
				t.Errorf("got %v, want %v", test.got, test.want)
			}
		})
	}
}
