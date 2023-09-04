package telegram

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	FUNDING_ALERT_INTERVAL = 8 * 60 * 60
	PRICE_ALERT_INTERVAL   = 15 * 60
)

var (
	tgApiToken = os.Getenv("TELEGRAM_APITOKEN")
	debug, _   = strconv.ParseBool(os.Getenv("DEBUG"))
)

func NewTelegramBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(tgApiToken)
	if err != nil {
		panic(err)
	}
	return bot
}

func (handler *telegramHandler) sendMsg(id int64, msg string) {
	reply := tgbotapi.NewMessage(id, msg)
	if _, err := handler.tgbot.Send(reply); err != nil {
		fmt.Println("cannot send message in telegram\n", err)
	}
}

func (handler *telegramHandler) Run() {
	handler.tgbot.Debug = debug
	handler.setCommand()

	go handler.notify()
	handler.commandReply()

}

func (handler *telegramHandler) notify() {
	for {
		now := time.Now().Unix()
		if now%FUNDING_ALERT_INTERVAL == 0 {
			handler.fundingNotify()
		}
		if now%PRICE_ALERT_INTERVAL == 0 {
			handler.priceAlert()
		}
		time.Sleep(1 * time.Second)
	}
}

func getCommandFromFile() []tgbotapi.BotCommand {
	file, err := os.Open("commands.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cmds := []tgbotapi.BotCommand{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ": ")
		cmds = append(cmds, tgbotapi.BotCommand{
			Command:     text[0],
			Description: text[1],
		})
	}
	return cmds
}

func (handler *telegramHandler) setCommand() {
	deleteCommand := tgbotapi.NewDeleteMyCommands()
	if _, err := handler.tgbot.Request(deleteCommand); err != nil {
		panic(err)
	}

	cmds := getCommandFromFile()
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

		text := update.Message.Text
		chatID := update.Message.Chat.ID

		msg := ""
		if _, exist := handler.msgHistory[chatID]; !exist {
			handler.msgHistory[chatID] = []string{}
		}
		handler.msgHistory[chatID] = append(handler.msgHistory[chatID], text)
		// switch command {
		// case "start":
		// 	msg = handler.start(chatID)
		// case "funding":
		// 	msg = handler.funding(chatID, text)
		// case "perp":
		// 	msg = handler.perp(chatID, text)
		// // case "addfunding":
		// // 	msg = handler.addFunding(chatID, text)
		// // case "addprice":
		// // 	msg = handler.addPerp(chatID, text)
		// case "show":
		// 	msg = handler.show(chatID)
		// case "removefunding":
		// 	msg = handler.removeFunding(chatID, text)
		// case "removeprice":
		// 	msg = handler.removePerp(chatID, text)
		// //
		// case "":
		// }

		msg = handler.commandStep(chatID)
		handler.sendMsg(chatID, msg)
		fmt.Println(handler.msgHistory[chatID])
	}
}

func (handler *telegramHandler) commandStep(id int64) string {
	history := handler.msgHistory[id]
	if len(history) == 0 {
		return "Please enter command first."
	}

	switch history[0] {
	case "start":
		return handler.start(id)
	case "/funding":
		return handler.funding(id)
	case "/perp":
		return handler.perp(id)
	case "/show":
		handler.msgHistory[id] = []string{}
		return handler.show(id)
	case "/add":
		return handler.add(id)
	case "/remove":
		return handler.remove(id)
	default:
		handler.msgHistory[id] = []string{}
		return "Error"
	}
}
