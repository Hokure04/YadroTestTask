package engine

import "impulse/internal/domain"

func (g *Game) ProcessEvent(event domain.Event) {
	player := g.getPlayer(event.PlayerID)

	if player.Disqualified {
		return
	}

	if player.Run != nil && player.Run.Finished {
		return
	}

	if event.EventID == domain.Register {
		g.processRegister(event, player)
		return
	}

	if !player.Registered {
		player.Disqualified = true
		g.event(event.TimeEventHappen, "Player [%d] is disqualified", player.ID)
		return
	}

	switch event.EventID {
	case domain.EnterDungeon:
		g.processEnterDungeon(event, player)
	case domain.KillMonster:
		g.processKillMonster(event, player)
	case domain.NextFloor:
		g.processNextFloor(event, player)
	case domain.PreviousFloor:
		g.processPreviousFloor(event, player)
	case domain.EnterBossFloor:
		g.processEnterBossFloor(event, player)
	case domain.KillBoss:
		g.processKillBoss(event, player)
	case domain.LeaveDungeon:
		g.processLeaveDungeon(event, player)
	case domain.CannotContinue:
		g.processCannotContinue(event, player)
	case domain.RestoreHealth:
		g.processRestoreHealth(event, player)
	case domain.ReceiveDamage:
		g.processRecieveDamage(event, player)
	default:
		g.impossible(event)
	}
}
