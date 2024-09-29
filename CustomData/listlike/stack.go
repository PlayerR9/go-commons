package listlike

// Queue is a queue of events that can be replayed.
type Queue[T any] struct {
	// elems is the queue of events.
	elems []T
}

// IsNil checks if the queue is nil.
//
// Returns:
//   - bool: True if the queue is nil, false otherwise.
func (q *Queue[T]) IsNil() bool {
	return q == nil
}

// NewQueue creates a new queue.
//
// Returns:
//   - *Queue[T]: The new queue. Never returns nil.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		elems: make([]T, 0),
	}
}

// Enqueue adds an element to the queue. Doesn't add nil elements or when the
// receiver is nil.
//
// Parameters:
//   - elem: The element to add to the queue.
//
// Returns:
//   - bool: True if the element was added to the queue, false otherwise.
func (q *Queue[T]) Enqueue(elem T) bool {
	if q == nil {
		return false
	}

	q.elems = append(q.elems, elem)

	return true
}

// EnqueueMany adds multiple elements to the queue. Doesn't add nil elements or
// when the receiver is nil.
//
// Parameters:
//   - elems: The elements to add to the queue.
//
// Returns:
//   - bool: True if the element was added to the queue, false otherwise.
func (q *Queue[T]) EnqueueMany(elems []T) bool {
	if q == nil {
		return false
	}

	q.elems = append(q.elems, elems...)

	return true
}

// Dequeue removes an element from the queue. Doesn't remove nil elements or when
// the receiver is nil.
//
// Returns:
//   - T: The removed element.
//   - bool: True if the element was removed from the queue, false otherwise.
func (q *Queue[T]) Dequeue() (T, bool) {
	if q == nil || len(q.elems) == 0 {
		return *new(T), false
	}

	elem := q.elems[0]
	q.elems = q.elems[1:]

	return elem, true
}
