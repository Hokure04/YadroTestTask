package engine

import (
	"fmt"
	"impulse/internal/config"
	"impulse/internal/domain"
	"time"
)

type Game struct {
	config      config.Config
	players     map[int]*domain.Player
	outputLines []string
	openAt      time.Time
	closeAt     time.Time
}

func NewGame(cfg config.Config) (*Game, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	openAt, err := time.Parse("15:04:05", cfg.OpenAt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse open at: %w", err)
	}

	return &Game{
		config:  cfg,
		players: make(map[int]*domain.Player),
		openAt:  openAt,
		closeAt: openAt.Add(time.Duration(cfg.Duration) * time.Hour),
	}, nil
}

func validateConfig(cfg config.Config) error {
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

func (g *Game) getPlayer(id int) *domain.Player {
	player, ok := g.players[id]
	if !ok {
		player = domain.NewPlayer(id)
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
			state := domain.Fail
			if g.isDungeonCompleted(player) {
				state = domain.Success
			}

			player.FinishRun(state, finishedAt)
		}
	}
}

func (g *Game) isDungeonCompleted(player *domain.Player) bool {
	if player.Run == nil {
		return false
	}
	return player.Run.Floor.ClearedFloors >= g.config.Floors && player.Run.Boss.Killed
}
