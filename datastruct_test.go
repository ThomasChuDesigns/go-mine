package gomine

import (
	"sync"
	"testing"
)

var results = map[string][]interface{}{
	"new":  []interface{}{nil, -1},
	"push": []interface{}{9},
	"pop":  []interface{}{0, 2},
}

func TestStack(t *testing.T) {
	s := NewStack()
	if s == results["new"][0] {
		t.Error("stack not initialized")
	}

	t.Logf("stack address %p", &s)

	var wg sync.WaitGroup

	wg.Add(1)
	go s.Push(1, &wg)
	wg.Wait()

	wg.Add(1)
	go s.Pop(&wg)
	wg.Wait()

	if s.current != results["new"][1] {
		t.Error("push/pop function not working")
	}
}

func TestStackPush(t *testing.T) {
	var wg sync.WaitGroup
	s := NewStack()

	wg.Add(10)
	for w := 0; w < 10; w++ {
		go s.Push(w, &wg)
	}
	wg.Wait()

	if s.current != results["push"][0] {
		t.Error("push function not working")
	}

	t.Log(s.array[:10])
}

func TestStackPop(t *testing.T) {
	var wg sync.WaitGroup
	s := NewStack("a", "b", "c", "d")

	// get to peek index = n/4 check if stack size shrinks to n/2
	wg.Add(3)
	go s.Pop(&wg)
	go s.Pop(&wg)
	go s.Pop(&wg)
	wg.Wait()

	if s.current != results["pop"][0] || s.size != results["pop"][1] {
		t.Errorf("got: %v, expect: %v\n", []int{s.current, s.size}, results["pop"])
	}

	// check if we can pop empty stack
	wg.Add(2)
	go s.Pop(&wg)
	go s.Pop(&wg)
	wg.Wait()

	if !s.IsEmpty() {
		t.Error("stack supposed to be empty")
	}
}
