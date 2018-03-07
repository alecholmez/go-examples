package main

import (
	"errors"
	"fmt"
)

// Stack data structure
type Stack struct {
	top    *node
	length int
}

type node struct {
	value interface{}
	prev  *node
}

// New will create a new stack in memory
func New(values ...interface{}) *Stack {
	if len(values) == 0 {
		return &Stack{
			top:    nil,
			length: 0,
		}
	}

	// If we receive values in the constructor, fill the stack and return it
	s := &Stack{
		// Create our top node so we don't get a nil pointer dereference
		top: &node{
			value: values[0],
			prev:  nil,
		},
		length: 0,
	}
	for i, v := range values {
		// Ignore the first value since we created that previously
		if i != 0 {
			s.Push(v)
		}
	}
	return s
}

// Len will return the length of the stack
func (s *Stack) Len() int {
	return s.length
}

// Peek will return the top of the stack
func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

// Pop will remove the top item of the stack and return it
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	// Retrieve the node and reset the top of the stack as well as the length
	node := s.top
	s.top = node.prev
	s.length--

	return node.value
}

// Push will add a new item to the top of the stack
func (s *Stack) Push(v interface{}) {
	node := &node{
		value: v,
		prev:  s.top,
	}

	// Reset the top stack node and increase the length
	s.top = node
	s.length++
}

// Print will traverse the stack and print out the values to os.Stdout
func (s *Stack) Print() {
	top := s.top

	for i := 0; i <= s.length; i++ {
		fmt.Printf("%v -> ", top.value)
		top = top.prev
	}

	fmt.Println()
}

// Search will traverse the stack and look for a specific value
func (s *Stack) Search(value interface{}) (interface{}, error) {
	// replicate the stack so we can modify the top node
	var err error

	top := s.top

	for i := 0; i <= s.length; i++ {
		switch {
		case top.value == value:
			return value, nil
		case top.value == nil:
			err = errors.New("Item not found in stack")
		}
		top = top.prev
	}

	return nil, err
}
