// https://bybit-exchange.github.io/docs/v5/intro
package bybit

import (
	"crypto-exchange/exchange/strategy"
	"strconv"

	"github.com/hirokisan/bybit/v2"
)

func (ex *Bybit) GetCrossExArbitrageInformation(coin string) (*strategy.CrossExArbitrageInformation, error) {
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
