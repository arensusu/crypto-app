package bybit

import (
	"crypto-exchange/exchange/asset"
	"strconv"

	"github.com/hirokisan/bybit/v2"
)

func (ex *Bybit) GetAllAsset() (*asset.ExchangeAsset, error) {
	// only can get balance of main account
	res, err := ex.Client.V5().Account().GetWalletBalance(bybit.AccountTypeUnified, nil)
	if err != nil {
		return nil, err
	}

	assets := []asset.Asset{}
	for _, record := range res.Result.List[0].Coin {
		balance, err := strconv.ParseFloat(record.WalletBalance, 64)
		if err != nil {
			return nil, err
		}

		assets = append(assets, asset.Asset{
			Coin:   string(record.Coin),
			Amount: balance,
		})
	}

	return &asset.ExchangeAsset{
		Name:   ex.Name,
		Assets: assets,
	}, nil
}
