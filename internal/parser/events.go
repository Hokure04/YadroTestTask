package parser

import (
	"bufio"
	"errors"
	"fmt"
	"impulse/internal/domain"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadEvents(path string) ([]domain.Event, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var events []domain.Event
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		event, err := ParseEvent(line)
		if err != nil {
			return nil, fmt.Errorf("Error parsing event %d: %w", lineNumber, err)
		}
		events = append(events, event)
	}
	return events, nil
}

func ParseEvent(line string) (domain.Event, error) {
	fields := strings.Fields(line)

	if len(fields) < 3 {
		return domain.Event{}, errors.New("Invalid event")
	}

	time := time.Now()

	playerID, err := strconv.Atoi(fields[1])
	if err != nil {
		return domain.Event{}, fmt.Errorf("Invalid player ID: %s", fields[1])
	}

	eventID, err := strconv.Atoi(fields[2])
	if err != nil {
		return domain.Event{}, fmt.Errorf("Invalid event ID: %s", fields[2])
	}

	return domain.Event{
		TimeEventHappen: time,
		PlayerID:        playerID,
		EventID:         eventID,
	}, nil
}
