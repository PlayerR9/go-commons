package team

// FilterFn is a filter function.
//
// Parameters:
//   - league: The league to filter.
//
// Returns:
//   - bool: True if the league should be kept, false otherwise.
type FilterFn[T interface {
	Equals(other T) bool
}] func(league League[T]) bool

// WithNTeams returns a filter function that keeps only the specified number of teams.
//
// Parameters:
//   - n: The number of teams to keep.
//
// Returns:
//   - FilterFn: The filter function.
func WithNTeams[T interface {
	Equals(other T) bool
}](n int) FilterFn[T] {
	if n < 0 {
		return nil
	}

	return func(league League[T]) bool {
		return len(league) == n
	}
}

// FilterSolutions filters solutions based on a filter function.
//
// Parameters:
//   - res: The solutions to filter.
//   - filter: The filter function.
//
// Returns:
//   - []League: The filtered solutions.
func FilterSolutions[T interface {
	Equals(other T) bool
}](res []League[T], filter FilterFn[T]) []League[T] {
	if len(res) == 0 || filter == nil {
		return nil
	}

	var top int

	for i := 0; i < len(res); i++ {
		league := res[i]

		ok := filter(league)
		if ok {
			res[top] = league
			top++
		}
	}

	return res[:top:top]
}
