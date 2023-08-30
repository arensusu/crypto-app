package domain

type Pair struct {
	Exchange string `json:"exchange"`
	Symbol   string `json:"symbol"`
}

type Watchlist struct {
	ID     uint `gorm:"primarykey"`
	ChatID int64
	Pair
}

type FundingWatchlist struct {
	Watchlist
}

type PerpetualWatchlist struct {
	Watchlist
	TargetPrice float64
}

type PrevPrice struct {
	Pair
	Price     float64
	Overdated bool
}

type WatchlistRepository interface {
	CreateFundingWatchlist(chatID int64, pair Pair) error
	RetrieveFundingWatchlists(chatID int64) ([]Pair, error)
	DeleteFundingWatchlist(chatID int64, pair Pair) error

	CreatePerpetualWatchlist(PerpetualWatchlist) error
	RetrievePerpetualWatchlists() ([]PerpetualWatchlist, error)
	RetrievePerpetualWatchlistsOfUser(chatID int64) ([]PerpetualWatchlist, error)
	DeletePerpetualWatchlist(PerpetualWatchlist) error

	CreatePerpPrevPrice(PrevPrice) error
	RetrievePerpPrevPrice(Pair) (PrevPrice, error)
	UpdatePerpPrevPrice(PrevPrice) error
}

type WatchlistUsecase interface {
	AddWatchlist(chatID int64, exchange, symbol string) error
	GetWatchlists(chatID int64) ([]Pair, error)
	RemoveWatchlist(chatID int64, exchange, symbol string) error

	AddPerpetualWatchlist(PerpetualWatchlist) error
	GetPerpetualWatchlists() ([]PerpetualWatchlist, error)
	GetPerpetualWatchlistsOfUser(chatID int64) ([]PerpetualWatchlist, error)
	RemovePerpetualWatchlist(PerpetualWatchlist) error

	GetPerpPrevPrice(Pair) (float64, error)
	SetPerpPrevPrice(PrevPrice) error
}
