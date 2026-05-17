package main

import (
	"fmt"
	"impulse/internal/config"
	"impulse/internal/engine"
	"impulse/internal/parser"
	"os"
)

func main() {
	configPath := "config.json"
	eventsPath := "events"

	if len(os.Args) >= 2 {
		configPath = os.Args[1]
	}

	if len(os.Args) >= 3 {
		eventsPath = os.Args[2]
	}

	cfg, err := config.Read(configPath)
	if err != nil {
		fmt.Println(os.Stderr, "config error:", err)
		os.Exit(1)
	}

	events, err := parser.ReadEvents(eventsPath)
	if err != nil {
		fmt.Println(os.Stderr, "events error:", err)
		os.Exit(1)
	}

	game, err := engine.NewGame(cfg)
	if err != nil {
		fmt.Println(os.Stderr, "new game error:", err)
		os.Exit(1)
	}

	for _, event := range events {
		game.CloseExpiredDungeons(event.TimeEventHappen)
		game.ProcessEvent(event)
	}

	game.FinishOpenRuns()

	for _, line := range game.OutputLines() {
		fmt.Println(line)
	}

	fmt.Println("Final report:")
	for _, line := range game.CreateReport() {
		fmt.Println(line)
	}
}
