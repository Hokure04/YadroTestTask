package engine

import (
	"fmt"
	"impulse/internal/domain"
)

func (g *Game) event(format string, playerId int, eventId int) {
	message := fmt.Sprintf(format, playerId, eventId)
	line := fmt.Sprintf("[%s] %s", message)
	g.outputLines = append(g.outputLines, line)
}

func (g *Game) impossible(event domain.Event) {
	g.event("Player [%d] makes imposible move [%d]", event.PlayerID, event.EventID)
}
