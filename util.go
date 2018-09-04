package gomine

import (
	"sync"
)

// Stack object allows goroutines for
// pushing/popping data while
// handling race conditions
type Stack struct {
	Current *StackNode
	Size    int
	sync.Mutex
}

// StackNode is a linked list that handles the stacks dynamic allocation/deallocation
type StackNode struct {
	Val  interface{}
	Next *StackNode
	Prev *StackNode
}

// Push appends a value into the stack
func (s *Stack) Push(value interface{}) {
	// Handle race conditions when accessing same memory space
	s.Lock()
	defer s.Unlock()

	newNode := &StackNode{value, nil, s.Current}

	// stack is empty so this node is at the top
	if s.Current == nil {
		s.Current = newNode
	}

	// link current node to new node
	s.Current.Next = newNode
	s.Current = newNode

	// update size of stack
	s.Size++
}

// Pop takes the most recent value out of the stack
func (s *Stack) Pop() interface{} {
	// Handle race conditions when accessing same memory space
	s.Lock()
	defer s.Unlock()

	// nothing to pop
	if s.Current == nil {
		return nil
	}

	// take node from the top of stack and set current to previous node
	popVal := s.Current.Val
	s.Current = s.Current.Prev

	s.Size--
	return popVal
}

// AsyncPush appends a value into the stack
func (s *Stack) AsyncPush(value interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	s.Push(value)

}

// AsyncPop takes the most recent value out of the stack
func (s *Stack) AsyncPop(wg *sync.WaitGroup) interface{} {
	defer wg.Done()

	return s.Pop()
}

// NewStack initializes a new stack
func NewStack(a ...interface{}) *Stack {
	newStack := &Stack{Current: nil, Size: 0}

	// if params available use given values for stack
	if a != nil {
		// iterate through by pushing items to stack
		for _, val := range a {
			newStack.Push(val)
		}
	}

	return newStack
}

// IsEmpty checks if stack is empty
func (s *Stack) IsEmpty() bool {
	s.Lock()
	defer s.Unlock()

	return s.Size == 0
}
