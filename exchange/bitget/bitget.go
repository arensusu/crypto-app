// https://bitgetlimited.github.io/apidoc/en/mix/
package bitget

import (
	"github.com/arensusu/bitget-golang-sdk-api"
)

type Bitget struct {
	Name   string
	Client *bitget.Client
}

func New() *Bitget {
	client := bitget.NewClient()
	return &Bitget{
		Name:   "Bitget",
		Client: client,
	}
}
