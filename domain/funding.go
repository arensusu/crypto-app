package domain

import "funding-rate/coinglass"

type WatchList struct {
	ID     uint `gorm:"primarykey"`
	ChatID int64
	coinglass.Pair
}

type IFundingUseCase interface {
	NewFunding(chatID int64, message string) string
	Funding(chatID int64) string
	ShowFundingWatchList(chatID int64) string
	RemoveFromFundingWatchList(chatID int64, message string) string
}

type IFundingRepository interface {
	AddFundingWatchList(chatID int64, pair coinglass.Pair) error
	GetFundingWatchList(chatID int64) ([]coinglass.Pair, error)
	GetFundingHistory(pair coinglass.Pair) ([]float64, error)
	DeleteFundingWatchList(chatID int64, pair coinglass.Pair) error
}
