// Package simplestack implements a naive stack container that is thread safe and should be fine for many projects
package simplestack

import (
	"sync"
)

type stackEntry struct {
	value interface{}
}

// Stack is an instance of the stack container
type Stack struct {
	entries   []stackEntry
	sizelimit int
	lock      sync.Mutex
}

// NewStack provisions a new instance of the stack container
func NewStack(size int) *Stack {
	var s Stack
	s.sizelimit = size
	return &s
}

// Push adds a new entry to the stack, and drops the oldest entry if it has reached its size limit
func (s *Stack) Push(entry interface{}) {
	var e stackEntry
	e.value = entry
	s.lock.Lock()
	var tmp []stackEntry
	tmp = append(tmp, e)
	tmp = append(tmp, s.entries...)
	if len(tmp) >= s.sizelimit {
		tmp = tmp[0:s.sizelimit]
	}
	s.entries = tmp
	s.lock.Unlock()
}

// Pop grabs the newest entry from the stack
func (s *Stack) Pop() (interface{}, bool) {
	ok := false
	var tmp interface{}
	s.lock.Lock()
	if len(s.entries) > 0 {
		tmp = s.entries[0].value
		s.entries = s.entries[1:]
		ok = true
	}
	s.lock.Unlock()
	return tmp, ok
}

// Dump retrieves the entire stack's contents
func (s *Stack) Dump() []interface{} {
	s.lock.Lock()
	if len(s.entries)==0 {
		s.lock.Unlock();
		return nil
	}
	var tmp []interface{}
	for x := range s.entries {
		tmp = append(tmp, s.entries[x].value)
	}
	s.lock.Unlock()
	return tmp
}

// Flush removes all stack entries and continues with an empty stack
func (s *Stack) Flush() {
	s.lock.Lock()
	var tmp []stackEntry
	s.entries = tmp
	s.lock.Unlock()
}

// Count returns the current size of the stack
func (s *Stack) Count() int {
	s.lock.Lock()
	tmp := len(s.entries)
	s.lock.Unlock()
	return tmp
}

// VisitAll will call your function for each entry in the stack without modifying the stack
func (s *Stack) VisitAll(f func (interface{})) {
	for x:=range(s.entries) {
		f(s.entries[x].value)
	}
}