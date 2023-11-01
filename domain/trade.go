package domain

type MarketTradeParam struct {
	Symbol   string
	Side     string
	Quantity string
}

type MarketTrader interface {
	MarketTrade(MarketTradeParam) error
}
