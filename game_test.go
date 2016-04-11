package acquire

import (
	"testing"

	"github.com/svera/acquire/interfaces"
	"github.com/svera/acquire/mocks"
)

func TestNewGameWrongNumberPlayers(t *testing.T) {
	players, corporations, board, tileset := setup()
	players = players[:1]

	if _, err := New(board, players, corporations, tileset, &mocks.State{}); err.Error() != WrongNumberPlayers {
		t.Errorf("Game must not be created with less than 3 players, got %d", len(players))
	}
}

func TestNewGameNotUniqueCorpNames(t *testing.T) {
	players, corporations, board, tileset := setup()

	corporations[0] = &mocks.Corporation{FakeName: "A"}
	corporations[1] = &mocks.Corporation{FakeName: "A"}

	if _, err := New(board, players, corporations, tileset, &mocks.State{}); err.Error() != CorpNamesNotUnique {
		t.Errorf("Corporations must have unique names, expecting %s error, got %s", CorpNamesNotUnique, err.Error())
	}
}

func TestNewGameWrongNumberOfCorpsPerClass(t *testing.T) {
	players, corporations, board, tileset := setup()

	corporations[2] = &mocks.Corporation{FakeClass: 0}

	if _, err := New(board, players, corporations, tileset, &mocks.State{}); err.Error() != WrongNumberCorpsClass {
		t.Errorf("Game must catch wrong number of corporations per class")
	}
}

func TestNewGameInitsPlayersTilesets(t *testing.T) {
	players, corporations, board, tileset := setup()
	New(board, players, corporations, tileset, &mocks.State{})

	for i, player := range players {
		if len(player.Tiles()) != 6 {
			t.Errorf("Players must have 6 tiles at the beginning, player %d got %d", i, len(player.Tiles()))
		}
	}
}

func TestAreEndConditionsReached(t *testing.T) {
	players, corporations, board, tileset := setup()
	game, _ := New(board, players, corporations, tileset, &mocks.State{})

	if game.AreEndConditionsReached() {
		t.Errorf("End game conditions not reached (no active corporations) but detected as it were")
	}

	corporations[0].(*mocks.Corporation).FakeSize = 41
	corporations[0].(*mocks.Corporation).FakeIsActive = true

	if !game.AreEndConditionsReached() {
		t.Errorf("End game conditions reached (a corporation bigger than 40 tiles) but not detected")
	}

	corporations[0].(*mocks.Corporation).FakeSize = 11
	corporations[0].(*mocks.Corporation).FakeIsActive = true
	corporations[0].(*mocks.Corporation).FakeIsSafe = true

	if !game.AreEndConditionsReached() {
		t.Errorf("End game conditions reached (all active corporations safe) but not detected")
	}

	corporations[0].(*mocks.Corporation).FakeSize = 11
	corporations[1].(*mocks.Corporation).FakeSize = 2
	corporations[0].(*mocks.Corporation).FakeIsActive = true
	corporations[1].(*mocks.Corporation).FakeIsActive = true

	if game.AreEndConditionsReached() {
		t.Errorf("End game conditions not reached but detected as it were")
	}

}

func TestPlayTileFoundCorporation(t *testing.T) {
	players, corporations, bd, ts := setup()
	tileToPlay := &mocks.Tile{FakeNumber: 5, FakeLetter: "A"}
	bd.(*mocks.Board).FakeFoundCorporation = true
	players[0].(*mocks.Player).FakeHasTile = true

	game, _ := New(bd, players, corporations, ts, &mocks.State{FakeStateName: interfaces.PlayTileStateName, TimesCalled: map[string]int{}})
	game.PlayTile(tileToPlay)
	if game.state.(*mocks.State).TimesCalled["ToFoundCorp"] != 1 {
		t.Errorf("Game must change its state to FoundCorp")
	}
}

