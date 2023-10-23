package binance

import (
	"context"
	"crypto-exchange/exchange/asset"
	"strconv"
)

func (ex *Binance) GetAllAsset() ([]asset.Asset, error) {
	assets := []asset.Asset{}
	spotAsset, err := ex.GetAllAssetInSpot()
	if err != nil {
		return []asset.Asset{}, err
	}
	assets = append(assets, spotAsset...)

	return assets, nil
}

func (ex *Binance) GetAllAssetInSpot() ([]asset.Asset, error) {
	assets := []asset.Asset{}
	res, err := ex.Client.NewGetUserAsset().Do(context.Background())
	if err != nil {
		return assets, err
	}

	for _, record := range res {
		freeAmount, err := strconv.ParseFloat(record.Free, 64)
		if err != nil {
			return []asset.Asset{}, err
		}

		lockedAmount, err := strconv.ParseFloat(record.Locked, 64)
		if err != nil {
			return []asset.Asset{}, err
		}

		assets = append(assets, asset.Asset{
			Coin:   record.Asset,
			Amount: freeAmount + lockedAmount,
		})
	}

	return assets, nil
}
