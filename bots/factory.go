package bots

import (
	"errors"
	"github.com/svera/acquire/interfaces"
)

const (
	BotNotFound = "bot_not_found"
)

func Create(name string) (interfaces.Bot, error) {
	switch name {
	case "random":
		return NewRandom(), nil
	default:
		return &NullBot{}, errors.New(BotNotFound)
	}
}
