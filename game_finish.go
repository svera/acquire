package acquire

import (
	"errors"
)

// Finish terminates game as game manual states: "Majority and minority shareholders’ bonuses are paid out
// For all active corporations, and all stocks are sold back to the
// stock market bank at current prices. Stock in a corporation that is not on the board is worthless."
func (g *Game) finish() error {
	if g.state.Name() != "EndGame" {
		return errors.New(ActionNotAllowed)
	}
	for _, corp := range g.getActiveCorporations() {
		g.payBonuses(corp)
		for _, pl := range g.players {
			if pl.Shares(corp) > 0 {
				g.sell(pl, corp, pl.Shares(corp))
			}
		}
	}
	return nil
}