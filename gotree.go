package bst

// Tree represent a binary Tree structure
type Tree struct {
	Left  *Tree
	Value float64
	Right *Tree
}

// Build takes some array of inputs and constructs a tree
func Build(vals []float64) *Tree {
	t := &Tree{}
	var v float64
	v, vals = vals[0], vals[1:]

	t.Value = v

	for {
		if len(vals) == 0 {
			break
		}
		v, vals = vals[0], vals[1:]
		t.Insert(v)
	}

	return t
}

// Insert inserts a new node with value v into the tree
func (t *Tree) Insert(v float64) {

	if v <= t.Value {
		if t.Left != nil {
			t.Left.Insert(v)
		} else {
			t.Left = &Tree{}
			t.Left.Value = v
		}
	} else {
		if t.Right != nil {
			t.Right.Insert(v)
		} else {
			t.Right = &Tree{}
			t.Right.Value = v
		}
	}
}

// PreOrder traverses over the tree visiting and printing out nodes before branching
func PreOrder(t *Tree) {
	if t != nil {
		println(t.Value)
		PreOrder(t.Left)
		PreOrder(t.Right)
	}
}
