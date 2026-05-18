package main

import "time"

const (
	Register = iota + 1
	EnterDungeon
	KillMonster
	NextFloor
	PreviousFloor
	EnterBossFloor
	KillBoss
	LeaveDungeon
	CannotContinue
	RestoreHealth
	ReceiveDamage
)

type Event struct {
	TimeEventHappen time.Time
	PlayerID        int
	EventID         int
	ExtraParam      string
}
