package assets

import (
	"crypto-exchange/exchange/asset"
	"fmt"
)

type AssetsUsecase struct {
	Clients []asset.AssetGetter
}

func NewAssetsUsecase(clients []any) *AssetsUsecase {
	assetGetters := []asset.AssetGetter{}
	for _, c := range clients {
		if assetGetter, ok := c.(asset.AssetGetter); ok {
			assetGetters = append(assetGetters, assetGetter)
		}
	}

	return &AssetsUsecase{Clients: assetGetters}
}

func (u *AssetsUsecase) GetAssets() {
	assets := []asset.Asset{}
	for _, client := range u.Clients {
		asset, err := client.GetAllAsset()
		if err != nil {
			panic(fmt.Errorf("get asset fail: %w", err))
		}
		assets = append(assets, asset...)
	}

	fmt.Println(assets)
}
