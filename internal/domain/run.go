package domain

import "time"

const (
	Success = "SUCCESS"
	Fail    = "FAIL"
	Disqual = "DISQUAL"
)

type Run struct {
	State      string
	StartedAt  time.Time
	FinishedAt time.Time
	Finished   bool
	Floor      FloorProgress
	Boss       BossProgress
}

func NewRun(startedAt time.Time) *Run {
	return &Run{
		State:     Fail,
		StartedAt: startedAt,
		Floor: FloorProgress{
			CurrentFloor:        1,
			MonstersKilledCount: map[int]int{1: 0},
			FloorStartedAt:      map[int]time.Time{1: startedAt},
			FloorCleared:        make(map[int]bool),
		},
	}
}

func (r *Run) Finish(state string, finishedAt time.Time) {
	r.State = state
	r.Finished = true
	r.FinishedAt = finishedAt
}
