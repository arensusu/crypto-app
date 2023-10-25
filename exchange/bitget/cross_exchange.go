// https://bitgetlimited.github.io/apidoc/en/mix/
package bitget

import (
	"crypto-exchange/exchange/strategy"
	"strings"
)

func (ex *Bitget) GetCrossExArbitrageInformation() (strategy.SymbolExchangeFundingPrice, error) {
	tickers, err := ex.Client.NewMixMarketGetAllTickersService().Do("umcbl")
	if err != nil {
		return nil, err
	}

	results := strategy.SymbolExchangeFundingPrice{}
	for _, ticker := range tickers.Data {
		results.Set(ex.Name, strings.TrimSuffix(ticker.Symbol, "_UMCBL"), ticker.Last, ticker.FundingRate, "0")
	}

	return results, nil
}
