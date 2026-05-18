package main

import (
	"testing"
)

func TestDungeonCloseByTime(t *testing.T) {
	game := newTestGame(t)

	events := []Event{
		event(t, "14:00:00", 1, Register),
		event(t, "14:10:00", 1, EnterDungeon),
		event(t, "16:10:00", 1, KillMonster),
	}

	runEvents(game, events)
	expectedOutput := []string{
		"[14:00:00] Player [1] registered",
		"[14:10:00] Player [1] entered the dungeon",
		"[16:10:00] Player [1] makes imposible move [3]",
	}

	assertEqual(t, game.OutputLines(), expectedOutput)

	expectedReport := []string{
		"[FAIL] 1 [01:50:00, 00:00:00, 00:00:00] HP:100",
	}

	assertEqual(t, game.CreateReport(), expectedReport)
}
