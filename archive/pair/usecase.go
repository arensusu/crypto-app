package pair

import (
	"crypto-exchange/exchange/domain"
	"fmt"
	"math"
	"time"
)

type PairUsecase struct {
	watchlistRepo domain.WatchlistRepository
	fundingRepo   domain.PairRepository
}

func NewPairUsecase(watchlistRepo domain.WatchlistRepository, fundingRepo domain.PairRepository) domain.PairUsecase {
	return &PairUsecase{watchlistRepo, fundingRepo}
}

func (usecase *PairUsecase) GetFundingData(exchange, symbol string) (domain.FundingData, error) {
	history, err := usecase.fundingRepo.GetFundingHistory(exchange, symbol)
	if err != nil {
		return domain.FundingData{}, fmt.Errorf("get data from coinglass failed: %w", err)
	}

	history = history[:len(history)-1] // last one is current expect funding rate
	data := domain.FundingData{
		Pair:    domain.Pair{Exchange: exchange, Symbol: symbol},
		Last100: totalFundingRate(history, 100),
		Last30:  totalFundingRate(history, 30),
		Last:    totalFundingRate(history, 1),
	}
	return data, nil
}

func (usecase *PairUsecase) GetPerpData(exchange, symbol string) (domain.PerpData, error) {
	data, err := usecase.fundingRepo.GetPerpetualMarket(exchange, symbol)
	if err != nil {
		return domain.PerpData{}, fmt.Errorf("get data from coinglass failed: %w", err)
	}

	perp := domain.PerpData{
		Pair: domain.Pair{
			Exchange: exchange,
			Symbol:   data.OriginalSymbol,
		},
		Price:              data.Price,
		PriceChangePercent: data.PriceChangePercent,
		FundingRate:        domain.FundingRate(data.FundingRate),
		NextFundingTime:    time.UnixMilli(data.NextFundingTime),
	}
	return perp, nil
}

func (usecase *PairUsecase) GetFundingDataOfUser(chatID int64) ([]domain.FundingData, error) {
	pairs, err := usecase.watchlistRepo.RetrieveFundingWatchlists(chatID)
	if err != nil {
		return []domain.FundingData{}, fmt.Errorf("get data failed: %w", err)
	}

	datas := []domain.FundingData{}
	for _, pair := range pairs {
		data, err := usecase.GetFundingData(pair.Exchange, pair.Symbol)
		if err != nil {
			return []domain.FundingData{}, err
		}
		datas = append(datas, data)
	}
	return datas, nil
}

func (usecase *PairUsecase) GetFundingNotification(chatID int64) ([]domain.FundingNotification, error) {
	pairs, err := usecase.watchlistRepo.RetrieveFundingWatchlists(chatID)
	if err != nil {
		return []domain.FundingNotification{}, fmt.Errorf("get data failed: %w", err)
	}

	notifications := []domain.FundingNotification{}
	for _, pair := range pairs {
		history, err := usecase.fundingRepo.GetFundingHistory(pair.Exchange, pair.Symbol)
		if err != nil {
			return []domain.FundingNotification{}, fmt.Errorf("get data failed from coinglass: %w", err)
		}

		prev := history[len(history)-2]
		curr := history[len(history)-1]
		if math.Abs(prev+curr) < math.Abs(prev-curr) {
			notification := domain.FundingNotification{
				Pair:     pair,
				Previous: domain.FundingRate(prev),
				Current:  domain.FundingRate(curr),
			}
			notifications = append(notifications, notification)
		}
	}
	return notifications, nil
}

func (usecase *PairUsecase) AddFundingSearched(chatID int64, exchange, symbol string) error {
	pair := domain.Pair{Exchange: exchange, Symbol: symbol}
	err := usecase.fundingRepo.CreateFundingSearched(chatID, pair)
	if err != nil {
		return fmt.Errorf("create data failed: %w", err)
	}
	return nil
}

func (usecase *PairUsecase) GetLastFiveFundingSearched(chatID int64) ([]domain.Pair, error) {
	searched, err := usecase.fundingRepo.RetrieveFundingSearched(chatID)
	if err != nil {
		return []domain.Pair{}, fmt.Errorf("get data failed: %w", err)
	}

	return searched, nil
}

// func (usecase *FundingUseCase) Funding(chatID int64) string {
// 	pairs, err := usecase.fundingRepo.GetFundingWatchList(chatID)
// 	if err != nil {
// 		fmt.Println(err)
// 		return "Cannot get data of following pairs."
// 	}
// 	if len(pairs) == 0 {
// 		return "Please use /newfunding to follow a pair."
// 	}

// 	reply := "Funding Rate\n"
// 	for _, pair := range pairs {
// 		history, err := usecase.fundingRepo.GetFundingHistory(pair)
// 		if err != nil {
// 			reply += fmt.Sprintf("\nCannot obtain data of %s %s", pair.Exchange, pair.Symbol)
// 			fmt.Println(err)
// 			continue
// 		}
// 		reply += toMessage(pair, history)
// 	}
// 	return reply
// }

// func toMessage(pair coinglass.Pair, history []float64) string {
// 	length := len(history)
// 	msg := fmt.Sprintf("\n%s %s\n", pair.Exchange, pair.Symbol)

// 	if length >= 100 {
// 		total := totalFundingRate(history, 100)
// 		msg += fmt.Sprintf("Total of last 100: %.4f%%, APR: %.2f%%\n", total, total/100*3*365)
// 	}

// 	if length >= 30 {
// 		total := totalFundingRate(history, 30)
// 		msg += fmt.Sprintf("Total of last 30:  %.4f%%, APR: %.2f%%\n", total, total/30*3*365)
// 	}

// 	msg += fmt.Sprintf("Last: %.4f%%\n", history[length-1])

// 	return msg
// }

func totalFundingRate(data []float64, period int) domain.FundingRate {
	total := 0.0
	for _, fundingRate := range data[len(data)-period:] {
		total += fundingRate
	}
	return domain.FundingRate(total)
}

// func (usecase *UserUseCase) GetUsersNotification() []domain.Notification {
// 	users, _ := usecase.userRepo.RetrieveUsers()

// 	notifications := []domain.Notification{}
// 	for _, user := range users {
// 		watchlist, err := usecase.fundingRepo.GetFundingWatchList(user.ChatID)
// 		if err != nil {
// 			notifications = append(notifications, domain.Notification{ChatID: user.ChatID, Message: "Cannot get data of watchlist.\n"})
// 			continue
// 		}

// 		for _, pair := range watchlist {
// 			history, err := usecase.fundingRepo.GetFundingHistory(pair)
// 			if err != nil {
// 				notifications = append(notifications, domain.Notification{ChatID: user.ChatID, Message: "Cannot get data from coinglass.\n"})
// 				continue
// 			}

// 			prev := history[len(history)-2]
// 			curr := history[len(history)-1]
// 			if math.Abs(prev+curr) < math.Abs(prev-curr) {
// 				notifications = append(notifications, domain.Notification{ChatID: user.ChatID, Message: fmt.Sprintf("Alert: current funding rate of %s %s is flipped (%.4f to %.4f)\n", pair.Exchange, pair.Symbol, prev, curr)})
// 			}
// 		}
// 	}
// 	return notifications
// }
