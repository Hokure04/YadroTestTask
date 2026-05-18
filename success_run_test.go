package main

import (
	"testing"
)

func TestSuccessfulRun(t *testing.T) {
	game := newTestGame(t)

	events := []Event{
		event(t, "14:00:00", 1, Register),
		event(t, "14:10:00", 1, EnterDungeon),
		event(t, "14:12:00", 1, KillMonster),
		event(t, "14:15:00", 1, KillMonster),
		event(t, "14:16:00", 1, NextFloor),
		event(t, "14:18:00", 1, KillMonster),
		event(t, "14:20:00", 1, KillMonster),
		event(t, "14:21:00", 1, NextFloor),
		event(t, "14:22:00", 1, EnterBossFloor),
		event(t, "14:30:00", 1, KillBoss),
		event(t, "14:34:00", 1, LeaveDungeon),
	}

	runEvents(game, events)
	expectedOutput := []string{
		"[14:00:00] Player [1] registered",
		"[14:10:00] Player [1] entered the dungeon",
		"[14:12:00] Player [1] killed the monster",
		"[14:15:00] Player [1] killed the monster",
		"[14:16:00] Player [1] went to the next floor",
		"[14:18:00] Player [1] killed the monster",
		"[14:20:00] Player [1] killed the monster",
		"[14:21:00] Player [1] went to the next floor",
		"[14:22:00] Player [1] entered the boss's floor",
		"[14:30:00] Player [1] killed the boss",
		"[14:34:00] Player [1] left the dungeon",
	}

	assertEqual(t, game.OutputLines(), expectedOutput)
	expectedReport := []string{
		"[SUCCESS] 1 [00:24:00, 00:04:30, 00:08:00] HP:100",
	}

	assertEqual(t, game.CreateReport(), expectedReport)
}
