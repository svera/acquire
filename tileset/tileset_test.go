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
	tileset.Draw()
	if len(tileset.tiles) != 107 {
		t.Errorf("Tile must be extracted from tileset")
	}
}
