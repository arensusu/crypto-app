// https://bybit-exchange.github.io/docs/v5/intro
package bybitEx

import (
	"funding-rate/exchange/strategy"
	"os"
	"strconv"

	"github.com/hirokisan/bybit/v2"
)

type BybitExchange struct {
	Name   string
	Client *bybit.Client
}

func New() *BybitExchange {
	return &BybitExchange{
		Name:   "Bybit",
		Client: bybit.NewClient().WithAuth(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_API_SECRET")),
	}
}

func (ex *BybitExchange) GetCrossExArbitrageInformation(coin string) (*strategy.CrossExArbitrageInformation, error) {
	symbol := bybit.SymbolV5(coin + "USDT")
	param := bybit.V5GetTickersParam{
		Category: bybit.CategoryV5Linear,
		Symbol:   &symbol,
	}
	tickers, err := ex.Client.V5().Market().GetTickers(param)
	if err != nil {
		return nil, err
	}

	ticker := tickers.Result.LinearInverse.List[0]

	price, _ := strconv.ParseFloat(ticker.LastPrice, 64)
	fundingRate, _ := strconv.ParseFloat(ticker.FundingRate, 64)
	fundingTime, _ := strconv.ParseInt(ticker.NextFundingTime, 10, 64)
	return &strategy.CrossExArbitrageInformation{
		ExchangeName:    ex.Name,
		LastPrice:       price,
		FundingRate:     fundingRate,
		NextFundingTime: fundingTime,
	}, nil
}
