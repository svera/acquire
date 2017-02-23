package acquire

import "github.com/svera/acquire/interfaces"

type sortablePlayers struct {
	players []interfaces.Player
	corp    interfaces.Corporation
}

func (s sortablePlayers) Len() int { return len(s.players) }
func (s sortablePlayers) Less(i, j int) bool {
	return s.players[i].Shares(s.corp) < s.players[j].Shares(s.corp)
}
func (s sortablePlayers) Swap(i, j int) { s.players[i], s.players[j] = s.players[j], s.players[i] }
