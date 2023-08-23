package user

import (
	"fmt"
	"funding-rate/domain"
	"math"
)

type UserUseCase struct {
	userRepo    domain.IUserRepository
	fundingRepo domain.IFundingRepository
}

func NewUserUseCase(userRepo domain.IUserRepository, fundingRepo domain.IFundingRepository) domain.IUserUseCase {
	return &UserUseCase{userRepo, fundingRepo}
}

func (usecase *UserUseCase) NewUser(chatID int64) string {
	err := usecase.userRepo.AddUser(chatID)
	if err != nil {
		return "Already register."
	}
	return "Welcome."
}

func (usecase *UserUseCase) GetUsersNotification() []domain.Notification {
	users, _ := usecase.userRepo.RetrieveUsers()

	notifications := []domain.Notification{}
	for _, user := range users {
		watchlist, err := usecase.fundingRepo.GetFundingWatchList(user.ChatID)
		if err != nil {
			notifications = append(notifications, domain.Notification{ChatID: user.ChatID, Message: "Cannot get data of watchlist.\n"})
			continue
		}

		for _, pair := range watchlist {
			history, err := usecase.fundingRepo.GetFundingHistory(pair)
			if err != nil {
				notifications = append(notifications, domain.Notification{ChatID: user.ChatID, Message: "Cannot get data from coinglass.\n"})
				continue
			}

			prev := history[len(history)-2]
			curr := history[len(history)-1]
			if math.Abs(prev+curr) < math.Abs(prev-curr) {
				notifications = append(notifications, domain.Notification{ChatID: user.ChatID, Message: fmt.Sprintf("Alert: current funding rate of %s %s is flipped (%.4f to %.4f)\n", pair.Exchange, pair.Symbol, prev, curr)})
			}
		}
	}
	return notifications
}
