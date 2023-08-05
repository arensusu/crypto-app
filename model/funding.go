package model

import (
	"funding-rate/coinglass"
	"funding-rate/database"
)

type WatchList struct {
	ID     uint `gorm:"primarykey"`
	ChatID int64
	coinglass.Pair
}

func Contains(curList []coinglass.Pair, target coinglass.Pair) bool {
	for _, pair := range curList {
		if pair == target {
			return true
		}
	}
	return false
}

func AddFundingWatchList(chatID int64, exchange, symbol string) error {
	pairs, err := GetFundingWatchList(chatID)
	if err != nil {
		return err
	}
	pair := coinglass.Pair{Exchange: exchange, Symbol: symbol}
	if Contains(pairs, pair) {
		return nil
	}

	newWatchlist := WatchList{ChatID: chatID, Pair: pair}
	err = database.DB.Create(&newWatchlist).Error
	return err
}

func GetFundingWatchList(chatID int64) ([]coinglass.Pair, error) {
	var pairs []coinglass.Pair
	err := database.DB.Model(&WatchList{}).Where("chat_id=?", chatID).Find(&pairs).Error
	if err != nil {
		return []coinglass.Pair{}, err
	}

	return pairs, nil
}
