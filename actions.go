package main

func (g *Game) processRegister(event Event, player *Player) {
	if g.reject(event, player.Registered || player.Run != nil) {
		return
	}

	player.Registered = true
	g.event(event.TimeEventHappen, "Player [%d] registered", player.ID)
}

func (g *Game) processEnterDungeon(event Event, player *Player) {
	if g.reject(event, !g.canEnterDungeon(event, player)) {
		return
	}

	player.StartRun(event.TimeEventHappen)
	g.event(event.TimeEventHappen, "Player [%d] entered the dungeon", player.ID)
}

func (g *Game) processKillMonster(event Event, player *Player) {
	if g.reject(event, !g.canKillMonster(player)) {
		return
	}

	player.Run.Floor.KillMonster(g.config.Monsters, event.TimeEventHappen)
	g.event(event.TimeEventHappen, "Player [%d] killed the monster", player.ID)
}

func (g *Game) processNextFloor(event Event, player *Player) {
	if g.reject(event, !g.canMoveNextFloor(player)) {
		return
	}

	player.Run.Floor.MoveNext(event.TimeEventHappen, g.config.Floors)

	g.event(event.TimeEventHappen, "Player [%d] went to the next floor", player.ID)
}

func (g *Game) processPreviousFloor(event Event, player *Player) {
	if g.reject(event, !g.canMovePreviousFloor(player)) {
		return
	}

	player.Run.Floor.MovePrevious()
	g.event(event.TimeEventHappen, "Player [%d] went to the previous floor", player.ID)
}

func (g *Game) processEnterBossFloor(event Event, player *Player) {
	if g.reject(event, !g.canEnterBossFloor(player)) {
		return
	}

	player.Run.Boss.Enter(event.TimeEventHappen)
	g.event(event.TimeEventHappen, "Player [%d] entered the boss's floor", player.ID)
}

func (g *Game) processKillBoss(event Event, player *Player) {
	if g.reject(event, !g.canKillBoss(player)) {
		return
	}

	player.Run.Boss.Kill(event.TimeEventHappen)
	g.event(event.TimeEventHappen, "Player [%d] killed the boss", player.ID)
}

func (g *Game) processLeaveDungeon(event Event, player *Player) {
	if g.reject(event, !g.canAct(player)) {
		return
	}

	state := Fail
	if g.isDungeonCompleted(player) {
		state = Success
	}

	player.FinishRun(state, event.TimeEventHappen)
	g.event(event.TimeEventHappen, "Player [%d] left the dungeon", player.ID)
}

func (g *Game) processCannotContinue(event Event, player *Player) {
	if g.reject(event, !g.canAct(player)) {
		return
	}

	player.FinishRun(Disqual, event.TimeEventHappen)
	player.Disqualified = true

	g.event(event.TimeEventHappen, "Player [%d] cannot continue due to [%s]", player.ID, event.ExtraParam)
}

func (g *Game) processRestoreHealth(event Event, player *Player) {
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

func (g *Game) processReceiveDamage(event Event, player *Player) {
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
		player.FinishRun(Fail, event.TimeEventHappen)
		g.event(event.TimeEventHappen, "Player [%d] is dead", player.ID)
	}
}
