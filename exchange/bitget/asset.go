package bitget

import (
	"crypto-exchange/exchange/domain"
	"strconv"
)

func (ex *Bitget) GetAllAsset() (*domain.ExchangeAsset, error) {
	assets := []domain.Asset{}

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

	return &domain.ExchangeAsset{
		Name:   ex.Name(),
		Assets: assets,
	}, nil
}

func (ex *Bitget) GetSpotAssets() ([]domain.Asset, error) {
	spotRes, err := ex.Client.NewSpotAccountGetAccountAssetsLiteService().Do()
	if err != nil {
		return nil, err
	}

	assets := []domain.Asset{}
	for _, data := range spotRes.Data {
		availabe, err := strconv.ParseFloat(data.Available, 64)
		if err != nil {
			return nil, err
		}

		assets = append(assets, domain.Asset{
			Coin:   data.CoinName,
			Amount: availabe,
		})
	}

	return assets, nil
}

func (ex *Bitget) GetMixAssets() ([]domain.Asset, error) {
	spotRes, err := ex.Client.NewMixAccountGetAccountListService().ProductType("umcbl").Do()
	if err != nil {
		return nil, err
	}

	assets := []domain.Asset{}
	for _, data := range spotRes.Data {
		availabe, err := strconv.ParseFloat(data.Available, 64)
		if err != nil {
			return nil, err
		}

		assets = append(assets, domain.Asset{
			Coin:   data.MarginCoin,
			Amount: availabe,
		})
	}

	return assets, nil
}
