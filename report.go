package main

import (
	"fmt"
	"sort"
	"time"
)

func (g *Game) CreateReport() []string {
	playersId := make([]int, 0, len(g.players))
	for player := range g.players {
		playersId = append(playersId, player)
	}

	sort.Ints(playersId)
	report := make([]string, 0, len(playersId))

	for _, id := range playersId {
		player := g.players[id]
		state := Fail
		totalTime := time.Duration(0)
		averageTime := time.Duration(0)
		bossKillTime := time.Duration(0)

		if player.Disqualified {
			state = Disqual
		}

		if player.Run != nil {
			state = player.Run.State

			if !player.Run.FinishedAt.IsZero() {
				totalTime = player.Run.FinishedAt.Sub(player.Run.StartedAt)
			}

			if player.Run.Floor.ClearedFloors > 0 {
				averageTime = player.Run.Floor.TotalClearDuration /
					time.Duration(player.Run.Floor.ClearedFloors)
			}

			if player.Run.Boss.Killed {
				bossKillTime = player.Run.Boss.KillDuration
			}
		}

		line := fmt.Sprintf(
			"[%s] %d [%s, %s, %s] HP:%d",
			state,
			player.ID,
			formatDuration(totalTime),
			formatDuration(averageTime),
			formatDuration(bossKillTime),
			player.HP,
		)

		report = append(report, line)
	}
	return report
}

func formatDuration(duration time.Duration) string {
	totalSeconds := int(duration.Seconds())

	hours := totalSeconds / 3600
	minutes := totalSeconds % 3600 / 60
	seconds := totalSeconds % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
