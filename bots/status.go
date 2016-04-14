package bots

// Status is a struct used by bot implementations to know about the current
// status of a game.
type Status struct {
	Board      map[string]string
	State      string
	Hand       []HandData
	Corps      [7]CorpData
	TiedCorps  []string
	PlayerInfo PlayerData
	RivalsInfo []PlayerData
	LastTurn   bool
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

// HandData is a struct which holds data about a player hand in a game.
type HandData struct {
	Coords   string
	Playable bool
}
