package domain

import "time"

type BossProgress struct {
	Entered      bool
	Killed       bool
	EnteredAt    time.Time
	KillDuration time.Duration
}

func (boss *BossProgress) CanEnter() bool {
	return !boss.Entered
}

func (boss *BossProgress) Enter(eventTime time.Time) {
	boss.Entered = true
	boss.EnteredAt = eventTime
}

func (boss *BossProgress) CanKill() bool {
	return boss.Entered && !boss.Killed
}

func (boss *BossProgress) Kill(eventTime time.Time) {
	boss.Killed = true
	boss.KillDuration = eventTime.Sub(boss.EnteredAt)
}
