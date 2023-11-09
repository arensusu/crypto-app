// https://binance-docs.github.io/apidocs/futures/en/
package binance_future

import (
	"os"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

type BinanceFuture struct {
	Client *futures.Client
}

func New() *BinanceFuture {
	apiKey, apiSecret := os.Getenv("BINANCE_API_KEY"), os.Getenv("BINANCE_API_SECRET")
	futureClient := binance.NewFuturesClient(apiKey, apiSecret)
	return &BinanceFuture{Client: futureClient}
}

func (ex *BinanceFuture) Name() string { return "BinanceFuture" }
