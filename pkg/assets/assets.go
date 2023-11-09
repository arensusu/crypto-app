package assets

import (
	"crypto-exchange/exchange"
	"crypto-exchange/exchange/types"
	"fmt"
)

type AssetsFinder struct {
	Exchanges []exchange.Exchange
}

func NewAssetsFinder(exchanges ...exchange.Exchange) *AssetsFinder {
	return &AssetsFinder{Exchanges: exchanges}
}

func (f *AssetsFinder) GetAssets() map[string][]types.Asset {
	assets := map[string][]types.Asset{}
	for _, ex := range f.Exchanges {
		asset, err := ex.GetAssets()
		if err != nil {
			fmt.Println(fmt.Errorf("get asset fail: %w", err))
			continue
		}
		assets[ex.Name()] = asset
	}
	return assets
}
