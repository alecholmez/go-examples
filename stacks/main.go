package main

import (
	"fmt"
	"log"
)

func main() {

	// Create our stack with values
	s := New(1, 2, 3, 4, 5, 6, 7, 8)

	// Add another value
	s.Push(9)

	// Print out our full stack
	s.Print()

	// Search for a value in the stack
	v, err := s.Search(2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)
}
