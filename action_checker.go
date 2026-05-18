package main

import (
	"strconv"
)

func (g *Game) canEnterDungeon(event Event, player *Player) bool {
	return !player.IsInDungeon() &&
		!event.TimeEventHappen.Before(g.openAt) &&
		event.TimeEventHappen.Before(g.closeAt)
}

func (g *Game) canKillMonster(player *Player) bool {
	if !g.canAct(player) || player.Run.Boss.Entered {
		return false
	}

	return player.Run.Floor.CanKillMonster(g.config.Floors, g.config.Monsters)
}

func (g *Game) canMoveNextFloor(player *Player) bool {
	if !g.canAct(player) || player.Run.Boss.Entered {
		return false
	}

	return player.Run.Floor.CanMoveNext(g.config.Floors)
}

func (g *Game) canMovePreviousFloor(player *Player) bool {
	if !g.canAct(player) || player.Run.Boss.Entered {
		return false
	}
	return player.Run.Floor.CanMovePrevious()
}

func (g *Game) canEnterBossFloor(player *Player) bool {
	if !g.canAct(player) || !player.Run.Boss.CanEnter() {
		return false
	}
	return player.Run.Floor.CanEnterBossFloor(g.config.Floors)
}

func (g *Game) canKillBoss(player *Player) bool {
	return g.canAct(player) && player.Run.Boss.CanKill()
}

func (g *Game) canAct(player *Player) bool {
	return player.IsInDungeon() && player.IsAlive()
}

func (g *Game) reject(event Event, invalid bool) bool {
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
