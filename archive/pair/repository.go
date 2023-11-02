package pair

import (
	"crypto-exchange/coinglass"
	"crypto-exchange/exchange/domain"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type PairPostgresRepository struct {
	db  *gorm.DB
	api *coinglass.CoinglassApi
}

func NewPairPostgresRepository(db *gorm.DB, api *coinglass.CoinglassApi) domain.PairRepository {
	return &PairPostgresRepository{db, api}
}

func (repo *PairPostgresRepository) GetFundingHistory(exchange, symbol string) ([]float64, error) {
	responseData, err := repo.api.GetFundingRateUSDHistory(symbol, "h8")
	if err != nil {
		return []float64{}, err
	}

	dataList, isExist := responseData.DataMap[exchange]
	if !isExist {
		return []float64{}, errors.New("exchange is not exist")
	}

	return dataList, nil
}

func (repo *PairPostgresRepository) GetPerpetualMarket(exchange, symbol string) (coinglass.PerpetualMarket, error) {
	responseData, err := repo.api.GetPerpetualMarket(symbol)
	if err != nil {
		return coinglass.PerpetualMarket{}, err
	}

	for _, data := range responseData {
		if data.ExchangeName == exchange {
			return data, nil
		}
	}
	return coinglass.PerpetualMarket{}, errors.New("exchange is not exist")
}

func (repo *PairPostgresRepository) CreateFundingSearched(chatID int64, pair domain.Pair) error {
	searched := domain.FundingSearched{
		ChatID: chatID,
		Pair:   pair,
	}
	err := repo.db.Create(&searched).Error
	if err != nil {
		return fmt.Errorf("create data failed: %w", err)
	}
	return nil
}

func (repo *PairPostgresRepository) RetrieveFundingSearched(chatID int64) ([]domain.Pair, error) {
	var pairs []domain.Pair

	err := repo.db.Model(&domain.FundingSearched{}).Order("created_at desc").Where("chat_id=?", chatID).Limit(5).Find(&pairs).Error
	if err != nil {
		return []domain.Pair{}, err
	}
	return pairs, nil
}
