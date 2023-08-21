package funding

import (
	"fmt"
	"funding-rate/coinglass"
	"funding-rate/domain"
	"strconv"
	"strings"
)

type FundingUseCase struct {
	fundingRepo domain.IFundingRepository
}

func NewFundingUseCase(repo domain.IFundingRepository) domain.IFundingUseCase {
	return &FundingUseCase{repo}
}

func (usecase *FundingUseCase) Funding(chatID int64) string {
	pairs, err := usecase.fundingRepo.GetFundingWatchList(chatID)
	if err != nil {
		fmt.Println(err)
		return "Cannot get data of following pairs."
	}
	if len(pairs) == 0 {
		return "Please use /newfunding to follow a pair."
	}

	reply := "Funding Rate\n"
	for _, pair := range pairs {
		history, err := usecase.fundingRepo.GetFundingHistory(pair)
		if err != nil {
			reply += fmt.Sprintf("\nCannot obtain data of %s %s", pair.Exchange, pair.Symbol)
			fmt.Println(err)
			continue
		}
		reply += toMessage(pair, history)
	}
	return reply
}

func toMessage(pair coinglass.Pair, history []float64) string {
	length := len(history)
	msg := fmt.Sprintf("\n%s %s\n", pair.Exchange, pair.Symbol)

	if length >= 100 {
		total := totalFundingRate(history, 100)
		msg += fmt.Sprintf("Total of last 100: %.4f%%, APR: %.2f%%\n", total, total/100*3*365)
	}

	if length >= 30 {
		total := totalFundingRate(history, 30)
		msg += fmt.Sprintf("Total of last 30:  %.4f%%, APR: %.2f%%\n", total, total/30*3*365)
	}

	msg += fmt.Sprintf("Last: %.4f%%\n", history[length-1])

	return msg
}

func totalFundingRate(data []float64, period int) float64 {
	total := 0.0
	for _, fundingRate := range data[len(data)-period:] {
		total += fundingRate
	}
	return total
}

func (usecase *FundingUseCase) NewFunding(chatID int64, message string) string {
	msg := strings.Split(message, " ")

	if len(msg) < 3 {
		return "Invalid."
	}

	pair := coinglass.Pair{Exchange: msg[1], Symbol: msg[2]}
	history, err := usecase.fundingRepo.GetFundingHistory(pair)
	if err != nil {
		return "Cannot get data from Coinglass."
	}
	if len(history) == 0 {
		return "Pair is not exist."
	}

	isFollowing, err := usecase.isPairFollowing(chatID, pair)
	if err != nil {
		return "Cannot get data of following pairs."
	}
	if isFollowing {
		return "Already following."
	}

	if err := usecase.fundingRepo.AddFundingWatchList(chatID, pair); err != nil {
		fmt.Println(err)
		return "Added Failed."
	}
	return "Added Successfully."
}

func contains(curList []coinglass.Pair, target coinglass.Pair) bool {
	for _, pair := range curList {
		if pair == target {
			return true
		}
	}
	return false
}

func (usecase *FundingUseCase) isPairFollowing(chatID int64, pair coinglass.Pair) (bool, error) {
	pairs, err := usecase.fundingRepo.GetFundingWatchList(chatID)
	if err != nil {
		return false, err
	}
	return contains(pairs, pair), nil
}

func (usecase *FundingUseCase) ShowFundingWatchList(chatID int64) string {
	pairs, err := usecase.fundingRepo.GetFundingWatchList(chatID)
	if err != nil {
		fmt.Println(err)
		return "Cannot get data of following pairs."
	}
	if len(pairs) == 0 {
		return "You do not following any trading pair."
	}

	result := "Watchlist of funding rate:\n"
	for i, pair := range pairs {
		result += fmt.Sprintf("%d. %s %s\n", i+1, pair.Exchange, pair.Symbol)
	}
	return result
}

func (usecase *FundingUseCase) RemoveFromFundingWatchList(chatID int64, message string) string {
	index, err := strconv.Atoi(message)
	if err != nil {
		return "Invalid message. Please Enter a valid index again."
	}
	pairs, err := usecase.fundingRepo.GetFundingWatchList(chatID)
	if err != nil {
		fmt.Println(err)
		return "Cannot get data of following pairs."
	}
	if len(pairs) < index {
		return "Invalid index. Please Enter a valid index again."
	}

	pair := pairs[index-1]
	err = usecase.fundingRepo.DeleteFundingWatchList(chatID, pair)
	if err != nil {
		fmt.Println(err)
		return "Cannot delete the following pair."
	}
	return fmt.Sprintf("Remove %s %s from watchlist successfully.", pair.Exchange, pair.Symbol)
}
