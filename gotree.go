package gotree

import "sync"

// Tree represents a binary tree structure
type Tree struct {
	Parent *Tree
	Value  float32
	Left   *Tree
	Right  *Tree
}

// Build takes some slice of float32s and builds a binary search tree
func Build(values []float32) *Tree {
	var val float32
	val, values = values[0], values[1:]

	tree := &Tree{nil, val, nil, nil}
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
	if val <= parent.Value {
		parent.Left = tree
	} else {
		parent.Right = tree
	}
}

// InOrder traverses over the tree branching left, visiting the node, and then branching right
func InOrder(tree *Tree, level int) {

	if tree != nil {
		if level == 1 {
			wg := &sync.WaitGroup{}
			wg.Add(2)
			go func() {
				inOrderFast(tree.Right)
				wg.Done()
			}()
			go func() {
				inOrderFast(tree.Left)
				wg.Done()
			}()
			wg.Wait()
		} else {
			InOrder(tree.Right, level+1)
			InOrder(tree.Left, level+1)
		}
	}
}

func inOrderFast(tree *Tree) {
	if tree != nil {
		inOrderFast(tree.Right)
		inOrderFast(tree.Left)
	}
}

// Search does stuff
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
