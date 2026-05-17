package engine

import "impulse/internal/domain"

func (g *Game) processRegister(event domain.Event, player *domain.Player) {
	if player.Registered {
		g.impossible(event)
		return
	}
	player.Registered = true
	g.event("Player [%d] registered", player.ID, event.EventID)
}

func (g *Game) processEnterDungeon(event domain.Event, player *domain.Player) {
	player.StartRun(event.TimeEventHappen)
	g.event("Player [%d] entered dungeon", player.ID, event.EventID)
}

func (g *Game) processKillMonster(event domain.Event, player *domain.Player) {
	player.KillMonster(g.config.Monsters, event.TimeEventHappen)
	g.event("Player [%d] killed monster", player.ID, event.EventID)
}

func (g *Game) processNextFloor(event domain.Event, player *domain.Player) {
	player.MoveToNextFloor(event.TimeEventHappen)
	g.event("Player [%d] moved floor", player.ID, event.EventID)
}

func (g *Game) processEnterBossFloor(event domain.Event, player *domain.Player) {
	g.event("Player [%d] entered boss floor", player.ID, event.EventID)
}

func (g *Game) processKillBoss(event domain.Event, player *domain.Player) {
	g.event("Player [%d] killed boss", player.ID, event.EventID)
}

func (g *Game) processCannotContinue(event domain.Event, player *domain.Player) {
	player.FinishRun(domain.Disqual, event.TimeEventHappen)
	g.event("Player [%d] killed boss", player.ID, event.EventID)
}

func (g *Game) processRestoreHealth(event domain.Event, player *domain.Player) {
	health := 100
	player.RestoreHealth(health)
	g.event("Player [%d] restored health", player.ID, event.EventID)
}

func (g *Game) processRecieveDamage(event domain.Event, player *domain.Player) {
	damage := 50
	player.RecieveDamage(damage)
	g.event("Player [%d] recieved damage", player.ID, event.EventID)
	if !player.IsAlive() {
		player.FinishRun(domain.Fail, event.TimeEventHappen)
		g.event("Player [%d] is dead", player.ID, event.EventID)
	}
}
