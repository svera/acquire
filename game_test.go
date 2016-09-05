package acquire

import (
	"testing"

	"github.com/svera/acquire/interfaces"
	"github.com/svera/acquire/mocks"
)

func TestNewGameWrongNumberPlayers(t *testing.T) {
	players, optional := setup()
	players = players[:1]

	if _, err := New(players, optional); err.Error() != WrongNumberPlayers {
		t.Errorf("Game must not be created with less than 3 players, got %d", len(players))
	}
}

func TestNewGameInitsPlayersTilesets(t *testing.T) {
	players, optional := setup()
	New(players, optional)

	for i, player := range players {
		if len(player.Tiles()) != 6 {
			t.Errorf("Players must have 6 tiles at the beginning, player %d got %d", i, len(player.Tiles()))
		}
	}
}

func TestAreEndConditionsReached(t *testing.T) {
	players, optional := setup()
	game, _ := New(players, optional)

	if game.AreEndConditionsReached() {
		t.Errorf("End game conditions not reached (no active corporations) but detected as it were")
	}

	optional.Corporations[0].(*mocks.Corporation).FakeSize = 41
	optional.Corporations[0].(*mocks.Corporation).FakeIsActive = true

	if !game.AreEndConditionsReached() {
		t.Errorf("End game conditions reached (a corporation bigger than 40 tiles) but not detected")
	}

	optional.Corporations[0].(*mocks.Corporation).FakeSize = 11
	optional.Corporations[0].(*mocks.Corporation).FakeIsActive = true
	optional.Corporations[0].(*mocks.Corporation).FakeIsSafe = true

	if !game.AreEndConditionsReached() {
		t.Errorf("End game conditions reached (all active corporations safe) but not detected")
	}

	optional.Corporations[0].(*mocks.Corporation).FakeSize = 11
	optional.Corporations[1].(*mocks.Corporation).FakeSize = 2
	optional.Corporations[0].(*mocks.Corporation).FakeIsActive = true
	optional.Corporations[1].(*mocks.Corporation).FakeIsActive = true

	if game.AreEndConditionsReached() {
		t.Errorf("End game conditions not reached but detected as it were (only corporation 0 is safe)")
	}

	optional.Corporations[0].(*mocks.Corporation).FakeSize = 2
	optional.Corporations[1].(*mocks.Corporation).FakeSize = 41
	optional.Corporations[0].(*mocks.Corporation).FakeIsSafe = false
	optional.Corporations[1].(*mocks.Corporation).FakeIsSafe = true

	if !game.AreEndConditionsReached() {
		t.Errorf("End game conditions reached but not detected (corporation 1 is bigger than 40 tiles)")
	}

}

func TestPlayTileFoundCorporation(t *testing.T) {
	players, optional := setup()
	tileToPlay := &mocks.Tile{FakeNumber: 5, FakeLetter: "A"}
	optional.Board.(*mocks.Board).FakeFoundCorporation = true
	players[0].(*mocks.Player).FakeHasTile = true
	//optional.StateMachine = &mocks.StateMachine{FakeStateName: interfaces.PlayTileStateName, TimesCalled: map[string]int{}}
	game, _ := New(players, optional)
	game.currentPlayerNumber = 0
	game.PlayTile(tileToPlay)
	if game.stateMachine.(*mocks.StateMachine).TimesCalled["ToFoundCorp"] != 1 {
		t.Errorf("Game must change its state to FoundCorp")
	}
}

func TestFoundCorporation(t *testing.T) {
	players, optional := setup()
	game, _ := New(players, optional)
	game.currentPlayerNumber = 0

	players[0].(*mocks.Player).FakeHasTile = true
	optional.StateMachine = &mocks.StateMachine{FakeStateName: interfaces.PlayTileStateName, TimesCalled: map[string]int{}}
	if err := game.FoundCorporation(optional.Corporations[0]); err == nil {
		t.Errorf("Game in a state different than FoundCorp must not execute FoundCorporation()")
	}
	game.stateMachine.(*mocks.StateMachine).FakeStateName = interfaces.FoundCorpStateName
	newCorpTiles := []interfaces.Tile{
		&mocks.Tile{FakeNumber: 5, FakeLetter: "E"},
		&mocks.Tile{FakeNumber: 6, FakeLetter: "E"},
	}
	game.newCorpTiles = newCorpTiles
	optional.Corporations[0].(*mocks.Corporation).FakeStock = 25

	game.FoundCorporation(optional.Corporations[0])
	if game.stateMachine.(*mocks.StateMachine).TimesCalled["ToBuyStock"] != 1 {
		t.Errorf("Game must change its state to BuyStock")
	}
	if players[0].Shares(optional.Corporations[0]) != 1 {
		t.Errorf("Player must have 1 share of corporation stock, got %d", players[0].Shares(optional.Corporations[0]))
	}
	if optional.Corporations[0].Size() != 2 {
		t.Errorf("Corporation must have 2 tiles, got %d", optional.Corporations[0].Size())
	}
	if game.board.(*mocks.Board).TimesCalled["SetOwner"] != 1 {
		t.Errorf("Corporation tiles are not set on board")
	}
}

