package main

import "time"

type FloorProgress struct {
	CurrentFloor        int
	MonstersKilledCount map[int]int
	FloorStartedAt      map[int]time.Time
	FloorCleared        map[int]bool
	ClearedFloors       int
	TotalClearDuration  time.Duration
}

func NewFloorProgress(startedAt time.Time) FloorProgress {
	return FloorProgress{
		CurrentFloor:        1,
		MonstersKilledCount: map[int]int{1: 0},
		FloorStartedAt:      map[int]time.Time{1: startedAt},
		FloorCleared:        make(map[int]bool),
	}
}

func (floor *FloorProgress) CanKillMonster(totalFloors int, monstersPerFloor int) bool {
	currentFloor := floor.CurrentFloor

	return currentFloor >= 1 &&
		currentFloor <= totalFloors &&
		!floor.FloorCleared[currentFloor] &&
		floor.MonstersKilledCount[currentFloor] < monstersPerFloor
}

func (floor *FloorProgress) KillMonster(monstersPerFloor int, eventTime time.Time) {
	currentFloor := floor.CurrentFloor

	floor.MonstersKilledCount[currentFloor]++

	if floor.MonstersKilledCount[currentFloor] == monstersPerFloor {
		floor.FloorCleared[currentFloor] = true
		floor.ClearedFloors++

		startedAt := floor.FloorStartedAt[currentFloor]
		floor.TotalClearDuration += eventTime.Sub(startedAt)
	}
}

func (floor *FloorProgress) CanMoveNext(totalFloors int) bool {
	currentFloor := floor.CurrentFloor

	return currentFloor >= 1 &&
		currentFloor <= totalFloors &&
		floor.FloorCleared[currentFloor]
}

func (floor *FloorProgress) MoveNext(eventTime time.Time, totalFloors int) {
	floor.CurrentFloor++

	if floor.CurrentFloor <= totalFloors {
		if _, exists := floor.FloorStartedAt[floor.CurrentFloor]; !exists {
			floor.MonstersKilledCount[floor.CurrentFloor] = 0
			floor.FloorStartedAt[floor.CurrentFloor] = eventTime
			floor.FloorCleared[floor.CurrentFloor] = false
		}
	}
}

func (floor *FloorProgress) CanMovePrevious() bool {
	return floor.CurrentFloor > 1
}

func (floor *FloorProgress) MovePrevious() {
	floor.CurrentFloor--
}

func (floor *FloorProgress) CanEnterBossFloor(totalFloors int) bool {
	return floor.ClearedFloors >= totalFloors &&
		floor.CurrentFloor == totalFloors+1
}

func (floor *FloorProgress) AverageClearDuration() time.Duration {
	if floor.ClearedFloors == 0 {
		return 0
	}
	return floor.TotalClearDuration / time.Duration(floor.ClearedFloors)
}