func TestFoundCorporation(t *testing.T) {
	players, corporations, bd, ts := setup()
	game, _ := New(bd, players, corporations, ts, &mocks.State{FakeStateName: interfaces.PlayTileStateName, TimesCalled: map[string]int{}})
	if err := game.FoundCorporation(corporations[0]); err == nil {
		t.Errorf("Game in a state different than FoundCorp must not execute FoundCorporation()")
	}
	game.state.(*mocks.State).FakeStateName = interfaces.FoundCorpStateName
	newCorpTiles := []interfaces.Tile{
		&mocks.Tile{FakeNumber: 5, FakeLetter: "E"},
		&mocks.Tile{FakeNumber: 6, FakeLetter: "E"},
	}
	game.newCorpTiles = newCorpTiles
	corporations[0].(*mocks.Corporation).FakeStock = 25

	game.FoundCorporation(corporations[0])
	if game.state.(*mocks.State).TimesCalled["ToBuyStock"] != 1 {
		t.Errorf("Game must change its state to BuyStock")
	}
	if players[0].Shares(corporations[0]) != 1 {
		t.Errorf("Player must have 1 share of corporation stock, got %d", players[0].Shares(corporations[0]))
	}
	if corporations[0].Size() != 2 {
		t.Errorf("Corporation must have 2 tiles, got %d", corporations[0].Size())
	}
	if game.board.(*mocks.Board).TimesCalled["SetOwner"] != 1 {
		t.Errorf("Corporation tiles are not set on board")
	}
}

