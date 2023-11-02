package assets

import (
	"crypto-exchange/exchange/domain"
	"fmt"
)

type AssetsUsecase struct {
	Clients []domain.AssetGetter
}

func NewAssetsUsecase(clients []any) *AssetsUsecase {
	assetGetters := []domain.AssetGetter{}
	for _, c := range clients {
		if assetGetter, ok := c.(domain.AssetGetter); ok {
			assetGetters = append(assetGetters, assetGetter)
		}
	}

	return &AssetsUsecase{Clients: assetGetters}
}

func (u *AssetsUsecase) GetAssets() {
	for _, client := range u.Clients {
		asset, err := client.GetAllAsset()
		if err != nil {
			panic(fmt.Errorf("get asset fail: %w", err))
		}
		fmt.Println(asset)
	}
}
