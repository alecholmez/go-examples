package main

import "errors"

// Node - basis for binary search tree
type Node struct {
	Index int64
	Data  interface{}
	Left  *Node
	Right *Node
}

// Insert will inject a new node
func (n *Node) Insert(index int64, data interface{}) error {
	if n == nil {
		// We can insert anything into a tree with no root node. Return an error if that's the case
		return errors.New("No root node found")
	}

	switch {
	// If the data already exists in the tree, return
	case index == n.Index:
		return nil
	case index < n.Index:
		// If there is no left subtree, assume we can insert a node here
		if n.Left == nil {
			n.Left = &Node{
				Index: index,
				Data:  data,
			}
			return nil
		}
		// If there is a left subtree, recursively call the insert method
		return n.Left.Insert(index, data)
	case index > n.Index:
		// If there is no right subtree, assume we can insert a node here
		if n.Right == nil {
			n.Right = &Node{
				Index: index,
				Data:  data,
			}
			return nil
		}
		// If there is a right subtree, recursively call the insert method
		return n.Right.Insert(index, data)
	}

	return nil
}

// Find will traverse through the tree and search for the requested value
func (n *Node) Find(i int64) (interface{}, error) {
	if n == nil {
		return nil, errors.New("No root node exists")
	}

	// If the value is not at the currently selected node, traverse down the height of the tree
	// using a recursive method
	switch {
	case i == n.Index:
		return n.Data, nil
	case i < n.Index:
		return n.Left.Find(i)
	default:
		return n.Right.Find(i)
	}
}

func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}

func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == nil {
		return errors.New("replace node not allowed on a nil node")
	}

	if n == parent.Left {
		parent.Left = replacement
		return nil
	}

	parent.Right = replacement
	return nil
}

// Delete will remove an item from the search tree
func (n *Node) Delete(i int64, parent *Node) error {
	if n == nil {
		return errors.New("value to be deleted does not exist in tree")
	}

	switch {
	case i < n.Index:
		return n.Left.Delete(i, n)
	case i > n.Index:
		return n.Right.Delete(i, n)
	default:
		switch {
		case n.Left == nil && n.Right == nil:
			n.replaceNode(parent, nil)
			return nil
		case n.Left == nil:
			n.replaceNode(parent, n.Right)
			return nil
		case n.Right == nil:
			n.replaceNode(parent, n.Left)
			return nil
		}

		replacement, replParent := n.Left.findMax(n)

		n.Index = replacement.Index
		n.Data = replacement.Data

		return replacement.Delete(replacement.Index, replParent)
	}
}
