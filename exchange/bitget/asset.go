package bitget

import (
	"crypto-exchange/exchange/asset"
	"strconv"
)

func (ex *Bitget) GetAllAsset() (*asset.ExchangeAsset, error) {
	assets := []asset.Asset{}

	spotAssets, err := ex.GetSpotAssets()
	if err != nil {
		return nil, err
	}
	assets = append(assets, spotAssets...)

	mixAssets, err := ex.GetMixAssets()
	if err != nil {
		return nil, err
	}
	assets = append(assets, mixAssets...)

	return &asset.ExchangeAsset{
		Name:   ex.Name,
		Assets: assets,
	}, nil
}

func (ex *Bitget) GetSpotAssets() ([]asset.Asset, error) {
	spotRes, err := ex.Client.NewSpotAccountGetAccountAssetsLiteService().Do()
	if err != nil {
		return nil, err
	}

	assets := []asset.Asset{}
	for _, data := range spotRes.Data {
		availabe, err := strconv.ParseFloat(data.Available, 64)
		if err != nil {
			return nil, err
		}

		assets = append(assets, asset.Asset{
			Coin:   data.CoinName,
			Amount: availabe,
		})
	}

	return assets, nil
}

func (ex *Bitget) GetMixAssets() ([]asset.Asset, error) {
	spotRes, err := ex.Client.NewMixAccountGetAccountListService().ProductType("umcbl").Do()
	if err != nil {
		return nil, err
	}

	assets := []asset.Asset{}
	for _, data := range spotRes.Data {
		availabe, err := strconv.ParseFloat(data.Available, 64)
		if err != nil {
			return nil, err
		}

		assets = append(assets, asset.Asset{
			Coin:   data.MarginCoin,
			Amount: availabe,
		})
	}

	return assets, nil
}
