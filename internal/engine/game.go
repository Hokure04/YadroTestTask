package engine

import (
	"fmt"
	"impulse/internal/config"
	"impulse/internal/domain"
)

type Game struct {
	config      config.Config
	players     map[int]*domain.Player
	outputLines []string
}

func NexGame(cfg config.Config) (*Game, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	return &Game{
		config:  cfg,
		players: make(map[int]*domain.Player),
	}, nil
}

func validateConfig(cfg config.Config) error {
	if cfg.Floors < 0 {
		return fmt.Errorf("invalid floors: %d", cfg.Floors)
	}

	if cfg.Monsters < 0 {
		return fmt.Errorf("invalid monsters: %d", cfg.Monsters)
	}

	if cfg.Duration < 0 {
		return fmt.Errorf("invalid duration: %d", cfg.Duration)
	}
	return nil
}

func (g *Game) OutputLines() []string {
	return g.outputLines
}
