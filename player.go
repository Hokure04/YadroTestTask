package main

import "time"

type Player struct {
	ID           int
	HP           int
	Registered   bool
	Disqualified bool
	Run          *Run
}

const minHealth = 0
const maxHealth = 100

func NewPlayer(id int) *Player {
	return &Player{
		ID: id,
		HP: maxHealth,
	}
}

func (player *Player) IsAlive() bool {
	return player.HP > 0
}

func (player *Player) IsInDungeon() bool {
	return player.Run != nil && !player.Run.Finished
}

func (player *Player) StartRun(startedAt time.Time) {
	player.Run = NewRun(startedAt)
}

func (player *Player) FinishRun(state Result, finishedAt time.Time) {
	if player.Run == nil || player.Run.Finished {
		return
	}
	player.Run.State = state
	player.Run.Finished = true
	player.Run.FinishedAt = finishedAt
}

func (player *Player) ReceiveDamage(damage int) {
	player.HP -= damage
	if player.HP < minHealth {
		player.HP = minHealth
	}
}

func (player *Player) RestoreHealth(health int) {
	player.HP += health
	if player.HP > maxHealth {
		player.HP = maxHealth
	}
}
