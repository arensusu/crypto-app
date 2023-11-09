package bitget

import (
	"crypto-exchange/exchange/types"
	"strconv"
)

func (ex *Bitget) GetAssets() ([]types.Asset, error) {
	assets := []types.Asset{}

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

	return assets, nil
}

func (ex *Bitget) GetSpotAssets() ([]types.Asset, error) {
	spotRes, err := ex.Client.NewSpotAccountGetAccountAssetsLiteService().Do()
	if err != nil {
		return nil, err
	}

	assets := []types.Asset{}
	for _, data := range spotRes.Data {
		availabe, err := strconv.ParseFloat(data.Available, 64)
		if err != nil {
			return nil, err
		}

		assets = append(assets, types.Asset{
			Coin:   data.CoinName,
			Amount: availabe,
		})
	}

	return assets, nil
}

func (ex *Bitget) GetMixAssets() ([]types.Asset, error) {
	spotRes, err := ex.Client.NewMixAccountGetAccountListService().ProductType("umcbl").Do()
	if err != nil {
		return nil, err
	}

	assets := []types.Asset{}
	for _, data := range spotRes.Data {
		availabe, err := strconv.ParseFloat(data.Available, 64)
		if err != nil {
			return nil, err
		}

		assets = append(assets, types.Asset{
			Coin:   data.MarginCoin,
			Amount: availabe,
		})
	}

	return assets, nil
}
