package strategy

type CrossExArbitrageInformation struct {
	ExchangeName    string
	LastPrice       float64
	FundingRate     float64
	NextFundingTime int64
}

type CrossExArbitrager interface {
	GetCrossExArbitrageInformation(string) (*CrossExArbitrageInformation, error)
}

type CrossExArbitrageResult struct {
	ExchangePair     string
	PriceDiffPercent float64
	FundingRateDiff  float64
	NextFundingTime  string
}

type CrossExArbitrageUsecase interface {
	CalculateCrossExArbitrage(string) ([]CrossExArbitrageResult, error)
}
