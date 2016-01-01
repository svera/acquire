package acquire

import (
	"errors"
)

// Finish terminates game as game manual states: "Majority and minority shareholders’ bonuses are paid out
// For all active corporations, and all stocks are sold back to the
// stock market bank at current prices. Stock in a corporation that is not on the board is worthless."
// TODO finish
func (g *Game) finish() error {
	if g.state.Name() != "EndGame" {
		return errors.New(ActionNotAllowed)
	}
	/*
		for pl := range g.players {

		}
	*/
	return nil
}
