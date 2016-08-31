package bots

import (
	"errors"

	"github.com/svera/acquire/interfaces"
)

const (
	// BotNotFound is an error message returned when trying to instance un inexistent bot.
	BotNotFound = "bot_not_found"
)

// Create returns a new instance of a bot.
func Create(level string) (interfaces.Bot, error) {
	switch level {
	case "chaotic":
		return NewChaotic(), nil
	default:
		return nil, errors.New(BotNotFound)
	}
}
