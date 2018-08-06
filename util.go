package gomine

import (
	"math"
	"sync"
)

// Stack object allows goroutines for
// pushing/popping data while
// handling race conditions
type Stack struct {
	Array   []interface{}
	Current int
	Size    int
	sync.Mutex
}

func (s *Stack) resize(newSize int) {
	// copy current array to new array with newSize
	newArray := make([]interface{}, newSize)
	copy(newArray, s.Array)

	s.Array = newArray
	s.Size = newSize
}

// Push appends a value into the stack
func (s *Stack) Push(value interface{}) {
	// Handle race conditions when accessing same memory space
	s.Lock()
	defer s.Unlock()

	s.Current++
	s.Array[s.Current] = value

	// if current index + 1 == n then expand array by 2n
	if s.Current+1 == s.Size {
		s.resize(2 * s.Size)
	}

}

// Pop takes the most recent value out of the stack
func (s *Stack) Pop() interface{} {
	// Handle race conditions when accessing same memory space
	s.Lock()
	defer s.Unlock()

	if s.Current == -1 {
		return nil
	}

	// replace peek value with nil
	var value = s.Array[s.Current]
	s.Array[s.Current] = nil

	//if current == n/4 then reduce array to n/2
	resizeValue := int(math.Ceil(float64(s.Size) / 4.0))
	if s.Current == resizeValue {
		newSize := int(math.Ceil(float64(s.Size) / 2.0))
		s.resize(newSize)
	}

	s.Current--
	return value
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
	// default values for stack
	stackArray := make([]interface{}, 4)
	stackSize := 4
	stackCurrent := -1

	// if params available use given values for stack
	if a != nil {
		stackArray = a
		stackSize = len(a)
		stackCurrent = len(a) - 1
	}
	return &Stack{Array: stackArray, Current: stackCurrent, Size: stackSize}
}

// IsEmpty checks if stack is empty
func (s *Stack) IsEmpty() bool {
	s.Lock()
	defer s.Unlock()

	if s.Current == -1 {
		return true
	}
	return false
}
