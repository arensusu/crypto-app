package binance_future

import (
	"context"
	"crypto-exchange/exchange/domain"
	"strconv"
)

func (ex *BinanceFuture) GetAllAsset() (*domain.ExchangeAsset, error) {
	assets := []domain.Asset{}
	res, err := ex.Client.NewGetBalanceService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	for _, record := range res {
		amount, err := strconv.ParseFloat(record.Balance, 64)
		if err != nil {
			return nil, err
		}
		assets = append(assets, domain.Asset{
			Coin:   record.Asset,
			Amount: amount,
		})
	}

	return &domain.ExchangeAsset{
		Name:   ex.Name,
		Assets: assets,
	}, nil
}
