package bots

// Status is a struct used by bot implementations to know about the current
// status of a game.
type Status struct {
	Board map[string]string
	State string
	// Hand is a map that stores all player's tiles and whether each is playable or not
	Hand        map[string]bool
	Corps       [7]CorpData
	TiedCorps   []int
	PlayerInfo  PlayerData
	RivalsInfo  []PlayerData
	IsLastRound bool
}

// CorpData is a struct which holds data about a corporation in a game.
type CorpData struct {
	Name            string
	Price           int
	MajorityBonus   int
	MinorityBonus   int
	RemainingShares int
	Size            int
	Defunct         bool
}

// PlayerData is a struct which holds data about a player in a game.
type PlayerData struct {
	Name        string
	Cash        int
	OwnedShares [7]int
}
