package watchlist

import (
	"crypto-exchange/exchange/domain"

	"gorm.io/gorm"
)

type WatchlistPostgresRepository struct {
	db *gorm.DB
}

func NewWatchlistPostgresRepository(db *gorm.DB) domain.WatchlistRepository {
	return &WatchlistPostgresRepository{db}
}

func (repo *WatchlistPostgresRepository) CreateFundingWatchlist(chatID int64, pair domain.Pair) error {
	watchlist := domain.Watchlist{
		ChatID: chatID,
		Pair:   pair,
	}

	if err := repo.db.Create(&watchlist).Error; err != nil {
		return err
	}
	return nil
}

func (repo *WatchlistPostgresRepository) RetrieveFundingWatchlists(chatID int64) ([]domain.Pair, error) {
	var pairs []domain.Pair

	err := repo.db.Model(&domain.Watchlist{}).Where("chat_id=?", chatID).Find(&pairs).Error
	if err != nil {
		return []domain.Pair{}, err
	}
	return pairs, nil
}

func (repo *WatchlistPostgresRepository) DeleteFundingWatchlist(chatID int64, pair domain.Pair) error {
	watchlist := domain.Watchlist{}
	err := repo.db.Where("chat_id = ? and exchange = ? and symbol = ?", chatID, pair.Exchange, pair.Symbol).Delete(&watchlist).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *WatchlistPostgresRepository) CreatePerpetualWatchlist(w domain.PerpetualWatchlist) error {
	err := repo.db.Create(&w).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *WatchlistPostgresRepository) RetrievePerpetualWatchlists() ([]domain.PerpetualWatchlist, error) {
	var watchlists []domain.PerpetualWatchlist
	err := repo.db.Find(&watchlists).Error
	if err != nil {
		return nil, err
	}
	return watchlists, nil
}

func (repo *WatchlistPostgresRepository) RetrievePerpetualWatchlistsOfUser(chatID int64) ([]domain.PerpetualWatchlist, error) {
	var watchlists []domain.PerpetualWatchlist
	err := repo.db.Where("chat_id = ?", chatID).Find(&watchlists).Error
	if err != nil {
		return nil, err
	}
	return watchlists, nil
}

func (repo *WatchlistPostgresRepository) DeletePerpetualWatchlist(watchlist domain.PerpetualWatchlist) error {
	err := repo.db.Where("chat_id=? and exchange=? and symbol=? and target_price=?", watchlist.ChatID, watchlist.Exchange, watchlist.Symbol, watchlist.TargetPrice).Delete(&watchlist).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *WatchlistPostgresRepository) CreatePerpPrevPrice(prevPrice domain.PrevPrice) error {
	err := repo.db.Create(&prevPrice).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *WatchlistPostgresRepository) RetrievePerpPrevPrice(pair domain.Pair) (domain.PrevPrice, error) {
	var prevPrice domain.PrevPrice
	err := repo.db.Where("exchange=? and symbol=?", pair.Exchange, pair.Symbol).First(&prevPrice).Error
	if err != nil {
		return domain.PrevPrice{}, err
	}
	return prevPrice, nil
}

func (repo *WatchlistPostgresRepository) UpdatePerpPrevPrice(prevPrice domain.PrevPrice) error {
	err := repo.db.Model(&domain.PrevPrice{}).Where("exchange=? and symbol=?", prevPrice.Exchange, prevPrice.Symbol).Update("price", prevPrice.Price).Error
	if err != nil {
		return err
	}
	return nil
}
