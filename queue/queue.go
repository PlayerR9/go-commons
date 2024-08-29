package queue

// Queue represents a queue of elements.
type Queue[T any] struct {
	// elems is the queue of elements
	elems []T
}

// NewQueue creates a new empty queue.
//
// Returns:
//   - *Queue[T]: The created queue. Never returns nil.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		elems: make([]T, 0),
	}
}

// NewQueueWithElems creates a new queue with the given elements.
//
// Parameters:
//   - elems: The elements to add to the queue.
//
// Returns:
//   - *Queue[T]: The created queue. Never returns nil.
func NewQueueWithElems[T any](elems []T) *Queue[T] {
	return &Queue[T]{
		elems: elems,
	}
}

// Peek returns the first element in the queue.
//
// Returns:
//   - T: The first element in the queue.
//   - bool: True if the queue is not empty, false otherwise.
func (q Queue[T]) Peek() (T, bool) {
	if len(q.elems) == 0 {
		return *new(T), false
	}

	return q.elems[0], true
}

// Enqueue adds an element to the queue.
//
// Parameters:
//   - elem: The element to add to the queue.
func (q *Queue[T]) Enqueue(elem T) {
	q.elems = append(q.elems, elem)
}

// Dequeue removes and returns the first element in the queue.
//
// Returns:
//   - T: The first element in the queue.
//   - bool: True if the queue is not empty, false otherwise.
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.elems) == 0 {
		return *new(T), false
	}

	elem := q.elems[0]
	q.elems = q.elems[1:]

	return elem, true
}
