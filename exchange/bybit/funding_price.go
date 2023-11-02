// https://bybit-exchange.github.io/docs/v5/intro
package bybit

import (
	"crypto-exchange/exchange/domain"
	"errors"

	"github.com/hirokisan/bybit/v2"
)

func (ex *Bybit) GetFundingAndPrice(symbol string) (*domain.FundingPrice, error) {
	param := bybit.V5GetTickersParam{
		Category: bybit.CategoryV5Linear,
		Symbol:   (*bybit.SymbolV5)(&symbol),
	}
	res, err := ex.Client.V5().Market().GetTickers(param)

	if err != nil {
		return nil, err
	}

	tickers := res.Result.LinearInverse.List
	if len(tickers) != 1 {
		return nil, errors.New("response data length error")
	}

	ticker := tickers[0]

	result := domain.NewFundingPrice(ex.Name, ticker.LastPrice, ticker.FundingRate, ticker.NextFundingTime)

	return result, nil
}

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
