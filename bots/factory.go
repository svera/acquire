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
func Create(name string) (interfaces.Bot, error) {
	switch name {
	case "random":
		return NewRandom(), nil
	default:
		return nil, errors.New(BotNotFound)
	}
}
