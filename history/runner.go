package history

import (
	"iter"

	gers "github.com/PlayerR9/go-errors"
)

// execute_until executes events from the history until the subject is done, or
// until all events have been exhausted. If the subject is done due to an error,
// the function returns false.
//
// The function returns a slice of histories, where each history is a branching
// path of the original history. The returned slice may be empty if no branching
// paths were found.
//
// Parameters:
//   - history: The history to execute events from.
//   - subject: The subject to apply events to.
//
// Returns:
//   - []*History[E]: The branching paths of the history.
//   - bool: True if the subject is done due to exhaustion of the history, false
//     otherwise.
func execute_until[E any](history *History[E], subject Subjecter[E]) ([]*History[E], bool) {
	gers.AssertNotNil(history, "history")
	gers.AssertNotNil(subject, "subject")

	var paths []*History[E]

	for {
		nexts := subject.NextEvents()
		if subject.HasError() {
			return paths, false
		} else if len(nexts) == 0 {
			break
		}

		for _, next := range nexts[1:] {
			path := history.Copy()
			path.AddEvent(next)
			path.Restart()

			paths = append(paths, path)
		}

		history.AddEvent(nexts[0])
		is_done, err := history.AdvanceOne(subject)
		if err != nil {
			return paths, false
		} else if is_done {
			break
		}
	}

	return paths, true
}

// Execute returns a sequence of subjects that can be reached by applying events
// from the history to the subjects, where the history is updated by adding events
// to the history, and the subjects are updated by applying events to the subjects.
//
// The function takes a function that returns a subject as an argument. This
// function is called once for each subject in the sequence.
//
// The sequence is infinite, and will continue to generate subjects until the
// subject is done due to exhaustion of the history, or until the subject is
// done due to an error.
//
// The function will panic if an error occurs while executing the sequence.
func Execute[E any, S Subjecter[E]](init_fn func() S) iter.Seq[S] {
	if init_fn == nil {
		return func(yield func(S) bool) {}
	}

	fn := func(yield func(S) bool) {
		var invalid_subjects []S

		var queue []*Pair[E]

		subject := init_fn()
		if subject.IsNil() {
			invalid_subjects = append(invalid_subjects, subject)
			return
		}

		pair := gers.AssertNew(
			NewPair(nil, subject),
		)

		queue = append(queue, pair)

		for len(queue) > 0 {
			first := queue[0]
			queue = queue[1:]

			gers.AssertNotNil(first, "top")
			subject = gers.AssertConv[S](first.subject, "top.subject")

			err := first.history.Align(subject)
			if err == SubjectHasError {
				invalid_subjects = append(invalid_subjects, subject)
				continue
			} else if err == SubjectIsDone {
				if !yield(subject) {
					return
				} else {
					continue
				}
			} else if err != nil {
				panic(err)
			}

			paths, ok := execute_until(first.history, subject)

			if len(paths) > 0 {
				new_paths := make([]*Pair[E], 0, len(paths))

				for _, path := range paths {
					gers.AssertNotNil(path, "path")

					subject := init_fn()
					if subject.IsNil() {
						invalid_subjects = append(invalid_subjects, subject)
						continue
					}

					pair = gers.AssertNew(
						NewPair(path, subject),
					)

					new_paths = append(new_paths, pair)
				}

				queue = append(queue, new_paths...)
			}

			if !ok {
				invalid_subjects = append(invalid_subjects, subject)
			} else if !yield(subject) {
				return
			}
		}

		for _, subject := range invalid_subjects {
			if !yield(subject) {
				return
			}
		}
	}

	return fn
}
