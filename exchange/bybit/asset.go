package bybit

import (
	"crypto-exchange/exchange/domain"
	"strconv"

	"github.com/hirokisan/bybit/v2"
)

func (ex *Bybit) GetAllAsset() (*domain.ExchangeAsset, error) {
	// only can get balance of main account
	res, err := ex.Client.V5().Account().GetWalletBalance(bybit.AccountTypeUnified, nil)
	if err != nil {
		return nil, err
	}

	assets := []domain.Asset{}
	for _, record := range res.Result.List[0].Coin {
		balance, err := strconv.ParseFloat(record.WalletBalance, 64)
		if err != nil {
			return nil, err
		}

		assets = append(assets, domain.Asset{
			Coin:   string(record.Coin),
			Amount: balance,
		})
	}

	return &domain.ExchangeAsset{
		Name:   ex.Name,
		Assets: assets,
	}, nil
}
