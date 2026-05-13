package engine

import "impulse/internal/domain"

func (g *Game) ProcessEvent(event domain.Event) {
	player := g.getPlayer(event.PlayerID)

	if !player.IsInDungeon() {
		g.impossible(event)
		return
	}

	if event.EventID != domain.Register && !player.Registered {
		player.FinishRun(domain.Disqual, event.TimeEventHappen)
		g.event("Player [%d] is disqualified", player.ID, event.EventID)
		return
	}

	switch event.EventID {
	case domain.Register:
	case domain.EnterDungeon:
	case domain.KillMonster:
	case domain.NextFloor:
	case domain.PreviousFloor:
	case domain.EnterBossFloor:
	case domain.KillBoss:
	case domain.LeaveDungeon:
	case domain.CannotContinue:
	case domain.RestoreHealth:
	case domain.ReceiveDamage:
	default:
		g.impossible(event)
	}
}
