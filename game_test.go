package acquire

import (
	"fmt"
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/fsm"
	"github.com/svera/acquire/interfaces"
	"github.com/svera/acquire/player"
	"github.com/svera/acquire/tile"
	"github.com/svera/acquire/tileset"
	"testing"
)

func TestNewGameWrongNumberPlayers(t *testing.T) {
	players, corporations, board, tileset := setup()
	players = players[:1]

	if _, err := New(board, players, corporations, tileset, &fsm.PlayTile{}); err.Error() != WrongNumberPlayers {
		t.Errorf("Game must not be created with less than 3 players, got %d", len(players))
	}
}

func TestNewGameNotUniqueCorpNames(t *testing.T) {
	players, corporations, board, tileset := setup()

	corporations[0] = &interfaces.CorporationMock{FakeName: "A"}
	corporations[1] = &interfaces.CorporationMock{FakeName: "A"}

	if _, err := New(board, players, corporations, tileset, &fsm.PlayTile{}); err.Error() != CorpNamesNotUnique {
		t.Errorf("Corporations must have unique names, expecting %s error, got %s", CorpNamesNotUnique, err.Error())
	}
}

func TestNewGameWrongNumberOfCorpsPerClass(t *testing.T) {
	players, corporations, board, tileset := setup()

	corporations[2] = &interfaces.CorporationMock{FakeClass: 0}

	if _, err := New(board, players, corporations, tileset, &fsm.PlayTile{}); err.Error() != WrongNumberCorpsClass {
		t.Errorf("Game must catch wrong number of corporations per class")
	}
}

func TestNewGameInitsPlayersTilesets(t *testing.T) {
	players, corporations, board, tileset := setup()
	New(board, players, corporations, tileset, &fsm.PlayTile{})

	for i, player := range players {
		if len(player.Tiles()) != 6 {
			t.Errorf("Players must have 6 tiles at the beginning, player %d got %d", i, len(player.Tiles()))
		}
	}
}

func TestAreEndConditionsReached(t *testing.T) {
	players, corporations, board, tileset := setup()
	game, _ := New(board, players, corporations, tileset, &fsm.PlayTile{})

	if game.AreEndConditionsReached() {
		t.Errorf("End game conditions not reached (no active corporations) but detected as it were")
	}

	corporations[0].(*interfaces.CorporationMock).FakeSize = 41
	corporations[0].(*interfaces.CorporationMock).FakeIsActive = true

	if !game.AreEndConditionsReached() {
		t.Errorf("End game conditions reached (a corporation bigger than 40 tiles) but not detected")
	}

	corporations[0].(*interfaces.CorporationMock).FakeSize = 11
	corporations[0].(*interfaces.CorporationMock).FakeIsActive = true
	corporations[0].(*interfaces.CorporationMock).FakeIsSafe = true

	if !game.AreEndConditionsReached() {
		t.Errorf("End game conditions reached (all active corporations safe) but not detected")
	}

	corporations[0].(*interfaces.CorporationMock).FakeSize = 11
	corporations[1].(*interfaces.CorporationMock).FakeSize = 2
	corporations[0].(*interfaces.CorporationMock).FakeIsActive = true
	corporations[1].(*interfaces.CorporationMock).FakeIsActive = true

	if game.AreEndConditionsReached() {
		t.Errorf("End game conditions not reached but detected as it were")
	}

}

func TestPlayTileFoundCorporation(t *testing.T) {
	players, corporations, bd, ts := setup()
	tileToPlay := &interfaces.TileMock{FakeNumber: 5, FakeLetter: "A"}
	bd.(*interfaces.BoardMock).FakeFoundCorporation = true
	players[0].(*interfaces.PlayerMock).FakeHasTile = true

	game, _ := New(bd, players, corporations, ts, &fsm.PlayTile{})
	game.PlayTile(tileToPlay)
	if game.state.Name() != interfaces.FoundCorpStateName {
		t.Errorf("Game must be in state FoundCorp, got %s", game.state.Name())
	}
}

