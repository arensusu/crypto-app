package binance_future

import (
	"context"
	"crypto-exchange/domain"
	"errors"
	"strconv"
)

func (ex *BinanceFuture) GetFundingAndPrice(symbol string) (*domain.FundingPrice, error) {
	symbolPrices, err := ex.Client.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return nil, err
	}

	premiumIndexes, err := ex.Client.NewPremiumIndexService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return nil, err
	}

	if len(symbolPrices) != 1 || len(premiumIndexes) != 1 {
		return nil, errors.New("response data length error")
	}

	symbolPrice := symbolPrices[0]
	premiumIndex := premiumIndexes[0]

	result := domain.NewFundingPrice(ex.Name, symbolPrice.Price, premiumIndex.LastFundingRate,
		strconv.FormatInt(premiumIndex.NextFundingTime, 10))

	return result, nil
}

func (ex *BinanceFuture) GetFundingAndPrices() (domain.FundingPricesOfSymbol, error) {
	symbolPrices, err := ex.Client.NewListPricesService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	premiumIndexes, err := ex.Client.NewPremiumIndexService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	results := domain.FundingPricesOfSymbol{}
	for _, symbolPrice := range symbolPrices {
		results.Set(ex.Name, string(symbolPrice.Symbol), symbolPrice.Price, "0", "0")
	}

	for _, premium := range premiumIndexes {
		results.SetSpecial(ex.Name, premium.Symbol, premium.LastFundingRate, strconv.FormatInt(premium.NextFundingTime, 10))
	}

	return results, nil
}
