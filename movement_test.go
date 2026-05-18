package main

import (
	"testing"
)

func TestMovingBackDoesNotResetFloorProgress(t *testing.T) {
	game := newTestGame(t)
	events := []Event{
		event(t, "14:00:00", 1, Register),
		event(t, "14:00:00", 1, EnterDungeon),
		event(t, "14:02:00", 1, KillMonster),
		event(t, "14:04:00", 1, KillMonster),
		event(t, "14:05:00", 1, NextFloor),
		event(t, "14:07:00", 1, KillMonster),
		event(t, "14:08:00", 1, PreviousFloor),
		event(t, "14:09:00", 1, NextFloor),
		event(t, "14:10:00", 1, KillMonster),
		event(t, "14:11:00", 1, NextFloor),
		event(t, "14:12:00", 1, EnterBossFloor),
		event(t, "14:14:00", 1, KillBoss),
		event(t, "14:15:00", 1, LeaveDungeon),
	}

	runEvents(game, events)
	expectedReport := []string{
		"[SUCCESS] 1 [00:15:00, 00:04:30, 00:02:00] HP:100",
	}

	assertEqual(t, game.CreateReport(), expectedReport)
}
