package helpers

// WeightedElement is a type that represents an element with a weight.
type WeightedElement[O any] struct {
	// elem is the element.
	elem O

	// weight is the weight of the element.
	weight float64
}

// IsNil checks if the WeightedElement is nil.
//
// Returns:
//   - bool: True if the WeightedElement is nil, false otherwise.
func (we *WeightedElement[O]) IsNil() bool {
	return we == nil
}

// Data implements the Helperer interface.
func (we WeightedElement[O]) Data() (O, error) {
	return we.elem, nil
}

// Weight returns the weight of the element.
//
// Returns:
//   - float64: The weight of the element.
func (we WeightedElement[O]) Weight() float64 {
	return we.weight
}

// NewWeightedElement creates a new WeightedElement with the given element and weight.
//
// Parameters:
//   - elem: The element.
//   - weight: The weight of the element.
//
// Returns:
//   - *WeightedElement: The new WeightedElement. Never returns nil.
func NewWeightedElement[O any](elem O, weight float64) *WeightedElement[O] {
	we := &WeightedElement[O]{
		elem:   elem,
		weight: weight,
	}

	return we
}
