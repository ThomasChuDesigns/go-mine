package gomine

import (
	"math"
	"sync"
)

// Stack object allows goroutines for
// pushing/popping data while
// handling race conditions
type Stack struct {
	array   []interface{}
	current int
	size    int
	sync.Mutex
}

func (s *Stack) resize(newSize int) {
	// copy current array to new array with newSize
	newArray := make([]interface{}, newSize)
	copy(newArray, s.array)

	s.array = newArray
	s.size = newSize
}

// Push appends a value into the stack
func (s *Stack) Push(value interface{}, wg *sync.WaitGroup) {
	// Handle race conditions when accessing same memory space
	s.Lock()
	defer s.Unlock()
	defer wg.Done()

	s.current++
	s.array[s.current] = value

	// if current index + 1 == n then expand array by 2n
	if s.current+1 == s.size {
		s.resize(2 * s.size)
	}

}

// Pop takes the most recent value out of the stack
func (s *Stack) Pop(wg *sync.WaitGroup) interface{} {
	// Handle race conditions when accessing same memory space
	s.Lock()
	defer s.Unlock()
	defer wg.Done()

	if s.current == -1 {
		return nil
	}

	// replace peek value with nil
	var value = s.array[s.current]
	s.array[s.current] = nil

	//if current == n/4 then reduce array to n/2
	resizeValue := int(math.Ceil(float64(s.size) / 4.0))
	if s.current == resizeValue {
		newSize := int(math.Ceil(float64(s.size) / 2.0))
		s.resize(newSize)
	}

	s.current--
	return value
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
	return &Stack{array: stackArray, current: stackCurrent, size: stackSize}
}

// IsEmpty checks if stack is empty
func (s *Stack) IsEmpty() bool {
	if s.current == -1 {
		return true
	}
	return false
}
