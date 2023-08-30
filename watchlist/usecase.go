package watchlist

import (
	"errors"
	"fmt"
	"funding-rate/domain"
)

type WatchlistUsecase struct {
	watchlistRepo domain.WatchlistRepository
}

func NewWatchlistUsecase(repo domain.WatchlistRepository) domain.WatchlistUsecase {
	return &WatchlistUsecase{repo}
}

func (usecase *WatchlistUsecase) AddWatchlist(chatID int64, exchange, symbol string) error {
	pair := domain.Pair{Exchange: exchange, Symbol: symbol}
	pairs, err := usecase.GetWatchlists(chatID)
	if err != nil {
		return err
	}
	if contains(pairs, pair) {
		return errors.New("pair is existed")
	}

	if err := usecase.watchlistRepo.CreateFundingWatchlist(chatID, pair); err != nil {
		return fmt.Errorf("create data failed: %w", err)
	}
	return nil
}

func (u *WatchlistUsecase) AddPerpetualWatchlist(watchlist domain.PerpetualWatchlist) error {
	err := u.watchlistRepo.CreatePerpetualWatchlist(watchlist)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *WatchlistUsecase) GetWatchlists(chatID int64) ([]domain.Pair, error) {
	pairs, err := usecase.watchlistRepo.RetrieveFundingWatchlists(chatID)
	if err != nil {
		return []domain.Pair{}, fmt.Errorf("get data failed: %w", err)
	}
	return pairs, nil
}

func (usecase *WatchlistUsecase) GetPerpetualWatchlists() ([]domain.PerpetualWatchlist, error) {
	watchlists, err := usecase.watchlistRepo.RetrievePerpetualWatchlists()
	if err != nil {
		return []domain.PerpetualWatchlist{}, fmt.Errorf("get data failed: %w", err)
	}
	return watchlists, nil
}

func (usecase *WatchlistUsecase) GetPerpetualWatchlistsOfUser(chatID int64) ([]domain.PerpetualWatchlist, error) {
	watchlists, err := usecase.watchlistRepo.RetrievePerpetualWatchlistsOfUser(chatID)
	if err != nil {
		return []domain.PerpetualWatchlist{}, fmt.Errorf("get data failed: %w", err)
	}
	return watchlists, nil
}

func (usecase *WatchlistUsecase) RemoveWatchlist(chatID int64, exchange, symbol string) error {
	pair := domain.Pair{Exchange: exchange, Symbol: symbol}

	err := usecase.watchlistRepo.DeleteFundingWatchlist(chatID, pair)
	if err != nil {
		return fmt.Errorf("delete data failed: %w", err)
	}
	return nil
}

func (u *WatchlistUsecase) RemovePerpetualWatchlist(watchlist domain.PerpetualWatchlist) error {
	err := u.watchlistRepo.DeletePerpetualWatchlist(watchlist)
	if err != nil {
		return err
	}
	return nil
}

func contains(curList []domain.Pair, target domain.Pair) bool {
	for _, pair := range curList {
		if pair == target {
			return true
		}
	}
	return false
}

func (u *WatchlistUsecase) GetPerpPrevPrice(pair domain.Pair) (float64, error) {
	prevPrice, err := u.watchlistRepo.RetrievePerpPrevPrice(pair)
	if err != nil {
		return 0.0, err
	}

	if prevPrice.Overdated {
		return 0.0, nil
	}
	return prevPrice.Price, nil
}

func (u *WatchlistUsecase) SetPerpPrevPrice(prevPrice domain.PrevPrice) error {
	_, err := u.GetPerpPrevPrice(prevPrice.Pair)
	if err != nil {
		return u.watchlistRepo.CreatePerpPrevPrice(prevPrice)
	}
	return u.watchlistRepo.UpdatePerpPrevPrice(prevPrice)
}
