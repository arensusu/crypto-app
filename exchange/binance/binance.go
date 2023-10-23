// https://binance-docs.github.io/apidocs/futures/en/
package binance

import (
	"os"

	"github.com/adshao/go-binance/v2"
)

type Binance struct {
	Name   string
	Client *binance.Client
}

func New() *Binance {
	apiKey, apiSecret := os.Getenv("BINANCE_API_KEY"), os.Getenv("BINANCE_API_SECRET")
	client := binance.NewClient(apiKey, apiSecret)
	return &Binance{
		Name:   "Binance",
		Client: client,
	}
}
