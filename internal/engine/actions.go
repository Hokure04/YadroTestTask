package engine

import "impulse/internal/domain"

func (g *Game) processRegister(event domain.Event, player *domain.Player) {
	if g.reject(event, player.Registered || player.Run != nil) {
		return
	}

	player.Registered = true
	g.event(event.TimeEventHappen, "Player [%d] registered", player.ID)
}

func (g *Game) processEnterDungeon(event domain.Event, player *domain.Player) {
	if g.reject(event, !g.canEnterDungeon(event, player)) {
		return
	}

	player.StartRun(event.TimeEventHappen)
	g.event(event.TimeEventHappen, "Player [%d] entered the dungeon", player.ID)
}

func (g *Game) processKillMonster(event domain.Event, player *domain.Player) {
	if g.reject(event, !g.canKillMonster(player)) {
		return
	}

	floor := &player.Run.Floor
	currentFloor := floor.CurrentFloor
	floor.MonstersKilledCount[currentFloor]++

	g.event(event.TimeEventHappen, "Player [%d] killed the monster", player.ID)

	if floor.MonstersKilledCount[currentFloor] == g.config.Monsters {
		floor.FloorCleared[currentFloor] = true
		floor.ClearedFloors++
		startedAt := floor.FloorStartedAt[currentFloor]
		floor.ClearDuration = append(
			floor.ClearDuration,
			event.TimeEventHappen.Sub(startedAt),
		)
	}
}

func (g *Game) processNextFloor(event domain.Event, player *domain.Player) {
	if g.reject(event, !g.canMoveNextFloor(player)) {
		return
	}

	floor := &player.Run.Floor
	floor.CurrentFloor++

	if floor.CurrentFloor <= g.config.Floors {
		floor.MonstersKilledCount[floor.CurrentFloor] = 0
		floor.FloorStartedAt[floor.CurrentFloor] = event.TimeEventHappen
		floor.FloorCleared[floor.CurrentFloor] = false
	}

	g.event(event.TimeEventHappen, "Player [%d] went to the next floor", player.ID)
}

func (g *Game) processPreviousFloor(event domain.Event, player *domain.Player) {
	if g.reject(event, !g.canMovePreviousFloor(player)) {
		return
	}

	player.Run.Floor.CurrentFloor--
	g.event(event.TimeEventHappen, "Player [%d] went to the previous floor", player.ID)
}

func (g *Game) processEnterBossFloor(event domain.Event, player *domain.Player) {
	if g.reject(event, !g.canEnterBossFloor(player)) {
		return
	}

	player.Run.Boss.Entered = true
	player.Run.Boss.EnteredAt = event.TimeEventHappen

	g.event(event.TimeEventHappen, "Player [%d] entered the boss's floor", player.ID)
}

func (g *Game) processKillBoss(event domain.Event, player *domain.Player) {
	if g.reject(event, !g.canKillBoss(player)) {
		return
	}

	player.Run.Boss.Killed = true
	player.Run.Boss.KilledAt = event.TimeEventHappen.Sub(player.Run.Boss.EnteredAt)

	g.event(event.TimeEventHappen, "Player [%d] killed the boss", player.ID)
}

func (g *Game) processLeaveDungeon(event domain.Event, player *domain.Player) {
	if g.reject(event, !g.canLeaveDungeon(player)) {
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
	if g.reject(event, !g.canAct(player)) {
		return
	}

	player.FinishRun(domain.Disqual, event.TimeEventHappen)
	player.Disqualified = true

	g.event(event.TimeEventHappen, "Player [%d] cannot continue due to [%s]", player.ID, event.ExtraParam)
}

func (g *Game) processRestoreHealth(event domain.Event, player *domain.Player) {
	if g.reject(event, !g.canAct(player)) {
		return
	}

	health, ok := parseNonNegativeInt(event.ExtraParam)
	if g.reject(event, !ok) {
		return
	}

	player.RestoreHealth(health)
	g.event(event.TimeEventHappen, "Player [%d] has restored [%d] of health", player.ID, health)
}

func (g *Game) processReceiveDamage(event domain.Event, player *domain.Player) {
	if g.reject(event, !g.canAct(player)) {
		return
	}

	damage, ok := parseNonNegativeInt(event.ExtraParam)
	if g.reject(event, !ok) {
		return
	}

	player.ReceiveDamage(damage)
	g.event(event.TimeEventHappen, "Player [%d] recieved [%d] of damage", player.ID, damage)

	if !player.IsAlive() {
		player.FinishRun(domain.Fail, event.TimeEventHappen)
		g.event(event.TimeEventHappen, "Player [%d] is dead", player.ID)
	}
}
