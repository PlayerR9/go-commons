package history

import gers "github.com/PlayerR9/go-errors"

// Pair is a pair of a subject and a history.
type Pair[E any] struct {
	// subject is the subject.
	subject Subjecter[E]

	// history is the history.
	history *History[E]
}

// IsNil implements the errors.Pointer interface.
func (p *Pair[E]) IsNil() bool {
	return p == nil
}

// NewPair creates a new pair from the given history and subject.
//
// If the given history is nil, a new one is created.
//
// Parameters:
//   - history: The history for the pair. If nil, a new one is created.
//   - subject: The subject for the pair. Must not be nil.
//
// Returns:
//   - *Pair[E]: The new pair.
//   - error: If the subject is nil.
func NewPair[E any](history *History[E], subject Subjecter[E]) (*Pair[E], error) {
	if subject == nil {
		err := gers.NewErrNilParameter("subject")
		err.AddFrame("NewPair()")

		return nil, err
	}

	if history == nil {
		history = NewHistory[E]()
	}

	return &Pair[E]{
		subject: subject,
		history: history,
	}, nil
}
