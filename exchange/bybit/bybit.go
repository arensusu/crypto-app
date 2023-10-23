// https://bybit-exchange.github.io/docs/v5/intro
package bybit

import (
	"os"

	"github.com/hirokisan/bybit/v2"
)

type Bybit struct {
	Name   string
	Client *bybit.Client
}

func New() *Bybit {
	return &Bybit{
		Name:   "Bybit",
		Client: bybit.NewClient().WithAuth(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET")),
	}
}
