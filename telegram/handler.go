package telegram

import (
	"fmt"
	"funding-rate/domain"
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
	tgbot          *tgbotapi.BotAPI
	msgHistory     []string
	FundingUseCase domain.IFundingUseCase
	UserUseCase    domain.IUserUseCase
}

func NewTelegramBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(tgApiToken)
	if err != nil {
		panic(err)
	}
	return bot
}

func NewTelegramHandler(tgbot *tgbotapi.BotAPI, fundingCase domain.IFundingUseCase, userCase domain.IUserUseCase) telegramHandler {
	return telegramHandler{tgbot, []string{}, fundingCase, userCase}
}

func (handler *telegramHandler) sendMsg(id int64, msg string) error {
	reply := tgbotapi.NewMessage(id, msg)
	if _, err := handler.tgbot.Send(reply); err != nil {
		return fmt.Errorf("cannot send message in telegram: %w", err)
	}
	return nil
}

func (handler *telegramHandler) start(id int64) {
	msg := handler.UserUseCase.NewUser(id)
	if err := handler.sendMsg(id, msg); err != nil {
		fmt.Println(err)
	}
}

func (handler *telegramHandler) showfunding(id int64) {
	msg := handler.FundingUseCase.Funding(id)
	if err := handler.sendMsg(id, msg); err != nil {
		fmt.Println(err)
	}
}

func (handler *telegramHandler) preRemove(id int64) {
	handler.msgHistory = append(handler.msgHistory, "remove")
	msg := handler.FundingUseCase.ShowFundingWatchList(id)
	msg += "\nWhich trading pair do you want to remove? Please enter a index of pair.\n"
	if err := handler.sendMsg(id, msg); err != nil {
		fmt.Println(err)
	}
}

func (handler *telegramHandler) remove(chatID int64, msg string) {
	handler.msgHistory = []string{}

	index, err := strconv.Atoi(msg)
	reply := ""
	if err != nil {
		reply = "Invalid message. Please Enter a valid index again."
	} else {
		reply = handler.FundingUseCase.RemoveFromFundingWatchList(chatID, index)
	}
	if err := handler.sendMsg(chatID, reply); err != nil {
		fmt.Println(err)
	}
}

func (handler *telegramHandler) getFunding(id int64) {
	msg := handler.FundingUseCase.Funding(id)
	if err := handler.sendMsg(id, msg); err != nil {
		fmt.Println(err)
	}
}

func (handler *telegramHandler) add(id int64, msg string) {
	words := strings.Split(msg, " ")

	reply := ""
	if len(msg) < 3 {
		reply = "Invalid."
	} else {
		reply = handler.FundingUseCase.NewFunding(id, words[1], words[2])
	}
	if err := handler.sendMsg(id, reply); err != nil {
		fmt.Println(err)
	}
}
