package funding

import (
	"fmt"
	"funding-rate/coinglass"
	"funding-rate/domain"
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
		fundingRates, err := pair.GetFundingRate("h8", 100)
		if err != nil {
			reply += fmt.Sprintf("\nCannot obtain data of %s %s", pair.Exchange, pair.Symbol)
			fmt.Println(err)
			continue
		}
		reply += pair.ToMessage(&fundingRates)
	}
	return reply
}

func (usecase *FundingUseCase) NewFunding(chatID int64, message string) string {
	msg := strings.Split(message, " ")

	if len(msg) < 3 {
		return "Invalid."

	}

	pair := coinglass.Pair{Exchange: msg[1], Symbol: msg[2]}
	isExist, err := pair.IsExist()
	if err != nil {
		return "Cannot get data from Coinglass."

	}
	if !isExist {
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
