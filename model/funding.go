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

func AddFundingWatchList(chatID int64, pair coinglass.Pair) error {
	newWatchlist := WatchList{ChatID: chatID, Pair: pair}
	err := database.DB.Create(&newWatchlist).Error
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
