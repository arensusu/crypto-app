package binance_future

import (
	"context"
	"errors"
	"strconv"

	"github.com/arensusu/crypto-app/exchange/types"
)

func (ex *BinanceFuture) GetFundingAndPrice(symbol string) (*types.FundingFeeArbitrage, error) {
	symbolPrices, err := ex.Client.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return nil, err
	}

	premiumIndexes, err := ex.Client.NewPremiumIndexService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return nil, err
	}

	orderTickers, err := ex.Client.NewListBookTickersService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return nil, err
	}

	if len(symbolPrices) != 1 || len(premiumIndexes) != 1 || len(orderTickers) != 1 {
		return nil, errors.New("response data length error")
	}

	symbolPrice := symbolPrices[0]
	premiumIndex := premiumIndexes[0]
	orderTicker := orderTickers[0]

	result := &types.FundingFeeArbitrage{
		LastPrice:   symbolPrice.Price,
		FundingRate: premiumIndex.LastFundingRate,
		FundingTime: strconv.FormatInt(premiumIndex.NextFundingTime, 10),
		Bid1Price:   orderTicker.BidPrice,
		Bid1Size:    orderTicker.BidQuantity,
		Ask1Price:   orderTicker.AskPrice,
		Ask1Size:    orderTicker.AskQuantity,
	}

	return result, nil
}
