package domain

import "time"

type BossProgress struct {
	Entered   bool
	Killed    bool
	EnteredAt time.Time
	KilledAt  time.Duration
}
