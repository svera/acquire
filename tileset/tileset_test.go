package tileset

import (
	"testing"

	"github.com/svera/acquire/interfaces"
	"github.com/svera/acquire/mocks"
)

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

func TestDrawWhenOnlyOneTileLeft(t *testing.T) {
	tileset := &Tileset{
		tiles: []interfaces.Tile{
			&mocks.Tile{FakeNumber: 1, FakeLetter: "A"},
		},
	}
	tl, _ := tileset.Draw()
	if tl.Number() != 1 || tl.Letter() != "A" {
		t.Errorf("Tile extracted must have been 1A")
	}
}

func TestDrawFromEmptyTileset(t *testing.T) {
	tileset := &Tileset{}
	if _, err := tileset.Draw(); err == nil {
		t.Errorf("Trying to get a tile from an empty tileset must return an error")
	}
}
