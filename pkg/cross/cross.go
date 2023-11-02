package cross

type CrossExchangeStrategy interface {
	GetCrossExchangeFundingPrice(symbol string)
}