func TestPlayTileGrowCorporation(t *testing.T) {
	players, corporations, bd, ts := setup()
	tileToPlay := &mocks.Tile{FakeNumber: 6, FakeLetter: "E"}

	game, _ := New(bd, players, corporations, ts, &mocks.State{FakeStateName: interfaces.PlayTileStateName, TimesCalled: map[string]int{}})
	bd.(*mocks.Board).FakeGrowCorporation = true
	bd.(*mocks.Board).FakeGrowCorporationTiles = []interfaces.Tile{
		&mocks.Tile{FakeNumber: 7, FakeLetter: "E"},
		&mocks.Tile{FakeNumber: 8, FakeLetter: "E"},
	}
	bd.(*mocks.Board).FakeGrowCorporationCorp = corporations[0]
	players[0].(*mocks.Player).FakeHasTile = true

	game.PlayTile(tileToPlay)
	if game.state.(*mocks.State).TimesCalled["ToBuyStock"] != 1 {
		t.Errorf("Game must change its state to BuyStock")
	}
	if corporations[0].(*mocks.Corporation).TimesCalled["Grow"] != 1 {
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
	tileToPlay := &mocks.Tile{FakeNumber: 6, FakeLetter: "E"}

	game, _ := New(bd, players, corporations, ts, &mocks.State{FakeStateName: interfaces.PlayTileStateName, TimesCalled: map[string]int{}})

	players[0].(*mocks.Player).FakeShares[corporations[0]] = 6
	players[0].(*mocks.Player).FakeHasTile = true
	players[1].(*mocks.Player).FakeShares[corporations[0]] = 6
	players[2].(*mocks.Player).FakeShares[corporations[0]] = 4

	corporations[0].(*mocks.Corporation).FakeMajorityBonus = 2000
	corporations[0].(*mocks.Corporation).FakeMinorityBonus = 1000

	game.PlayTile(tileToPlay)

	expectedPlayer0Cash := 7500
	if players[0].Cash() != expectedPlayer0Cash {
		t.Errorf("Player hasn't received the correct bonus, must have %d$, got %d$", expectedPlayer0Cash, players[0].Cash())
	}
	expectedPlayer1Cash := 7500
	if players[1].Cash() != expectedPlayer1Cash {
		t.Errorf("Player hasn't received the correct bonus, must have %d$, got %d$", expectedPlayer1Cash, players[1].Cash())
	}
	expectedPlayer2Cash := 6000
	if players[2].Cash() != expectedPlayer2Cash {
		t.Errorf("Player hasn't received the correct bonus, must have %d$, got %d$", expectedPlayer2Cash, players[2].Cash())
	}
}

// Testing this merge:
//   4 5 6 7 8 9
// E [][]><[][][]
//
// In this case, player 0 is the majority shareholder and players 1 and 2 are the minority ones
// of corporation 0, with a size of 2
func TestPlayTileMergeCorporationsMultipleMinorityShareholders(t *testing.T) {
	players, corporations, bd, ts := setup()
	setupPlayTileMerge(corporations, bd)
	tileToPlay := &mocks.Tile{FakeNumber: 6, FakeLetter: "E"}

	game, _ := New(bd, players, corporations, ts, &mocks.State{FakeStateName: interfaces.PlayTileStateName, TimesCalled: map[string]int{}})

	players[0].(*mocks.Player).FakeShares[corporations[0]] = 6
	players[0].(*mocks.Player).FakeHasTile = true
	players[1].(*mocks.Player).FakeShares[corporations[0]] = 4
	players[2].(*mocks.Player).FakeShares[corporations[0]] = 4

	corporations[0].(*mocks.Corporation).FakeMajorityBonus = 2000
	corporations[0].(*mocks.Corporation).FakeMinorityBonus = 1000

	game.PlayTile(tileToPlay)
	expectedPlayer0Cash := 8000
	if players[0].Cash() != expectedPlayer0Cash {
		t.Errorf("Player hasn't received the correct bonus, must have %d$, got %d$", expectedPlayer0Cash, players[0].Cash())
	}
	expectedPlayer1Cash := 6500
	if players[1].Cash() != expectedPlayer1Cash {
		t.Errorf("Player hasn't received the correct bonus, must have %d$, got %d$", expectedPlayer1Cash, players[1].Cash())
	}
	expectedPlayer2Cash := 6500
	if players[2].Cash() != expectedPlayer2Cash {
		t.Errorf("Player hasn't received the correct bonus, must have %d$, got %d$", expectedPlayer2Cash, players[2].Cash())
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
	tileToPlay := &mocks.Tile{FakeNumber: 6, FakeLetter: "E"}

	game, _ := New(bd, players, corporations, ts, &mocks.State{FakeStateName: interfaces.PlayTileStateName, TimesCalled: map[string]int{}})

	players[0].(*mocks.Player).FakeShares[corporations[0]] = 6
	players[0].(*mocks.Player).FakeHasTile = true

	corporations[0].(*mocks.Corporation).FakeMajorityBonus = 2000
	corporations[0].(*mocks.Corporation).FakeMinorityBonus = 1000

	game.PlayTile(tileToPlay)
	expectedPlayerCash := 9000
	if players[0].Cash() != expectedPlayerCash {
		t.Errorf("Player hasn't received the correct bonus, must have %d$, got %d$", expectedPlayerCash, players[0].Cash())
	}
}

// Same as previous test, but including the actual change of ownership of tiles on board
// (the complete merge flow)
func TestPlayTileMergeCorporationsComplete(t *testing.T) {
	players, corporations, bd, ts := setup()
	setupPlayTileMerge(corporations, bd)
	tileToPlay := &mocks.Tile{FakeNumber: 6, FakeLetter: "E"}

	game, _ := New(bd, players, corporations, ts, &mocks.State{FakeStateName: interfaces.PlayTileStateName, TimesCalled: map[string]int{}})

	players[0].(*mocks.Player).FakeShares[corporations[0]] = 6
	players[0].(*mocks.Player).FakeHasTile = true

	corporations[0].(*mocks.Corporation).FakeMajorityBonus = 2000
	corporations[0].(*mocks.Corporation).FakeMinorityBonus = 1000
	corporations[0].(*mocks.Corporation).FakeStockPrice = 200

	game.PlayTile(tileToPlay)
	sell := map[interfaces.Corporation]int{corporations[0]: 6}
	trade := map[interfaces.Corporation]int{}
	game.state.(*mocks.State).FakeStateName = interfaces.SellTradeStateName
	game.SellTrade(sell, trade)

	if game.state.(*mocks.State).TimesCalled["ToBuyStock"] != 1 {
		t.Errorf("Game must change its state to BuyStock")
	}
	if game.corporations[0].Size() != 0 {
		t.Errorf("Wrong size for corporation 0, expected %d, got %d", 0, game.corporations[0].Size())
	}
	if game.corporations[1].Size() != 6 {
		t.Errorf("Wrong size for corporation 1, expected %d, got %d", 6, game.corporations[1].Size())
	}
	if game.board.(*mocks.Board).TimesCalled["ChangeOwner"] != 1 {
		t.Errorf("Corporation tiles does not change ownership")
	}
	// 6000$ (base cash) + 3000 (majority and minority bonus for class 0 corporation) + (200 * 6) (6 shares owned of the defunct corporation,
	// 200$ per share) = 10200$
	if players[0].Cash() != 10200 {
		t.Errorf("Wrong cash amount for player, expected %d, got %d", 10200, players[0].Cash())
	}
	if players[0].Shares(game.corporations[0]) != 0 {
		t.Errorf("Wrong stock shares amount for player, expected %d, got %d", 0, players[0].Shares(game.corporations[0]))
	}
}

// Testing merge as this:
//   4 5 6 7 8 9
// E [][][]><[][]
// F       []
//
// This is a special case in which, after the merger, tile 7F must be included
// in the resulting corporation as it make it grow
func TestPlayTileMergeCorporationsAndGrow(t *testing.T) {
	players, corporations, bd, ts := setup()
	setupPlayTileMerge(corporations, bd)
	tileToPlay := &mocks.Tile{FakeNumber: 7, FakeLetter: "E"}

	game, _ := New(bd, players, corporations, ts, &mocks.State{FakeStateName: interfaces.PlayTileStateName, TimesCalled: map[string]int{}})

	bd.PutTile(&mocks.Tile{FakeNumber: 7, FakeLetter: "F"})

	game.PlayTile(tileToPlay)

	bd.(*mocks.Board).FakeAdjacentCells = []interfaces.Owner{
		corporations[0],
		corporations[1],
		&mocks.Empty{},
		&mocks.Tile{FakeNumber: 7, FakeLetter: "F"},
	}

	players[0].(*mocks.Player).FakeShares[corporations[0]] = 6
	players[0].(*mocks.Player).FakeHasTile = true
	players[2].(*mocks.Player).FakeShares[corporations[0]] = 4

	sell := map[interfaces.Corporation]int{corporations[0]: 6}
	trade := map[interfaces.Corporation]int{}
	game.state.(*mocks.State).FakeStateName = interfaces.SellTradeStateName
	game.mergeCorps = map[string][]interfaces.Corporation{
		"acquirer": []interfaces.Corporation{corporations[1]},
		"defunct":  []interfaces.Corporation{corporations[0]},
	}
	game.lastPlayedTile = tileToPlay
	game.SellTrade(sell, trade)

	if game.corporations[1].Size() != 7 {
		t.Errorf("Corporation 1 must have a size of %d, got %d", 7, corporations[1].Size())
	}
}

func TestSellTradeTurnPassing(t *testing.T) {
	players, corporations, bd, ts := setup()
	setupPlayTileMerge(corporations, bd)
	tileToPlay := &mocks.Tile{FakeNumber: 6, FakeLetter: "E"}

	game, _ := New(bd, players, corporations, ts, &mocks.State{FakeStateName: interfaces.PlayTileStateName, TimesCalled: map[string]int{}})

	players[0].(*mocks.Player).FakeShares[corporations[0]] = 6
	players[0].(*mocks.Player).FakeHasTile = true
	players[2].(*mocks.Player).FakeShares[corporations[0]] = 4

	bd.(*mocks.Board).FakeAdjacentCells = []interfaces.Owner{}
	bd.(*mocks.Board).FakeMergeCorporations = true

	game.PlayTile(tileToPlay)

	sell := map[interfaces.Corporation]int{corporations[0]: 6}
	trade := map[interfaces.Corporation]int{}
	game.state.(*mocks.State).FakeStateName = interfaces.SellTradeStateName
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
	bd.(*mocks.Board).FakeMergeCorporations = true
	bd.(*mocks.Board).FakeMergeCorporationsCorps = map[string][]interfaces.Corporation{
		"acquirer": []interfaces.Corporation{corporations[1]},
		"defunct":  []interfaces.Corporation{corporations[0]},
	}

	corporations[0].Grow(2)
	corporations[1].Grow(3)
}

func TestBuyStock(t *testing.T) {
	players, corporations, bd, ts := setup()
	corporations[0].Grow(2)
	buys := map[interfaces.Corporation]int{corporations[0]: 2}
	expectedAvailableStock := 23
	expectedPlayerStock := 2
	game, _ := New(bd, players, corporations, ts, &mocks.State{FakeStateName: interfaces.BuyStockStateName, TimesCalled: map[string]int{}})
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
	players[0].(*mocks.Player).FakeCash = 100
	corporations[0].(*mocks.Corporation).FakeStockPrice = 200

	corporations[0].Grow(2)

	buys := map[interfaces.Corporation]int{corporations[0]: 2}
	game, _ := New(bd, players, corporations, ts, &mocks.State{FakeStateName: interfaces.BuyStockStateName, TimesCalled: map[string]int{}})
	err := game.BuyStock(buys)
	if err == nil {
		t.Errorf("Trying to buy stock shares without enough money must throw error")
	}
}

func TestBuyStockAndEndGame(t *testing.T) {
	players, corporations, bd, ts := setup()
	corporations[0].Grow(42)
	corporations[0].(*mocks.Corporation).FakeIsActive = true
	corporations[0].(*mocks.Corporation).FakeIsSafe = true
	corporations[0].(*mocks.Corporation).FakeMajorityBonus = 10000
	corporations[0].(*mocks.Corporation).FakeMinorityBonus = 5000
	corporations[0].(*mocks.Corporation).FakeStockPrice = 1000
	buys := map[interfaces.Corporation]int{}
	// Remember, every active corporation has always at least one shareholder
	players[0].AddShares(corporations[0], 2)
	game, _ := New(bd, players, corporations, ts, &mocks.State{FakeStateName: interfaces.BuyStockStateName, TimesCalled: map[string]int{}})
	game.ClaimEndGame()
	game.BuyStock(buys)

	if game.state.(*mocks.State).TimesCalled["ToEndGame"] != 1 {
		t.Errorf("Game must change its state to EndGame")
	}
	game.state.(*mocks.State).FakeStateName = interfaces.EndGameStateName
	game.finish()
	// 6000$ (base cash) + 15000 (majority and minority bonus for class 0 corporation) + (1000 * 2) (2 shares owned of the defunct corporation,
	// 1000$ per share) = 23000$
	if players[0].Cash() != 23000 {
		t.Errorf("Player final cash must be %d, got %d", 23000, players[0].Cash())
	}
}

// Testing that if player has an permanently unplayable tile, it is detected:
// In the following example, tile 6D is unplayable because it would merge safe
// corporations 0 and 1
//
//    5 6  7 8
// D [0]><[1]
func TestIsTilePlayable(t *testing.T) {
	players, corporations, bd, ts := setup()
	corporations[0].(*mocks.Corporation).FakeIsSafe = true
	corporations[1].(*mocks.Corporation).FakeIsSafe = true
	unplayableTile := &mocks.Tile{FakeNumber: 6, FakeLetter: "D"}
	bd.(*mocks.Board).FakeAdjacentCorporations = []interfaces.Corporation{corporations[0], corporations[1]}

	game, _ := New(bd, players, corporations, ts, &mocks.State{FakeStateName: interfaces.BuyStockStateName, TimesCalled: map[string]int{}})
	if game.IsTilePlayable(unplayableTile) {
		t.Errorf("Unplayable tile not detected")
	}
}

// Testing that if player has a neighbour tile of a safe corporation, that tile
// is still playable:
// In the following example, tile 6D is playable because it would make safe
// corporations 0 grow
//
//    5 6  7 8
// D [0]><[0]
func TestIsGrowTilePlayable(t *testing.T) {
	players, corporations, bd, ts := setup()
	corporations[0].(*mocks.Corporation).FakeIsSafe = true
	playableTile := &mocks.Tile{FakeNumber: 6, FakeLetter: "D"}
	bd.(*mocks.Board).FakeAdjacentCorporations = []interfaces.Corporation{corporations[0], corporations[1]}

	game, _ := New(bd, players, corporations, ts, &mocks.State{FakeStateName: interfaces.BuyStockStateName, TimesCalled: map[string]int{}})
	if !game.IsTilePlayable(playableTile) {
		t.Errorf("Playable tile detected as unplayable")
	}
}

func TestUntieMerge(t *testing.T) {
	players, corporations, bd, ts := setup()
	game, _ := New(bd, players, corporations, ts, &mocks.State{FakeStateName: interfaces.UntieMergeStateName, TimesCalled: map[string]int{}})
	game.mergeCorps = map[string][]interfaces.Corporation{
		"acquirer": []interfaces.Corporation{corporations[0], corporations[1], corporations[2]},
		"defunct":  []interfaces.Corporation{corporations[3]},
	}
	game.lastPlayedTile = &mocks.Tile{FakeNumber: 5, FakeLetter: "E"}
	players[0].(*mocks.Player).FakeShares[corporations[0]] = 6
	players[0].(*mocks.Player).FakeShares[corporations[2]] = 6
	players[0].(*mocks.Player).FakeShares[corporations[3]] = 6

	game.UntieMerge(corporations[1])
	if game.mergeCorps["acquirer"][0] != corporations[1] {
		t.Errorf("Tied merge not untied, expected acquirer to be %s, got %s", corporations[1].Name(), game.mergeCorps["acquirer"][0].Name())
	}
	if len(game.mergeCorps["defunct"]) != 3 {
		t.Errorf("Wrong number of defunct corporations after merge untie, expected %d, got %d", 3, len(game.mergeCorps["defunct"]))
	}
	if game.state.(*mocks.State).TimesCalled["ToSellTrade"] != 1 {
		t.Errorf("Game must change its state to SellTrade")
	}
}

func setup() ([]interfaces.Player, [7]interfaces.Corporation, interfaces.Board, interfaces.Tileset) {
	players := []interfaces.Player{
		&mocks.Player{FakeShares: map[interfaces.Corporation]int{}, FakeCash: 6000, TimesCalled: map[string]int{}},
		&mocks.Player{FakeShares: map[interfaces.Corporation]int{}, FakeCash: 6000, TimesCalled: map[string]int{}},
		&mocks.Player{FakeShares: map[interfaces.Corporation]int{}, FakeCash: 6000, TimesCalled: map[string]int{}},
	}

	corporations := [7]interfaces.Corporation{
		&mocks.Corporation{FakeName: "A", FakeClass: 0, TimesCalled: map[string]int{}, FakeStock: 25},
		&mocks.Corporation{FakeName: "B", FakeClass: 0, TimesCalled: map[string]int{}, FakeStock: 25},
		&mocks.Corporation{FakeName: "C", FakeClass: 1, TimesCalled: map[string]int{}, FakeStock: 25},
		&mocks.Corporation{FakeName: "D", FakeClass: 1, TimesCalled: map[string]int{}, FakeStock: 25},
		&mocks.Corporation{FakeName: "E", FakeClass: 1, TimesCalled: map[string]int{}, FakeStock: 25},
		&mocks.Corporation{FakeName: "F", FakeClass: 2, TimesCalled: map[string]int{}, FakeStock: 25},
		&mocks.Corporation{FakeName: "G", FakeClass: 2, TimesCalled: map[string]int{}, FakeStock: 25},
	}

	board := &mocks.Board{TimesCalled: map[string]int{}}
	tileset := &mocks.Tileset{FakeTile: &mocks.Tile{FakeNumber: 1, FakeLetter: "A"}}
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
