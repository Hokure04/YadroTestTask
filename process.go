package main

func (g *Game) ProcessEvent(event Event) {
	player := g.getPlayer(event.PlayerID)

	if player.Disqualified {
		g.impossible(event)
		return
	}

	if player.Run != nil && player.Run.Finished {
		g.impossible(event)
		return
	}

	if event.EventID == Register {
		g.processRegister(event, player)
		return
	}

	if !player.Registered {
		player.Disqualified = true
		g.event(event.TimeEventHappen, "Player [%d] is disqualified", player.ID)
		return
	}

	switch event.EventID {
	case EnterDungeon:
		g.processEnterDungeon(event, player)
	case KillMonster:
		g.processKillMonster(event, player)
	case NextFloor:
		g.processNextFloor(event, player)
	case PreviousFloor:
		g.processPreviousFloor(event, player)
	case EnterBossFloor:
		g.processEnterBossFloor(event, player)
	case KillBoss:
		g.processKillBoss(event, player)
	case LeaveDungeon:
		g.processLeaveDungeon(event, player)
	case CannotContinue:
		g.processCannotContinue(event, player)
	case RestoreHealth:
		g.processRestoreHealth(event, player)
	case ReceiveDamage:
		g.processReceiveDamage(event, player)
	default:
		g.impossible(event)
	}
}
