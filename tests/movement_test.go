package tests

import (
	"impulse/internal/domain"
	"testing"
)

func TestMovingBackDoesNotResetFloorProgress(t *testing.T) {
	game := newTestGame(t)
	events := []domain.Event{
		event(t, "14:00:00", 1, domain.Register),
		event(t, "14:00:00", 1, domain.EnterDungeon),
		event(t, "14:02:00", 1, domain.KillMonster),
		event(t, "14:04:00", 1, domain.KillMonster),
		event(t, "14:05:00", 1, domain.NextFloor),
		event(t, "14:07:00", 1, domain.KillMonster),
		event(t, "14:08:00", 1, domain.PreviousFloor),
		event(t, "14:09:00", 1, domain.NextFloor),
		event(t, "14:10:00", 1, domain.KillMonster),
		event(t, "14:11:00", 1, domain.NextFloor),
		event(t, "14:12:00", 1, domain.EnterBossFloor),
		event(t, "14:14:00", 1, domain.KillBoss),
		event(t, "14:15:00", 1, domain.LeaveDungeon),
	}

	runEvents(game, events)
	expectedReport := []string{
		"[SUCCESS] 1 [00:15:00, 00:04:30, 00:02:00] HP:100",
	}

	assertEqual(t, game.CreateReport(), expectedReport)
}
