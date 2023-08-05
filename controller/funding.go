package controller

import (
	"fmt"
	"funding-rate/model"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Funding(bot *tgbotapi.BotAPI, chatID int64) {
	reply := tgbotapi.NewMessage(chatID, "")
	reply.Text = "Funding Rate\n"

	pairs, err := model.GetFundingWatchList(chatID)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, pair := range pairs {
		fundingRates, err := pair.GetFundingRate("h8", 100)
		if err != nil {
			reply.Text += fmt.Sprintf("\nCannot obtain data of %s %s", pair.Exchange, pair.Symbol)
			fmt.Println(err)
			continue
		}
		reply.Text += pair.ToMessage(&fundingRates)
	}
	if _, err := bot.Send(reply); err != nil {
		fmt.Println(err)
	}
}

func NewFunding(bot *tgbotapi.BotAPI, chatID int64, message string) {
	reply := tgbotapi.NewMessage(chatID, "")
	msg := strings.Split(message, " ")
	if len(msg) < 3 {
		reply.Text = "Invalid."
	} else {
		if err := model.AddFundingWatchList(chatID, msg[1], msg[2]); err != nil {
			fmt.Println(err)
			reply.Text = "Added Failed."
		} else {
			reply.Text = "Added Successfully."
		}
		if _, err := bot.Send(reply); err != nil {
			fmt.Println(err)
		}
	}
}
