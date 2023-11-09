package binance_future

import (
	"context"
	"strconv"

	"github.com/arensusu/crypto-app/exchange/types"
)

func (ex *BinanceFuture) GetAssets() ([]types.Asset, error) {
	assets := []types.Asset{}
	res, err := ex.Client.NewGetBalanceService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	for _, record := range res {
		amount, err := strconv.ParseFloat(record.Balance, 64)
		if err != nil {
			return nil, err
		}
		assets = append(assets, types.Asset{
			Coin:   record.Asset,
			Amount: amount,
		})
	}

	return assets, nil
}
