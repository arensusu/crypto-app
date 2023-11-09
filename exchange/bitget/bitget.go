// https://bitgetlimited.github.io/apidoc/en/mix/
package bitget

import (
	"github.com/arensusu/crypto-app/exchange"

	"github.com/arensusu/bitget-golang-sdk-api"
)

type Bitget struct {
	Client *bitget.Client
}

func New() exchange.Exchange {
	return &Bitget{Client: bitget.NewClient()}
}

func (ex *Bitget) Name() string { return "Bitget" }
