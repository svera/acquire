package interfaces

// Board is an interface that declares all methods to be implemented by a board implementation
type Board interface {
	Cell(number int, letter string) Owner
	TileFoundCorporation(t Tile) (bool, []Tile)
	TileMergeCorporations(t Tile) (bool, map[string][]Corporation)
	TileGrowCorporation(t Tile) (bool, []Tile, Corporation)
	PutTile(t Tile) Board
	AdjacentCells(number int, letter string) []Owner
	SetOwner(cp Corporation, tiles []Tile) Board
	ChangeOwner(oldOwner Corporation, newOwner Corporation) Board
}
