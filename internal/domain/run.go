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
		State:     Success,
		StartedAt: startedAt,
		Floor: FloorProgress{
			CurrentFloor: 1,
		},
	}
}

func (r *Run) Finish(state string, finishedAt time.Time) {
	r.State = state
	r.Finished = true
	r.FinishedAt = finishedAt
}
