package tileset

import "testing"

func TestNewTileSet(t *testing.T) {
	tileset := New()
	if len(tileset.tiles) != 108 {
		t.Errorf("Tileset must have exactly 108 tiles, got %d", len(tileset.tiles))
	}
}

func TestDraw(t *testing.T) {
	tileset := New()
	tile := tileset.Draw()
	if tile.Number < 1 || tile.Number > 13 {
		t.Errorf("Drawn tile value is not valid")
	}
	if tile.Letter < "A" || tile.Letter > "I" {
		t.Errorf("Drawn tile letter is not valid")
	}
	if len(tileset.tiles) != 107 {
		t.Errorf("Tile must be extracted from tileset")
	}
}
