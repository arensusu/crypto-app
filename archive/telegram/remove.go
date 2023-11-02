package telegram

import (
	"crypto-exchange/exchange/domain"
	"fmt"
	"strconv"
)

func (handler *telegramHandler) remove(id int64) string {
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
			return handler.removeFunding(id, history[2], history[3])
		}
		return "Target price?"
	case 5:
		handler.msgHistory[id] = []string{}
		return handler.removePrice(id, history[2], history[3], history[4])
	default:
		handler.msgHistory[id] = []string{}
		return "add error"
	}
}

func (handler *telegramHandler) removeFunding(id int64, exchange, symbol string) string {
	err := handler.watchlistUsecase.RemoveWatchlist(id, exchange, symbol)
	if err != nil {
		return "Cannot edit the watchlist. Please try again."
	}
	return fmt.Sprintf("%s %s has been removed.", exchange, symbol)
}

func (handler *telegramHandler) removePrice(id int64, exchange, symbol, targetString string) string {
	target, err := strconv.ParseFloat(targetString, 64)
	if err != nil {
		return "Please enter a valid target price."
	}

	watchlist := domain.PerpetualWatchlist{
		Watchlist: domain.Watchlist{
			ChatID: id,
			Pair: domain.Pair{
				Exchange: exchange,
				Symbol:   symbol,
			},
		},
		TargetPrice: target,
	}
	if err := handler.watchlistUsecase.RemovePerpetualWatchlist(watchlist); err != nil {
		return "Cannot add the watchlist. Please try again."
	}

	return fmt.Sprintf("Alert when %s touches %f on %s is removed", symbol, target, exchange)
}
