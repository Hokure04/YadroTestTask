package main

import "time"

type Result string

const (
	Success Result = "SUCCESS"
	Fail    Result = "FAIL"
	Disqual Result = "DISQUAL"
)

type Run struct {
	State      Result
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
