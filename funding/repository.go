package funding

import (
	"funding-rate/coinglass"
	"funding-rate/domain"

	"gorm.io/gorm"
)

type PostgresFundingRepository struct {
	db  *gorm.DB
	api *coinglass.CoinglassApi
}

func NewPostgresFundingRepository(db *gorm.DB, api *coinglass.CoinglassApi) domain.IFundingRepository {
	return &PostgresFundingRepository{db, api}
}

func (repo *PostgresFundingRepository) AddFundingWatchList(chatID int64, pair coinglass.Pair) error {
	newWatchlist := domain.WatchList{ChatID: chatID, Pair: pair}
	err := repo.db.Create(&newWatchlist).Error
	return err
}

func (repo *PostgresFundingRepository) GetFundingWatchList(chatID int64) ([]coinglass.Pair, error) {
	var pairs []coinglass.Pair
	err := repo.db.Model(&domain.WatchList{}).Where("chat_id=?", chatID).Find(&pairs).Error
	if err != nil {
		return []coinglass.Pair{}, err
	}

	return pairs, nil
}

func (repo *PostgresFundingRepository) GetFundingHistory(pair coinglass.Pair) ([]float64, error) {
	response, err := repo.api.GetFundingRate(pair, "h8", 100)
	if err != nil {
		return []float64{}, err
	}

	history := []float64{}
	for _, record := range response.Data {
		history = append(history, record.Rate)
	}
	return history, nil
}

func (repo *PostgresFundingRepository) DeleteFundingWatchList(chatID int64, pair coinglass.Pair) error {
	watchlist := domain.WatchList{}
	err := repo.db.Where("chat_id = ? and exchange = ? and symbol = ?", chatID, pair.Exchange, pair.Symbol).Delete(&watchlist).Error
	return err
}
