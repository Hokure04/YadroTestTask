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
		Floor:     NewFloorProgress(startedAt),
	}
}
