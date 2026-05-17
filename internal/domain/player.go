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

func (player *Player) KillMonster(monsterPerFloor int, eventTime time.Time) {
	floor := &player.Run.Floor
	floor.MonstersKilledCount++

	if floor.MonstersKilledCount >= monsterPerFloor {
		floor.FloorCleared = true
		floor.ClearedFloors++
	}
}

func (player Player) MoveToNextFloor(eventTime time.Time) {
	floor := &player.Run.Floor

	floor.CurrentFloor++
	floor.MonstersKilledCount = 0
	floor.FloorCleared = false
	floor.FloorStartedAt = eventTime
}

func (player *Player) RecieveDamage(damage int) {
	player.HP -= damage
}
func (player *Player) RestoreHealth(health int) {
	player.HP += health
	if player.HP > 100 {
		player.HP = 100
	}
}
