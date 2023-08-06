package controller

import (
	"fmt"
	"funding-rate/coinglass"
	"funding-rate/model"
	"strings"
)

func Funding(chatID int64) string {
	pairs, err := model.GetFundingWatchList(chatID)
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

func NewFunding(chatID int64, message string) string {
	msg := strings.Split(message, " ")

	if len(msg) < 3 {
		return "Invalid."

	}

	pair := coinglass.Pair{Exchange: msg[1], Symbol: msg[2]}
	isExist, err := isPairExist(pair)
	if err != nil {
		return "Cannot get data from Coinglass."

	}
	if !isExist {
		return "Pair is not exist."

	}

	isFollowing, err := isPairFollowing(chatID, pair)
	if err != nil {
		return "Cannot get data of following pairs."

	}
	if isFollowing {
		return "Already following."

	}

	if err := model.AddFundingWatchList(chatID, pair); err != nil {
		fmt.Println(err)
		return "Added Failed."

	}
	return "Added Successfully."

}

func isPairExist(pair coinglass.Pair) (bool, error) {
	response, err := pair.GetFundingRate("h8", 1)
	if err != nil {
		return false, err
	}
	return response.Msg != "pair unknown", nil
}

func contains(curList []coinglass.Pair, target coinglass.Pair) bool {
	for _, pair := range curList {
		if pair == target {
			return true
		}
	}
	return false
}

func isPairFollowing(chatID int64, pair coinglass.Pair) (bool, error) {
	pairs, err := model.GetFundingWatchList(chatID)
	if err != nil {
		return false, err
	}
	return contains(pairs, pair), nil
}