func TestFoundCorporation(t *testing.T) {
	players, corporations, bd, ts := setup()
	game, _ := New(bd, players, corporations, ts, &fsm.PlayTile{})
	if err := game.FoundCorporation(corporations[0]); err == nil {
		t.Errorf("Game in a state different than FoundCorp must not execute FoundCorporation()")
	}
	game.state = &fsm.FoundCorp{}
	newCorpTiles := []interfaces.Tile{
		&interfaces.TileMock{FakeNumber: 5, FakeLetter: "E"},
		&interfaces.TileMock{FakeNumber: 6, FakeLetter: "E"},
	}
	game.newCorpTiles = newCorpTiles
	corporations[0].(*interfaces.CorporationMock).FakeStock = 25

	game.FoundCorporation(corporations[0])
	if game.state.Name() != interfaces.BuyStockStateName {
		t.Errorf("Game must be in state BuyStock, got %s", game.state.Name())
	}
	if players[0].Shares(corporations[0]) != 1 {
		t.Errorf("Player must have 1 share of corporation stock, got %d", players[0].Shares(corporations[0]))
	}
	if corporations[0].Size() != 2 {
		t.Errorf("Corporation must have 2 tiles, got %d", corporations[0].Size())
	}
	if game.board.(*interfaces.BoardMock).TimesCalled["SetOwner"] != 1 {
		t.Errorf("Corporation tiles are not set on board")
	}
}

func TestPlayTileGrowCorporation(t *testing.T) {
	players, corporations, bd, ts := setup()
	tileToPlay := &interfaces.TileMock{FakeNumber: 6, FakeLetter: "E"}

	game, _ := New(bd, players, corporations, ts, &fsm.PlayTile{})
	bd.(*interfaces.BoardMock).FakeGrowCorporation = true
	bd.(*interfaces.BoardMock).FakeGrowCorporationTiles = []interfaces.Tile{
		&interfaces.TileMock{FakeNumber: 7, FakeLetter: "E"},
		&interfaces.TileMock{FakeNumber: 8, FakeLetter: "E"},
	}
	bd.(*interfaces.BoardMock).FakeGrowCorporationCorp = corporations[0]
	players[0].(*interfaces.PlayerMock).FakeHasTile = true

	err := game.PlayTile(tileToPlay)
	fmt.Println(err)
	if game.state.Name() != interfaces.BuyStockStateName {
		t.Errorf("Game must be in state BuyStock, got %s", game.state.Name())
	}
	if corporations[0].(*interfaces.CorporationMock).TimesCalled["Grow"] != 1 {
		t.Errorf("Corporation not grown")
	}
}

// Testing this merge:
//   4 5 6 7 8 9
// E [][]><[][][]
//
// In this case, players 0 and 1 are the majority shareholders and player 2 is the minority one
// of corporation 0, with a size of 2
func TestPlayTileMergeCorporationsMultipleMajorityShareholders(t *testing.T) {
	players, corporations, bd, ts := setup()
	setupPlayTileMerge(corporations, bd)
	tileToPlay := &interfaces.TileMock{FakeNumber: 6, FakeLetter: "E"}

	game, _ := New(bd, players, corporations, ts, &fsm.PlayTile{})

	players[0].(*interfaces.PlayerMock).FakeShares[corporations[0]] = 6
	players[1].(*interfaces.PlayerMock).FakeShares[corporations[0]] = 6
	players[2].(*interfaces.PlayerMock).FakeShares[corporations[0]] = 4

	corporations[0].(*interfaces.CorporationMock).FakeMajorityBonus = 6000
	corporations[0].(*interfaces.CorporationMock).FakeMinorityBonus = 3000
	game.PlayTile(tileToPlay)
	expectedPlayer0Cash := 7500
	if players[0].Cash() != expectedPlayer0Cash {
		t.Errorf("Player haven't received the correct bonus, must have %d$, got %d$", expectedPlayer0Cash, players[0].Cash())
	}
	expectedPlayer1Cash := 7500
	if players[1].Cash() != expectedPlayer1Cash {
		t.Errorf("Player haven't received the correct bonus, must have %d$, got %d$", expectedPlayer1Cash, players[1].Cash())
	}
	expectedPlayer2Cash := 6000
	if players[2].Cash() != expectedPlayer2Cash {
		t.Errorf("Player haven't received the correct bonus, must have %d$, got %d$", expectedPlayer2Cash, players[2].Cash())
	}
}

