package user

import (
	"funding-rate/coinglass"
	"funding-rate/domain"
	"funding-rate/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewUser(t *testing.T) {
	mockUserRepo := mocks.NewIUserRepository(t)
	mockFundingRepo := mocks.NewIFundingRepository(t)
	mockUserRepo.On("AddUser", int64(1)).Return(nil)
	usecase := NewUserUseCase(mockUserRepo, mockFundingRepo)

	result := usecase.NewUser(1)

	assert.Equal(t, "Welcome.", result)
}

func Test_GetUserNotification(t *testing.T) {
	users := []domain.User{{ChatID: int64(1)}}
	pairs := []coinglass.Pair{{Exchange: "Bybit", Symbol: "ETHUSDT"}}
	history := []float64{0.1, -0.1}
	mockUserRepo := mocks.NewIUserRepository(t)
	mockFundingRepo := mocks.NewIFundingRepository(t)
	mockUserRepo.On("RetrieveUsers").Return(users, nil)
	mockFundingRepo.On("GetFundingWatchList", int64(1)).Return(pairs, nil)
	mockFundingRepo.On("GetFundingHistory", pairs[0]).Return(history, nil)
	usecase := NewUserUseCase(mockUserRepo, mockFundingRepo)

	result := usecase.GetUsersNotification()

	assert.Equal(t, 1, len(result))
	assert.Equal(t, int64(1), result[0].ChatID)
	assert.Equal(t, "Alert: current funding rate of Bybit ETHUSDT is flipped (0.1000 to -0.1000)\n", result[0].Message)
}
