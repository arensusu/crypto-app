package domain

import (
	"fmt"
	"funding-rate/coinglass"
	"time"
)

type FundingSearched struct {
	ID     uint `gorm:"primarykey"`
	ChatID int64
	Pair
	CreatedAt time.Time
}

type FundingRepository interface {
	GetFundingHistory(exchange, symbol string) ([]float64, error)
	GetPerpetualMarket(exchange, symbol string) (coinglass.PerpetualMarket, error)

	CreateFundingSearched(chatID int64, pair Pair) error
	RetrieveFundingSearched(chatID int64) ([]Pair, error)
}

type FundingRate float64

func (f FundingRate) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%.4f", float64(f))
	return []byte(s), nil
}

type FundingData struct {
	Pair
	Last100 FundingRate `json:"last100"`
	Last30  FundingRate `json:"last30"`
	Last    FundingRate `json:"last"`
}

type FundingNotification struct {
	Pair
	Previous FundingRate
	Current  FundingRate
}

type PerpData struct {
	Pair
	Price              float64     `json:"price"`
	PriceChangePercent float64     `json:"changePercent"`
	FundingRate        FundingRate `json:"fundingRate"`
	NextFundingTime    time.Time   `json:"nextFundingTime"`
}

type FundingUsecase interface {
	GetFundingData(exchange, symbol string) (FundingData, error)
	GetPerpData(exchange, symbol string) (PerpData, error)
	GetFundingDataOfUser(chatID int64) ([]FundingData, error)
	GetFundingNotification(chatID int64) ([]FundingNotification, error)

	GetLastFiveFundingSearched(chatID int64) ([]Pair, error)
	AddFundingSearched(chatID int64, exchange, symbol string) error
}
