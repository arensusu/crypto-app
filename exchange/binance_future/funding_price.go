package binance_future

import (
	"context"
	"crypto-exchange/domain"
	"strconv"
)

func (ex *BinanceFuture) GetFundingAndPrices() (domain.FundingPricesOfSymbol, error) {
	symbolPrices, err := ex.Client.NewListPricesService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	premiumIndexes, err := ex.Client.NewPremiumIndexService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	//results := []domain.CrossExArbitrageInformation{}

	results := domain.FundingPricesOfSymbol{}
	for _, symbolPrice := range symbolPrices {
		results.Set(ex.Name, string(symbolPrice.Symbol), symbolPrice.Price, "0", "0")
	}

	for _, premium := range premiumIndexes {
		results.SetSpecial(ex.Name, premium.Symbol, premium.LastFundingRate, strconv.FormatInt(premium.NextFundingTime, 10))
	}

	return results, nil
}
