package domain

import "time"

type FloorProgress struct {
	CurrentFloor        int
	MonstersKilledCount int
	FloorStartedAt      time.Time
	FloorCleared        bool
	ClearedFloors       int
}
