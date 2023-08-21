package telegram

import (
	"fmt"
	"funding-rate/domain"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	NOTIFY_INTERVAL = 3600 * 8
)

var (
	tgApiToken = os.Getenv("TELEGRAM_APITOKEN")
	debug      = true
)

type telegramHandler struct {
	FundingUseCase domain.IFundingUseCase
	UserUseCase    domain.IUserUseCase
}

func NewTelegramHandler(fundingCase domain.IFundingUseCase, userCase domain.IUserUseCase) telegramHandler {
	return telegramHandler{fundingCase, userCase}
}

func (handler *telegramHandler) Start() {
	bot, err := tgbotapi.NewBotAPI(tgApiToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = debug

	go handler.Notify(bot)
	handler.CommandReply(bot)

}

func (handler *telegramHandler) FundingNotify(bot *tgbotapi.BotAPI) {
	notifications := handler.UserUseCase.GetUsersNotification()
	for _, notification := range notifications {
		reply := tgbotapi.NewMessage(notification.ChatID, notification.Message)
		if _, err := bot.Send(reply); err != nil {
			fmt.Println(err)
		}
	}
}

func (handler *telegramHandler) Notify(bot *tgbotapi.BotAPI) {
	for {
		now := time.Now().Unix()
		if now%NOTIFY_INTERVAL == 0 {
			handler.FundingNotify(bot)
		}
		time.Sleep(1 * time.Second)
	}
}

func (handler *telegramHandler) CommandReply(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	msgHistory := []string{}
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		reply := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		// Extract the command from the Message.
		command := update.Message.Command()
		switch command {
		case "start":
			reply.Text = handler.UserUseCase.NewUser(update.Message.Chat.ID)
		case "help":
			reply.Text = "I understand /sayhi and /status."
		case "show":
			reply.Text = handler.FundingUseCase.ShowFundingWatchList(update.Message.Chat.ID)
		case "remove":
			msgHistory = append(msgHistory, "remove")
			reply.Text = handler.FundingUseCase.ShowFundingWatchList(update.Message.Chat.ID)
			reply.Text += "\nWhich trading pair do you want to remove? Please enter a index of pair.\n"
		case "funding":
			reply.Text = handler.FundingUseCase.Funding(update.Message.Chat.ID)
		case "newfunding":
			reply.Text = handler.FundingUseCase.NewFunding(update.Message.Chat.ID, update.Message.Text)
		case "":
			if len(msgHistory) == 0 {
				continue
			}
			if msgHistory[0] == "remove" {
				msgHistory = []string{}
				reply.Text = handler.FundingUseCase.RemoveFromFundingWatchList(update.Message.Chat.ID, update.Message.Text)
			}

		default:
			reply.Text = "I don't know that command"
		}

		if _, err := bot.Send(reply); err != nil {
			fmt.Println(err)
		}
	}

}
