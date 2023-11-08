// https://bitgetlimited.github.io/apidoc/en/mix/
package bitget

import (
	"crypto-exchange/exchange"

	"github.com/arensusu/bitget-golang-sdk-api"
)

type Bitget struct {
	Client *bitget.Client
}

func New() exchange.Exchange {
	return &Bitget{Client: bitget.NewClient()}
}

func (ex *Bitget) Name() string { return "Bitget" }
