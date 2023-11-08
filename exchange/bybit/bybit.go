// https://bybit-exchange.github.io/docs/v5/intro
package bybit

import (
	"crypto-exchange/exchange"
	"os"

	"github.com/hirokisan/bybit/v2"
)

type Bybit struct {
	Client *bybit.Client
}

func New() exchange.Exchange {
	return &Bybit{Client: bybit.NewClient().WithAuth(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET"))}
}

func (ex *Bybit) Name() string { return "Bybit" }
