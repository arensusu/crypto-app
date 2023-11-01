package domain

type Asset struct {
	Coin   string
	Amount float64
}

type ExchangeAsset struct {
	Name   string
	Assets []Asset
}

type AssetGetter interface {
	//GetAsset(string) Asset
	GetAllAsset() (ExchangeAsset, error)
}
