package gotree

import (
	"sync"
)

// Tree represents a binary tree structure
type Tree struct {
	Parent *Tree
	Value  float32
	Left   *Tree
	Right  *Tree
	Level  int
}

// Build takes some slice of float32s and builds a binary search tree
func Build(values []float32) *Tree {
	var val float32
	val, values = values[0], values[1:]

	tree := &Tree{nil, val, nil, nil, 0}
	pool := make([]Tree, len(values))

	for i := 0; i < len(values); i++ {
		tree.Insert(values[i], &(pool[i]))
	}

	return tree
}

// Insert inserts a new node with passed value into the tree
//
// Values <= the current node's Value will branch left, while values > the current node's value will branch right
func (root *Tree) Insert(val float32, tree *Tree) {
	var parent *Tree
	for root != nil {
		if val <= root.Value {
			parent = root
			root = root.Left
		} else {
			parent = root
			root = root.Right
		}
	}
	tree.Parent = parent
	tree.Value = val
	tree.Level = parent.Level + 1
	if val <= parent.Value {
		parent.Left = tree
	} else {
		parent.Right = tree
	}
}

// InOrder traverses over the tree branching left, visiting the node, and then branching right
func (root *Tree) InOrder() {

	if root != nil {
		root.Left.InOrder()
		//Print or visit node
		root.Right.InOrder()
	}
}

// InOrderParallel performs in in order traversal of the tree making use of multiple threads
func (root *Tree) InOrderParallel() {

	if root != nil {
		if root.Level <= 8 {
			wg := &sync.WaitGroup{}
			wg.Add(2)
			go func() {
				defer wg.Done()
				root.Left.InOrderParallel()
			}()
			//Print or visit node
			go func() {
				defer wg.Done()
				root.Right.InOrderParallel()
			}()
			wg.Wait()
		} else {
			root.Left.InOrderParallel()
			//Print or visit node
			root.Right.InOrderParallel()
		}
	}
}

// Search performs a search on tree given some float32, and returns the node that contains that val.
//
// If the tree contains duplicates this will return only the first found
func (root *Tree) Search(val float32) *Tree {

	for root != nil && root.Value != val {
		if val <= root.Value {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return root
}
