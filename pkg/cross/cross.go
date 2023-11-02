package cross

import "crypto-exchange/exchange/domain"

type CrossExchangeStrategy interface {
	GetCrossExchangeFundingPrice(symbol string) (SingleSymbolResult, error)
}

type SingleSymbolResult struct {
	Data []*domain.FundingPrice
	Diff []FundingPriceDiff
}

type FundingPriceDiff struct {
	ExchangeBuy     string
	ExchangeSell    string
	PriceDiff       float64
	FundingRateDiff float64
	FundingTime     string
}
