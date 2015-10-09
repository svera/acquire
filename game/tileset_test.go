package game

import "testing"

func TestNewTileSet(t *testing.T) {
	tileset := NewTileset()
	if len(tileset.tiles) != 108 {
		t.Errorf("Tileset must have exactly 108 tiles, got %d", len(tileset.tiles))
	}
}

func TestDraw(t *testing.T) {
	tileset := NewTileset()
	tile := tileset.Draw()
	if tile.number < 1 || tile.number > 13 {
		t.Errorf("Drawn tile value is not valid")
	}
	if tile.letter < "A" || tile.letter > "I" {
		t.Errorf("Drawn tile letter is not valid")
	}
	if len(tileset.tiles) != 107 {
		t.Errorf("Tile must be extracted from tileset")
	}
}
