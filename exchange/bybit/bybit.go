package bybitEx

import (
	"funding-rate/exchange/strategy"

	"github.com/hirokisan/bybit/v2"
)

type BybitExchange struct {
	Name   string
	Client *bybit.Client
}

func New(apiKey, apiSecret string) *BybitExchange {
	return &BybitExchange{
		Name:   "Bybit",
		Client: bybit.NewClient().WithAuth(apiKey, apiSecret),
	}
}

func (ex *BybitExchange) GetCrossExArbitrageResponse(coin string) (*strategy.CrossExArbitrageResponse, error) {
	symbol := bybit.SymbolV5(coin + "USDT")
	param := bybit.V5GetTickersParam{
		Category: bybit.CategoryV5Linear,
		Symbol:   &symbol,
	}
	ticker, err := ex.Client.V5().Market().GetTickers(param)
	if err != nil {
		return nil, err
	}

	return strategy.CrossExArbitrageEncoder(ex.Name, ticker.Result.LinearInverse.List[0])
}
