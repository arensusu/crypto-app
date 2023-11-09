package binance

import (
	"context"
	"strconv"

	"github.com/arensusu/crypto-app/exchange/types"
)

func (ex *Binance) GetAssets() ([]types.Asset, error) {
	assets := []types.Asset{}
	res, err := ex.Client.NewGetUserAsset().Do(context.Background())
	if err != nil {
		return nil, err
	}

	for _, record := range res {
		freeAmount, err := strconv.ParseFloat(record.Free, 64)
		if err != nil {
			return nil, err
		}

		lockedAmount, err := strconv.ParseFloat(record.Locked, 64)
		if err != nil {
			return nil, err
		}

		assets = append(assets, types.Asset{
			Coin:   record.Asset,
			Amount: freeAmount + lockedAmount,
		})
	}

	return assets, nil
}
