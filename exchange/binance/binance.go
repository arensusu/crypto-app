// https://binance-docs.github.io/apidocs/futures/en/
package binanceEx

import (
	"context"
	"funding-rate/exchange/strategy"
	"strconv"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

type BinanceExchange struct {
	Name          string
	FuturesClient *futures.Client
}

func New(apiKey, apiSecret string) *BinanceExchange {
	client := binance.NewFuturesClient(apiKey, apiSecret)
	return &BinanceExchange{
		Name:          "Binance",
		FuturesClient: client,
	}
}

func (ex *BinanceExchange) GetCrossExArbitrageResponse(coin string) (*strategy.CrossExArbitrageResponse, error) {
	symbol := coin + "USDT"
	prices, err := ex.FuturesClient.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return nil, err
	}

	premiumIndexes, err := ex.FuturesClient.NewPremiumIndexService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return nil, err
	}

	price := prices[0]
	premiumIndex := premiumIndexes[0]

	return &strategy.CrossExArbitrageResponse{
		ExchangeName:    ex.Name,
		LastPrice:       price.Price,
		FundingRate:     premiumIndex.LastFundingRate,
		NextFundingTime: strconv.FormatInt((premiumIndex.NextFundingTime), 10),
	}, nil
}
