package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadEvents(path string) ([]Event, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var events []Event
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
			return nil, fmt.Errorf("error parsing event %d: %w", lineNumber, err)
		}
		events = append(events, event)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

func ParseEvent(line string) (Event, error) {
	fields := strings.Fields(line)

	if len(fields) < 3 {
		return Event{}, errors.New("invalid event")
	}

	rawTime := strings.Trim(fields[0], "[]")
	eventTime, err := time.Parse("15:04:05", rawTime)
	if err != nil {
		return Event{}, fmt.Errorf("error parsing event time: %s", fields[0])
	}

	playerID, err := strconv.Atoi(fields[1])
	if err != nil {
		return Event{}, fmt.Errorf("invalid player ID: %s", fields[1])
	}

	eventID, err := strconv.Atoi(fields[2])
	if err != nil {
		return Event{}, fmt.Errorf("invalid event ID: %s", fields[2])
	}

	extraParam := ""
	if len(fields) > 3 {
		extraParam = strings.Join(fields[3:], " ")
	}

	return Event{
		TimeEventHappen: eventTime,
		PlayerID:        playerID,
		EventID:         eventID,
		ExtraParam:      extraParam,
	}, nil
}
