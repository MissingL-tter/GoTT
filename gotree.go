package gotree

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

	for len(values) != 0 {
		val, values = values[0], values[1:]

		tree.Insert(val)
	}

	return tree
}

// Insert inserts a new node with passed value into the tree
//
// Values <= the current node's Value will branch left, while values > the current node's value will branch right
func (tree *Tree) Insert(val float32) {

	var parent *Tree
	for tree != nil {
		if val <= tree.Value {
			parent = tree
			tree = tree.Left
		} else {
			parent = tree
			tree = tree.Right
		}
	}
	tree = &Tree{parent, val, nil, nil}
	if val <= parent.Value {
		parent.Left = tree
	} else {
		parent.Right = tree
	}
}

// InOrder traverses over the tree branching left, visiting the node, and then branching right
func InOrder(tree *Tree) {
	if tree != nil {
		InOrder(tree.Left)
		//fmt.Printf("%v\n", tree.Value)
		tree.Value += 0
		InOrder(tree.Right)
	}
}

// Search does stuff
func (tree *Tree) Search(val float32) *Tree {
	for tree != nil && tree.Value != val {
		if val <= tree.Value {
			tree = tree.Left
		} else {
			tree = tree.Right
		}
	}
	return tree
}
