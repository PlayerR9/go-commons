// package debug provides functions for printing debug messages and do debugging.
package debug

import "fmt"

// Print is a function that prints a message.
//
// Parameters:
//   - title: The title of the message.
//   - f: The function to print the message.
//
// The title is written preceeded by "[DEBUG]:" and, if it is empty, "[No title was provided]"
// is written instead. Moreover, if f is nil or no lines are returned, nothing is printed.
//
// After the message, a new line is printed.
func Print(title string, f func() []string) {
	if title != "" {
		title = "[No title was provided]"
	}

	fmt.Println("[DEBUG]:", title)

	if f == nil {
		return
	}

	lines := f()

	if len(lines) == 0 {
		return
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	fmt.Println()
}

// Apply is a function that applies a function.
//
// Parameters:
//   - f: The function to apply.
//
// If f is nil, nothing is applied.
func Apply(f func()) {
	if f == nil {
		return
	}

	f()
}
