package main

import (
	"fmt"
	"log"
)

func main() {

	// Some test data
	indexes := []int64{1, 2, 3, 4, 5, 6, 7, 8}
	data := []string{"alpha", "bravo", "charlie", "delta", "echo", "foxtrot", "golf", "hotel"}

	// Create our tree
	tree := &Tree{}

	for i, index := range indexes {
		err := tree.Insert(index, data[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Print("Sorted values: | ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Index, ": ", n.Data, " | ") })
	fmt.Println()

	search := int64(4)
	fmt.Print("Find node '", search, "': ")
	d, err := tree.Find(search)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)

	err = tree.Delete(search)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Sorted values: | ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Index, ": ", n.Data, " | ") })
	fmt.Println()
}
