package acquire

// Error messages returned by the game
const (
	// Returned when action not allowed at current state
	ActionNotAllowed = "action_%s_not_allowed_expecting_%s"
	// Returned when stock shares from a corporation not on board are not buyable
	StockSharesNotBuyable = "stock_shares_not_buyable"
	// Returned when not enough stock shares of a corporation to buy
	NotEnoughStockShares = "not_enough_stock_shares"
	// Returned when tile temporarily unplayable
	TileTemporarilyUnplayable = "tile_temporarily_unplayable"
	// Returned when tile permanently unplayable
	TilePermanentlyUnplayable = "tile_permanently_unplayable"
	// Returned when player has not enough cash to buy stock shares
	NotEnoughCash = "not_enough_cash"
	// Returned when player can not buy more than 3 stock shares per round
	TooManyStockSharesToBuy = "too_many_stock_shares_to_buy"
	// Returned when some corporation names are repeated
	CorpNamesNotUnique = "corp_names_not_unique"
	// Returned when corporations classes do not fit rules
	WrongNumberCorpsClass = "wrong_number_corps_class"
	// Returned when corporation is already on board and cannot be founded
	CorporationAlreadyOnBoard = "corporation_already_on_board"
	// Returned when there must be between 3 and 6 players
	WrongNumberPlayers = "wrong_number_players"
	// Returned when player does not own stock shares of a certain corporation
	NoCorporationSharesOwned = "no_corporation_shares_owned"
	// Returned when player does not own enough stock shares of a certain corporation
	NotEnoughCorporationSharesOwned = "not_enough_corporation_shares_owned"
	// Returned when player does not have tile on hand
	TileNotOnHand = "tile_not_on_hand"
	// Returned when corporation is not the acquirer in a merge
	NotAnAcquirerCorporation = "not_an_acquirer_corporation"
	// Returned when number of stock shares is not even in a trade
	TradeAmountNotEven = "trade_amount_not_even"
)
