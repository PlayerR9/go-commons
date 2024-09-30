package team

import (
	"github.com/PlayerR9/go-errors/assert"
)

// AreEnemyFunc is a function that checks if two members are enemies.
//
// Parameters:
//   - m1: The first member.
//   - m2: The second member.
//
// Returns:
//   - bool: True if the members are enemies, false otherwise.
//   - error: An error if the members are not valid.
type AreEnemyFunc[T interface {
	Equals(other T) bool
}] func(m1, m2 T) (bool, error)

// Team is a team of members.
type Team[T interface {
	Equals(other T) bool
}] []T

// Equals checks whether two teams are equal.
//
// Two teams are equal if they have the same members using a loose equality.
//
// Parameters:
//   - other: The other team to compare to.
//
// Returns:
//   - bool: True if the teams are equal, false otherwise.
func (t Team[T]) Equals(other Team[T]) bool {
	if other == nil {
		return false
	}

	if len(t) != len(other) {
		return false
	}

	for _, m := range t {
		if !other.HasMember(m) {
			return false
		}
	}

	return true
}

// HasMember checks whether the team has a member.
//
// Parameters:
//   - member: the member to check.
//
// Returns:
//   - bool: true if the team has the member, false otherwise.
func (t Team[T]) HasMember(member T) bool {
	if len(t) == 0 {
		return false
	}

	for _, m := range t {
		if m.Equals(member) {
			return true
		}
	}

	return false
}

// is_candidable checks whether a pair is candidable.
//
// Parameters:
//   - pair: The pair to check.
//   - enemy_fn: The function to use to check for enemies.
//
// Returns:
//   - bool: True if the pair is candidable, false otherwise.
//   - error: An error if the pair is not valid.
func (team Team[T]) is_candidable(pair T, enemy_fn AreEnemyFunc[T]) (bool, error) {
	assert.NotNil(enemy_fn, "enemy_fn")
	assert.NotNil(team, "team")

	for _, member := range team {
		ok, err := enemy_fn(member, pair)
		if err != nil {
			return false, err
		}

		if ok {
			return false, nil
		}
	}

	return true, nil
}

// League is a collection of teams.
type League[T interface {
	Equals(other T) bool
}] []Team[T]

// Equals checks whether two leagues are equal.
//
// Two leagues are equal if they contain the same teams regardless of order.
//
// Parameters:
//   - other: The other league to compare to.
//
// Returns:
//   - bool: True if the leagues are equal, false otherwise.
func (l League[T]) Equals(other League[T]) bool {
	if len(l) != len(other) {
		return false
	}

	for _, m1 := range l {
		if !other.ContainsTeam(m1) {
			return false
		}
	}

	return true
}

// ContainsTeam checks whether the league contains a team.
//
// Parameters:
//   - team: The team to check.
//
// Returns:
//   - bool: True if the league contains the team, false otherwise.
func (l League[T]) ContainsTeam(team Team[T]) bool {
	if team == nil || len(l) == 0 {
		return false
	}

	for _, t := range l {
		if t.Equals(team) {
			return true
		}
	}

	return false
}
