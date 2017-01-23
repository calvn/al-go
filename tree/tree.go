package tree

// BTree is a binary tree representation
type BTree struct {
	LTree *BTree
	RTree *BTree
	Value int
}

// New returns a new BTree. It populates the tree via
// level-order traversal.
func New(entries []int) *BTree {
	root := &BTree{}
	var next *BTree

	for i, val := range entries {
		if i == 0 {
			root.Value = val
			next = root
			continue
		}

		next = insert(next, val)
	}

	return root
}

func insert(t *BTree, val int) *BTree {
	// If left tree is empty, set val to left node, return same level tree
	if t.LTree == nil {
		t.LTree = &BTree{Value: val}
		return t
	}

	// If right node is empty, set val to right node, return left tree
	if t.RTree == nil {
		t.RTree = &BTree{Value: val}
		return t.LTree
	}

	// Otherwise return original tree
	return t
}
