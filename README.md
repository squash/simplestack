package simplestack // import "github.com/squash/simplestack"

git

TYPES

type Stack struct {
	// Has unexported fields.
}
    Stack is an instance of the stack container

func NewStack(size int) *Stack
    NewStack provisions a new instance of the stack container

func (s *Stack) Count() int
    Count returns the current size of the stack

func (s *Stack) Dump() []interface{}
    Dump retrieves the entire stack's contents

func (s *Stack) Flush()
    Flush removes all stack entries and continues with an empty stack

func (s *Stack) Pop() (interface{}, bool)
    Pop grabs the newest entry from the stack

func (s *Stack) Push(entry interface{})
    Push adds a new entry to the stack, and drops the oldest entry if it has
    reached its size limit