// Testing this merge:
//   4 5 6 7 8 9
// E [][]><[][][]
//
// In this case, player 0 is the majority shareholder and players 1 and 2 are the minority ones
// of corporation 0, with a size of 2
func TestPlayTileMergeCorporationsMultipleMinorityhareholders(t *testing.T) {
	players, corporations, bd, ts := setup()
	setupPlayTileMerge(corporations, bd)
	tileToPlay := tile.New(6, "E")

	game, _ := New(bd, players, corporations, ts, &fsm.PlayTile{})
	playerTiles := players[0].Tiles()
	players[0].
		DiscardTile(playerTiles[0]).
		PickTile(tileToPlay)
	players[0].(*player.Stub).SetShares(corporations[0], 6)
	players[1].(*player.Stub).SetShares(corporations[0], 4)
	players[2].(*player.Stub).SetShares(corporations[0], 4)

	game.PlayTile(tileToPlay)
	expectedPlayer0Cash := 8000
	if players[0].Cash() != expectedPlayer0Cash {
		t.Errorf("Player havent received the correct bonus, must have %d$, got %d$", expectedPlayer0Cash, players[0].Cash())
	}
	expectedPlayer1Cash := 6500
	if players[1].Cash() != expectedPlayer1Cash {
		t.Errorf("Player havent received the correct bonus, must have %d$, got %d$", expectedPlayer1Cash, players[1].Cash())
	}
	expectedPlayer2Cash := 6500
	if players[2].Cash() != expectedPlayer2Cash {
		t.Errorf("Player havent received the correct bonus, must have %d$, got %d$", expectedPlayer2Cash, players[2].Cash())
	}
}

// Testing this merge:
//    4  5 6  7  8  9
// E [0][0]><[1][1][1]
//
// In this case, only player 0 has shares of the defunct corp 0, which has
// a size of 2, thus getting both majority and minority bonuses
func TestPlayTileMergeCorporationsOneShareholder(t *testing.T) {
	players, corporations, bd, ts := setup()
	setupPlayTileMerge(corporations, bd)
	tileToPlay := tile.New(6, "E")

	game, _ := New(bd, players, corporations, ts, &fsm.PlayTile{})
	playerTiles := players[0].Tiles()
	players[0].
		DiscardTile(playerTiles[0]).
		PickTile(tileToPlay)
	players[0].(*player.Stub).SetShares(corporations[0], 6)

	game.PlayTile(tileToPlay)
	expectedPlayerCash := 9000
	if players[0].Cash() != expectedPlayerCash {
		t.Errorf("Player havent received the correct bonus, must have %d$, got %d$", expectedPlayerCash, players[0].Cash())
	}
}

// Same as previous test, but including the actual change of ownership of tiles on board
// (the complete merge flow)
func TestPlayTileMergeCorporationsComplete(t *testing.T) {
	players, corporations, bd, ts := setup()
	setupPlayTileMerge(corporations, bd)
	tileToPlay := tile.New(6, "E")

	game, _ := New(bd, players, corporations, ts, &fsm.PlayTile{})
	playerTiles := players[0].Tiles()
	players[0].
		DiscardTile(playerTiles[0]).
		PickTile(tileToPlay)
	players[0].(*player.Stub).SetShares(corporations[0], 6)

	game.PlayTile(tileToPlay)
	sell := map[interfaces.Corporation]int{corporations[0]: 6}
	trade := map[interfaces.Corporation]int{}
	game.SellTrade(sell, trade)
	if game.state.Name() != interfaces.BuyStockStateName {
		t.Errorf("Wrong game state after merge, expected %s, got %s", interfaces.BuyStockStateName, game.state.Name())
	}
	if game.corporations[0].Size() != 0 {
		t.Errorf("Wrong size for corporation 0, expected %d, got %d", 0, game.corporations[0].Size())
	}
	if game.corporations[1].Size() != 6 {
		t.Errorf("Wrong size for corporation 1, expected %d, got %d", 6, game.corporations[1].Size())
	}
	if game.board.Cell(6, "E") != corporations[1] {
		t.Errorf("Wrong owner for cell %d%s, expected %s, got %s", 6, "E", "corporation", game.board.Cell(6, "E").Type())
	}
	// 6000$ (base cash) + 3000 (majority and minority bonus for class 0 corporation) + (200 * 6) (6 shares owned of the defunct corporation,
	// 200$ per share) = 10200$
	if players[0].Cash() != 10200 {
		t.Errorf("Wrong cash amount for player, expected %d, got %d", 6000, players[0].Cash())
	}
	if players[0].Shares(game.corporations[0]) != 0 {
		t.Errorf("Wrong stock shares amount for player, expected %d, got %d", 0, players[0].Shares(game.corporations[0]))
	}
}

