package history

import (
	"errors"
	"iter"

	gers "github.com/PlayerR9/go-errors"
	"github.com/PlayerR9/go-errors/assert"
)

// History is a stack of events that can be replayed.
type History[E any] struct {
	// timeline is the list of events that have been applied to the subject.
	timeline []E

	// arrow is the current position in the timeline.
	arrow int
}

// IsNil checks if the history is nil.
//
// Returns:
//   - bool: True if the history is nil, false otherwise.
func (h *History[E]) IsNil() bool {
	return h == nil
}

// NewHistory creates a new history.
//
// Returns:
//   - *History: The new history. Never returns nil.
func NewHistory[E any]() *History[E] {
	return &History[E]{
		timeline: make([]E, 0),
		arrow:    0,
	}
}

// Copy creates a copy of the history.
//
// Returns:
//   - *History: The copy of the history. Never returns nil.
func (h History[E]) Copy() *History[E] {
	new_timeline := make([]E, len(h.timeline))
	copy(new_timeline, h.timeline)

	return &History[E]{
		timeline: new_timeline,
		arrow:    h.arrow,
	}
}

// Restart resets the history without removing any events.
func (h *History[E]) Restart() {
	if h == nil {
		return
	}

	h.arrow = 0
}

// AddEvent adds an event to the history.
//
// Parameters:
//   - event: The event to add to the history.
func (h *History[E]) AddEvent(event E) {
	assert.NotNil(h, "h")

	h.timeline = append(h.timeline, event)
}

// Event returns a sequence of events in the history that have not been applied to the
// subject yet.
//
// Returns:
//   - iter.Seq[E]: The sequence of events.
func (h *History[E]) Event() iter.Seq[E] {
	if h == nil {
		return func(yield func(E) bool) {}
	}

	return func(yield func(E) bool) {
		for i := h.arrow; i < len(h.timeline); i++ {
			event := h.timeline[i]

			if !yield(event) {
				h.arrow = i
				break
			}
		}

		h.arrow = len(h.timeline)
	}
}

var (
	// SubjectHasError is returned when the subject has an error.
	SubjectHasError error

	// SubjectIsDone is returned when the subject is done.
	SubjectIsDone error

	// HistoryEnded is returned when the history has ended.
	HistoryEnded error
)

func init() {
	SubjectHasError = errors.New("subject has an error")
	SubjectIsDone = errors.New("subject is done")
	HistoryEnded = errors.New("history ended")
}

// AdvanceOne advances the history by one event and applies the event to the subject.
//
// Parameters:
//   - history: The history to advance.
//   - subject: The subject to apply the event to.
//
// Returns:
//   - bool: True if the subject is done, false otherwise.
//   - error: If the subject is nil, or if an error occurs while applying the event.
//
// Errors:
//   - SubjectHasError: If the subject is done due to an error.
//   - SubjectIsDone: If the subject is done due to internal success.
//   - HistoryEnded: If the history has ended.
func (h *History[E]) AdvanceOne(subject Subjecter[E]) (bool, error) {
	if h == nil {
		err := gers.NewErrNilReceiver()
		err.AddFrame("history.AdvanceOne()")

		return false, err
	} else if subject == nil {
		err := gers.NewErrNilParameter("subject")
		err.AddFrame("history.AdvanceOne()")

		return false, err
	}

	if h.arrow >= len(h.timeline) {
		return false, HistoryEnded
	}

	event := h.timeline[h.arrow]
	h.arrow++

	is_done := subject.ApplyEvent(event)
	if subject.HasError() {
		return true, SubjectHasError
	} else if is_done {
		return true, nil
	}

	return false, nil
}

// Align applies events from the history to the subject until the subject is
// at the end of the history.
//
// Parameters:
//   - history: The history to apply events from. Assumed not nil.
//   - subject: The subject to apply events to. Assumed not nil.
//
// Returns:
//   - error: If the subject is nil, or if an error occurs while applying events.
//
// Errors:
//   - SubjectHasError: If the subject is done due to an error.
//   - SubjectIsDone: If the subject is done due to internal success.
func (h *History[E]) Align(subject Subjecter[E]) error {
	if h == nil {
		err := gers.NewErrNilReceiver()
		err.AddFrame("history.Align()")

		return err
	} else if subject == nil {
		err := gers.NewErrNilParameter("subject")
		err.AddFrame("history.Align()")

		return err
	}

	for h.arrow < len(h.timeline) {
		event := h.timeline[h.arrow]
		h.arrow++

		is_done := subject.ApplyEvent(event)
		if subject.HasError() {
			return SubjectHasError
		} else if is_done {
			return SubjectIsDone
		}
	}

	return nil
}

/*
// ApplyEvents applies events from the history to the subject until either the subject
// is done, or the specified number of events have been applied.
//
// Parameters:
//   - history: The history to apply events from. Assumed not nil.
//   - subject: The subject to apply events to. Assumed not nil.
//   - n: The number of events to apply. If negative, applies all events.
//
// Returns:
//   - bool: True if the subject is done, false otherwise.
//   - error: If the subject is nil, or if an error occurs while applying events.
func ApplyEvents[E any, S interface {
	Subjecter[E]
}](history *History[E], subject S, n int) (bool, error) {
	if history == nil || n == 0 {
		return false, nil
	} else if subject.IsNil() {
		return false, gers.NewErrNilParameter("subject")
	}

	if n < 0 {
		for event := range history.Event() {
			is_done := subject.ApplyEvent(event)
			if subject.HasError() {
				return true, nil
			} else if is_done {
				return true, IsDone
			}
		}
	} else {
		var count int

		for event := range history.Event() {
			is_done := subject.ApplyEvent(event)
			if subject.HasError() {
				return true, nil
			} else if is_done {
				return true, IsDone
			}

			count++

			if count == n {
				break
			}
		}

	}

	return false, nil
}
*/
