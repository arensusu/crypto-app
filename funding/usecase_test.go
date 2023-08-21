package funding

import (
	"funding-rate/coinglass"
	"funding-rate/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Funding(t *testing.T) {
	pairs := []coinglass.Pair{{Exchange: "Bybit", Symbol: "ETHUSDT"}}
	history := getFakeHistory(100)
	mockRepo := mocks.NewIFundingRepository(t)
	mockRepo.On("GetFundingWatchList", int64(1)).Return(pairs, nil)
	mockRepo.On("GetFundingHistory", pairs[0]).Return(history, nil)
	fundingUseCase := NewFundingUseCase(mockRepo)

	result := fundingUseCase.Funding(1)

	expectResult := "Funding Rate\n\nBybit ETHUSDT\nTotal of last 100: 1.0000%, APR: 10.95%\nTotal of last 30:  0.3000%, APR: 10.95%\nLast: 0.0100%\n"
	assert.Equal(t, expectResult, result)
}

func getFakeHistory(period int) []float64 {
	result := make([]float64, period)
	for i := range result {
		result[i] = 0.01
	}
	return result
}

func Test_NewFunding(t *testing.T) {
	newPair := coinglass.Pair{Exchange: "Binance", Symbol: "BTCUSDT"}
	followingPairs := []coinglass.Pair{{Exchange: "Bybit", Symbol: "ETHUSDT"}}
	mockRepo := mocks.NewIFundingRepository(t)
	mockRepo.On("GetFundingHistory", newPair).Return([]float64{0.1, 0.1}, nil)
	mockRepo.On("GetFundingWatchList", int64(1)).Return(followingPairs, nil)
	mockRepo.On("AddFundingWatchList", int64(1), newPair).Return(nil)
	fundingUseCase := NewFundingUseCase(mockRepo)

	result := fundingUseCase.NewFunding(1, "/newfunding Binance BTCUSDT")

	assert.Equal(t, "Added Successfully.", result)
}

func Test_ShowFunding(t *testing.T) {
	pairs := []coinglass.Pair{{Exchange: "Bybit", Symbol: "ETHUSDT"}}
	mockRepo := mocks.NewIFundingRepository(t)
	mockRepo.On("GetFundingWatchList", int64(1)).Return(pairs, nil)
	fundingUseCase := NewFundingUseCase(mockRepo)

	result := fundingUseCase.ShowFundingWatchList(1)

	expectResult := "Watchlist of funding rate:\n1. Bybit ETHUSDT\n"
	assert.Equal(t, expectResult, result)
}

func Test_DeleteFunding(t *testing.T) {
	pairs := []coinglass.Pair{{Exchange: "Bybit", Symbol: "ETHUSDT"}}
	mockRepo := mocks.NewIFundingRepository(t)
	mockRepo.On("GetFundingWatchList", int64(1)).Return(pairs, nil)
	mockRepo.On("DeleteFundingWatchList", int64(1), pairs[0]).Return(nil)
	fundingUseCase := NewFundingUseCase(mockRepo)

	result := fundingUseCase.RemoveFromFundingWatchList(1, "1")

	assert.Equal(t, "Remove Bybit ETHUSDT from watchlist successfully.", result)
}
