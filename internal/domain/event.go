package domain

import "time"

const (
	Register       = 1
	EnterDungeon   = 2
	KillMonster    = 3
	NextFloor      = 4
	PreviousFloor  = 5
	EnterBossFloor = 6
	KillBoss       = 7
	LeaveDungeon   = 8
	CannotContinue = 9
	RestoreHealth  = 10
	ReceiveDamage  = 11
)

type Event struct {
	TimeEventHappen time.Time
	PlayerID        int
	EventID         int
}
