package main

import (
	"fmt"
	"time"
)

type Game struct {
	config      Config
	players     map[int]*Player
	outputLines []string
	openAt      time.Time
	closeAt     time.Time
}

func NewGame(cfg Config) (*Game, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	openAt, err := time.Parse(clockLayout, cfg.OpenAt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse open at: %w", err)
	}

	return &Game{
		config:  cfg,
		players: make(map[int]*Player),
		openAt:  openAt,
		closeAt: openAt.Add(time.Duration(cfg.Duration) * time.Hour),
	}, nil
}

func validateConfig(cfg Config) error {
	if cfg.Floors <= 0 {
		return fmt.Errorf("invalid floors: %d", cfg.Floors)
	}

	if cfg.Monsters <= 0 {
		return fmt.Errorf("invalid monsters: %d", cfg.Monsters)
	}

	if cfg.Duration <= 0 {
		return fmt.Errorf("invalid duration: %d", cfg.Duration)
	}
	return nil
}

func (g *Game) OutputLines() []string {
	return g.outputLines
}

func (g *Game) getPlayer(id int) *Player {
	player, ok := g.players[id]
	if !ok {
		player = NewPlayer(id)
		g.players[id] = player
	}
	return player
}

func (g *Game) CloseExpiredDungeons(currentTime time.Time) {
	if currentTime.Before(g.closeAt) {
		return
	}
	g.finishOpenRuns(g.closeAt)
}

func (g *Game) FinishOpenRuns() {
	g.finishOpenRuns(g.closeAt)
}

func (g *Game) finishOpenRuns(finishedAt time.Time) {
	for _, player := range g.players {
		if player.IsInDungeon() {
			state := Fail
			if g.isDungeonCompleted(player) {
				state = Success
			}

			player.FinishRun(state, finishedAt)
		}
	}
}

func (g *Game) isDungeonCompleted(player *Player) bool {
	if player.Run == nil {
		return false
	}
	return player.Run.Floor.ClearedFloors >= g.config.Floors && player.Run.Boss.Killed
}