func TestPlayTileGrowCorporation(t *testing.T) {
	players, optional := setup()
	tileToPlay := &mocks.Tile{FakeNumber: 6, FakeLetter: "E"}
	optional.StateMachine = &mocks.StateMachine{FakeStateName: interfaces.PlayTileStateName, TimesCalled: map[string]int{}}

	game, _ := New(players, optional)
	game.currentPlayerNumber = 0
	optional.Board.(*mocks.Board).FakeGrowCorporation = true
	optional.Board.(*mocks.Board).FakeGrowCorporationTiles = []interfaces.Tile{
		&mocks.Tile{FakeNumber: 7, FakeLetter: "E"},
		&mocks.Tile{FakeNumber: 8, FakeLetter: "E"},
	}
	optional.Board.(*mocks.Board).FakeGrowCorporationCorp = optional.Corporations[0]
	players[0].(*mocks.Player).FakeHasTile = true

	game.PlayTile(tileToPlay)
	if game.stateMachine.(*mocks.StateMachine).TimesCalled["ToBuyStock"] != 1 {
		t.Errorf("Game must change its state to BuyStock")
	}
	if optional.Corporations[0].(*mocks.Corporation).TimesCalled["Grow"] != 1 {
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
	players, optional := setup()
	setupPlayTileMerge(optional.Corporations, optional.Board)
	tileToPlay := &mocks.Tile{FakeNumber: 6, FakeLetter: "E"}
	optional.StateMachine = &mocks.StateMachine{FakeStateName: interfaces.PlayTileStateName, TimesCalled: map[string]int{}}

	game, _ := New(players, optional)
	game.currentPlayerNumber = 0

	players[0].(*mocks.Player).FakeShares[optional.Corporations[0]] = 6
	players[0].(*mocks.Player).FakeHasTile = true
	players[1].(*mocks.Player).FakeShares[optional.Corporations[0]] = 6
	players[2].(*mocks.Player).FakeShares[optional.Corporations[0]] = 4

	optional.Corporations[0].(*mocks.Corporation).FakeMajorityBonus = 2000
	optional.Corporations[0].(*mocks.Corporation).FakeMinorityBonus = 1000

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
	players, optional := setup()
	setupPlayTileMerge(optional.Corporations, optional.Board)
	tileToPlay := &mocks.Tile{FakeNumber: 6, FakeLetter: "E"}

	game, _ := New(players, optional)
	game.currentPlayerNumber = 0

	players[0].(*mocks.Player).FakeShares[optional.Corporations[0]] = 6
	players[0].(*mocks.Player).FakeHasTile = true
	players[1].(*mocks.Player).FakeShares[optional.Corporations[0]] = 4
	players[2].(*mocks.Player).FakeShares[optional.Corporations[0]] = 4

	optional.Corporations[0].(*mocks.Corporation).FakeMajorityBonus = 2000
	optional.Corporations[0].(*mocks.Corporation).FakeMinorityBonus = 1000

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
	players, optional := setup()
	setupPlayTileMerge(optional.Corporations, optional.Board)
	tileToPlay := &mocks.Tile{FakeNumber: 6, FakeLetter: "E"}

	game, _ := New(players, optional)
	game.currentPlayerNumber = 0

	players[0].(*mocks.Player).FakeShares[optional.Corporations[0]] = 6
	players[0].(*mocks.Player).FakeHasTile = true

	optional.Corporations[0].(*mocks.Corporation).FakeMajorityBonus = 2000
	optional.Corporations[0].(*mocks.Corporation).FakeMinorityBonus = 1000

	game.PlayTile(tileToPlay)
	expectedPlayerCash := 9000
	if players[0].Cash() != expectedPlayerCash {
		t.Errorf("Player hasn't received the correct bonus, must have %d$, got %d$", expectedPlayerCash, players[0].Cash())
	}
}

// Same as previous test, but including the actual change of ownership of tiles on board
// (the complete merge flow)
func TestPlayTileMergeCorporationsComplete(t *testing.T) {
	players, optional := setup()
	setupPlayTileMerge(optional.Corporations, optional.Board)
	tileToPlay := &mocks.Tile{FakeNumber: 6, FakeLetter: "E"}

	game, _ := New(players, optional)
	game.currentPlayerNumber = 0

	players[0].(*mocks.Player).FakeShares[optional.Corporations[0]] = 6
	players[0].(*mocks.Player).FakeHasTile = true

	optional.Corporations[0].(*mocks.Corporation).FakeMajorityBonus = 2000
	optional.Corporations[0].(*mocks.Corporation).FakeMinorityBonus = 1000
	optional.Corporations[0].(*mocks.Corporation).FakeStockPrice = 200

	game.PlayTile(tileToPlay)
	sell := map[interfaces.Corporation]int{optional.Corporations[0]: 6}
	trade := map[interfaces.Corporation]int{}
	game.stateMachine.(*mocks.StateMachine).FakeStateName = interfaces.SellTradeStateName
	game.SellTrade(sell, trade)

	if game.stateMachine.(*mocks.StateMachine).TimesCalled["ToBuyStock"] != 1 {
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
	players, optional := setup()
	setupPlayTileMerge(optional.Corporations, optional.Board)
	tileToPlay := &mocks.Tile{FakeNumber: 7, FakeLetter: "E"}

	game, _ := New(players, optional)
	game.currentPlayerNumber = 0

	optional.Board.PutTile(&mocks.Tile{FakeNumber: 7, FakeLetter: "F"})

	game.PlayTile(tileToPlay)

	optional.Board.(*mocks.Board).FakeAdjacentCells = []interfaces.Owner{
		optional.Corporations[0],
		optional.Corporations[1],
		&mocks.Empty{},
		&mocks.Tile{FakeNumber: 7, FakeLetter: "F"},
	}

	players[0].(*mocks.Player).FakeShares[optional.Corporations[0]] = 6
	players[0].(*mocks.Player).FakeHasTile = true
	players[2].(*mocks.Player).FakeShares[optional.Corporations[0]] = 4

	sell := map[interfaces.Corporation]int{optional.Corporations[0]: 6}
	trade := map[interfaces.Corporation]int{}
	game.stateMachine.(*mocks.StateMachine).FakeStateName = interfaces.SellTradeStateName
	game.mergeCorps = map[string][]interfaces.Corporation{
		"acquirer": []interfaces.Corporation{optional.Corporations[1]},
		"defunct":  []interfaces.Corporation{optional.Corporations[0]},
	}
	game.lastPlayedTile = tileToPlay
	game.SellTrade(sell, trade)

	if game.corporations[1].Size() != 7 {
		t.Errorf("Corporation 1 must have a size of %d, got %d", 7, optional.Corporations[1].Size())
	}
}

func TestSellTradeTurnPassing(t *testing.T) {
	players, optional := setup()
	setupPlayTileMerge(optional.Corporations, optional.Board)
	tileToPlay := &mocks.Tile{FakeNumber: 6, FakeLetter: "E"}

	game, _ := New(players, optional)
	game.currentPlayerNumber = 0

	players[0].(*mocks.Player).FakeShares[optional.Corporations[0]] = 6
	players[0].(*mocks.Player).FakeHasTile = true
	players[2].(*mocks.Player).FakeShares[optional.Corporations[0]] = 4

	optional.Board.(*mocks.Board).FakeAdjacentCells = []interfaces.Owner{}
	optional.Board.(*mocks.Board).FakeMergeCorporations = true

	game.PlayTile(tileToPlay)

	sell := map[interfaces.Corporation]int{optional.Corporations[0]: 6}
	trade := map[interfaces.Corporation]int{}
	game.stateMachine.(*mocks.StateMachine).FakeStateName = interfaces.SellTradeStateName
	game.SellTrade(sell, trade)

	if game.stateMachine.CurrentStateName() != interfaces.SellTradeStateName {
		t.Errorf("Wrong game state after merge, expected %s, got %s", interfaces.SellTradeStateName, game.stateMachine.CurrentStateName())
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
	players, optional := setup()
	optional.Corporations[0].Grow(2)
	buys := map[interfaces.Corporation]int{optional.Corporations[0]: 2}
	expectedAvailableStock := 23
	expectedPlayerStock := 2
	optional.StateMachine = &mocks.StateMachine{FakeStateName: interfaces.BuyStockStateName, TimesCalled: map[string]int{}}
	game, _ := New(players, optional)
	game.currentPlayerNumber = 0

	game.BuyStock(buys)

	if optional.Corporations[0].Stock() != expectedAvailableStock {
		t.Errorf("Corporation stock shares have not decreased, must be %d, got %d", expectedAvailableStock, optional.Corporations[0].Stock())
	}
	if players[0].Shares(optional.Corporations[0]) != expectedPlayerStock {
		t.Errorf("Player stock shares have not increased, must be %d, got %d", expectedPlayerStock, players[0].Shares(optional.Corporations[0]))
	}
}

func TestBuyStockWithNotEnoughCash(t *testing.T) {
	players, optional := setup()
	players[0].(*mocks.Player).FakeCash = 100
	optional.Corporations[0].(*mocks.Corporation).FakeStockPrice = 200

	optional.Corporations[0].Grow(2)

	buys := map[interfaces.Corporation]int{optional.Corporations[0]: 2}
	optional.StateMachine = &mocks.StateMachine{FakeStateName: interfaces.BuyStockStateName, TimesCalled: map[string]int{}}
	game, _ := New(players, optional)
	game.currentPlayerNumber = 0

	err := game.BuyStock(buys)
	if err == nil {
		t.Errorf("Trying to buy stock shares without enough money must throw error")
	}
}

func TestBuyStockAndEndGame(t *testing.T) {
	players, optional := setup()
	optional.Corporations[0].Grow(42)
	optional.Corporations[0].(*mocks.Corporation).FakeIsActive = true
	optional.Corporations[0].(*mocks.Corporation).FakeIsSafe = true
	optional.Corporations[0].(*mocks.Corporation).FakeMajorityBonus = 10000
	optional.Corporations[0].(*mocks.Corporation).FakeMinorityBonus = 5000
	optional.Corporations[0].(*mocks.Corporation).FakeStockPrice = 1000
	buys := map[interfaces.Corporation]int{}
	// Remember, every active corporation has always at least one shareholder
	players[0].AddShares(optional.Corporations[0], 2)
	optional.StateMachine = &mocks.StateMachine{FakeStateName: interfaces.BuyStockStateName, TimesCalled: map[string]int{}}
	game, _ := New(players, optional)
	game.ClaimEndGame()
	game.BuyStock(buys)

	if game.stateMachine.(*mocks.StateMachine).TimesCalled["ToEndGame"] != 1 {
		t.Errorf("Game must change its state to EndGame")
	}
	game.stateMachine.(*mocks.StateMachine).FakeStateName = interfaces.EndGameStateName
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
	players, optional := setup()
	optional.Corporations[0].(*mocks.Corporation).FakeIsSafe = true
	optional.Corporations[1].(*mocks.Corporation).FakeIsSafe = true
	unplayableTile := &mocks.Tile{FakeNumber: 6, FakeLetter: "D"}
	optional.Board.(*mocks.Board).FakeAdjacentCorporations = []interfaces.Corporation{optional.Corporations[0], optional.Corporations[1]}
	optional.StateMachine = &mocks.StateMachine{FakeStateName: interfaces.BuyStockStateName, TimesCalled: map[string]int{}}

	game, _ := New(players, optional)
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
	players, optional := setup()
	optional.Corporations[0].(*mocks.Corporation).FakeIsSafe = true
	playableTile := &mocks.Tile{FakeNumber: 6, FakeLetter: "D"}
	optional.Board.(*mocks.Board).FakeAdjacentCorporations = []interfaces.Corporation{optional.Corporations[0], optional.Corporations[1]}
	optional.StateMachine = &mocks.StateMachine{FakeStateName: interfaces.BuyStockStateName, TimesCalled: map[string]int{}}

	game, _ := New(players, optional)
	if !game.IsTilePlayable(playableTile) {
		t.Errorf("Playable tile detected as unplayable")
	}
}

func TestUntieMerge(t *testing.T) {
	players, optional := setup()
	optional.StateMachine = &mocks.StateMachine{FakeStateName: interfaces.UntieMergeStateName, TimesCalled: map[string]int{}}

	game, _ := New(players, optional)
	game.mergeCorps = map[string][]interfaces.Corporation{
		"acquirer": []interfaces.Corporation{optional.Corporations[0], optional.Corporations[1], optional.Corporations[2]},
		"defunct":  []interfaces.Corporation{optional.Corporations[3]},
	}
	game.lastPlayedTile = &mocks.Tile{FakeNumber: 5, FakeLetter: "E"}
	players[0].(*mocks.Player).FakeShares[optional.Corporations[0]] = 6
	players[0].(*mocks.Player).FakeShares[optional.Corporations[2]] = 6
	players[0].(*mocks.Player).FakeShares[optional.Corporations[3]] = 6

	game.UntieMerge(optional.Corporations[1])
	if game.mergeCorps["acquirer"][0] != optional.Corporations[1] {
		t.Errorf("Tied merge not untied")
	}
	if len(game.mergeCorps["defunct"]) != 3 {
		t.Errorf("Wrong number of defunct corporations after merge untie, expected %d, got %d", 3, len(game.mergeCorps["defunct"]))
	}
	if game.stateMachine.(*mocks.StateMachine).TimesCalled["ToSellTrade"] != 1 {
		t.Errorf("Game must change its state to SellTrade")
	}
}

func TestDeactivatePlayer(t *testing.T) {
	players, optional := setup()
	optional.StateMachine = &mocks.StateMachine{FakeStateName: interfaces.UntieMergeStateName, TimesCalled: map[string]int{}}

	game, _ := New(players, optional)
	game.DeactivatePlayer(players[1])
	if players[1].Active() == true {
		t.Errorf("Deactivated player must be not active")
	}
	if players[1].Cash() != 0 {
		t.Errorf("Deactivated player expected to have no money, got %d", players[1].Cash())
	}
	if game.stateMachine.(*mocks.StateMachine).TimesCalled["ToInsufficientPlayers"] != 1 {
		t.Errorf("Game must change its state to InsufficientPlayers as it has less than 3 active players")
	}
}

// Testing that a completely unplayable hand (both corporations 0 and 1 are safe)
// is detected and replaced
func TestUnplayableHandIsReplaced(t *testing.T) {
	players, optional := setup()
	optional.Corporations[0].(*mocks.Corporation).FakeIsSafe = true
	optional.Corporations[1].(*mocks.Corporation).FakeIsSafe = true
	optional.Board.(*mocks.Board).FakeAdjacentCorporations = []interfaces.Corporation{optional.Corporations[0], optional.Corporations[1]}

	game, _ := New(players, optional)
	game.currentPlayerNumber = 0
	tileToPlay := &mocks.Tile{FakeNumber: 5, FakeLetter: "A"}
	game.PlayTile(tileToPlay)
	if players[1].(*mocks.Player).TimesCalled["PickTile"] != 6 {
		t.Errorf("Player's hand should have been completely replaced")
	}
}

func setup() ([]interfaces.Player, Optional) {
	players := []interfaces.Player{
		&mocks.Player{FakeShares: map[interfaces.Corporation]int{}, FakeCash: 6000, TimesCalled: map[string]int{}, FakeActive: true},
		&mocks.Player{FakeShares: map[interfaces.Corporation]int{}, FakeCash: 6000, TimesCalled: map[string]int{}, FakeActive: true},
		&mocks.Player{FakeShares: map[interfaces.Corporation]int{}, FakeCash: 6000, TimesCalled: map[string]int{}, FakeActive: true},
	}

	corporations := [7]interfaces.Corporation{
		&mocks.Corporation{TimesCalled: map[string]int{}, FakeStock: 25},
		&mocks.Corporation{TimesCalled: map[string]int{}, FakeStock: 25},
		&mocks.Corporation{TimesCalled: map[string]int{}, FakeStock: 25},
		&mocks.Corporation{TimesCalled: map[string]int{}, FakeStock: 25},
		&mocks.Corporation{TimesCalled: map[string]int{}, FakeStock: 25},
		&mocks.Corporation{TimesCalled: map[string]int{}, FakeStock: 25},
		&mocks.Corporation{TimesCalled: map[string]int{}, FakeStock: 25},
	}

	board := &mocks.Board{TimesCalled: map[string]int{}}
	tileset := &mocks.Tileset{FakeTile: &mocks.Tile{FakeNumber: 1, FakeLetter: "A"}, TimesCalled: map[string]int{}}

	optional := Optional{
		Board:        board,
		Corporations: corporations,
		Tileset:      tileset,
		StateMachine: &mocks.StateMachine{FakeStateName: interfaces.PlayTileStateName, TimesCalled: map[string]int{}},
	}

	return players, optional
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
