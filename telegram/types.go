package telegram

import (
	"fmt"
	"funding-rate/domain"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type telegramHandler struct {
	tgbot      *tgbotapi.BotAPI
	msgHistory map[int64][]string

	userUsecase      domain.UserUsecase
	watchlistUsecase domain.WatchlistUsecase
	fundingUsecase   domain.FundingUsecase
}

func NewTelegramHandler(tgbot *tgbotapi.BotAPI, userUsecase domain.UserUsecase, watchlistUsecase domain.WatchlistUsecase, fundingUsecase domain.FundingUsecase) telegramHandler {
	return telegramHandler{tgbot, map[int64][]string{}, userUsecase, watchlistUsecase, fundingUsecase}
}

func (handler *telegramHandler) start(id int64) string {
	err := handler.userUsecase.AddUser(id)
	if err != nil {
		return "Register failed. Please enter '/start' again."
	}
	return "Welcome!"
}

func (handler *telegramHandler) fundingNotify() {
	users, err := handler.userUsecase.GetUsers()
	if err != nil {
		log.Print(err)
		return
	}

	for _, user := range users {
		notifications, err := handler.fundingUsecase.GetFundingNotification(user.ChatID)
		if err != nil {
			fmt.Println(err)
		}

		for _, notify := range notifications {
			msg := fmt.Sprintf("Alert: current funding rate of %s %s is flipped (%.4f to %.4f)\n", notify.Exchange, notify.Symbol, notify.Previous, notify.Current)
			handler.sendMsg(user.ChatID, msg)
		}
	}
}

func (handler *telegramHandler) priceAlert() {
	watchlists, err := handler.watchlistUsecase.GetPerpetualWatchlists()
	if err != nil {
		log.Print(err)
		return
	}

	updateList := map[domain.Pair]float64{}
	for _, watchlist := range watchlists {
		perpData, err := handler.fundingUsecase.GetPerpData(watchlist.Exchange, watchlist.Symbol)
		if err != nil {
			log.Print(err)
			continue
		}
		prevPrice, err := handler.watchlistUsecase.GetPerpPrevPrice(watchlist.Pair)
		if err != nil {
			log.Print(err)
			continue
		}
		updateList[watchlist.Pair] = perpData.Price

		fmt.Println(perpData.Price)
		if prevPrice < watchlist.TargetPrice && perpData.Price > watchlist.TargetPrice {
			msg := fmt.Sprintf("Alert: %s %s is crossing UP %f", watchlist.Exchange, watchlist.Symbol, watchlist.TargetPrice)
			handler.sendMsg(watchlist.ChatID, msg)
		} else if prevPrice > watchlist.TargetPrice && perpData.Price < watchlist.TargetPrice {
			msg := fmt.Sprintf("Alert: %s %s is crossing DOWN %f", watchlist.Exchange, watchlist.Symbol, watchlist.TargetPrice)
			handler.sendMsg(watchlist.ChatID, msg)
		}
	}

	for pair, price := range updateList {
		if err := handler.watchlistUsecase.SetPerpPrevPrice(domain.PrevPrice{Pair: pair, Price: price}); err != nil {
			log.Print(err)
		}
	}
}
