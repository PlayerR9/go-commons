package tree

import "iter"

// Tree is the tree data structure.
type Tree[T interface {
	Child() iter.Seq[T]

	BackwardChild() iter.Seq[T]

	TreeNoder
}] struct {
	// root is the root node of the tree.
	root T
}

// GoString implements the fmt.GoStringer interface.
func (t Tree[T]) GoString() string {
	trav := PrintFn[T]()

	info, err := ApplyDFS(&t, trav)
	if err != nil {
		panic(err.Error())
	}

	return info.String()
}

// NewTree creates a new tree with the given root node.
//
// Parameters:
//   - root: The root node of the tree.
//
// Returns:
//   - *Tree[T]: The new tree. Never returns nil.
func NewTree[T interface {
	Child() iter.Seq[T]
	BackwardChild() iter.Seq[T]

	TreeNoder
}](root T) *Tree[T] {
	return &Tree[T]{
		root: root,
	}
}

// Root returns the root node of the tree.
//
// Returns:
//   - T: The root node of the tree.
func (t Tree[T]) Root() T {
	return t.root
}
