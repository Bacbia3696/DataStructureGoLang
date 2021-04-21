package linklist

import (
	"testing"
)

func TestLength(t *testing.T) {
	l := LinkedList{}
	if l.Length() != 0 {
		t.Errorf("Expect 0 length, got %v\n", l.Length())
	}
}

func TestInsertBegin(t *testing.T) {
	l := LinkedList{}
	l.InsertBegin(10)
	_ = l.Display()
}
