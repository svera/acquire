package bots

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

type CorpData struct {
	Name            string
	Price           int
	MajorityBonus   int
	MinorityBonus   int
	RemainingShares int
	Size            int
	Defunct         bool
}

type PlayerData struct {
	Name        string
	Cash        int
	OwnedShares [7]int
}

type HandData struct {
	Coords   string
	Playable bool
}
