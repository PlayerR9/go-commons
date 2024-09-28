package helpers

// FilterByPositiveWeight is a function that iterates over weight results and
// returns the elements with the maximum weight.
//
// Parameters:
//   - slices: slice of weight results.
//
// Returns:
//   - []H: slice of elements with the maximum weight.
//
// Behaviors:
//   - If S is empty, the function returns a nil slice.
//   - If multiple elements have the same maximum weight, they are all returned.
//   - If S contains only one element, that element is returned.
func FilterByPositiveWeight[H Helperer[O], O any](slices []H) []H {
	if len(slices) == 0 {
		return nil
	}

	max_weight := slices[0].Weight()
	indices := []int{0}

	for i := 1; i < len(slices); i++ {
		elem := slices[i]

		weight := elem.Weight()

		if weight > max_weight {
			max_weight = weight
			indices = []int{i}
		} else if weight == max_weight {
			indices = append(indices, i)
		}
	}

	solution := make([]H, 0, len(indices))

	for _, index := range indices {
		solution = append(solution, slices[index])
	}

	return solution
}
