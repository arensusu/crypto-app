// https://bybit-exchange.github.io/docs/v5/intro
package bybit

import (
	"crypto-exchange/domain"

	"github.com/hirokisan/bybit/v2"
)

func (ex *Bybit) GetFundingAndPrices() (domain.FundingPricesOfSymbol, error) {
	param := bybit.V5GetTickersParam{
		Category: bybit.CategoryV5Linear,
	}
	tickers, err := ex.Client.V5().Market().GetTickers(param)

	if err != nil {
		return nil, err
	}

	results := domain.FundingPricesOfSymbol{}
	for _, ticker := range tickers.Result.LinearInverse.List {
		results.Set(ex.Name, string(ticker.Symbol), ticker.LastPrice, ticker.FundingRate, ticker.NextFundingTime)
	}

	return results, nil
}
