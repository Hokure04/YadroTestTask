package domain

import "time"

type Player struct {
	ID           int
	HP           int
	Registered   bool
	Disqualified bool
	Run          *Run
}

func NewPlayer(id int) *Player {
	return &Player{
		ID: id,
		HP: 100,
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

func (player *Player) FinishRun(state string, finishedAt time.Time) {
	if player.Run == nil || player.Run.Finished {
		return
	}
	player.Run.State = state
	player.Run.Finished = true
	player.Run.FinishedAt = finishedAt
}

func (player *Player) ReceiveDamage(damage int) {
	player.HP -= damage
	if player.HP < 0 {
		player.HP = 0
	}
}

func (player *Player) RestoreHealth(health int) {
	player.HP += health
	if player.HP > 100 {
		player.HP = 100
	}
}
