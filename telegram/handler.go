package telegram

import (
	"fmt"
	"funding-rate/domain"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	tgApiToken = os.Getenv("TELEGRAM_APITOKEN")
	debug, _   = strconv.ParseBool(os.Getenv("DEBUG"))
)

type telegramHandler struct {
	tgbot      *tgbotapi.BotAPI
	msgHistory []string

	userUsecase      domain.UserUsecase
	watchlistUsecase domain.WatchlistUsecase
	fundingUsecase   domain.FundingUsecase
}

func NewTelegramBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(tgApiToken)
	if err != nil {
		panic(err)
	}
	return bot
}

func NewTelegramHandler(tgbot *tgbotapi.BotAPI, userUsecase domain.UserUsecase, watchlistUsecase domain.WatchlistUsecase, fundingUsecase domain.FundingUsecase) telegramHandler {
	return telegramHandler{tgbot, []string{}, userUsecase, watchlistUsecase, fundingUsecase}
}

func (handler *telegramHandler) sendMsg(id int64, msg string) {
	reply := tgbotapi.NewMessage(id, msg)
	if _, err := handler.tgbot.Send(reply); err != nil {
		fmt.Println("cannot send message in telegram\n", err)
	}
}

func (handler *telegramHandler) start(id int64) string {
	err := handler.userUsecase.AddUser(id)
	if err != nil {
		return "Register failed. Please enter '/start' again."
	}
	return "Welcome!"
}

// func (handler *telegramHandler) funding(id int64) string {
// 	fundingDatas, err := handler.fundingUsecase.GetFundingDataOfUser(id)
// 	if err != nil {
// 		fmt.Println(err)
// 		return "Cannot get funding rate. Please try again later."
// 	}

// 	msg := "Funding Rate\n"
// 	for _, data := range fundingDatas {
// 		msg += fmt.Sprintf("\n%s %s\n", data.Exchange, data.Symbol)
// 		msg += fmt.Sprintf("Total of last 100: %.4f%%\n", data.Last100)
// 		msg += fmt.Sprintf("Total of last 30: %.4f%%\n", data.Last30)
// 		msg += fmt.Sprintf("Last: %.4f%%\n", data.Last)
// 	}
// 	return msg
// }

func (handler *telegramHandler) perp(id int64, text string) string {
	words := strings.Split(text, " ")
	if len(words) < 3 {
		return "Invalid input. Please enter again."
	}

	perp, err := handler.fundingUsecase.GetPerpData(words[1], words[2])
	if err != nil {
		fmt.Println(err)
		return "Cannot get the perpetual data. Please try again later."
	}

	msg := fmt.Sprintf("%s %s\n", perp.Exchange, perp.Symbol)
	msg += fmt.Sprintf("Price: %f (%.2f%%)\n", perp.Price, perp.PriceChangePercent)
	msg += fmt.Sprintf("Next Funding: %.4f", perp.FundingRate)
	return msg
}

func (handler *telegramHandler) show(id int64) string {
	pairs, err := handler.watchlistUsecase.GetWatchlists(id)
	if err != nil {
		return "Cannot get watchlist. Please try again later."
	}

	msg := "Funding Watchlist:\n"
	for i, pair := range pairs {
		msg += fmt.Sprintf("%d. %s %s\n", i+1, pair.Exchange, pair.Symbol)
	}

	watchlists, err := handler.watchlistUsecase.GetPerpetualWatchlistsOfUser(id)
	if err != nil {
		return "Cannot get watchlist. Please try again later."
	}
	msg += "Perpetual Watchlist:\n"
	for i, watchlist := range watchlists {
		msg += fmt.Sprintf("%d. %s %s %f\n", i+1, watchlist.Exchange, watchlist.Symbol, watchlist.TargetPrice)
	}
	return msg
}

func (handler *telegramHandler) removeFunding(chatID int64, msg string) string {
	words := strings.Split(msg, " ")

	if len(words) < 3 {
		return "Invalid input. Please enter again."
	}
	err := handler.watchlistUsecase.RemoveWatchlist(chatID, words[1], words[2])
	if err != nil {
		return "Cannot edit the watchlist. Please try again."
	}
	return fmt.Sprintf("%s %s has been removed.", words[1], words[2])
}

