package gotree

// Tree represent a binary tree structure
type Tree struct {
	Left  *Tree
	Value float32
	Right *Tree
}

// Build takes some array of inputs and builds a binary search tree
func Build(values []float32) *Tree {

	var val float32
	val, values = values[0], values[1:]

	tree := &Tree{nil, val, nil}

	for len(values) != 0 {
		val, values = values[0], values[1:]

		tree.Insert(val)
	}

	return tree
}

// Insert inserts a new node with passed value into the tree
func (tree *Tree) Insert(val float32) {

	for tree.Left != nil && val <= tree.Value {
		tree = tree.Left
	}
	for tree.Right != nil && val > tree.Value {
		tree = tree.Right
	}
	if val <= tree.Value {
		tree.Left = &Tree{nil, val, nil}
	} else {
		tree.Right = &Tree{nil, val, nil}
	}
}

// InOrder traverses over the tree branching left, visiting the node, and then branching right
func InOrder(tree *Tree) {
	if tree != nil {
		InOrder(tree.Left)
		// fmt.Printf("%v\n", tree.Value)
		tree.Value += 0
		InOrder(tree.Right)
	}
}
