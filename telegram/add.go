package telegram

import (
	"fmt"
	"funding-rate/domain"
	"log"
	"strconv"
)

func (handler *telegramHandler) add(id int64) string {
	history := handler.msgHistory[id]
	switch len(history) {
	case 1:
		return "funding or perp?"
	case 2:
		return "Exchange?"
	case 3:
		return "Symbol?"
	case 4:
		if history[1] == "funding" {
			handler.msgHistory[id] = []string{}
			return handler.addFunding(id, history[2], history[3])
		}
		return "Target price?"
	case 5:
		handler.msgHistory[id] = []string{}
		return handler.addPrice(id, history[2], history[3], history[4])
	default:
		handler.msgHistory[id] = []string{}
		return "add error"
	}
}

func (handler *telegramHandler) addFunding(id int64, exchange, symbol string) string {
	err := handler.watchlistUsecase.AddWatchlist(id, exchange, symbol)
	if err != nil {
		return "Cannot edit the watchlist. Please try again."
	}
	return fmt.Sprintf("%s %s has been added.", exchange, symbol)
}

func (handler *telegramHandler) addPrice(chatID int64, exchange, symbol, targetString string) string {
	target, err := strconv.ParseFloat(targetString, 64)
	if err != nil {
		return "Please enter a valid target price."
	}

	watchlist := domain.PerpetualWatchlist{
		Watchlist: domain.Watchlist{
			ChatID: chatID,
			Pair: domain.Pair{
				Exchange: exchange,
				Symbol:   symbol,
			},
		},
		TargetPrice: target,
	}
	if err := handler.watchlistUsecase.AddPerpetualWatchlist(watchlist); err != nil {
		return "Cannot add the watchlist. Please try again."
	}

	perpData, err := handler.fundingUsecase.GetPerpData(watchlist.Exchange, watchlist.Symbol)
	if err != nil {
		log.Print(err)
		return "Cannot get perpetual data."
	}
	if err := handler.watchlistUsecase.SetPerpPrevPrice(domain.PrevPrice{Pair: watchlist.Pair, Price: perpData.Price}); err != nil {
		log.Print(err)
		return "Cannot set perpetual price data."
	}

	return fmt.Sprintf("Alert when %s touches %f on %s", symbol, target, exchange)
}
