// https://binance-docs.github.io/apidocs/futures/en/
package binanceEx

import (
	"context"
	"funding-rate/exchange/strategy"
	"os"
	"strconv"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

type BinanceExchange struct {
	Name          string
	FuturesClient *futures.Client
}

func New() *BinanceExchange {
	client := binance.NewFuturesClient(os.Getenv("BINANCE_API_KEY"), os.Getenv("BINANCE_API_SECRET"))
	return &BinanceExchange{
		Name:          "Binance",
		FuturesClient: client,
	}
}

func (ex *BinanceExchange) GetCrossExArbitrageInformation(coin string) (*strategy.CrossExArbitrageInformation, error) {
	symbol := coin + "USDT"
	symbolPrices, err := ex.FuturesClient.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return nil, err
	}

	premiumIndexes, err := ex.FuturesClient.NewPremiumIndexService().Symbol(symbol).Do(context.Background())
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
