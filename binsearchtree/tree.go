package main

import "errors"

// Tree ...
type Tree struct {
	Root *Node
}

// Insert will inject a new node into the binary search tree
func (t *Tree) Insert(index int64, data interface{}) error {
	if t.Root == nil {
		t.Root = &Node{
			Index: index,
			Data:  data,
		}
		return nil
	}

	return t.Root.Insert(index, data)
}

// Find will recursively search it's node structure
func (t *Tree) Find(index int64) (interface{}, error) {
	if t.Root == nil {
		return nil, errors.New("Missing root node")
	}
	return t.Root.Find(index)
}

// Delete will remove an item from the binary tree
func (t *Tree) Delete(index int64) error {
	if t.Root == nil {
		return errors.New("Missing root node")
	}

	fakeParent := &Node{Right: t.Root}
	err := t.Root.Delete(index, fakeParent)
	if err != nil {
		return err
	}

	if fakeParent.Right == nil {
		t.Root = nil
	}

	return nil
}

// Traverse will go down the tree in left to right order (smallest to largest)
func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}
