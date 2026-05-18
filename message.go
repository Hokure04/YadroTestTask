package main

import (
	"fmt"
	"time"
)

const clockLayout = time.TimeOnly

func (g *Game) event(eventTime time.Time, format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	line := fmt.Sprintf("[%s] %s", eventTime.Format(clockLayout), message)
	g.outputLines = append(g.outputLines, line)
}

func (g *Game) impossible(event Event) {
	g.event(event.TimeEventHappen, "Player [%d] makes imposible move [%d]", event.PlayerID, event.EventID)
}
