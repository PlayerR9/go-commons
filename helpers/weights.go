package helpers

// WeightFunc is a type for a function that assigns a weight to an element.
//
// Parameters:
//   - elem: The element to assign a weight to. Assumed to be non-nil.
//
// Returns:
//   - float64: The weight of the element.
//   - bool: True if the weight is valid, otherwise false.
type WeightFunc[O any] func(elem O) (float64, bool)

// ApplyWeightFunc applies the weight function to each element in the slice.
// If the weight function returns false, the element is not included in the result.
//
// Parameters:
//   - slice: slice of elements.
//   - f: the weight function.
//
// Returns:
//   - []*WeightedElement[O]: slice of WeightedElement.
//
// If the slice contains nil elements, then said elements will be passed to the weight function.
func ApplyWeightFunc[O any](slice []O, f WeightFunc[O]) []*WeightedElement[O] {
	if len(slice) == 0 || f == nil {
		return nil
	}

	weights := make([]*WeightedElement[O], 0, len(slice))

	for i := 0; i < len(slice); i++ {
		elem := slice[i]

		weight, ok := f(elem)
		if ok {
			we := NewWeightedElement(elem, weight)
			weights = append(weights, we)
		}
	}

	return weights[:len(weights):len(weights)]
}
