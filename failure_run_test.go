package main

import (
	"testing"
)

func TestUnregisteredPlayerIsDisqualified(t *testing.T) {
	game := newTestGame(t)

	events := []Event{
		event(t, "14:10:00", 3, KillMonster),
	}

	runEvents(game, events)
	expectedOutput := []string{
		"[14:10:00] Player [3] is disqualified",
	}

	assertEqual(t, game.OutputLines(), expectedOutput)
	expectedReport := []string{
		"[DISQUAL] 3 [00:00:00, 00:00:00, 00:00:00] HP:100",
	}
	assertEqual(t, game.CreateReport(), expectedReport)
}

func TestHealthRestoreAndDeath(t *testing.T) {
	game := newTestGame(t)
	events := []Event{
		event(t, "14:00:00", 1, Register),
		event(t, "14:10:00", 1, EnterDungeon),
		event(t, "14:12:00", 1, ReceiveDamage, "60"),
		event(t, "14:13:00", 1, RestoreHealth, "80"),
		event(t, "14:20:00", 1, ReceiveDamage, "100"),
	}

	runEvents(game, events)
	expectedOutput := []string{
		"[14:00:00] Player [1] registered",
		"[14:10:00] Player [1] entered the dungeon",
		"[14:12:00] Player [1] recieved [60] of damage",
		"[14:13:00] Player [1] has restored [80] of health",
		"[14:20:00] Player [1] recieved [100] of damage",
		"[14:20:00] Player [1] is dead",
	}

	assertEqual(t, game.OutputLines(), expectedOutput)

	expectedReport := []string{
		"[FAIL] 1 [00:10:00, 00:00:00, 00:00:00] HP:0",
	}

	assertEqual(t, game.CreateReport(), expectedReport)
}

func TestImpossibleToMove(t *testing.T) {
	game := newTestGame(t)
	events := []Event{
		event(t, "14:00:00", 1, Register),
		event(t, "14:10:00", 1, EnterDungeon),
		event(t, "14:11:00", 1, NextFloor),
	}

	runEvents(game, events)
	expectedOutput := []string{
		"[14:00:00] Player [1] registered",
		"[14:10:00] Player [1] entered the dungeon",
		"[14:11:00] Player [1] makes imposible move [4]",
	}
	assertEqual(t, game.OutputLines(), expectedOutput)
}
