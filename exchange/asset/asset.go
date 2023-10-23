package asset

type Asset struct {
	Coin   string
	Amount float64
}

type AssetGetter interface {
	//GetAsset(string) Asset
	GetAllAsset() []Asset
}
