package tree

import (
	"iter"
	"slices"
)

// Traverser is the traverser that holds the traversal logic.
type Traverser[N interface {
	Child() iter.Seq[N]

	BackwardChild() iter.Seq[N]

	TreeNoder
}, I any] struct {
	// InitFn is the function that initializes the traversal info.
	//
	// Parameters:
	//   - root: The root node of the tree.
	//
	// Returns:
	//   - I: The initial traversal info.
	InitFn func(root N) I

	// DoFn is the function that performs the traversal logic.
	//
	// Parameters:
	//   - info: The traversal info.
	//
	// Returns:
	//   - []I: The next traversal info.
	//   - error: The error that might occur during the traversal.
	DoFn func(info I) ([]I, error)
}

// ApplyDFS applies the DFS traversal logic to the tree.
//
// Parameters:
//   - t: The tree to apply the traversal logic to.
//   - trav: The traverser that holds the traversal logic.
//
// Returns:
//   - error: The error that might occur during the traversal.
func ApplyDFS[N interface {
	Child() iter.Seq[N]
	BackwardChild() iter.Seq[N]

	TreeNoder
}, I any](t *Tree[N], trav Traverser[N, I]) (I, error) {
	if t == nil {
		return *new(I), nil
	}

	info := trav.InitFn(t.root)

	stack := []I{info}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		nexts, err := trav.DoFn(top)
		if err != nil {
			return info, err
		}

		if len(nexts) > 0 {
			slices.Reverse(nexts)
			stack = append(stack, nexts...)
		}
	}

	return info, nil
}
