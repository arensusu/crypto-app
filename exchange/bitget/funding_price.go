// https://bitgetlimited.github.io/apidoc/en/mix/
package bitget

import (
	"crypto-exchange/domain"
	"strings"
)

func (ex *Bitget) GetFundingAndPrices() (domain.FundingPricesOfSymbol, error) {
	tickers, err := ex.Client.NewMixMarketGetAllTickersService().Do("umcbl")
	if err != nil {
		return nil, err
	}

	results := domain.FundingPricesOfSymbol{}
	for _, ticker := range tickers.Data {
		results.Set(ex.Name, strings.TrimSuffix(ticker.Symbol, "_UMCBL"), ticker.Last, ticker.FundingRate, "0")
	}

	return results, nil
}
