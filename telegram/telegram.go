package telegram

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	NOTIFY_INTERVAL = 3600 * 8
)

func (handler *telegramHandler) Run() {
	handler.tgbot.Debug = debug

	go handler.notify()
	handler.commandReply()

}

func (handler *telegramHandler) fundingNotify() {
	notifications := handler.UserUseCase.GetUsersNotification()
	for _, notification := range notifications {
		reply := tgbotapi.NewMessage(notification.ChatID, notification.Message)
		if _, err := handler.tgbot.Send(reply); err != nil {
			fmt.Println(err)
		}
	}
}

func (handler *telegramHandler) notify() {
	for {
		now := time.Now().Unix()
		if now%NOTIFY_INTERVAL == 0 {
			handler.fundingNotify()
		}
		time.Sleep(1 * time.Second)
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

		// Extract the command from the Message.
		command := update.Message.Command()
		chatID := update.Message.Chat.ID
		text := update.Message.Text
		switch command {
		case "start":
			handler.start(chatID)
		case "showfunding":
			handler.showfunding(chatID)
		case "remove":
			handler.preRemove(chatID)
		case "getfunding":
			handler.getFunding(chatID)
		case "add":
			handler.add(chatID, text)
		case "":
			if len(handler.msgHistory) == 0 {
				continue
			}
			if handler.msgHistory[0] == "remove" {
				handler.remove(chatID, text)
			}
		}
	}
}
