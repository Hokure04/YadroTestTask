package domain

import "time"

type FloorProgress struct {
	CurrentFloor        int
	MonstersKilledCount map[int]int
	FloorStartedAt      map[int]time.Time
	FloorCleared        map[int]bool
	ClearedFloors       int
	ClearDuration       []time.Duration
}
