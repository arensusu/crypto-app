package exchange

import "github.com/arensusu/crypto-app/exchange/types"

type Exchange interface {
	Name() string

	GetFundingAndPrice(symbol string) (*types.FundingFeeArbitrage, error)
	GetAssets() ([]types.Asset, error)
}
