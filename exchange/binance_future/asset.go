package binance_future

import (
	"context"
	"crypto-exchange/exchange/asset"
	"fmt"
	"strconv"
)

func (ex *BinanceFuture) GetAllAsset() ([]asset.Asset, error) {
	assets := []asset.Asset{}
	res, err := ex.Client.NewGetBalanceService().Do(context.Background())
	if err != nil {
		return assets, err
	}

	for _, record := range res {
		fmt.Println(record)
		amount, err := strconv.ParseFloat(record.Balance, 64)
		if err != nil {
			return []asset.Asset{}, err
		}
		assets = append(assets, asset.Asset{
			Coin:   record.Asset,
			Amount: amount,
		})
	}

	return assets, nil
}
