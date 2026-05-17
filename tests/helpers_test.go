package tests

import (
	"impulse/internal/config"
	"impulse/internal/domain"
	"impulse/internal/engine"
	"reflect"
	"strings"
	"testing"
	"time"
)

func newTestGame(t *testing.T) *engine.Game {
	t.Helper()
	cfg := config.Config{
		Floors:   2,
		Monsters: 2,
		OpenAt:   "14:00:00",
		Duration: 2,
	}

	game, err := engine.NewGame(cfg)
	if err != nil {
		t.Fatalf("error while creating game: %v", err)
	}
	return game
}

func runEvents(game *engine.Game, events []domain.Event) {
	for _, event := range events {
		game.CloseExpiredDungeons(event.TimeEventHappen)
		game.ProcessEvent(event)
	}
	game.FinishOpenRuns()
}

func event(t *testing.T, timeText string, playerId int, eventId int, extraParam ...string) domain.Event {
	t.Helper()
	parsedTime, err := time.Parse("15:04:05", timeText)
	if err != nil {
		t.Fatalf("error while parsing time %s: %v", timeText, err)
	}

	return domain.Event{
		TimeEventHappen: parsedTime,
		PlayerID:        playerId,
		EventID:         eventId,
		ExtraParam:      strings.Join(extraParam, ","),
	}
}

func assertEqual[T any](t *testing.T, actual T, expected T) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}
