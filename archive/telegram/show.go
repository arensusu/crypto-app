package telegram

import "fmt"

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
