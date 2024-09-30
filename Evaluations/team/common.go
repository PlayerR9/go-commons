package team

import (
	hst "github.com/PlayerR9/go-commons/Evaluations/history"
	gers "github.com/PlayerR9/go-errors"
)

// EvaluateTeams evaluates the teams.
//
// Parameters:
//   - members: the list of members to evaluate.
//   - enemy_fn: the function to use to check for enemies.
//
// Returns:
//   - []League[T]: the list of teams. Never returns nil.
//   - error: an error if the evaluation fails.
func EvaluateTeams[T interface {
	Equals(other T) bool
}](members []T, enemy_fn AreEnemyFunc[T]) ([]League[T], error) {
	if enemy_fn == nil {
		return nil, gers.NewErrNilParameter("enemy_fn")
	}

	tm := &_Global[T]{
		enemy_fn: enemy_fn,
		members:  members,
	}

	init_fn := func() *_Active[T] {
		return &_Active[T]{
			global: tm,
			league: nil,
			pos:    0,
		}
	}

	var success []*_Active[T]

	for res := range hst.Execute(init_fn) {
		if res.HasError() {
			break
		}

		success = append(success, res)
	}

	if len(success) == 0 {
		return nil, gers.New(gers.OperationFail, "no teams found")
	}

	var leagues []League[T]

	for _, s := range success {
		leagues = append(leagues, s.league)
	}

	return leagues, nil
}
