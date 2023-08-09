package user

import (
	"fmt"
	"funding-rate/domain"
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
			res, err := pair.GetFundingRate("h8", 3)
			if err != nil {
				notifications = append(notifications, domain.Notification{ChatID: user.ChatID, Message: "Cannot get data from coinglass.\n"})
				continue
			}

			data := res.Data
			rate := data[len(data)-1].Rate
			if rate < 0.0 {
				notifications = append(notifications, domain.Notification{ChatID: user.ChatID, Message: fmt.Sprintf("Alert: current funding rate of %s %s is %.4f\n", pair.Exchange, pair.Symbol, rate)})
			}
		}
	}
	return notifications
}
