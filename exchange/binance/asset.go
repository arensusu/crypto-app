package binance

import (
	"context"
	"crypto-exchange/domain"
	"strconv"
)

func (ex *Binance) GetAllAsset() ([]domain.Asset, error) {
	assets := []domain.Asset{}
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

		assets = append(assets, domain.Asset{
			Coin:   record.Asset,
			Amount: freeAmount + lockedAmount,
		})
	}

	return assets, nil
}
