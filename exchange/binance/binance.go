// https://binance-docs.github.io/apidocs/futures/en/
package binance

import (
	"crypto-exchange/exchange"
	"crypto-exchange/exchange/types"
	"fmt"
	"os"

	"github.com/adshao/go-binance/v2"
)

type Binance struct {
	Client *binance.Client
}

func New() exchange.Exchange {
	apiKey, apiSecret := os.Getenv("BINANCE_API_KEY"), os.Getenv("BINANCE_API_SECRET")
	client := binance.NewClient(apiKey, apiSecret)
	return &Binance{Client: client}
}

func (ex *Binance) Name() string { return "Binance" }

func (ex *Binance) GetFundingAndPrice(symbol string) (*types.FundingFeeArbitrage, error) {
	return nil, fmt.Errorf("not support")
}
