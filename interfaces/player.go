package interfaces

// Player is an interface that declares all methods to be implemented by a player implementation
type Player interface {
	Shares(c Corporation) int
	AddShares(corp Corporation, amount int) Player
	RemoveShares(corp Corporation, amount int) Player
	PickTile(t Tile) Player
	Tiles() []Tile
	DiscardTile(t Tile) Player
	HasTile(t Tile) bool
	Cash() int
	AddCash(amount int) Player
	RemoveCash(amount int) Player
}
