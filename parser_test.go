package main

import (
	"testing"
)

func TestParseEventWithExtraParam(t *testing.T) {
	event, err := ParseEvent("[14:10:15] 42 9 out of strength")
	if err != nil {
		t.Fatalf("error parsing event: %v", err)
	}

	if event.TimeEventHappen.Format("15:04:05") != "14:10:15" {
		t.Fatalf("unexpected time: %s", event.TimeEventHappen.Format("15:04:05"))
	}

	if event.PlayerID != 42 {
		t.Fatalf("unexpected player: %d", event.PlayerID)
	}

	if event.EventID != 9 {
		t.Fatalf("unexpected event: %d", event.EventID)
	}

	if event.ExtraParam != "out of strength" {
		t.Fatalf("unexpected extraParam: %s", event.ExtraParam)
	}
}

func TestParseEventInvalidFormat(t *testing.T) {
	_, err := ParseEvent("[14:10:15] 42")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
