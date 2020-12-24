package simplestack

import (
	"log"
	"testing"
)

func TestStack(t *testing.T) {
	c := NewStack(200)
	for x := 0; x <= 259; x++ {
		c.Push(x)
	}
	tmp, ok := c.Pop()
	if !ok {
		t.Fatal("Pop returned no entry")
	}
	x := tmp.(int)
	if x != 259 {
		t.Fatalf("Pop returned a value that was not the most recently pushed value: %d", x)
	}
	tmp, ok = c.Pop()
	if !ok {
		t.Fatal("Pop returned no entry")
	}
	x = tmp.(int)
	if x != 258 {
		t.Fatal("Pop returned a value that was not the second most recently pushed value")
	}
	if c.Count() != 198 {
		t.Fatalf("Count returned the wrong number of entries: %d", c.Count())
	}
	z:=c.Dump()
	log.Printf("%#v", z)
	
}
