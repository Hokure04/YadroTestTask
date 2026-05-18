package main

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func newTestGame(t *testing.T) *Game {
	t.Helper()
	cfg := Config{
		Floors:   2,
		Monsters: 2,
		OpenAt:   "14:00:00",
		Duration: 2,
	}

	game, err := NewGame(cfg)
	if err != nil {
		t.Fatalf("error while creating game: %v", err)
	}
	return game
}

func runEvents(game *Game, events []Event) {
	for _, event := range events {
		game.CloseExpiredDungeons(event.TimeEventHappen)
		game.ProcessEvent(event)
	}
	game.FinishOpenRuns()
}

func event(t *testing.T, timeText string, playerId int, eventId int, extraParam ...string) Event {
	t.Helper()
	parsedTime, err := time.Parse("15:04:05", timeText)
	if err != nil {
		t.Fatalf("error while parsing time %s: %v", timeText, err)
	}

	return Event{
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
