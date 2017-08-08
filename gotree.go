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

// BuildParallel takes some slice of float32s and builds a binary search tree making use of multiple threads
func BuildParallel(values []float32) *Tree {
	var val float32
	val, values = values[0], values[1:]

	tree := &Tree{nil, val, nil, nil}
	pool := make([]Tree, len(values))

	i := 0
	for ; i < len(values); i++ {
		tree.Insert(values[i], &(pool[i]))
		if tree.Left != nil && tree.Right != nil {
			i++
			break
		}
	}
	if tree.Left != nil && tree.Right != nil {

		wg := sync.WaitGroup{}
		queuedAdder := func(subtree *Tree, q chan float32) {
			for v := range q {
				subtree.Insert(v, &Tree{})
				wg.Done()
			}
		}
		rightVals := make(chan float32, 1000)
		//rightPool := make(chan *Tree, 1000)
		leftVals := make(chan float32, 1000)
		//leftPool := make(chan *Tree, 1000)

		go queuedAdder(tree.Right, rightVals)
		go queuedAdder(tree.Left, leftVals)

		for ; i < len(values); i++ {
			wg.Add(1)
			if values[i] <= tree.Value {
				leftVals <- values[i]
			} else {
				rightVals <- values[i]
			}
		}
		wg.Wait()
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
func InOrder(tree *Tree) {

	if tree != nil {
		InOrder(tree.Right)
		InOrder(tree.Left)
	}
}

// InOrderParallel performs in in order traversal of the tree making use of multiple threads
func InOrderParallel(tree *Tree, level ...int) {

	if level == nil {
		level = []int{0}
	}

	if tree != nil {
		if level[0] == 2 {
			wg := &sync.WaitGroup{}
			wg.Add(2)
			go func() {
				InOrder(tree.Right)
				wg.Done()
			}()
			go func() {
				InOrder(tree.Left)
				wg.Done()
			}()
			wg.Wait()
		} else {
			InOrderParallel(tree.Right, level[0]+1)
			InOrderParallel(tree.Left, level[0]+1)
		}
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
