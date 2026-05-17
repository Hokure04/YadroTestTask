package engine

import (
	"impulse/internal/domain"
	"strconv"
)

func (g *Game) processRegister(event domain.Event, player *domain.Player) {
	if player.Registered || player.Run != nil {
		g.impossible(event)
		return
	}
	player.Registered = true
	g.event(event.TimeEventHappen, "Player [%d] registered", player.ID)
}

func (g *Game) processEnterDungeon(event domain.Event, player *domain.Player) {
	if player.IsInDungeon() {
		g.impossible(event)
		return
	}

	if event.TimeEventHappen.Before(g.openAt) || event.TimeEventHappen.After(g.closeAt) {
		g.impossible(event)
		return
	}

	player.StartRun(event.TimeEventHappen)
	g.event(event.TimeEventHappen, "Player [%d] entered dungeon", player.ID)
}

func (g *Game) processKillMonster(event domain.Event, player *domain.Player) {
	if !player.IsInDungeon() || !player.IsAlive() || g.isBossFloor(player) {
		g.impossible(event)
		return
	}

	floor := &player.Run.Floor
	currentFloor := floor.CurrentFloor

	floor.MonstersKilledCount[currentFloor]++
	if floor.MonstersKilledCount[currentFloor] >= g.config.Monsters {
		floor.FloorCleared[currentFloor] = true
		floor.ClearedFloors++

		startedAt := floor.FloorStartedAt[currentFloor]
		floor.ClearDuration = append(floor.ClearDuration, event.TimeEventHappen.Sub(startedAt))
	}

	g.event(event.TimeEventHappen, "Player [%d] killed the monster", player.ID)
}

func (g *Game) processNextFloor(event domain.Event, player *domain.Player) {
	if !player.IsInDungeon() || !player.IsAlive() {
		g.impossible(event)
		return
	}

	floor := &player.Run.Floor
	currentFloor := floor.CurrentFloor

	if currentFloor >= g.config.Floors {
		g.impossible(event)
		return
	}

	if currentFloor <= g.monsterFloorCount() && !floor.FloorCleared[currentFloor] {
		g.impossible(event)
		return
	}

	floor.ClearedFloors++

	if _, ok := floor.FloorStartedAt[floor.CurrentFloor]; ok {
		floor.FloorStartedAt[floor.CurrentFloor] = event.TimeEventHappen
	}

	g.event(event.TimeEventHappen, "Player [%d] went to the next floor", player.ID)
}

func (g *Game) processPreviousFloor(event domain.Event, player *domain.Player) {
	if !player.IsInDungeon() || !player.IsAlive() {
		g.impossible(event)
		return
	}

	floor := &player.Run.Floor

	if floor.CurrentFloor <= 1 {
		g.impossible(event)
		return
	}

	floor.ClearedFloors--

	g.event(event.TimeEventHappen, "Player [%d] went to the previous floor", player.ID)
}

func (g *Game) processEnterBossFloor(event domain.Event, player *domain.Player) {
	if !player.IsInDungeon() || !player.IsAlive() {
		g.impossible(event)
		return
	}

	if player.Run.Boss.Entered {
		g.impossible(event)
		return
	}

	player.Run.Boss.Entered = true
	player.Run.Boss.EnteredAt = event.TimeEventHappen

	g.event(event.TimeEventHappen, "Player [%d] entered boss's floor", player.ID)
}

func (g *Game) processKillBoss(event domain.Event, player *domain.Player) {
	if !player.IsInDungeon() || !player.IsAlive() {
		g.impossible(event)
		return
	}

	if !g.isBossFloor(player) || !player.Run.Boss.Entered || player.Run.Boss.Killed {
		g.impossible(event)
		return

	}

	player.Run.Boss.Killed = true
	player.Run.Boss.KilledAt = event.TimeEventHappen.Sub(player.Run.Boss.EnteredAt)

	g.event(event.TimeEventHappen, "Player [%d] killed boss", player.ID)
}

func (g *Game) processLeaveDungeon(event domain.Event, player *domain.Player) {
	if !player.IsInDungeon() {
		g.impossible(event)
		return
	}

	state := domain.Fail
	if g.isDungeonCompleted(player) {
		state = domain.Success
	}

	player.FinishRun(state, event.TimeEventHappen)

	g.event(event.TimeEventHappen, "Player [%d] left the dungeon", player.ID)
}

func (g *Game) processCannotContinue(event domain.Event, player *domain.Player) {
	if !player.IsInDungeon() {
		g.impossible(event)
		return
	}

	player.FinishRun(domain.Disqual, event.TimeEventHappen)
	player.Disqualified = true

	g.event(event.TimeEventHappen, "Player [%d] cannot continue due to [%s]", player.ID, event.ExtraParam)
}

func (g *Game) processRestoreHealth(event domain.Event, player *domain.Player) {
	if !player.IsInDungeon() || player.IsAlive() {
		g.impossible(event)
		return
	}

	health, err := strconv.Atoi(event.ExtraParam)
	if err != nil {
		g.impossible(event)
		return
	}

	player.RestoreHealth(health)
	g.event(event.TimeEventHappen, "Player [%d] restored health", player.ID, health)
}

func (g *Game) processReceiveDamage(event domain.Event, player *domain.Player) {
	if !player.IsInDungeon() || !player.IsAlive() {
		g.impossible(event)
		return
	}

	damage, err := strconv.Atoi(event.ExtraParam)
	if err != nil {
		g.impossible(event)
		return
	}

	player.RecieveDamage(damage)
	g.event(event.TimeEventHappen, "Player [%d] recieved damage", player.ID, damage)
	if !player.IsAlive() {
		player.FinishRun(domain.Fail, event.TimeEventHappen)
		g.event(event.TimeEventHappen, "Player [%d] is dead", player.ID)
	}
}
