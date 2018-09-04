package gomine

import (
	"testing"

	"github.com/ThomasChuDesigns/gomine"
)

var results = map[string][]interface{}{
	"new":  []interface{}{nil, -1},
	"push": []interface{}{9},
	"pop":  []interface{}{"a", 1},
}

func TestStack(t *testing.T) {
	// check if new stack is created
	s := gomine.NewStack()

	if s == results["new"][0] {
		t.Error("stack not initialized")
	}

}

func TestStackPush(t *testing.T) {
	s := gomine.NewStack()

	// testing concurrent pushing
	for w := 0; w < 10; w++ {
		s.Push(w)
	}

	if s.Current.Val.(int) != results["push"][0] {
		t.Error("push function not working")
	}
}

func TestStackPop(t *testing.T) {
	s := gomine.NewStack("a", "b", "c", "d")

	// get to peek index = n/4 check if stack size shrinks to n/2
	s.Pop()
	s.Pop()
	s.Pop()

	if s.Current.Val.(string) != results["pop"][0] || s.Size != results["pop"][1] {
		t.Errorf("got: %v, expect: %v\n", []interface{}{s.Current.Val.(string), s.Size}, results["pop"])
	}

	// check if we can pop empty stack
	s.Pop()
	s.Pop()

	if !s.IsEmpty() {
		t.Error("stack supposed to be empty")
	}
}
