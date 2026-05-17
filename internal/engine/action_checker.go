package engine

import (
	"impulse/internal/domain"
	"strconv"
)

func (g *Game) canEnterDungeon(event domain.Event, player *domain.Player) bool {
	return !player.IsInDungeon() &&
		!event.TimeEventHappen.Before(g.openAt) &&
		event.TimeEventHappen.Before(g.closeAt)
}

func (g *Game) canKillMonster(player *domain.Player) bool {
	if !player.IsInDungeon() || !player.IsAlive() || player.Run.Boss.Entered {
		return false
	}

	currentFloor := player.Run.Floor.CurrentFloor

	return currentFloor >= 1 &&
		currentFloor <= g.config.Floors &&
		!player.Run.Floor.FloorCleared[currentFloor] &&
		player.Run.Floor.MonstersKilledCount[currentFloor] < g.config.Monsters
}

func (g *Game) canMoveNextFloor(player *domain.Player) bool {
	if !player.IsInDungeon() || !player.IsAlive() || player.Run.Boss.Entered {
		return false
	}

	currentFloor := player.Run.Floor.CurrentFloor

	return currentFloor >= 1 &&
		currentFloor <= g.config.Floors &&
		player.Run.Floor.FloorCleared[currentFloor]
}

func (g *Game) canMovePreviousFloor(player *domain.Player) bool {
	if !player.IsInDungeon() || !player.IsAlive() || player.Run.Boss.Entered {
		return false
	}
	return player.Run.Floor.CurrentFloor > 1
}

func (g *Game) canEnterBossFloor(player *domain.Player) bool {
	if !player.IsInDungeon() || !player.IsAlive() || player.Run.Boss.Entered {
		return false
	}

	return player.Run.Floor.ClearedFloors >= g.config.Floors &&
		player.Run.Floor.CurrentFloor == g.config.Floors+1
}

func (g *Game) canKillBoss(player *domain.Player) bool {
	return player.IsInDungeon() &&
		player.IsAlive() &&
		player.Run.Boss.Entered &&
		!player.Run.Boss.Killed
}

func (g *Game) canLeaveDungeon(player *domain.Player) bool {
	return player.IsInDungeon() && player.IsAlive()
}

func (g *Game) canAct(player *domain.Player) bool {
	return player.IsInDungeon() && player.IsAlive()
}

func (g *Game) reject(event domain.Event, invalid bool) bool {
	if invalid {
		g.impossible(event)
		return true
	}

	return false
}

func parseNonNegativeInt(value string) (int, bool) {
	number, err := strconv.Atoi(value)
	if err != nil || number < 0 {
		return 0, false
	}
	return number, true
}