func (handler *telegramHandler) removePerp(chatID int64, text string) string {
	words := strings.Split(text, " ")
	if len(words) < 4 {
		return "Invalid input. Please enter again."
	}
	target, err := strconv.ParseFloat(words[3], 64)
	if err != nil {
		return "Please enter a valid target price."
	}

	watchlist := domain.PerpetualWatchlist{
		Watchlist: domain.Watchlist{
			ChatID: chatID,
			Pair: domain.Pair{
				Exchange: words[1],
				Symbol:   words[2],
			},
		},
		TargetPrice: target,
	}
	if err := handler.watchlistUsecase.RemovePerpetualWatchlist(watchlist); err != nil {
		return "Cannot add the watchlist. Please try again."
	}

	return fmt.Sprintf("Alert when %s touches %f on %s is removed", words[2], target, words[1])
}

func (handler *telegramHandler) funding(chatID int64, text string) string {
	words := strings.Split(text, " ")

	if len(words) < 3 {
		return "Invalid input. Please enter again."
	}
	data, err := handler.fundingUsecase.GetFundingData(words[1], words[2])
	if err != nil {
		return "Cannot get the funding rate. Please try again."
	}

	if err := handler.fundingUsecase.AddFundingSearched(chatID, words[1], words[2]); err != nil {
		fmt.Println(err)
	}

	msg := "Funding Rate\n"
	msg += fmt.Sprintf("\n%s %s\n", data.Exchange, data.Symbol)
	msg += fmt.Sprintf("Total of last 100: %.4f%%\n", data.Last100)
	msg += fmt.Sprintf("Total of last 30: %.4f%%\n", data.Last30)
	msg += fmt.Sprintf("Last: %.4f%%\n", data.Last)
	return msg
}

func (handler *telegramHandler) addFunding(chatID int64, text string) string {
	words := strings.Split(text, " ")

	if len(words) < 3 {
		return "Invalid input. Please enter again."
	}
	err := handler.watchlistUsecase.AddWatchlist(chatID, words[1], words[2])
	if err != nil {
		return "Cannot edit the watchlist. Please try again."
	}
	return fmt.Sprintf("%s %s has been added.", words[1], words[2])
}

func (handler *telegramHandler) addPerp(chatID int64, text string) string {
	words := strings.Split(text, " ")
	if len(words) < 4 {
		return "Invalid input. Please enter again."
	}
	target, err := strconv.ParseFloat(words[3], 64)
	if err != nil {
		return "Please enter a valid target price."
	}

	watchlist := domain.PerpetualWatchlist{
		Watchlist: domain.Watchlist{
			ChatID: chatID,
			Pair: domain.Pair{
				Exchange: words[1],
				Symbol:   words[2],
			},
		},
		TargetPrice: target,
	}
	if err := handler.watchlistUsecase.AddPerpetualWatchlist(watchlist); err != nil {
		return "Cannot add the watchlist. Please try again."
	}

	perpData, err := handler.fundingUsecase.GetPerpData(watchlist.Exchange, watchlist.Symbol)
	if err != nil {
		log.Fatal(err)
		return "Cannot get perpetual data."
	}
	if err := handler.watchlistUsecase.SetPerpPrevPrice(domain.PrevPrice{Pair: watchlist.Pair, Price: perpData.Price}); err != nil {
		log.Fatal(err)
		return "Cannot set perpetual price data."
	}

	return fmt.Sprintf("Alert when %s touches %f on %s", words[2], target, words[1])
}

func (handler *telegramHandler) fundingNotify() {
	users, err := handler.userUsecase.GetUsers()
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
		return
	}

	for _, watchlist := range watchlists {
		perpData, err := handler.fundingUsecase.GetPerpData(watchlist.Exchange, watchlist.Symbol)
		if err != nil {
			log.Fatal(err)
			continue
		}
		prevPrice, err := handler.watchlistUsecase.GetPerpPrevPrice(watchlist.Pair)
		if err != nil {
			log.Fatal(err)
			continue
		}

		if (prevPrice < watchlist.TargetPrice && perpData.Price > watchlist.TargetPrice) ||
			(prevPrice > watchlist.TargetPrice && perpData.Price < watchlist.TargetPrice) {
			msg := fmt.Sprintf("Alert: %s %s crossovers target price: %f", watchlist.Exchange, watchlist.Symbol, watchlist.TargetPrice)
			handler.sendMsg(watchlist.ChatID, msg)
		}

		if err := handler.watchlistUsecase.SetPerpPrevPrice(domain.PrevPrice{Pair: watchlist.Pair, Price: perpData.Price}); err != nil {
			log.Fatal(err)
		}
	}
}
