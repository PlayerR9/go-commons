package tree

import (
	"iter"
	"strings"
)

// TreeStringTraverser is the tree stringer.
type TreeStringTraverser[N interface {
	Child() iter.Seq[N]

	BackwardChild() iter.Seq[N]

	TreeNoder
}] struct {
	// lines is the lines of the tree stringer.
	lines []string

	// seen is the seen map of the tree stringer.
	seen map[N]bool
}

// String implements the fmt.Stringer interface.
func (tst TreeStringTraverser[N]) String() string {
	return strings.Join(tst.lines, "\n")
}

// IsSeen is a helper function that checks if the node is seen.
//
// Parameters:
//   - node: The node to check.
//
// Returns:
//   - bool: The result of the check.
func (tst TreeStringTraverser[N]) IsSeen(node N) bool {
	prev, ok := tst.seen[node]
	return ok && prev
}

// AppendLine is a helper function that appends a line to the tree stringer.
//
// Parameters:
//   - line: The line to append.
func (tst *TreeStringTraverser[N]) AppendLine(line string) {
	tst.lines = append(tst.lines, line)
}

// SetSeen is a helper function that sets the seen flag.
//
// Parameters:
//   - node: The node to set.
func (tst *TreeStringTraverser[N]) SetSeen(node N) {
	tst.seen[node] = true
}

// TreeStackElem is the stack element of the tree stringer.
type TreeStackElem[N interface {
	Child() iter.Seq[N]

	BackwardChild() iter.Seq[N]

	TreeNoder
}] struct {
	// global contains the global info of the tree stringer.
	global *TreeStringTraverser[N]

	// indent is the indentation string.
	indent string

	// node is the node of the tree.
	node N

	// is_last is the flag that indicates whether the node is the last node in the level.
	is_last bool

	// same_level is the flag that indicates whether the node is in the same level.
	same_level bool
}

// String implements the fmt.Stringer interface.
func (tse TreeStackElem[N]) String() string {
	return tse.global.String()
}

// set_is_last is a helper function that sets the is_last flag.
func (tse *TreeStackElem[N]) set_is_last() {
	tse.is_last = true
}

// set_same_level is a helper function that sets the same_level flag.
func (tse *TreeStackElem[N]) set_same_level() {
	tse.same_level = true
}

// PrintFn returns the print function of the tree stringer.
//
// Parameters:
//   - root: The root node of the tree.
//
// Returns:
//   - Traverser[N, *TreeStackElem[N]]: The print function of the tree stringer.
func PrintFn[N interface {
	Child() iter.Seq[N]

	BackwardChild() iter.Seq[N]

	TreeNoder
}]() Traverser[N, *TreeStackElem[N]] {
	init_fn := func(root N) *TreeStackElem[N] {
		return &TreeStackElem[N]{
			global: &TreeStringTraverser[N]{
				lines: make([]string, 0),
				seen:  make(map[N]bool),
			},
			indent:     "",
			node:       root,
			is_last:    true,
			same_level: false,
		}
	}

	fn := func(info *TreeStackElem[N]) ([]*TreeStackElem[N], error) {
		var builder strings.Builder

		if info.indent != "" {
			builder.WriteString(info.indent)

			if !info.node.IsLeaf() || info.is_last {
				builder.WriteString("└── ")
			} else {
				builder.WriteString("├── ")
			}
		}

		// Prevent cycles.
		ok := info.global.IsSeen(info.node)
		if ok {
			builder.WriteString("... WARNING: Cycle detected!")

			info.global.AppendLine(builder.String())

			return nil, nil
		}

		builder.WriteString(info.node.GoString())
		info.global.AppendLine(builder.String())

		info.global.SetSeen(info.node)

		if info.node.IsLeaf() {
			return nil, nil
		}

		var indent strings.Builder

		indent.WriteString(info.indent)

		if info.same_level && !info.is_last {
			indent.WriteString("│   ")
		} else {
			indent.WriteString("    ")
		}

		var elems []*TreeStackElem[N]

		for c := range info.node.Child() {
			se := &TreeStackElem[N]{
				global:     info.global,
				indent:     indent.String(),
				node:       c,
				is_last:    false,
				same_level: false,
			}

			elems = append(elems, se)
		}

		if len(elems) >= 2 {
			for i := 0; i < len(elems); i++ {
				elems[i].set_same_level()
			}
		}

		elems[len(elems)-1].set_is_last()

		return elems, nil
	}

	return Traverser[N, *TreeStackElem[N]]{
		InitFn: init_fn,
		DoFn:   fn,
	}
}
