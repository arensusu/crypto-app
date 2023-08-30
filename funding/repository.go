package funding

import (
	"errors"
	"fmt"
	"funding-rate/coinglass"
	"funding-rate/domain"

	"gorm.io/gorm"
)

type FundingPostgresRepository struct {
	db  *gorm.DB
	api *coinglass.CoinglassApi
}

func NewFundingPostgresRepository(db *gorm.DB, api *coinglass.CoinglassApi) domain.FundingRepository {
	return &FundingPostgresRepository{db, api}
}

func (repo *FundingPostgresRepository) GetFundingHistory(exchange, symbol string) ([]float64, error) {
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

func (repo *FundingPostgresRepository) GetPerpetualMarket(exchange, symbol string) (coinglass.PerpetualMarket, error) {
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

func (repo *FundingPostgresRepository) CreateFundingSearched(chatID int64, pair domain.Pair) error {
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

func (repo *FundingPostgresRepository) RetrieveFundingSearched(chatID int64) ([]domain.Pair, error) {
	var pairs []domain.Pair

	err := repo.db.Model(&domain.FundingSearched{}).Where("chat_id=?", chatID).Find(&pairs).Error
	if err != nil {
		return []domain.Pair{}, err
	}
	return pairs, nil
}
