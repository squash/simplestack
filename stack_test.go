package simplestack

import (
	"testing"
)

func TestStack(t *testing.T) {
	c := NewStack(10)
	for x := 0; x <= 20; x++ {
		c.Push(x)
	}
	tmp, ok := c.Pop()
	if !ok {
		t.Fatal("Pop returned no entry")
	}
	x := tmp.(int)
	if x != 20 {
		t.Fatalf("Pop returned a value that was not the most recently pushed value: %d", x)
	}
	tmp, ok = c.Pop()
	if !ok {
		t.Fatal("Pop returned no entry")
	}
	x = tmp.(int)
	if x != 19 {
		t.Fatal("Pop returned a value that was not the second most recently pushed value")
	}
	if c.Count() != 8 {
		t.Fatalf("Count returned the wrong number of entries: %d", c.Count())
	}

}