func TestSellTradeTurnPassing(t *testing.T) {
	players, corporations, bd, ts := setup()
	setupPlayTileMerge(corporations, bd)
	tileToPlay := tile.New(6, "E")

	game, _ := New(bd, players, corporations, ts, &fsm.PlayTile{})
	playerTiles := players[0].Tiles()
	players[0].
		DiscardTile(playerTiles[0]).
		PickTile(tileToPlay)
	players[0].(*player.Stub).SetShares(corporations[0], 6)
	players[2].(*player.Stub).SetShares(corporations[0], 4)

	game.PlayTile(tileToPlay)
	sell := map[interfaces.Corporation]int{corporations[0]: 6}
	trade := map[interfaces.Corporation]int{}
	game.SellTrade(sell, trade)

	if game.state.Name() != interfaces.SellTradeStateName {
		t.Errorf("Wrong game state after merge, expected %s, got %s", interfaces.SellTradeStateName, game.state.Name())
	}
	if game.CurrentPlayer() != players[2] {
		t.Errorf("Wrong active player, expected %d, got %d", 2, game.currentPlayerNumber)
	}
}

// Set ups the board this way for merge tests
//   4 5 6 7 8 9
// E [][]><[][][]
func setupPlayTileMerge(corporations [7]interfaces.Corporation, bd interfaces.Board) {
	corp0Tiles := []interfaces.Tile{
		tile.New(4, "E"),
		tile.New(5, "E"),
	}
	corp1Tiles := []interfaces.Tile{
		tile.New(7, "E"),
		tile.New(8, "E"),
		tile.New(9, "E"),
	}
	corporations[0].Grow(len(corp0Tiles))
	corporations[1].Grow(len(corp1Tiles))
	bd.SetOwner(corporations[0], corp0Tiles)
	bd.SetOwner(corporations[1], corp1Tiles)
}

func TestBuyStock(t *testing.T) {
	players, corporations, bd, ts := setup()
	corporations[0].Grow(2)
	buys := map[interfaces.Corporation]int{corporations[0]: 2}
	expectedAvailableStock := 23
	expectedPlayerStock := 2
	game, _ := New(bd, players, corporations, ts, &fsm.PlayTile{})
	game.state = &fsm.BuyStock{}
	game.BuyStock(buys)

	if corporations[0].Stock() != expectedAvailableStock {
		t.Errorf("Corporation stock shares have not decreased, must be %d, got %d", expectedAvailableStock, corporations[0].Stock())
	}
	if players[0].Shares(corporations[0]) != expectedPlayerStock {
		t.Errorf("Player stock shares have not increased, must be %d, got %d", expectedPlayerStock, players[0].Shares(corporations[0]))
	}
}

func TestBuyStockWithNotEnoughCash(t *testing.T) {
	players, corporations, bd, ts := setup()
	players[0].(*player.Stub).SetCash(100)

	corporations[0].Grow(2)

	buys := map[interfaces.Corporation]int{corporations[0]: 2}
	game, _ := New(bd, players, corporations, ts, &fsm.PlayTile{})
	err := game.BuyStock(buys)
	if err == nil {
		t.Errorf("Trying to buy stock shares without enough money must throw error")
	}
}

func TestBuyStockAndEndGame(t *testing.T) {
	players, corporations, bd, ts := setup()
	corporations[0].Grow(42)
	buys := map[interfaces.Corporation]int{}
	// Remember, every active corporation has always at least one shareholder
	players[0].AddShares(corporations[0], 2)
	game, _ := New(bd, players, corporations, ts, &fsm.PlayTile{})
	game.ClaimEndGame()
	game.state = &fsm.BuyStock{}
	game.BuyStock(buys)

	if game.state.Name() != interfaces.EndGameStateName {
		t.Errorf("End game was rightly claimed and game state must be %s, got %s", interfaces.EndGameStateName, game.state.Name())
	}
	// 6000$ (base cash) + 15000 (majority and minority bonus for class 0 corporation) + (1000 * 2) (2 shares owned of the defunct corporation,
	// 1000$ per share) = 23000$
	if players[0].Cash() != 23000 {
		t.Errorf("Player final cash must be %d, got %d", 23000, players[0].Cash())
	}
}

