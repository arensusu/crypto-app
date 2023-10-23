package binance_future

import (
	"context"
	"crypto-exchange/exchange/strategy"
	"strconv"
)

func (ex *BinanceFuture) GetCrossExArbitrageInformation(coin string) (*strategy.CrossExArbitrageInformation, error) {
	symbol := coin + "USDT"
	symbolPrices, err := ex.Client.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return nil, err
	}

	premiumIndexes, err := ex.Client.NewPremiumIndexService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return nil, err
	}

	symbolPrice := symbolPrices[0]
	premiumIndex := premiumIndexes[0]

	price, _ := strconv.ParseFloat(symbolPrice.Price, 64)
	fundingRate, _ := strconv.ParseFloat(premiumIndex.LastFundingRate, 64)
	return &strategy.CrossExArbitrageInformation{
		ExchangeName:    ex.Name,
		LastPrice:       price,
		FundingRate:     fundingRate,
		NextFundingTime: premiumIndex.NextFundingTime,
	}, nil
}
