package exchange

import "crypto-exchange/exchange/types"

type Exchange interface {
	Name() string

	GetFundingAndPrice(symbol string) (*types.FundingFeeArbitrage, error)
}
