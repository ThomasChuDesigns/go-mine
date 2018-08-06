package gomine

import (
	"sync"
	"testing"

	"github.com/ThomasChuDesigns/gomine"
)

var results = map[string][]interface{}{
	"new":  []interface{}{nil, -1},
	"push": []interface{}{9},
	"pop":  []interface{}{0, 2},
}

func TestStack(t *testing.T) {
	// check if new stack is created
	s := gomine.NewStack()
	if s == results["new"][0] {
		t.Error("stack not initialized")
	}

	t.Logf("stack address %p", &s)

	// testing push and pop function
	var wg sync.WaitGroup

	wg.Add(1)
	go s.AsyncPush(1, &wg)
	wg.Wait()

	wg.Add(1)
	go s.AsyncPop(&wg)
	wg.Wait()

	if s.Current != results["new"][1] {
		t.Error("push/pop function not working")
	}
}

func TestStackPush(t *testing.T) {
	var wg sync.WaitGroup
	s := gomine.NewStack()

	// testing concurrent pushing
	wg.Add(10)
	for w := 0; w < 10; w++ {
		go s.AsyncPush(w, &wg)
	}
	wg.Wait()

	if s.Current != results["push"][0] {
		t.Error("push function not working")
	}

	t.Log(s.Array[:10])
}

func TestStackPop(t *testing.T) {
	var wg sync.WaitGroup
	s := gomine.NewStack("a", "b", "c", "d")

	// get to peek index = n/4 check if stack size shrinks to n/2
	wg.Add(3)
	go s.AsyncPop(&wg)
	go s.AsyncPop(&wg)
	go s.AsyncPop(&wg)
	wg.Wait()

	if s.Current != results["pop"][0] || s.Size != results["pop"][1] {
		t.Errorf("got: %v, expect: %v\n", []int{s.Current, s.Size}, results["pop"])
	}

	// check if we can pop empty stack
	wg.Add(2)
	go s.AsyncPop(&wg)
	go s.AsyncPop(&wg)
	wg.Wait()

	if !s.IsEmpty() {
		t.Error("stack supposed to be empty")
	}
}
