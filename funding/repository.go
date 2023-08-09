package funding

import (
	"funding-rate/coinglass"
	"funding-rate/domain"

	"gorm.io/gorm"
)

type PostgresFundingRepository struct {
	db *gorm.DB
}

func NewPostgresFundingRepository(db *gorm.DB) domain.IFundingRepository {
	return &PostgresFundingRepository{db}
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
