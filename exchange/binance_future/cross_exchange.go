package binance_future

import (
	"context"
	"crypto-exchange/exchange/strategy"
	"strconv"
)

func (ex *BinanceFuture) GetCrossExArbitrageInformation() (strategy.SymbolExchangeFundingPrice, error) {
	symbolPrices, err := ex.Client.NewListPricesService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	premiumIndexes, err := ex.Client.NewPremiumIndexService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	//results := []strategy.CrossExArbitrageInformation{}

	results := strategy.SymbolExchangeFundingPrice{}
	for _, symbolPrice := range symbolPrices {
		results.Set(ex.Name, string(symbolPrice.Symbol), symbolPrice.Price, "0", "0")
	}

	for _, premium := range premiumIndexes {
		results.SetSpecial(ex.Name, premium.Symbol, premium.LastFundingRate, strconv.FormatInt(premium.NextFundingTime, 10))
	}

	return results, nil
}
