package engine

import (
	"fmt"
	"impulse/internal/domain"
	"time"
)

func (g *Game) FinalReport() []string {
	playersId := make([]int, 0, len(g.players))
	for player := range g.players {
		playersId = append(playersId, player)
	}

	report := make([]string, 0, len(playersId))

	for _, id := range playersId {
		player := g.players[id]
		state := domain.Fail
		totalTime := time.Duration(0)
		averageTime := time.Duration(0)
		bossKillTime := time.Duration(0)

		if player.Run != nil {
			state = player.Run.State

			if !player.Run.FinishedAt.IsZero() {
				totalTime = player.Run.FinishedAt.Sub(player.Run.StartedAt)
			}

			if player.Run.Boss.Killed {
				bossKillTime = player.Run.Boss.KilledAt
			}
		}

		line := fmt.Sprintf(
			"[%s] %d [%s, %s, %s] HP:%d",
			state,
			player.ID,
			totalTime,
			averageTime,
			bossKillTime,
			player.HP,
		)

		report = append(report, line)
	}
	return report
}
