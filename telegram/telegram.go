package telegram

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	NOTIFY_INTERVAL = 3600 * 8
)

func (handler *telegramHandler) Run() {
	handler.tgbot.Debug = debug
	handler.setCommand()

	go handler.notify()
	handler.commandReply()

}

func (handler *telegramHandler) notify() {
	for {
		now := time.Now().Unix()
		if now%NOTIFY_INTERVAL == 0 {
			handler.fundingNotify()
		}
		if now%60 == 0 {
			handler.priceAlert()
		}
		time.Sleep(1 * time.Second)
	}
}

func (handler *telegramHandler) setCommand() {
	deleteCommand := tgbotapi.NewDeleteMyCommands()
	if _, err := handler.tgbot.Request(deleteCommand); err != nil {
		panic(err)
	}

	cmds := []tgbotapi.BotCommand{
		{Command: "funding", Description: "Show funding rate of watchlist."},
		{Command: "remove", Description: "Remove a pair from watchlist."},
		{Command: "getfunding", Description: "Get funding rate of specific pair."},
		{Command: "add", Description: "Add a pair to watchlist."},
		{Command: "show", Description: "Show the watchlist."},
	}
	setCommand := tgbotapi.NewSetMyCommands(cmds...)
	if _, err := handler.tgbot.Request(setCommand); err != nil {
		panic(err)
	}
}

func (handler *telegramHandler) commandReply() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := handler.tgbot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		command := update.Message.Command()
		chatID := update.Message.Chat.ID
		text := update.Message.Text

		msg := ""
		switch command {
		case "start":
			msg = handler.start(chatID)
		case "funding":
			msg = handler.funding(chatID, text)
		case "perp":
			msg = handler.perp(chatID, text)
		case "addfunding":
			msg = handler.addFunding(chatID, text)
		case "addperp":
			msg = handler.addPerp(chatID, text)
		case "show":
			msg = handler.show(chatID)
		case "removefunding":
			msg = handler.removeFunding(chatID, text)
		case "removeperp":
			msg = handler.removePerp(chatID, text)
			// case "getfunding":
			// 	msg = handler.getFunding(chatID, text)
		}

		handler.sendMsg(chatID, msg)
	}
}
