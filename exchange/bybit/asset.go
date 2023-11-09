package bybit

import (
	"strconv"

	"github.com/arensusu/crypto-app/exchange/types"

	"github.com/hirokisan/bybit/v2"
)

func (ex *Bybit) GetAssets() ([]types.Asset, error) {
	// only can get balance of main account
	res, err := ex.Client.V5().Account().GetWalletBalance(bybit.AccountTypeUnified, nil)
	if err != nil {
		return nil, err
	}

	assets := []types.Asset{}
	for _, record := range res.Result.List[0].Coin {
		balance, err := strconv.ParseFloat(record.WalletBalance, 64)
		if err != nil {
			return nil, err
		}

		assets = append(assets, types.Asset{
			Coin:   string(record.Coin),
			Amount: balance,
		})
	}

	return assets, nil
}
