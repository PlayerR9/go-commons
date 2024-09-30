package team

import (
	"github.com/PlayerR9/go-errors/assert"
)

// _Active is an active team.
type _Active[T interface {
	Equals(other T) bool
}] struct {
	// global is the global teams maker.
	global *_Global[T]

	// league is the list of teams.
	league League[T]

	// pos is the current position in the list of teams.
	pos int

	// err is the error.
	err error
}

// ApplyEvent implements the Subjecter interface.
func (a *_Active[T]) ApplyEvent(event int) bool {
	assert.NotNil(a, "receiver")
	assert.NotNil(a.global, "a.global")
	assert.Cond(event >= -1 && event < len(a.league), "event out of range")

	member, ok := a.global.MemberAt(a.pos)
	assert.Ok(ok, "a.global.MemberAt(%d)", a.pos)

	if event == -1 {
		a.league = append(a.league, []T{member})
	} else {
		team := a.league[event]
		team = append(team, member)
		a.league[event] = team
	}

	a.pos++

	return a.pos == a.global.Size()
}

// HasError implements the Subjecter interface.
func (a _Active[T]) HasError() bool {
	return a.err != nil
}

// IsNil implements the Subjecter interface.
func (a *_Active[T]) IsNil() bool {
	return a == nil
}

// NextEvents implements the Subjecter interface.
func (a *_Active[T]) NextEvents() []int {
	assert.NotNil(a, "receiver")
	assert.NotNil(a.global, "a.global")

	member, ok := a.global.MemberAt(a.pos)
	if !ok {
		return nil
	}

	var indices []int

	for i, team := range a.league {
		ok, err := team.is_candidable(member, a.global.enemy_fn)
		if err != nil {
			a.err = err
		} else if ok {
			indices = append(indices, i)
		}
	}

	indices = append(indices, -1)

	return indices
}

// _Global is a teams maker.
type _Global[T interface {
	Equals(other T) bool
}] struct {
	// enemy_fn is the function to use to check for enemies.
	enemy_fn AreEnemyFunc[T]

	// members is the list of members.
	members []T
}

// MemberAt returns the member at the given index.
//
// Parameters:
//   - idx: the index of the member to return.
//
// Returns:
//   - T: the member at the given index.
//   - bool: true if the index is valid, false otherwise.
func (tm _Global[T]) MemberAt(idx int) (T, bool) {
	if idx < 0 || idx >= len(tm.members) {
		return *new(T), false
	}

	return tm.members[idx], true
}

// Size returns the number of members.
//
// Returns:
//   - int: the number of members. Never returns less than 0.
func (tm _Global[T]) Size() int {
	return len(tm.members)
}