// Testing that if player has an permanently unplayable tile, this is exchanged:
// In the following example, tile 6D is unplayable because it would merge safe
// corporations 0 and 1
//
//    5 6  7 8
// D [0]><[1]
func TestDrawTile(t *testing.T) {
	players, corporations, bd, ts := setup()
	corporations[0].(*corporation.Stub).SetSize(11)
	corporations[1].(*corporation.Stub).SetSize(15)
	unplayableTile := tile.New(6, "D")
	bd.SetOwner(corporations[0], []interfaces.Tile{tile.New(5, "D")})
	bd.SetOwner(corporations[1], []interfaces.Tile{tile.New(7, "D")})

	game, _ := New(bd, players, corporations, ts, &fsm.PlayTile{})
	players[0].(*player.Stub).SetTiles([]interfaces.Tile{unplayableTile})
	game.tileset.(*tileset.Stub).DiscardTile(unplayableTile)
	game.state = &fsm.BuyStock{}
	game.drawTile()
	for _, tile := range players[0].Tiles() {
		if tile.Number() == unplayableTile.Number() && tile.Letter() == unplayableTile.Letter() {
			t.Errorf("Unplayable tile not discarded after drawing new tile, got %v", players[0].Tiles())
		}
	}
}

func TestUntieMerge(t *testing.T) {
	players, corporations, bd, ts := setup()
	game, _ := New(bd, players, corporations, ts, &fsm.PlayTile{})
	game.mergeCorps = map[string][]interfaces.Corporation{
		"acquirer": []interfaces.Corporation{corporations[0], corporations[1], corporations[2]},
		"defunct":  []interfaces.Corporation{corporations[3]},
	}
	game.state = &fsm.UntieMerge{}
	game.lastPlayedTile = tile.New(5, "E")
	players[0].(*player.Stub).SetShares(corporations[0], 6)
	players[0].(*player.Stub).SetShares(corporations[2], 6)
	players[0].(*player.Stub).SetShares(corporations[3], 6)

	game.UntieMerge(corporations[1])
	if game.mergeCorps["acquirer"][0] != corporations[1] {
		t.Errorf("Tied merge not untied, expected acquirer to be %s, got %s", corporations[1].Name(), game.mergeCorps["acquirer"][0].Name())
	}
	if len(game.mergeCorps["defunct"]) != 3 {
		t.Errorf("Wrong number of defunct corporations after merge untie, expected %d, got %d", 3, len(game.mergeCorps["defunct"]))
	}
	if game.state.Name() != interfaces.SellTradeStateName {
		t.Errorf("Wrong game state after merge untie, expected %s, got %s", "SellTrade", game.state.Name())
	}
}

func setup() ([]interfaces.Player, [7]interfaces.Corporation, interfaces.Board, interfaces.Tileset) {
	players := []interfaces.Player{
		&interfaces.PlayerMock{FakeShares: map[interfaces.Corporation]int{}},
		&interfaces.PlayerMock{FakeShares: map[interfaces.Corporation]int{}},
		&interfaces.PlayerMock{FakeShares: map[interfaces.Corporation]int{}},
	}

	corporations := [7]interfaces.Corporation{
		&interfaces.CorporationMock{FakeName: "A", FakeClass: 0, TimesCalled: map[string]int{}},
		&interfaces.CorporationMock{FakeName: "B", FakeClass: 0, TimesCalled: map[string]int{}},
		&interfaces.CorporationMock{FakeName: "C", FakeClass: 1, TimesCalled: map[string]int{}},
		&interfaces.CorporationMock{FakeName: "D", FakeClass: 1, TimesCalled: map[string]int{}},
		&interfaces.CorporationMock{FakeName: "E", FakeClass: 1, TimesCalled: map[string]int{}},
		&interfaces.CorporationMock{FakeName: "F", FakeClass: 2, TimesCalled: map[string]int{}},
		&interfaces.CorporationMock{FakeName: "G", FakeClass: 2, TimesCalled: map[string]int{}},
	}

	board := &interfaces.BoardMock{TimesCalled: map[string]int{}}
	tileset := &interfaces.TilesetMock{}
	return players, corporations, board, tileset
}

func slicesSameContent(slice1 []interfaces.Player, slice2 []interfaces.Player) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	var inSlice bool
	for _, val1 := range slice1 {
		inSlice = false
		for _, val2 := range slice2 {
			if val1 == val2 {
				inSlice = true
				break
			}
		}
		if !inSlice {
			return false
		}
	}
	return true
}
