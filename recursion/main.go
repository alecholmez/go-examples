package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sort"

	"github.com/alecholmez/go-examples/recursion/mutual"
)

func main() {

	// Make our list of ints in mem
	list := make([]int, 100)

	// Fill the list with random numbers
	list = PopulateList(list)

	// Sort the list of integers
	sort.Ints(list)

	// Start at position 0
	err := Print(0, list)
	if err != nil {
		// Crash the program if we receive an error
		log.Fatal(err)
	}

	// Mutual Recursion
	// See package mutual for more information
	fmt.Println(mutual.IsEven(4))
	fmt.Println(mutual.IsEven(3))
	fmt.Println(mutual.IsOdd(1))
	fmt.Println(mutual.IsOdd(8))
}

// Print will recursively loop through the loop
func Print(index int, list []int) error {
	if index > len(list) {
		return errors.New("Index is out of list range")
	}

	// Print and increase the index
	// We are acting like a loop here
	fmt.Println(list[index])
	index++

	// Define our exit condition
	// We don't want to exceed the bounds of the array so we check here
	if index < len(list) {
		Print(index, list)
	}

	// If index is larger than the list length, exit the recursive loop
	return nil
}

// PopulateList will randomly fill the list
func PopulateList(list []int) []int {
	for i := 0; i < len(list); i++ {
		list[i] = rand.Int()
	}

	return list
}
