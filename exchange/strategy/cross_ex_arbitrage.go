package strategy

type CrossExArbitrageInformation struct {
	ExchangeName    string
	Symbol          string
	LastPrice       string
	FundingRate     string
	NextFundingTime string
}

type CrossExArbitrager interface {
	GetCrossExArbitrageInformation() (SymbolExchangeFundingPrice, error)
}

type FundingPriceDiff struct {
	ExchangeBuy     string
	ExchangeSell    string
	PriceDiff       float64
	FundingRateDiff float64
	FundingTime     string
}

type SymbolFundingPriceDiffs map[string][]FundingPriceDiff

type CrossExArbitrageUsecase interface {
	CalculateCrossExArbitrage() (SymbolFundingPriceDiffs, error)
}
