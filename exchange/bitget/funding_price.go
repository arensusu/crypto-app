// https://bitgetlimited.github.io/apidoc/en/mix/
package bitget

import (
	"crypto-exchange/domain"
	"strings"
)

func (ex *Bitget) GetFundingAndPrice(symbol string) (*domain.FundingPrice, error) {
	bitgetSymbol := symbol + "_UMCBL"
	tickerRes, err := ex.Client.MixMarket().GetTicker(bitgetSymbol)
	if err != nil {
		return nil, err
	}

	fundingTimeRes, err := ex.Client.MixMarket().GetNextFundingTime(bitgetSymbol)
	if err != nil {
		return nil, err
	}

	ticker := tickerRes.Data
	fundingTime := fundingTimeRes.Data

	result := domain.NewFundingPrice(ex.Name, ticker.Last, ticker.FundingRate, fundingTime.FundingTime)
	return result, nil
}

func (ex *Bitget) GetFundingAndPrices() (domain.FundingPricesOfSymbol, error) {
	tickers, err := ex.Client.MixMarket().GetTickers("umcbl")
	if err != nil {
		return nil, err
	}

	results := domain.FundingPricesOfSymbol{}
	for _, ticker := range tickers.Data {
		results.Set(ex.Name, strings.TrimSuffix(ticker.Symbol, "_UMCBL"), ticker.Last, ticker.FundingRate, "0")
	}

	return results, nil
}
